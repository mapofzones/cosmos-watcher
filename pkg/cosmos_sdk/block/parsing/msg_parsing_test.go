package cosmos

import (
	"testing"

	connectiontypes "github.com/cosmos/ibc-go/v5/modules/core/03-connection/types"
	"github.com/stretchr/testify/assert"
	types6 "github.com/tendermint/tendermint/abci/types"
)

func TestParseIDsFromResults(t *testing.T) {
	type args struct {
		txResult       *types6.ResponseDeliverTx
		expectedEvents []string
		attributeKeys  []string
	}
	tests := []struct {
		name     string
		args     args
		expected []string
	}{
		{"empty_args", args{}, nil},
		{
			"nil_arg_tx_result",
			args{
				nil,
				[]string{connectiontypes.EventTypeConnectionOpenInit},
				[]string{connectiontypes.AttributeKeyConnectionID},
			},
			nil,
		},
		{
			"empty_arg_tx_result",
			args{
				&types6.ResponseDeliverTx{},
				[]string{connectiontypes.EventTypeConnectionOpenInit},
				[]string{connectiontypes.AttributeKeyConnectionID},
			},
			nil,
		},
		{
			"single_result_id",
			args{
				&types6.ResponseDeliverTx{
					Events: []types6.Event{{
						Type: connectiontypes.EventTypeConnectionOpenInit,
						Attributes: []types6.EventAttribute{{
							Key:   []byte(connectiontypes.AttributeKeyConnectionID),
							Value: []byte("myConnectionID"),
							Index: true,
						}},
					}},
				},
				[]string{connectiontypes.EventTypeConnectionOpenInit},
				[]string{connectiontypes.AttributeKeyConnectionID},
			},
			[]string{"myConnectionID"},
		},
		{
			"multiple_result_id",
			args{
				&types6.ResponseDeliverTx{
					Events: []types6.Event{
						{
							Type: connectiontypes.EventTypeConnectionOpenInit,
							Attributes: []types6.EventAttribute{
								{
									Key:   []byte(connectiontypes.AttributeKeyConnectionID),
									Value: []byte("myConnectionID"),
									Index: true,
								},
								{
									Key:   []byte(connectiontypes.AttributeKeyClientID),
									Value: []byte("myClientID"),
									Index: true,
								},
							},
						},
						{
							Type: connectiontypes.EventTypeConnectionOpenTry,
							Attributes: []types6.EventAttribute{
								{
									Key:   []byte(connectiontypes.AttributeKeyCounterpartyClientID),
									Value: []byte("myCounterpartyClientID"),
									Index: true,
								},
								{
									Key:   []byte(connectiontypes.AttributeKeyCounterpartyConnectionID),
									Value: []byte("myCounterpartyConnectionID"),
									Index: true,
								},
							},
						},
					},
				}, []string{connectiontypes.EventTypeConnectionOpenInit, connectiontypes.EventTypeConnectionOpenTry}, []string{connectiontypes.AttributeKeyConnectionID, connectiontypes.AttributeKeyCounterpartyConnectionID},
			},
			[]string{"myConnectionID", "myCounterpartyConnectionID"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ParseIDsFromResults(tt.args.txResult, tt.args.expectedEvents, tt.args.attributeKeys, attributeFiler{})
			assert.Equal(t, tt.expected, actual)
		})
	}
}

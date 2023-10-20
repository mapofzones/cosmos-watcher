package cosmos

import (
	"testing"

	types6 "github.com/cometbft/cometbft/abci/types"
	connectiontypes "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	"github.com/stretchr/testify/assert"
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
							Key:   connectiontypes.AttributeKeyConnectionID,
							Value: "myConnectionID",
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
									Key:   connectiontypes.AttributeKeyConnectionID,
									Value: "myConnectionID",
									Index: true,
								},
								{
									Key:   connectiontypes.AttributeKeyClientID,
									Value: "myClientID",
									Index: true,
								},
							},
						},
						{
							Type: connectiontypes.EventTypeConnectionOpenTry,
							Attributes: []types6.EventAttribute{
								{
									Key:   connectiontypes.AttributeKeyCounterpartyClientID,
									Value: "myCounterpartyClientID",
									Index: true,
								},
								{
									Key:   connectiontypes.AttributeKeyCounterpartyConnectionID,
									Value: "myCounterpartyConnectionID",
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
			actual := ParseIDsFromResults(tt.args.txResult, tt.args.expectedEvents, tt.args.attributeKeys,
				attributeFiler{}, attributeFiler{}, attributeFiler{}, attributeFiler{})
			assert.Equal(t, tt.expected, actual)
		})
	}
}

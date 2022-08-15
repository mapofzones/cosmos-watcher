package cosmos

import (
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	"github.com/stretchr/testify/assert"
	types6 "github.com/tendermint/tendermint/abci/types"
	"testing"
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
				&types6.ResponseDeliverTx{Events: []types6.Event{{
					connectiontypes.EventTypeConnectionOpenInit,
					[]types6.EventAttribute{{[]byte(connectiontypes.AttributeKeyConnectionID), []byte("myConnectionID"), true}},
				}}},
				[]string{connectiontypes.EventTypeConnectionOpenInit},
				[]string{connectiontypes.AttributeKeyConnectionID},
			},
			[]string{"myConnectionID"},
		},
		{
			"multiple_result_id",
			args{
				&types6.ResponseDeliverTx{Events: []types6.Event{
					{
						connectiontypes.EventTypeConnectionOpenInit,
						[]types6.EventAttribute{
							{[]byte(connectiontypes.AttributeKeyConnectionID), []byte("myConnectionID"), true},
							{[]byte(connectiontypes.AttributeKeyClientID), []byte("myClientID"), true},
						},
					},
					{
						connectiontypes.EventTypeConnectionOpenTry,
						[]types6.EventAttribute{
							{[]byte(connectiontypes.AttributeKeyCounterpartyClientID), []byte("myCounterpartyClientID"), true},
							{[]byte(connectiontypes.AttributeKeyCounterpartyConnectionID), []byte("myCounterpartyConnectionID"), true},
						},
					},
				}},
				[]string{connectiontypes.EventTypeConnectionOpenInit, connectiontypes.EventTypeConnectionOpenTry},
				[]string{connectiontypes.AttributeKeyConnectionID, connectiontypes.AttributeKeyCounterpartyConnectionID},
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

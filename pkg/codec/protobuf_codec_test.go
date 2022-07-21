package watcher

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_functionxRegisterInterfaces(t *testing.T) {
	assert.NotPanics(t, func() {
		interfaceRegistry := codectypes.NewInterfaceRegistry()
		functionxRegisterInterfaces(interfaceRegistry)
	})
}

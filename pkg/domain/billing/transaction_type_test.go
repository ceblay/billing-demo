package billing_test

import (
	"testing"

	bl "github.com/ceblay/billing-demo/pkg/domain/billing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTransactionTypeFromString_valid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name     string
		TypeName string
	}{
		{
			Name:     "Deposit Type",
			TypeName: "DEPOSIT",
		},
		{
			Name:     "Debitit Type",
			TypeName: "DEBIT",
		},
		{
			Name:     "Credit Type",
			TypeName: "CREDIT",
		},
		{
			Name:     "Transfer Type",
			TypeName: "TRANSFER",
		},
		{
			Name:     "Subscription TYpe",
			TypeName: "SUBSCRIPTION",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			kind, err := bl.NewTransactionTypeFromString(testCase.TypeName)

			require.NoError(t, err)
			assert.Equal(t, testCase.TypeName, kind.String())
		})
	}
}

func TestNewTransactionTypeFromString_invalid(t *testing.T) {
	kind := "NONEXISTENT TYPE"
	_, err := bl.NewTransactionTypeFromString(kind)

	assert.Error(t, err)
}

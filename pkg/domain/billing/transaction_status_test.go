package billing_test

import (
	"testing"

	bl "github.com/ceblay/billing-demo/pkg/domain/billing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTransactionStatusFromString_valid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name       string
		StatusName string
	}{
		{
			Name:       "Pending Status",
			StatusName: "PENDING",
		},
		{
			Name:       "Completed Status",
			StatusName: "COMPLETED",
		},
		{
			Name:       "Failed Status",
			StatusName: "FAILED",
		},
		{
			Name:       "Canceled Status",
			StatusName: "CANCELED",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			status, err := bl.NewTransactionStatusFromString(testCase.StatusName)

			require.NoError(t, err)
			assert.Equal(t, testCase.StatusName, status.String())
		})
	}
}

func TestNewTransactionStatusFromString_invalid(t *testing.T) {
	status := "NONEXISTENT STATUS"
	_, err := bl.NewTransactionStatusFromString(status)

	assert.Error(t, err)
}

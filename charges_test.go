package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// mockCharges creates a Charges
func mockCharges() *Charges {
	c := NewCharges()
	c.ChargeDetails = "B"
	c.SendersChargesOne = "USD0,99"
	c.SendersChargesTwo = "USD2,99"
	c.SendersChargesThree = "USD3,99"
	c.SendersChargesFour = "USD1,00"
	return c
}

// TestMockCharges validates mockCharges
func TestMockCharges(t *testing.T) {
	c := mockCharges()

	require.NoError(t, c.Validate(), "mockCharges does not validate and will break other tests")
}

// TestChargeDetailsValid validates ChargeDetails is valid
func TestPaymentNotificationIndicatorValid(t *testing.T) {
	c := mockCharges()
	c.ChargeDetails = "F"

	err := c.Validate()

	require.EqualError(t, err, fieldError("ChargeDetails", ErrChargeDetails, c.ChargeDetails).Error())
}

func TestChargesEmpty(t *testing.T) {
	c := &Charges{}
	c.Parse("{3700}B****") // invalid, caused a fuzz crash
	require.Equal(t, "B", c.ChargeDetails)
	require.Empty(t, c.SendersChargesOne)
	require.Empty(t, c.SendersChargesTwo)
	require.Empty(t, c.SendersChargesThree)
	require.Empty(t, c.SendersChargesFour)
}

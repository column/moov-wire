package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIPaymentMethodToBeneficiary creates a FIPaymentMethodToBeneficiary
func mockFIPaymentMethodToBeneficiary() *FIPaymentMethodToBeneficiary {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = "CHECK"
	pm.AdditionalInformation = "Additional Information"
	return pm
}

// TestMockFIPaymentMethodToBeneficiary validates mockFIPaymentMethodToBeneficiary
func TestMockFIPaymentMethodToBeneficiary(t *testing.T) {
	pm := mockFIPaymentMethodToBeneficiary()

	require.NoError(t, pm.Validate(), "mockFIPaymentMethodToBeneficiary does not validate and will break other tests")
}

// TestPaymentMethodValid validates FIPaymentMethodToBeneficiary PaymentMethod
func TestPaymentMethodValid(t *testing.T) {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = ""

	err := pm.Validate()

	require.EqualError(t, err, fieldError("PaymentMethod", ErrFieldInclusion, pm.PaymentMethod).Error())
}

// TestAdditionalInformationAlphaNumeric validates FIPaymentMethodToBeneficiary AdditionalInformation is alphanumeric
func TestAdditionalInformationAlphaNumeric(t *testing.T) {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.AdditionalInformation = "®"

	err := pm.Validate()

	require.EqualError(t, err, fieldError("AdditionalInformation", ErrNonAlphanumeric, pm.AdditionalInformation).Error())
}

// TestParseFIPaymentMethodToBeneficiaryReaderParseError parses a wrong FIPaymentMethodToBeneficiary reader parse error
func TestParseFIPaymentMethodToBeneficiaryReaderParseError(t *testing.T) {
	var line = "{6420}CHECK®dditional Information        *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIPaymentMethodToBeneficiary()

	expected := r.parseError(fieldError("AdditionalInformation", ErrNonAlphanumeric, "®dditional Information")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("AdditionalInformation", ErrNonAlphanumeric, "®dditional Information")).Error()
	require.EqualError(t, err, expected)
}

// TestFIPaymentMethodToBeneficiaryTagError validates a FIPaymentMethodToBeneficiary tag
func TestFIPaymentMethodToBeneficiaryTagError(t *testing.T) {
	pm := mockFIPaymentMethodToBeneficiary()
	pm.tag = "{9999}"

	err := pm.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, pm.tag).Error())
}

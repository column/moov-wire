package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBeneficiaryCustomer creates a BeneficiaryCustomer
func mockBeneficiaryCustomer() *BeneficiaryCustomer {
	bc := NewBeneficiaryCustomer()
	bc.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	bc.CoverPayment.SwiftLineOne = "Swift Line One"
	bc.CoverPayment.SwiftLineTwo = "Swift Line Two"
	bc.CoverPayment.SwiftLineThree = "Swift Line Three"
	bc.CoverPayment.SwiftLineFour = "Swift Line Four"
	bc.CoverPayment.SwiftLineFive = "Swift Line Five"
	return bc
}

// TestMockBeneficiaryCustomer validates mockBeneficiaryCustomer
func TestMockBeneficiaryCustomer(t *testing.T) {
	bc := mockBeneficiaryCustomer()

	require.NoError(t, bc.Validate(), "mockBeneficiaryCustomer does not validate and will break other tests")
}

// TestBeneficiaryCustomerSwiftFieldTagAlphaNumeric validates BeneficiaryCustomer SwiftFieldTag is alphanumeric
func TestBeneficiaryCustomerSwiftFieldTagAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftFieldTag = "®"

	err := bc.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, bc.CoverPayment.SwiftFieldTag).Error())
}

// TestBeneficiaryCustomerSwiftLineOneAlphaNumeric validates BeneficiaryCustomer SwiftLineOne is alphanumeric
func TestBeneficiaryCustomerSwiftLineOneAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineOne = "®"

	err := bc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineOne", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineOne).Error())
}

// TestBeneficiaryCustomerSwiftLineTwoAlphaNumeric validates BeneficiaryCustomer SwiftLineTwo is alphanumeric
func TestBeneficiaryCustomerSwiftLineTwoAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineTwo = "®"

	err := bc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineTwo", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineTwo).Error())
}

// TestBeneficiaryCustomerSwiftLineThreeAlphaNumeric validates BeneficiaryCustomer SwiftLineThree is alphanumeric
func TestBeneficiaryCustomerSwiftLineThreeAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineThree = "®"

	err := bc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineThree", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineThree).Error())
}

// TestBeneficiaryCustomerSwiftLineFourAlphaNumeric validates BeneficiaryCustomer SwiftLineFour is alphanumeric
func TestBeneficiaryCustomerSwiftLineFourAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineFour = "®"

	err := bc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFour", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineFour).Error())
}

// TestBeneficiaryCustomerSwiftLineFiveAlphaNumeric validates BeneficiaryCustomer SwiftLineFive is alphanumeric
func TestBeneficiaryCustomerSwiftLineFiveAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineFive = "®"

	err := bc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFive", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineFive).Error())
}

// TestBeneficiaryCustomerSwiftLineSixAlphaNumeric validates BeneficiaryCustomer SwiftLineSix is alphanumeric
func TestBeneficiaryCustomerSwiftLineSixAlphaNumeric(t *testing.T) {
	sr := mockBeneficiaryCustomer()
	sr.CoverPayment.SwiftLineSix = "Test"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SwiftLineSix", ErrInvalidProperty, sr.CoverPayment.SwiftLineSix).Error())
}

// TestParseBeneficiaryCustomerReaderParseError parses a wrong BeneficiaryCustomer reader parse error
func TestParseBeneficiaryCustomerReaderParseError(t *testing.T) {
	var line = "{7059}Swift*Swift ®ine One"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryCustomer()

	expected := r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error()
	require.EqualError(t, err, expected)
}

// TestBeneficiaryCustomerTagError validates a BeneficiaryCustomer tag
func TestBeneficiaryCustomerTagError(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.tag = "{9999}"

	err := bc.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, bc.tag).Error())
}

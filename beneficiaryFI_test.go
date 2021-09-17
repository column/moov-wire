package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBeneficiaryFI creates a BeneficiaryFI
func mockBeneficiaryFI() *BeneficiaryFI {
	bfi := NewBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	bfi.FinancialInstitution.Identifier = "123456789"
	bfi.FinancialInstitution.Name = "FI Name"
	bfi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bfi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bfi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return bfi
}

// TestMockBeneficiaryFI validates mockBeneficiaryFI
func TestMockBeneficiaryFI(t *testing.T) {
	bfi := mockBeneficiaryFI()

	require.NoError(t, bfi.Validate(), "mockBeneficiaryFI does not validate and will break other tests")
}

// TestBeneficiaryFIIdentificationCodeValid validates BeneficiaryFI IdentificationCode
func TestBeneficiaryFIIdentificationCodeValid(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = "Football Card ID"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, bfi.FinancialInstitution.IdentificationCode).Error())
}

// TestBeneficiaryFIIdentificationCodeFI validates BeneficiaryFI IdentificationCode is an FI code
func TestBeneficiaryFIIdentificationCodeFI(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = "1"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, bfi.FinancialInstitution.IdentificationCode).Error())
}

// TestBeneficiaryFIIdentifierAlphaNumeric validates BeneficiaryFI Identifier is alphanumeric
func TestBeneficiaryFIIdentifierAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Identifier = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrNonAlphanumeric, bfi.FinancialInstitution.Identifier).Error())
}

// TestBeneficiaryFINameAlphaNumeric validates BeneficiaryFI Name is alphanumeric
func TestBeneficiaryFINameAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Name = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, bfi.FinancialInstitution.Name).Error())
}

// TestBeneficiaryFIAddressLineOneAlphaNumeric validates BeneficiaryFI AddressLineOne is alphanumeric
func TestBeneficiaryFIAddressLineOneAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Address.AddressLineOne = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, bfi.FinancialInstitution.Address.AddressLineOne).Error())
}

// TestBeneficiaryFIAddressLineTwoAlphaNumeric validates BeneficiaryFI AddressLineTwo is alphanumeric
func TestBeneficiaryFIAddressLineTwoAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Address.AddressLineTwo = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, bfi.FinancialInstitution.Address.AddressLineTwo).Error())
}

// TestBeneficiaryFIAddressLineThreeAlphaNumeric validates BeneficiaryFI AddressLineThree is alphanumeric
func TestBeneficiaryFIAddressLineThreeAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Address.AddressLineThree = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, bfi.FinancialInstitution.Address.AddressLineThree).Error())
}

// TestBeneficiaryFIIdentificationCodeRequired validates BeneficiaryFI IdentificationCode is required
func TestBeneficiaryFIIdentificationCodeRequired(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = ""

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrFieldRequired).Error())
}

// TestBeneficiaryFIIdentifierRequired validates BeneficiaryFI Identifier is required
func TestBeneficiaryFIIdentifierRequired(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Identifier = ""

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrFieldRequired).Error())
}

// TestParseBeneficiaryFIReaderParseError parses a wrong BeneficiaryFI reader parse error
func TestParseBeneficiaryFIReaderParseError(t *testing.T) {
	var line = "{4100}D123456789                         *F® Name                            *Address One                        *Address Two                        *Address Three                      *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryFI()

	expected := r.parseError(fieldError("Name", ErrNonAlphanumeric, "F® Name")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Name", ErrNonAlphanumeric, "F® Name")).Error()
	require.EqualError(t, err, expected)
}

// TestBeneficiaryFITagError validates a BeneficiaryFI tag
func TestBeneficiaryFITagError(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.tag = "{9999}"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, bfi.tag).Error())
}

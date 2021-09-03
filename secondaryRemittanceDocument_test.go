package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// SecondaryRemittanceDocument creates a SecondaryRemittanceDocument
func mockSecondaryRemittanceDocument() *SecondaryRemittanceDocument {
	srd := NewSecondaryRemittanceDocument()
	srd.DocumentTypeCode = StatementAccount
	srd.ProprietaryDocumentTypeCode = ""
	srd.DocumentIdentificationNumber = "222222"
	srd.Issuer = "Issuer 2"
	return srd
}

// TestMockSecondaryRemittanceDocument validates mockSecondaryRemittanceDocument
func TestMockSecondaryRemittanceDocument(t *testing.T) {
	srd := mockSecondaryRemittanceDocument()

	require.NoError(t, srd.Validate(), "mockSecondaryRemittanceDocument does not validate and will break other tests")
}

// TestSRDDocumentTypeCodeValid validates SecondaryRemittanceDocument DocumentTypeCode
func TestSRDDocumentTypeCodeValid(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentTypeCode = "ZZZZ"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("DocumentTypeCode", ErrDocumentTypeCode, prd.DocumentTypeCode).Error())
}

// TestSRDProprietaryDocumentTypeCodeAlphaNumeric validates SecondaryRemittanceDocument ProprietaryDocumentTypeCode is alphanumeric
func TestSRDProprietaryDocumentTypeCodeAlphaNumeric(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentTypeCode = ProprietaryDocumentType
	prd.ProprietaryDocumentTypeCode = "®"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("ProprietaryDocumentTypeCode", ErrNonAlphanumeric, prd.ProprietaryDocumentTypeCode).Error())
}

// TestSRDDocumentIdentificationNumberAlphaNumeric validates SecondaryRemittanceDocument DocumentIdentificationNumber is alphanumeric
func TestSRDDocumentIdentificationNumberAlphaNumeric(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentIdentificationNumber = "®"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("DocumentIdentificationNumber", ErrNonAlphanumeric, prd.DocumentIdentificationNumber).Error())
}

// TestSRDIssuerAlphaNumeric validates SecondaryRemittanceDocument Issuer is alphanumeric
func TestSRDIssuerAlphaNumeric(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.Issuer = "®"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("Issuer", ErrNonAlphanumeric, prd.Issuer).Error())
}

// TestSRDProprietaryDocumentTypeCodeRequired validates SecondaryRemittanceDocument ProprietaryDocumentTypeCode is required
func TestSRDProprietaryDocumentTypeCodeRequired(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentTypeCode = ProprietaryDocumentType
	prd.ProprietaryDocumentTypeCode = ""

	err := prd.Validate()

	require.EqualError(t, err, fieldError("ProprietaryDocumentTypeCode", ErrFieldRequired).Error())
}

// TestSRDDocumentIdentificationNumberRequired validates SecondaryRemittanceDocument DocumentIdentificationNumber is required
func TestSRDDocumentIdentificationNumberRequired(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentIdentificationNumber = ""

	err := prd.Validate()

	require.EqualError(t, err, fieldError("DocumentIdentificationNumber", ErrFieldRequired).Error())
}

// TestSRDProprietaryDocumentTypeCodeInvalid validates SecondaryRemittanceDocument ProprietaryDocumentTypeCode is invalid
func TestSRDProprietaryDocumentTypeCodeInvalid(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentTypeCode = AccountsReceivableOpenItem
	prd.ProprietaryDocumentTypeCode = "Proprietary"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("ProprietaryDocumentTypeCode", ErrInvalidProperty, prd.ProprietaryDocumentTypeCode).Error())
}

// TestParseSecondaryRemittanceDocumentWrongLength parses a wrong SecondaryRemittanceDocument record length
func TestParseSecondaryRemittanceDocumentWrongLength(t *testing.T) {
	var line = "{8700}SOA"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSecondaryRemittanceDocument()

	require.EqualError(t, err, "line:0 record:SecondaryRemittanceDocument wire.TagWrongLengthErr must be [10, 118] characters and found 9")
}

// TestParseSecondaryRemittanceDocumentReaderParseError parses a wrong SecondaryRemittanceDocument reader parse error
func TestParseSecondaryRemittanceDocumentReaderParseError(t *testing.T) {
	var line = "{8700}ZZZZ*222222*Issuer 2*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSecondaryRemittanceDocument()

	require.EqualError(t, err, r.parseError(fieldError("DocumentTypeCode", ErrDocumentTypeCode, "ZZZZ")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("DocumentTypeCode", ErrDocumentTypeCode, "ZZZZ")).Error())
}

// TestSecondaryRemittanceDocumentTagError validates a SecondaryRemittanceDocument tag
func TestSecondaryRemittanceDocumentTagError(t *testing.T) {
	srd := mockSecondaryRemittanceDocument()
	srd.tag = "{9999}"

	require.EqualError(t, srd.Validate(), fieldError("tag", ErrValidTagForType, srd.tag).Error())
}

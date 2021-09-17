package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBeneficiaryReference creates a BeneficiaryReference
func mockBeneficiaryReference() *BeneficiaryReference {
	br := NewBeneficiaryReference()
	br.BeneficiaryReference = "Reference"
	return br
}

// TestMockBeneficiary validates mockBeneficiaryReference
func TestMockBeneficiaryReference(t *testing.T) {
	br := mockBeneficiaryReference()

	require.NoError(t, br.Validate(), "mockBeneficiaryReference does not validate and will break other tests")
}

// TestBeneficiaryReferenceAlphaNumeric validates BeneficiaryReference is alphanumeric
func TestBeneficiaryReferenceAlphaNumeric(t *testing.T) {
	br := mockBeneficiaryReference()
	br.BeneficiaryReference = "®"

	err := br.Validate()

	require.EqualError(t, err, fieldError("BeneficiaryReference", ErrNonAlphanumeric, br.BeneficiaryReference).Error())
}

// TestParseBeneficiaryReferenceReaderParseError parses a wrong BeneficiaryReference reader parse error
func TestParseBeneficiaryReferenceReaderParseError(t *testing.T) {
	var line = "{4320}Reference®      *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryReference()

	expected := r.parseError(fieldError("BeneficiaryReference", ErrNonAlphanumeric, "Reference®")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("BeneficiaryReference", ErrNonAlphanumeric, "Reference®")).Error()
	require.EqualError(t, err, expected)
}

// TestBeneficiaryReferenceTagError validates a BeneficiaryReference tag
func TestBeneficiaryReferenceTagError(t *testing.T) {
	br := mockBeneficiaryReference()
	br.tag = "{9999}"

	err := br.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, br.tag).Error())
}

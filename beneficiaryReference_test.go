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

// TestParseBeneficiaryReferenceWrongLength parses a wrong BeneficiaryReference record length
func TestParseBeneficiaryReferenceWrongLength(t *testing.T) {
	var line = "{4320}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryReference()

	require.EqualError(t, err, "line:0 record:BeneficiaryReference wire.TagWrongLengthErr must be [8, 23] characters and found 6")
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

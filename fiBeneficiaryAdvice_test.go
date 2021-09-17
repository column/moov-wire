package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIBeneficiaryAdvice creates a FIBeneficiaryAdvice
func mockFIBeneficiaryAdvice() *FIBeneficiaryAdvice {
	fiba := NewFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = AdviceCodeLetter
	fiba.Advice.LineOne = "Line One"
	fiba.Advice.LineTwo = "Line Two"
	fiba.Advice.LineThree = "Line Three"
	fiba.Advice.LineFour = "Line Four"
	fiba.Advice.LineFive = "Line Five"
	fiba.Advice.LineSix = "Line Six"
	return fiba
}

// TestMockFIBeneficiaryAdvice validates mockFIBeneficiaryAdvice
func TestMockFIBeneficiaryAdvice(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()

	require.NoError(t, fiba.Validate(), "mockFIBeneficiaryAdvice does not validate and will break other tests")
}

// TestFIBeneficiaryAdviceCodeValid validates FIBeneficiaryAdvice AdviceCode is alphanumeric
func TestFIBeneficiaryAdviceCodeValid(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = "Z"

	err := fiba.Validate()
	require.EqualError(t, err, fieldError("AdviceCode", ErrAdviceCode, fiba.Advice.AdviceCode).Error())
}

// TestFIBeneficiaryAdviceLineOneAlphaNumeric validates FIBeneficiaryAdvice LineOne is alphanumeric
func TestFIBeneficiaryAdviceLineOneAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineOne = "®"

	err := fiba.Validate()
	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, fiba.Advice.LineOne).Error())
}

// TestFIBeneficiaryAdviceLineTwoAlphaNumeric validates FIBeneficiaryAdvice LineTwo is alphanumeric
func TestFIBeneficiaryAdviceLineTwoAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineTwo = "®"

	err := fiba.Validate()
	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, fiba.Advice.LineTwo).Error())
}

// TestFIBeneficiaryAdviceLineThreeAlphaNumeric validates FIBeneficiaryAdvice LineThree is alphanumeric
func TestFIBeneficiaryAdviceLineThreeAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineThree = "®"

	err := fiba.Validate()
	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, fiba.Advice.LineThree).Error())
}

// TestFIBeneficiaryAdviceLineFourAlphaNumeric validates FIBeneficiaryAdvice LineFour is alphanumeric
func TestFIBeneficiaryAdviceLineFourAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineFour = "®"

	err := fiba.Validate()
	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, fiba.Advice.LineFour).Error())
}

// TestFIBeneficiaryAdviceLineFiveAlphaNumeric validates FIBeneficiaryAdvice LineFive is alphanumeric
func TestFIBeneficiaryAdviceLineFiveAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineFive = "®"

	err := fiba.Validate()
	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, fiba.Advice.LineFive).Error())
}

// TestFIBeneficiaryAdviceLineSixAlphaNumeric validates FIBeneficiaryAdvice LineSix is alphanumeric
func TestFIBeneficiaryAdviceLineSixAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineSix = "®"

	err := fiba.Validate()
	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, fiba.Advice.LineSix).Error())
}

// TestParseFIBeneficiaryAdviceReaderParseError parses a wrong FIBeneficiaryAdvice reader parse error
func TestParseFIBeneficiaryAdviceReaderParseError(t *testing.T) {
	var line = "{6410}LTRLine ®ne*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiaryAdvice()

	expected := r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line ®ne")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line ®ne")).Error()
	require.EqualError(t, err, expected)
}

// TestFIBeneficiaryAdviceTagError validates a FIBeneficiaryAdvice tag
func TestFIBeneficiaryAdviceTagError(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.tag = "{9999}"
	err := fiba.Validate()
	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, fiba.tag).Error())
}

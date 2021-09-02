package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockLocalInstrument creates a LocalInstrument
func mockLocalInstrument() *LocalInstrument {
	li := NewLocalInstrument()
	li.LocalInstrumentCode = ANSIX12format
	li.ProprietaryCode = ""
	return li
}

// TestMockLocalInstrument validates mockLocalInstrument
func TestMockLocalInstrument(t *testing.T) {
	li := mockLocalInstrument()

	require.NoError(t, li.Validate(), "mockLocalInstrument does not validate and will break other tests")
}

// TestLocalInstrumentCodeValid validates LocalInstrumentCode
func TestLocalInstrumentCodeValid(t *testing.T) {
	li := mockLocalInstrument()
	li.LocalInstrumentCode = "Chestnut"

	err := li.Validate()

	require.EqualError(t, err, fieldError("LocalInstrumentCode", ErrLocalInstrumentCode, li.LocalInstrumentCode).Error())
}

// TestProprietaryCodeValid validates ProprietaryCode
func TestProprietaryCodeValid(t *testing.T) {
	li := mockLocalInstrument()
	li.ProprietaryCode = "Proprietary"

	err := li.Validate()

	require.EqualError(t, err, fieldError("ProprietaryCode", ErrInvalidProperty, li.ProprietaryCode).Error())
}

// TestProprietaryCodeAlphaNumeric validates ProprietaryCode is alphanumeric
func TestProprietaryCodeAlphaNumeric(t *testing.T) {
	li := mockLocalInstrument()
	li.LocalInstrumentCode = ProprietaryLocalInstrumentCode
	li.ProprietaryCode = "®"

	err := li.Validate()

	require.EqualError(t, err, fieldError("ProprietaryCode", ErrNonAlphanumeric, li.ProprietaryCode).Error())
}

// TestParseLocalInstrumentWrongLength parses a wrong LocalInstrumente record length
func TestParseLocalInstrumentWrongLength(t *testing.T) {
	var line = "{3610}ANS"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseLocalInstrument()

	require.EqualError(t, err, "line:0 record:LocalInstrument wire.TagWrongLengthErr must be [10, 46] characters and found 9")
}

// TestParseLocalInstrumentReaderParseError parses a wrong LocalInstrumente reader parse error
func TestParseLocalInstrumentReaderParseError(t *testing.T) {
	var line = "{3610}ABCD                                   "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseLocalInstrument()

	require.EqualError(t, err, r.parseError(fieldError("LocalInstrumentCode", ErrLocalInstrumentCode, "ABCD")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("LocalInstrumentCode", ErrLocalInstrumentCode, "ABCD")).Error())
}

// TestLocalInstrumentTagError validates a LocalInstrument tag
func TestLocalInstrumentTagError(t *testing.T) {
	li := mockLocalInstrument()
	li.tag = "{9999}"

	require.EqualError(t, li.Validate(), fieldError("tag", ErrValidTagForType, li.tag).Error())
}

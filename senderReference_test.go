package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockSenderReference creates a SenderReference
func mockSenderReference() *SenderReference {
	sr := NewSenderReference()
	sr.SenderReference = "Sender Reference"
	return sr
}

// TestMockSenderReference validates mockSenderReference
func TestMockSenderReference(t *testing.T) {
	sr := mockSenderReference()

	require.NoError(t, sr.Validate(), "mockSenderReference does not validate and will break other tests")
}

// TestSenderReferenceAlphaNumeric validates SenderReference is alphanumeric
func TestSenderReferenceAlphaNumeric(t *testing.T) {
	sr := mockSenderReference()
	sr.SenderReference = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SenderReference", ErrNonAlphanumeric, sr.SenderReference).Error())
}

// TestParseSenderReferenceReaderParseError parses a wrong SenderReference reader parse error
func TestParseSenderReferenceReaderParseError(t *testing.T) {
	var line = "{3320}Sender®Reference*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderReference()

	require.EqualError(t, err, r.parseError(fieldError("SenderReference", ErrNonAlphanumeric, "Sender®Reference")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SenderReference", ErrNonAlphanumeric, "Sender®Reference")).Error())
}

// TestSenderReferenceTagError validates a SenderReference tag
func TestSenderReferenceTagError(t *testing.T) {
	sr := mockSenderReference()
	sr.tag = "{9999}"

	require.EqualError(t, sr.Validate(), fieldError("tag", ErrValidTagForType, sr.tag).Error())
}

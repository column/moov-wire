// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

// LocalInstrument is the LocalInstrument of the wire
type LocalInstrument struct {
	// tag
	tag string
	// LocalInstrumentCode is local instrument code
	LocalInstrumentCode string `json:"LocalInstrument,omitempty"`
	// ProprietaryCode is proprietary code
	ProprietaryCode string `json:"proprietaryCode,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewLocalInstrument returns a new LocalInstrument
func NewLocalInstrument() *LocalInstrument {
	li := &LocalInstrument{
		tag: TagLocalInstrument,
	}
	return li
}

// Parse takes the input string and parses the LocalInstrument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (li *LocalInstrument) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 10 || dataLen > 46 {
		return TagWrongLengthErr{
			Message:   fmt.Sprintf("must be [10, 46] characters and found %d", dataLen),
			TagLength: 46,
			Length:    dataLen,
		}
	}
	li.tag = record[:6]
	li.LocalInstrumentCode = li.parseStringField(record[6:10])
	if delim := strings.IndexByte(record, '*'); delim > 0 {
		li.ProprietaryCode = li.parseStringField(record[10:delim])
	}
	return nil
}

func (li *LocalInstrument) UnmarshalJSON(data []byte) error {
	type Alias LocalInstrument
	aux := struct {
		*Alias
	}{
		(*Alias)(li),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	li.tag = TagLocalInstrument
	return nil
}

// String writes LocalInstrument
func (li *LocalInstrument) String() string {
	var buf strings.Builder
	buf.Grow(45)
	buf.WriteString(li.tag)
	buf.WriteString(li.LocalInstrumentCodeField())
	if li.ProprietaryCode != "" {
		buf.WriteString(strings.TrimSpace(li.ProprietaryCodeField()) + "*")
	}
	return buf.String()
}

// Validate performs WIRE format rule checks on LocalInstrument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (li *LocalInstrument) Validate() error {
	if err := li.fieldInclusion(); err != nil {
		return err
	}
	if li.tag != TagLocalInstrument {
		return fieldError("tag", ErrValidTagForType, li.tag)
	}
	if err := li.isLocalInstrumentCode(li.LocalInstrumentCode); err != nil {
		return fieldError("LocalInstrumentCode", err, li.LocalInstrumentCode)
	}
	if err := li.isAlphanumeric(li.ProprietaryCode); err != nil {
		return fieldError("ProprietaryCode", err, li.ProprietaryCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
// ProprietaryCode is only allowed if LocalInstrument Code is PROP
func (li *LocalInstrument) fieldInclusion() error {
	if li.LocalInstrumentCode != ProprietaryLocalInstrumentCode && li.ProprietaryCode != "" {
		return fieldError("ProprietaryCode", ErrInvalidProperty, li.ProprietaryCode)
	}
	return nil
}

// LocalInstrumentCodeField gets a string of LocalInstrumentCode field
func (li *LocalInstrument) LocalInstrumentCodeField() string {
	return li.alphaField(li.LocalInstrumentCode, 4)
}

// ProprietaryCodeField gets a string of ProprietaryCode field
func (li *LocalInstrument) ProprietaryCodeField() string {
	return li.alphaField(li.ProprietaryCode, 35)
}

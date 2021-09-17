// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// BeneficiaryReference is a reference for the beneficiary
type BeneficiaryReference struct {
	// tag
	tag string
	// BeneficiaryReference
	BeneficiaryReference string `json:"beneficiaryReference,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryReference returns a new BeneficiaryReference
func NewBeneficiaryReference() *BeneficiaryReference {
	br := &BeneficiaryReference{
		tag: TagBeneficiaryReference,
	}
	return br
}

// Parse takes the input string and parses the BeneficiaryReference values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (br *BeneficiaryReference) Parse(record string) error {
	br.tag = record[:6]

	if delim := strings.IndexByte(record, '*'); delim > 0 {
		br.BeneficiaryReference = br.parseStringField(record[6:delim])
	} else {
		br.BeneficiaryReference = br.parseStringField(record[6:])
	}
	return nil
}

func (br *BeneficiaryReference) UnmarshalJSON(data []byte) error {
	type Alias BeneficiaryReference
	aux := struct {
		*Alias
	}{
		(*Alias)(br),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	br.tag = TagBeneficiaryReference
	return nil
}

// String writes BeneficiaryReference
func (br *BeneficiaryReference) String() string {
	var buf strings.Builder
	buf.Grow(22)
	buf.WriteString(br.tag)
	buf.WriteString(strings.TrimSpace(br.BeneficiaryReferenceField()) + "*")
	return br.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on BeneficiaryReference and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (br *BeneficiaryReference) Validate() error {
	if br.tag != TagBeneficiaryReference {
		return fieldError("tag", ErrValidTagForType, br.tag)
	}
	if err := br.isAlphanumeric(br.BeneficiaryReference); err != nil {
		return fieldError("BeneficiaryReference", err, br.BeneficiaryReference)
	}
	return nil
}

// BeneficiaryReferenceField gets a string of the BeneficiaryReference field
func (br *BeneficiaryReference) BeneficiaryReferenceField() string {
	return br.alphaField(br.BeneficiaryReference, 16)
}

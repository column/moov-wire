// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// Beneficiary is the beneficiary of the wire
type Beneficiary struct {
	// tag
	tag string
	// Personal
	Personal Personal `json:"personal,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiary returns a new Beneficiary
func NewBeneficiary() *Beneficiary {
	ben := &Beneficiary{
		tag: TagBeneficiary,
	}
	return ben
}

// Parse takes the input string and parses the Beneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ben *Beneficiary) Parse(record string) error {
	ben.tag = record[:6]
	ben.Personal.IdentificationCode = ben.parseStringField(record[6:7])

	optionalFields := strings.Split(record[7:], "*")
	if len(optionalFields) >= 1 {
		ben.Personal.Identifier = ben.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		ben.Personal.Name = ben.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		ben.Personal.Address.AddressLineOne = ben.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		ben.Personal.Address.AddressLineTwo = ben.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		ben.Personal.Address.AddressLineThree = ben.parseStringField(optionalFields[4])
	}
	return nil
}

func (ben *Beneficiary) UnmarshalJSON(data []byte) error {
	type Alias Beneficiary
	aux := struct {
		*Alias
	}{
		(*Alias)(ben),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ben.tag = TagBeneficiary
	return nil
}

// String writes Beneficiary
func (ben *Beneficiary) String() string {
	var buf strings.Builder
	buf.Grow(181)
	buf.WriteString(ben.tag)
	buf.WriteString(ben.IdentificationCodeField())
	buf.WriteString(strings.TrimSpace(ben.IdentifierField()) + "*")
	buf.WriteString(strings.TrimSpace(ben.NameField()) + "*")
	buf.WriteString(strings.TrimSpace(ben.AddressLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(ben.AddressLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(ben.AddressLineThreeField()) + "*")
	return ben.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on Beneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// If ID Code is present, Identifier is mandatory and vice versa.
func (ben *Beneficiary) Validate() error {
	if err := ben.fieldInclusion(); err != nil {
		return err
	}
	if ben.tag != TagBeneficiary {
		return fieldError("tag", ErrValidTagForType, ben.tag)
	}
	// Can be any Identification Code
	if err := ben.isIdentificationCode(ben.Personal.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, ben.Personal.IdentificationCode)
	}
	if err := ben.isAlphanumeric(ben.Personal.Identifier); err != nil {
		return fieldError("Identifier", err, ben.Personal.Identifier)
	}
	if err := ben.isAlphanumeric(ben.Personal.Name); err != nil {
		return fieldError("Name", err, ben.Personal.Name)
	}
	if err := ben.isAlphanumeric(ben.Personal.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, ben.Personal.Address.AddressLineOne)
	}
	if err := ben.isAlphanumeric(ben.Personal.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, ben.Personal.Address.AddressLineTwo)
	}
	if err := ben.isAlphanumeric(ben.Personal.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, ben.Personal.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ben *Beneficiary) fieldInclusion() error {
	if ben.Personal.IdentificationCode != "" && ben.Personal.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if ben.Personal.IdentificationCode == "" && ben.Personal.Identifier != "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (ben *Beneficiary) IdentificationCodeField() string {
	return ben.alphaField(ben.Personal.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (ben *Beneficiary) IdentifierField() string {
	return ben.alphaField(ben.Personal.Identifier, 34)
}

// NameField gets a string of the Name field
func (ben *Beneficiary) NameField() string {
	return ben.alphaField(ben.Personal.Name, 35)
}

// AddressLineOneField gets a string of AddressLineOne field
func (ben *Beneficiary) AddressLineOneField() string {
	return ben.alphaField(ben.Personal.Address.AddressLineOne, 35)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (ben *Beneficiary) AddressLineTwoField() string {
	return ben.alphaField(ben.Personal.Address.AddressLineTwo, 35)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (ben *Beneficiary) AddressLineThreeField() string {
	return ben.alphaField(ben.Personal.Address.AddressLineThree, 35)
}

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

// AccountDebitedDrawdown is the account which is debited in a drawdown
type AccountDebitedDrawdown struct {
	// tag
	tag string
	// Identification Code * `D` - Debit
	IdentificationCode string `json:"identificationCode"`
	// Identifier
	Identifier string `json:"identifier"`
	// Name
	Name    string  `json:"name"`
	Address Address `json:"address,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAccountDebitedDrawdown returns a new AccountDebitedDrawdown
func NewAccountDebitedDrawdown() *AccountDebitedDrawdown {
	debitDD := &AccountDebitedDrawdown{
		tag: TagAccountDebitedDrawdown,
	}
	return debitDD
}

// Parse takes the input string and parses the AccountDebitedDrawdown values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (debitDD *AccountDebitedDrawdown) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 14 || dataLen > 186 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [14, 186] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	debitDD.tag = record[:6]
	debitDD.IdentificationCode = debitDD.parseStringField(record[6:7])

	optionalFields := strings.Split(record[7:], "*")
	if len(optionalFields) >= 1 {
		debitDD.Identifier = debitDD.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		debitDD.Name = debitDD.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		debitDD.Address.AddressLineOne = debitDD.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		debitDD.Address.AddressLineTwo = debitDD.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		debitDD.Address.AddressLineThree = debitDD.parseStringField(optionalFields[4])
	}
	return nil
}

func (debitDD *AccountDebitedDrawdown) UnmarshalJSON(data []byte) error {
	type Alias AccountDebitedDrawdown
	aux := struct {
		*Alias
	}{
		(*Alias)(debitDD),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	debitDD.tag = TagAccountDebitedDrawdown
	return nil
}

// String writes AccountDebitedDrawdown
func (debitDD *AccountDebitedDrawdown) String() string {
	var buf strings.Builder
	buf.Grow(181)
	buf.WriteString(debitDD.tag)
	buf.WriteString(debitDD.IdentificationCodeField())
	buf.WriteString(strings.TrimSpace(debitDD.IdentifierField()) + "*")
	buf.WriteString(strings.TrimSpace(debitDD.NameField()) + "*")
	buf.WriteString(strings.TrimSpace(debitDD.AddressLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(debitDD.AddressLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(debitDD.AddressLineThreeField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on AccountDebitedDrawdown and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (debitDD *AccountDebitedDrawdown) Validate() error {
	if err := debitDD.fieldInclusion(); err != nil {
		return err
	}
	if debitDD.tag != TagAccountDebitedDrawdown {
		return fieldError("tag", ErrValidTagForType, debitDD.tag)
	}
	if err := debitDD.isIdentificationCode(debitDD.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, debitDD.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch debitDD.IdentificationCode {
	case
		DemandDepositAccountNumber:
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, debitDD.IdentificationCode)
	}
	if err := debitDD.isAlphanumeric(debitDD.Identifier); err != nil {
		return fieldError("Identifier", err, debitDD.Identifier)
	}
	if err := debitDD.isAlphanumeric(debitDD.Name); err != nil {
		return fieldError("Name", err, debitDD.Name)
	}
	if err := debitDD.isAlphanumeric(debitDD.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, debitDD.Address.AddressLineOne)
	}
	if err := debitDD.isAlphanumeric(debitDD.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, debitDD.Address.AddressLineTwo)
	}
	if err := debitDD.isAlphanumeric(debitDD.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, debitDD.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (debitDD *AccountDebitedDrawdown) fieldInclusion() error {
	if debitDD.IdentificationCode == "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	if debitDD.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if debitDD.Name == "" {
		return fieldError("Name", ErrFieldRequired)
	}
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (debitDD *AccountDebitedDrawdown) IdentificationCodeField() string {
	return debitDD.alphaField(debitDD.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (debitDD *AccountDebitedDrawdown) IdentifierField() string {
	return debitDD.alphaField(debitDD.Identifier, 34)
}

// NameField gets a string of the Name field
func (debitDD *AccountDebitedDrawdown) NameField() string {
	return debitDD.alphaField(debitDD.Name, 35)
}

// AddressLineOneField gets a string of AddressLineOne field
func (debitDD *AccountDebitedDrawdown) AddressLineOneField() string {
	return debitDD.alphaField(debitDD.Address.AddressLineOne, 35)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (debitDD *AccountDebitedDrawdown) AddressLineTwoField() string {
	return debitDD.alphaField(debitDD.Address.AddressLineTwo, 35)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (debitDD *AccountDebitedDrawdown) AddressLineThreeField() string {
	return debitDD.alphaField(debitDD.Address.AddressLineThree, 35)
}

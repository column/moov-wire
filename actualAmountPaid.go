// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// ActualAmountPaid is the actual amount paid
type ActualAmountPaid struct {
	// tag
	tag string
	// RemittanceAmount is remittance amounts
	RemittanceAmount RemittanceAmount `json:"remittanceAmount,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewActualAmountPaid returns a new ActualAmountPaid
func NewActualAmountPaid() *ActualAmountPaid {
	aap := &ActualAmountPaid{
		tag: TagActualAmountPaid,
	}
	return aap
}

// Parse takes the input string and parses the ActualAmountPaid values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (aap *ActualAmountPaid) Parse(record string) error {
	aap.tag = record[:6]
	aap.RemittanceAmount.CurrencyCode = aap.parseStringField(record[6:9])
	if delim := strings.IndexByte(record, '*'); delim > 0 {
		aap.RemittanceAmount.Amount = aap.parseStringField(record[9:delim])
	} else {
		aap.RemittanceAmount.Amount = aap.parseStringField(record[9:])
	}

	return nil
}

func (aap *ActualAmountPaid) UnmarshalJSON(data []byte) error {
	type Alias ActualAmountPaid
	aux := struct {
		*Alias
	}{
		(*Alias)(aap),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	aap.tag = TagActualAmountPaid
	return nil
}

// String writes ActualAmountPaid
func (aap *ActualAmountPaid) String() string {
	var buf strings.Builder
	buf.Grow(28)
	buf.WriteString(aap.tag)
	buf.WriteString(aap.CurrencyCodeField())
	buf.WriteString(strings.TrimSpace(aap.AmountField()) + "*")
	return aap.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on ActualAmountPaid and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// Currency Code and Amount are mandatory for each set of remittance data.
func (aap *ActualAmountPaid) Validate() error {
	if err := aap.fieldInclusion(); err != nil {
		return err
	}
	if aap.tag != TagActualAmountPaid {
		return fieldError("tag", ErrValidTagForType, aap.tag)
	}
	if err := aap.isCurrencyCode(aap.RemittanceAmount.CurrencyCode); err != nil {
		return fieldError("CurrencyCode", err, aap.RemittanceAmount.CurrencyCode)
	}
	if err := aap.isAmount(aap.RemittanceAmount.Amount); err != nil {
		return fieldError("Amount", err, aap.RemittanceAmount.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (aap *ActualAmountPaid) fieldInclusion() error {
	if aap.RemittanceAmount.Amount == "" {
		return fieldError("Amount", ErrFieldRequired)
	}
	if aap.RemittanceAmount.CurrencyCode == "" {
		return fieldError("CurrencyCode", ErrFieldRequired)

	}
	return nil
}

// CurrencyCodeField gets a string of the CurrencyCode field
func (aap *ActualAmountPaid) CurrencyCodeField() string {
	return aap.alphaField(aap.RemittanceAmount.CurrencyCode, 3)
}

// AmountField gets a string of the Amount field
func (aap *ActualAmountPaid) AmountField() string {
	return aap.alphaField(aap.RemittanceAmount.Amount, 19)
}

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

// FIBeneficiaryFIAdvice is the financial institution beneficiary financial institution
type FIBeneficiaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFIAdvice returns a new FIBeneficiaryFIAdvice
func NewFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	fibfia := &FIBeneficiaryFIAdvice{
		tag: TagFIBeneficiaryFIAdvice,
	}
	return fibfia
}

// Parse takes the input string and parses the FIBeneficiaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfia *FIBeneficiaryFIAdvice) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 10 || dataLen > 206 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [10, 206] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	fibfia.tag = record[:6]
	fibfia.Advice.AdviceCode = fibfia.parseStringField(record[6:9])

	optionalFields := strings.Split(record[9:], "*")
	fibfia.Advice.LineOne = fibfia.parseStringField(optionalFields[0])
	if len(optionalFields) >= 2 {
		fibfia.Advice.LineTwo = fibfia.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		fibfia.Advice.LineThree = fibfia.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		fibfia.Advice.LineFour = fibfia.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		fibfia.Advice.LineFive = fibfia.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		fibfia.Advice.LineSix = fibfia.parseStringField(optionalFields[5])
	}
	return nil
}

func (fibfia *FIBeneficiaryFIAdvice) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryFIAdvice
	aux := struct {
		*Alias
	}{
		(*Alias)(fibfia),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fibfia.tag = TagFIBeneficiaryFIAdvice
	return nil
}

// String writes FIBeneficiaryFIAdvice
func (fibfia *FIBeneficiaryFIAdvice) String() string {
	var buf strings.Builder
	buf.Grow(206)
	buf.WriteString(fibfia.tag)
	buf.WriteString(fibfia.AdviceCodeField())
	buf.WriteString(strings.TrimSpace(fibfia.LineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfia.LineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfia.LineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfia.LineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfia.LineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfia.LineSixField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfia *FIBeneficiaryFIAdvice) Validate() error {
	if fibfia.tag != TagFIBeneficiaryFIAdvice {
		return fieldError("tag", ErrValidTagForType, fibfia.tag)
	}
	if err := fibfia.isAdviceCode(fibfia.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fibfia.Advice.AdviceCode)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, fibfia.Advice.LineOne)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfia.Advice.LineTwo)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, fibfia.Advice.LineThree)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, fibfia.Advice.LineFour)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, fibfia.Advice.LineFive)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, fibfia.Advice.LineSix)
	}
	return nil
}

// AdviceCodeField gets a string of the AdviceCode field
func (fibfia *FIBeneficiaryFIAdvice) AdviceCodeField() string {
	return fibfia.alphaField(fibfia.Advice.AdviceCode, 3)
}

// LineOneField gets a string of the LineOne field
func (fibfia *FIBeneficiaryFIAdvice) LineOneField() string {
	return fibfia.alphaField(fibfia.Advice.LineOne, 26)
}

// LineTwoField gets a string of the LineTwo field
func (fibfia *FIBeneficiaryFIAdvice) LineTwoField() string {
	return fibfia.alphaField(fibfia.Advice.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fibfia *FIBeneficiaryFIAdvice) LineThreeField() string {
	return fibfia.alphaField(fibfia.Advice.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fibfia *FIBeneficiaryFIAdvice) LineFourField() string {
	return fibfia.alphaField(fibfia.Advice.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fibfia *FIBeneficiaryFIAdvice) LineFiveField() string {
	return fibfia.alphaField(fibfia.Advice.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fibfia *FIBeneficiaryFIAdvice) LineSixField() string {
	return fibfia.alphaField(fibfia.Advice.LineSix, 33)
}

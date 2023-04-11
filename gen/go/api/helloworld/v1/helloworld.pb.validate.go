// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: helloworld/v1/helloworld.proto

package helloworldv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SayHelloRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SayHelloRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayHelloRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SayHelloRequestMultiError, or nil if none found.
func (m *SayHelloRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SayHelloRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return SayHelloRequestMultiError(errors)
	}

	return nil
}

// SayHelloRequestMultiError is an error wrapping multiple validation errors
// returned by SayHelloRequest.ValidateAll() if the designated constraints
// aren't met.
type SayHelloRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayHelloRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayHelloRequestMultiError) AllErrors() []error { return m }

// SayHelloRequestValidationError is the validation error returned by
// SayHelloRequest.Validate if the designated constraints aren't met.
type SayHelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHelloRequestValidationError) ErrorName() string { return "SayHelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e SayHelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHelloRequestValidationError{}

// Validate checks the field values on SayHelloResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SayHelloResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayHelloResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SayHelloResponseMultiError, or nil if none found.
func (m *SayHelloResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SayHelloResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Error

	// no validation rules for Msg

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SayHelloResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SayHelloResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SayHelloResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SayHelloResponseMultiError(errors)
	}

	return nil
}

// SayHelloResponseMultiError is an error wrapping multiple validation errors
// returned by SayHelloResponse.ValidateAll() if the designated constraints
// aren't met.
type SayHelloResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayHelloResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayHelloResponseMultiError) AllErrors() []error { return m }

// SayHelloResponseValidationError is the validation error returned by
// SayHelloResponse.Validate if the designated constraints aren't met.
type SayHelloResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHelloResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHelloResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHelloResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHelloResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHelloResponseValidationError) ErrorName() string { return "SayHelloResponseValidationError" }

// Error satisfies the builtin error interface
func (e SayHelloResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHelloResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHelloResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHelloResponseValidationError{}

// Validate checks the field values on SayHiRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SayHiRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayHiRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SayHiRequestMultiError, or
// nil if none found.
func (m *SayHiRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SayHiRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return SayHiRequestMultiError(errors)
	}

	return nil
}

// SayHiRequestMultiError is an error wrapping multiple validation errors
// returned by SayHiRequest.ValidateAll() if the designated constraints aren't met.
type SayHiRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayHiRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayHiRequestMultiError) AllErrors() []error { return m }

// SayHiRequestValidationError is the validation error returned by
// SayHiRequest.Validate if the designated constraints aren't met.
type SayHiRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHiRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHiRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHiRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHiRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHiRequestValidationError) ErrorName() string { return "SayHiRequestValidationError" }

// Error satisfies the builtin error interface
func (e SayHiRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHiRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHiRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHiRequestValidationError{}

// Validate checks the field values on SayHiResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SayHiResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayHiResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SayHiResponseMultiError, or
// nil if none found.
func (m *SayHiResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SayHiResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Error

	// no validation rules for Msg

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SayHiResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SayHiResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SayHiResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SayHiResponseMultiError(errors)
	}

	return nil
}

// SayHiResponseMultiError is an error wrapping multiple validation errors
// returned by SayHiResponse.ValidateAll() if the designated constraints
// aren't met.
type SayHiResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayHiResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayHiResponseMultiError) AllErrors() []error { return m }

// SayHiResponseValidationError is the validation error returned by
// SayHiResponse.Validate if the designated constraints aren't met.
type SayHiResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHiResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHiResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHiResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHiResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHiResponseValidationError) ErrorName() string { return "SayHiResponseValidationError" }

// Error satisfies the builtin error interface
func (e SayHiResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHiResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHiResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHiResponseValidationError{}

// Validate checks the field values on SayHelloResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SayHelloResponse_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayHelloResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SayHelloResponse_DataMultiError, or nil if none found.
func (m *SayHelloResponse_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *SayHelloResponse_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for AgeNumber

	// no validation rules for Sex

	// no validation rules for Metadata

	if len(errors) > 0 {
		return SayHelloResponse_DataMultiError(errors)
	}

	return nil
}

// SayHelloResponse_DataMultiError is an error wrapping multiple validation
// errors returned by SayHelloResponse_Data.ValidateAll() if the designated
// constraints aren't met.
type SayHelloResponse_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayHelloResponse_DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayHelloResponse_DataMultiError) AllErrors() []error { return m }

// SayHelloResponse_DataValidationError is the validation error returned by
// SayHelloResponse_Data.Validate if the designated constraints aren't met.
type SayHelloResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHelloResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHelloResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHelloResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHelloResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHelloResponse_DataValidationError) ErrorName() string {
	return "SayHelloResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e SayHelloResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHelloResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHelloResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHelloResponse_DataValidationError{}

// Validate checks the field values on SayHiResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SayHiResponse_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayHiResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SayHiResponse_DataMultiError, or nil if none found.
func (m *SayHiResponse_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *SayHiResponse_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for AgeNumber

	if len(errors) > 0 {
		return SayHiResponse_DataMultiError(errors)
	}

	return nil
}

// SayHiResponse_DataMultiError is an error wrapping multiple validation errors
// returned by SayHiResponse_Data.ValidateAll() if the designated constraints
// aren't met.
type SayHiResponse_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayHiResponse_DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayHiResponse_DataMultiError) AllErrors() []error { return m }

// SayHiResponse_DataValidationError is the validation error returned by
// SayHiResponse_Data.Validate if the designated constraints aren't met.
type SayHiResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHiResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHiResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHiResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHiResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHiResponse_DataValidationError) ErrorName() string {
	return "SayHiResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e SayHiResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHiResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHiResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHiResponse_DataValidationError{}
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tango_api_service/service.proto

package tango_api_service

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

// Validate checks the field values on Employee with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Employee) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Employee with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmployeeMultiError, or nil
// if none found.
func (m *Employee) ValidateAll() error {
	return m.validate(true)
}

func (m *Employee) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return EmployeeMultiError(errors)
	}

	return nil
}

// EmployeeMultiError is an error wrapping multiple validation errors returned
// by Employee.ValidateAll() if the designated constraints aren't met.
type EmployeeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmployeeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmployeeMultiError) AllErrors() []error { return m }

// EmployeeValidationError is the validation error returned by
// Employee.Validate if the designated constraints aren't met.
type EmployeeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmployeeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmployeeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmployeeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmployeeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmployeeValidationError) ErrorName() string { return "EmployeeValidationError" }

// Error satisfies the builtin error interface
func (e EmployeeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmployee.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmployeeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmployeeValidationError{}

// Validate checks the field values on GetEmployeeV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetEmployeeV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEmployeeV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEmployeeV1RequestMultiError, or nil if none found.
func (m *GetEmployeeV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEmployeeV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetEmployeeId() <= 0 {
		err := GetEmployeeV1RequestValidationError{
			field:  "EmployeeId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetEmployeeV1RequestMultiError(errors)
	}

	return nil
}

// GetEmployeeV1RequestMultiError is an error wrapping multiple validation
// errors returned by GetEmployeeV1Request.ValidateAll() if the designated
// constraints aren't met.
type GetEmployeeV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEmployeeV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEmployeeV1RequestMultiError) AllErrors() []error { return m }

// GetEmployeeV1RequestValidationError is the validation error returned by
// GetEmployeeV1Request.Validate if the designated constraints aren't met.
type GetEmployeeV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEmployeeV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEmployeeV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEmployeeV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEmployeeV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEmployeeV1RequestValidationError) ErrorName() string {
	return "GetEmployeeV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEmployeeV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEmployeeV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEmployeeV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEmployeeV1RequestValidationError{}

// Validate checks the field values on GetEmployeeV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetEmployeeV1Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEmployeeV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEmployeeV1ResponseMultiError, or nil if none found.
func (m *GetEmployeeV1Response) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEmployeeV1Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEmployee()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetEmployeeV1ResponseValidationError{
					field:  "Employee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetEmployeeV1ResponseValidationError{
					field:  "Employee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEmployee()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEmployeeV1ResponseValidationError{
				field:  "Employee",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetEmployeeV1ResponseMultiError(errors)
	}

	return nil
}

// GetEmployeeV1ResponseMultiError is an error wrapping multiple validation
// errors returned by GetEmployeeV1Response.ValidateAll() if the designated
// constraints aren't met.
type GetEmployeeV1ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEmployeeV1ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEmployeeV1ResponseMultiError) AllErrors() []error { return m }

// GetEmployeeV1ResponseValidationError is the validation error returned by
// GetEmployeeV1Response.Validate if the designated constraints aren't met.
type GetEmployeeV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEmployeeV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEmployeeV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEmployeeV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEmployeeV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEmployeeV1ResponseValidationError) ErrorName() string {
	return "GetEmployeeV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEmployeeV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEmployeeV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEmployeeV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEmployeeV1ResponseValidationError{}

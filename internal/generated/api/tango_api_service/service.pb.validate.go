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

// Validate checks the field values on GetTangoParamsV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTangoParamsV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTangoParamsV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTangoParamsV1RequestMultiError, or nil if none found.
func (m *GetTangoParamsV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTangoParamsV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetSystemName()) < 1 {
		err := GetTangoParamsV1RequestValidationError{
			field:  "SystemName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetParameterName()) < 1 {
		err := GetTangoParamsV1RequestValidationError{
			field:  "ParameterName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_GetTangoParamsV1Request_StartTime_Pattern.MatchString(m.GetStartTime()) {
		err := GetTangoParamsV1RequestValidationError{
			field:  "StartTime",
			reason: "value does not match regex pattern \"^\\\\d{2,4}\\\\-\\\\d{1,2}\\\\-\\\\d{1,2}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.MemberName != nil {

		if utf8.RuneCountInString(m.GetMemberName()) < 1 {
			err := GetTangoParamsV1RequestValidationError{
				field:  "MemberName",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.EndTime != nil {

		if !_GetTangoParamsV1Request_EndTime_Pattern.MatchString(m.GetEndTime()) {
			err := GetTangoParamsV1RequestValidationError{
				field:  "EndTime",
				reason: "value does not match regex pattern \"^\\\\d{2,4}\\\\-\\\\d{1,2}\\\\-\\\\d{1,2}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return GetTangoParamsV1RequestMultiError(errors)
	}

	return nil
}

// GetTangoParamsV1RequestMultiError is an error wrapping multiple validation
// errors returned by GetTangoParamsV1Request.ValidateAll() if the designated
// constraints aren't met.
type GetTangoParamsV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTangoParamsV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTangoParamsV1RequestMultiError) AllErrors() []error { return m }

// GetTangoParamsV1RequestValidationError is the validation error returned by
// GetTangoParamsV1Request.Validate if the designated constraints aren't met.
type GetTangoParamsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTangoParamsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTangoParamsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTangoParamsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTangoParamsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTangoParamsV1RequestValidationError) ErrorName() string {
	return "GetTangoParamsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTangoParamsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTangoParamsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTangoParamsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTangoParamsV1RequestValidationError{}

var _GetTangoParamsV1Request_StartTime_Pattern = regexp.MustCompile("^\\d{2,4}\\-\\d{1,2}\\-\\d{1,2}$")

var _GetTangoParamsV1Request_EndTime_Pattern = regexp.MustCompile("^\\d{2,4}\\-\\d{1,2}\\-\\d{1,2}$")

// Validate checks the field values on GetTangoParamsV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTangoParamsV1Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTangoParamsV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTangoParamsV1ResponseMultiError, or nil if none found.
func (m *GetTangoParamsV1Response) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTangoParamsV1Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ParamType

	for idx, item := range m.GetScalarParams() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTangoParamsV1ResponseValidationError{
						field:  fmt.Sprintf("ScalarParams[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTangoParamsV1ResponseValidationError{
						field:  fmt.Sprintf("ScalarParams[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTangoParamsV1ResponseValidationError{
					field:  fmt.Sprintf("ScalarParams[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetArrayParmas() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTangoParamsV1ResponseValidationError{
						field:  fmt.Sprintf("ArrayParmas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTangoParamsV1ResponseValidationError{
						field:  fmt.Sprintf("ArrayParmas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTangoParamsV1ResponseValidationError{
					field:  fmt.Sprintf("ArrayParmas[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTangoParamsV1ResponseMultiError(errors)
	}

	return nil
}

// GetTangoParamsV1ResponseMultiError is an error wrapping multiple validation
// errors returned by GetTangoParamsV1Response.ValidateAll() if the designated
// constraints aren't met.
type GetTangoParamsV1ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTangoParamsV1ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTangoParamsV1ResponseMultiError) AllErrors() []error { return m }

// GetTangoParamsV1ResponseValidationError is the validation error returned by
// GetTangoParamsV1Response.Validate if the designated constraints aren't met.
type GetTangoParamsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTangoParamsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTangoParamsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTangoParamsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTangoParamsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTangoParamsV1ResponseValidationError) ErrorName() string {
	return "GetTangoParamsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTangoParamsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTangoParamsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTangoParamsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTangoParamsV1ResponseValidationError{}

// Validate checks the field values on GetTangoParamsV1Response_ScalarParam
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *GetTangoParamsV1Response_ScalarParam) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTangoParamsV1Response_ScalarParam
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetTangoParamsV1Response_ScalarParamMultiError, or nil if none found.
func (m *GetTangoParamsV1Response_ScalarParam) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTangoParamsV1Response_ScalarParam) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RawValueR

	// no validation rules for RawValueW

	// no validation rules for DataTime

	if len(errors) > 0 {
		return GetTangoParamsV1Response_ScalarParamMultiError(errors)
	}

	return nil
}

// GetTangoParamsV1Response_ScalarParamMultiError is an error wrapping multiple
// validation errors returned by
// GetTangoParamsV1Response_ScalarParam.ValidateAll() if the designated
// constraints aren't met.
type GetTangoParamsV1Response_ScalarParamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTangoParamsV1Response_ScalarParamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTangoParamsV1Response_ScalarParamMultiError) AllErrors() []error { return m }

// GetTangoParamsV1Response_ScalarParamValidationError is the validation error
// returned by GetTangoParamsV1Response_ScalarParam.Validate if the designated
// constraints aren't met.
type GetTangoParamsV1Response_ScalarParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTangoParamsV1Response_ScalarParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTangoParamsV1Response_ScalarParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTangoParamsV1Response_ScalarParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTangoParamsV1Response_ScalarParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTangoParamsV1Response_ScalarParamValidationError) ErrorName() string {
	return "GetTangoParamsV1Response_ScalarParamValidationError"
}

// Error satisfies the builtin error interface
func (e GetTangoParamsV1Response_ScalarParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTangoParamsV1Response_ScalarParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTangoParamsV1Response_ScalarParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTangoParamsV1Response_ScalarParamValidationError{}

// Validate checks the field values on GetTangoParamsV1Response_ArrayParam with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetTangoParamsV1Response_ArrayParam) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTangoParamsV1Response_ArrayParam
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetTangoParamsV1Response_ArrayParamMultiError, or nil if none found.
func (m *GetTangoParamsV1Response_ArrayParam) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTangoParamsV1Response_ArrayParam) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DataTime

	if len(errors) > 0 {
		return GetTangoParamsV1Response_ArrayParamMultiError(errors)
	}

	return nil
}

// GetTangoParamsV1Response_ArrayParamMultiError is an error wrapping multiple
// validation errors returned by
// GetTangoParamsV1Response_ArrayParam.ValidateAll() if the designated
// constraints aren't met.
type GetTangoParamsV1Response_ArrayParamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTangoParamsV1Response_ArrayParamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTangoParamsV1Response_ArrayParamMultiError) AllErrors() []error { return m }

// GetTangoParamsV1Response_ArrayParamValidationError is the validation error
// returned by GetTangoParamsV1Response_ArrayParam.Validate if the designated
// constraints aren't met.
type GetTangoParamsV1Response_ArrayParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTangoParamsV1Response_ArrayParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTangoParamsV1Response_ArrayParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTangoParamsV1Response_ArrayParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTangoParamsV1Response_ArrayParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTangoParamsV1Response_ArrayParamValidationError) ErrorName() string {
	return "GetTangoParamsV1Response_ArrayParamValidationError"
}

// Error satisfies the builtin error interface
func (e GetTangoParamsV1Response_ArrayParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTangoParamsV1Response_ArrayParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTangoParamsV1Response_ArrayParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTangoParamsV1Response_ArrayParamValidationError{}

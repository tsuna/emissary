// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/csrf/v2/csrf.proto

package csrfv2

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

// Validate checks the field values on CsrfPolicy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CsrfPolicy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CsrfPolicy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CsrfPolicyMultiError, or
// nil if none found.
func (m *CsrfPolicy) ValidateAll() error {
	return m.validate(true)
}

func (m *CsrfPolicy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetFilterEnabled() == nil {
		err := CsrfPolicyValidationError{
			field:  "FilterEnabled",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetFilterEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CsrfPolicyValidationError{
					field:  "FilterEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CsrfPolicyValidationError{
					field:  "FilterEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilterEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CsrfPolicyValidationError{
				field:  "FilterEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetShadowEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CsrfPolicyValidationError{
					field:  "ShadowEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CsrfPolicyValidationError{
					field:  "ShadowEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShadowEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CsrfPolicyValidationError{
				field:  "ShadowEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAdditionalOrigins() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CsrfPolicyValidationError{
						field:  fmt.Sprintf("AdditionalOrigins[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CsrfPolicyValidationError{
						field:  fmt.Sprintf("AdditionalOrigins[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CsrfPolicyValidationError{
					field:  fmt.Sprintf("AdditionalOrigins[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CsrfPolicyMultiError(errors)
	}
	return nil
}

// CsrfPolicyMultiError is an error wrapping multiple validation errors
// returned by CsrfPolicy.ValidateAll() if the designated constraints aren't met.
type CsrfPolicyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CsrfPolicyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CsrfPolicyMultiError) AllErrors() []error { return m }

// CsrfPolicyValidationError is the validation error returned by
// CsrfPolicy.Validate if the designated constraints aren't met.
type CsrfPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CsrfPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CsrfPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CsrfPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CsrfPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CsrfPolicyValidationError) ErrorName() string { return "CsrfPolicyValidationError" }

// Error satisfies the builtin error interface
func (e CsrfPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCsrfPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CsrfPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CsrfPolicyValidationError{}

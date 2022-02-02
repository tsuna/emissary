// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/auth/tls.proto

package auth

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

// Validate checks the field values on UpstreamTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpstreamTlsContext) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpstreamTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpstreamTlsContextMultiError, or nil if none found.
func (m *UpstreamTlsContext) ValidateAll() error {
	return m.validate(true)
}

func (m *UpstreamTlsContext) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonTlsContext()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamTlsContextValidationError{
					field:  "CommonTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamTlsContextValidationError{
					field:  "CommonTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonTlsContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamTlsContextValidationError{
				field:  "CommonTlsContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetSni()) > 255 {
		err := UpstreamTlsContextValidationError{
			field:  "Sni",
			reason: "value length must be at most 255 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for AllowRenegotiation

	if all {
		switch v := interface{}(m.GetMaxSessionKeys()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamTlsContextValidationError{
					field:  "MaxSessionKeys",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamTlsContextValidationError{
					field:  "MaxSessionKeys",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaxSessionKeys()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamTlsContextValidationError{
				field:  "MaxSessionKeys",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpstreamTlsContextMultiError(errors)
	}
	return nil
}

// UpstreamTlsContextMultiError is an error wrapping multiple validation errors
// returned by UpstreamTlsContext.ValidateAll() if the designated constraints
// aren't met.
type UpstreamTlsContextMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpstreamTlsContextMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpstreamTlsContextMultiError) AllErrors() []error { return m }

// UpstreamTlsContextValidationError is the validation error returned by
// UpstreamTlsContext.Validate if the designated constraints aren't met.
type UpstreamTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamTlsContextValidationError) ErrorName() string {
	return "UpstreamTlsContextValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamTlsContextValidationError{}

// Validate checks the field values on DownstreamTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownstreamTlsContext) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownstreamTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownstreamTlsContextMultiError, or nil if none found.
func (m *DownstreamTlsContext) ValidateAll() error {
	return m.validate(true)
}

func (m *DownstreamTlsContext) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonTlsContext()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownstreamTlsContextValidationError{
					field:  "CommonTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownstreamTlsContextValidationError{
					field:  "CommonTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonTlsContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownstreamTlsContextValidationError{
				field:  "CommonTlsContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRequireClientCertificate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownstreamTlsContextValidationError{
					field:  "RequireClientCertificate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownstreamTlsContextValidationError{
					field:  "RequireClientCertificate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequireClientCertificate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownstreamTlsContextValidationError{
				field:  "RequireClientCertificate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRequireSni()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownstreamTlsContextValidationError{
					field:  "RequireSni",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownstreamTlsContextValidationError{
					field:  "RequireSni",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequireSni()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownstreamTlsContextValidationError{
				field:  "RequireSni",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetSessionTimeout(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = DownstreamTlsContextValidationError{
				field:  "SessionTimeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			lt := time.Duration(4294967296*time.Second + 0*time.Nanosecond)
			gte := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur < gte || dur >= lt {
				err := DownstreamTlsContextValidationError{
					field:  "SessionTimeout",
					reason: "value must be inside range [0s, 1193046h28m16s)",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	switch m.SessionTicketKeysType.(type) {

	case *DownstreamTlsContext_SessionTicketKeys:

		if all {
			switch v := interface{}(m.GetSessionTicketKeys()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DownstreamTlsContextValidationError{
						field:  "SessionTicketKeys",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DownstreamTlsContextValidationError{
						field:  "SessionTicketKeys",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSessionTicketKeys()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DownstreamTlsContextValidationError{
					field:  "SessionTicketKeys",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DownstreamTlsContext_SessionTicketKeysSdsSecretConfig:

		if all {
			switch v := interface{}(m.GetSessionTicketKeysSdsSecretConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DownstreamTlsContextValidationError{
						field:  "SessionTicketKeysSdsSecretConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DownstreamTlsContextValidationError{
						field:  "SessionTicketKeysSdsSecretConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSessionTicketKeysSdsSecretConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DownstreamTlsContextValidationError{
					field:  "SessionTicketKeysSdsSecretConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DownstreamTlsContext_DisableStatelessSessionResumption:
		// no validation rules for DisableStatelessSessionResumption

	}

	if len(errors) > 0 {
		return DownstreamTlsContextMultiError(errors)
	}
	return nil
}

// DownstreamTlsContextMultiError is an error wrapping multiple validation
// errors returned by DownstreamTlsContext.ValidateAll() if the designated
// constraints aren't met.
type DownstreamTlsContextMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownstreamTlsContextMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownstreamTlsContextMultiError) AllErrors() []error { return m }

// DownstreamTlsContextValidationError is the validation error returned by
// DownstreamTlsContext.Validate if the designated constraints aren't met.
type DownstreamTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownstreamTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownstreamTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownstreamTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownstreamTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownstreamTlsContextValidationError) ErrorName() string {
	return "DownstreamTlsContextValidationError"
}

// Error satisfies the builtin error interface
func (e DownstreamTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownstreamTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownstreamTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownstreamTlsContextValidationError{}

// Validate checks the field values on CommonTlsContext with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CommonTlsContext) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonTlsContextMultiError, or nil if none found.
func (m *CommonTlsContext) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonTlsContext) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTlsParams()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommonTlsContextValidationError{
					field:  "TlsParams",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonTlsContextValidationError{
					field:  "TlsParams",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTlsParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContextValidationError{
				field:  "TlsParams",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTlsCertificates() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  fmt.Sprintf("TlsCertificates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  fmt.Sprintf("TlsCertificates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  fmt.Sprintf("TlsCertificates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetTlsCertificateSdsSecretConfigs()) > 1 {
		err := CommonTlsContextValidationError{
			field:  "TlsCertificateSdsSecretConfigs",
			reason: "value must contain no more than 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetTlsCertificateSdsSecretConfigs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  fmt.Sprintf("TlsCertificateSdsSecretConfigs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  fmt.Sprintf("TlsCertificateSdsSecretConfigs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  fmt.Sprintf("TlsCertificateSdsSecretConfigs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.ValidationContextType.(type) {

	case *CommonTlsContext_ValidationContext:

		if all {
			switch v := interface{}(m.GetValidationContext()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  "ValidationContext",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  "ValidationContext",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetValidationContext()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "ValidationContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CommonTlsContext_ValidationContextSdsSecretConfig:

		if all {
			switch v := interface{}(m.GetValidationContextSdsSecretConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  "ValidationContextSdsSecretConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  "ValidationContextSdsSecretConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetValidationContextSdsSecretConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "ValidationContextSdsSecretConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CommonTlsContext_CombinedValidationContext:

		if all {
			switch v := interface{}(m.GetCombinedValidationContext()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  "CombinedValidationContext",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonTlsContextValidationError{
						field:  "CombinedValidationContext",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCombinedValidationContext()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "CombinedValidationContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CommonTlsContextMultiError(errors)
	}
	return nil
}

// CommonTlsContextMultiError is an error wrapping multiple validation errors
// returned by CommonTlsContext.ValidateAll() if the designated constraints
// aren't met.
type CommonTlsContextMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonTlsContextMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonTlsContextMultiError) AllErrors() []error { return m }

// CommonTlsContextValidationError is the validation error returned by
// CommonTlsContext.Validate if the designated constraints aren't met.
type CommonTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonTlsContextValidationError) ErrorName() string { return "CommonTlsContextValidationError" }

// Error satisfies the builtin error interface
func (e CommonTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonTlsContextValidationError{}

// Validate checks the field values on
// CommonTlsContext_CombinedCertificateValidationContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonTlsContext_CombinedCertificateValidationContext) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CommonTlsContext_CombinedCertificateValidationContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonTlsContext_CombinedCertificateValidationContextMultiError, or nil if
// none found.
func (m *CommonTlsContext_CombinedCertificateValidationContext) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonTlsContext_CombinedCertificateValidationContext) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDefaultValidationContext() == nil {
		err := CommonTlsContext_CombinedCertificateValidationContextValidationError{
			field:  "DefaultValidationContext",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDefaultValidationContext()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommonTlsContext_CombinedCertificateValidationContextValidationError{
					field:  "DefaultValidationContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonTlsContext_CombinedCertificateValidationContextValidationError{
					field:  "DefaultValidationContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDefaultValidationContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContext_CombinedCertificateValidationContextValidationError{
				field:  "DefaultValidationContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetValidationContextSdsSecretConfig() == nil {
		err := CommonTlsContext_CombinedCertificateValidationContextValidationError{
			field:  "ValidationContextSdsSecretConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetValidationContextSdsSecretConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommonTlsContext_CombinedCertificateValidationContextValidationError{
					field:  "ValidationContextSdsSecretConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonTlsContext_CombinedCertificateValidationContextValidationError{
					field:  "ValidationContextSdsSecretConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValidationContextSdsSecretConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContext_CombinedCertificateValidationContextValidationError{
				field:  "ValidationContextSdsSecretConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CommonTlsContext_CombinedCertificateValidationContextMultiError(errors)
	}
	return nil
}

// CommonTlsContext_CombinedCertificateValidationContextMultiError is an error
// wrapping multiple validation errors returned by
// CommonTlsContext_CombinedCertificateValidationContext.ValidateAll() if the
// designated constraints aren't met.
type CommonTlsContext_CombinedCertificateValidationContextMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonTlsContext_CombinedCertificateValidationContextMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonTlsContext_CombinedCertificateValidationContextMultiError) AllErrors() []error {
	return m
}

// CommonTlsContext_CombinedCertificateValidationContextValidationError is the
// validation error returned by
// CommonTlsContext_CombinedCertificateValidationContext.Validate if the
// designated constraints aren't met.
type CommonTlsContext_CombinedCertificateValidationContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) ErrorName() string {
	return "CommonTlsContext_CombinedCertificateValidationContextValidationError"
}

// Error satisfies the builtin error interface
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonTlsContext_CombinedCertificateValidationContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonTlsContext_CombinedCertificateValidationContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonTlsContext_CombinedCertificateValidationContextValidationError{}

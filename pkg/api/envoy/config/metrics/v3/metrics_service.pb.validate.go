// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/metrics/v3/metrics_service.proto

package metricsv3

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

	v3 "github.com/datawire/ambassador/v2/pkg/api/envoy/config/core/v3"
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

	_ = v3.ApiVersion(0)
)

// Validate checks the field values on MetricsServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MetricsServiceConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MetricsServiceConfigMultiError, or nil if none found.
func (m *MetricsServiceConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsServiceConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetGrpcService() == nil {
		err := MetricsServiceConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetGrpcService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetricsServiceConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetricsServiceConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetricsServiceConfigValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := v3.ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		err := MetricsServiceConfigValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetReportCountersAsDeltas()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetricsServiceConfigValidationError{
					field:  "ReportCountersAsDeltas",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetricsServiceConfigValidationError{
					field:  "ReportCountersAsDeltas",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReportCountersAsDeltas()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetricsServiceConfigValidationError{
				field:  "ReportCountersAsDeltas",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for EmitTagsAsLabels

	if len(errors) > 0 {
		return MetricsServiceConfigMultiError(errors)
	}
	return nil
}

// MetricsServiceConfigMultiError is an error wrapping multiple validation
// errors returned by MetricsServiceConfig.ValidateAll() if the designated
// constraints aren't met.
type MetricsServiceConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsServiceConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsServiceConfigMultiError) AllErrors() []error { return m }

// MetricsServiceConfigValidationError is the validation error returned by
// MetricsServiceConfig.Validate if the designated constraints aren't met.
type MetricsServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsServiceConfigValidationError) ErrorName() string {
	return "MetricsServiceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsServiceConfigValidationError{}

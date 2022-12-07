// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cerbos/audit/v1/audit.proto

package auditv1

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

// Validate checks the field values on AccessLogEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AccessLogEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccessLogEntry with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AccessLogEntryMultiError,
// or nil if none found.
func (m *AccessLogEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *AccessLogEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CallId

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccessLogEntryValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccessLogEntryValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogEntryValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPeer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccessLogEntryValidationError{
					field:  "Peer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccessLogEntryValidationError{
					field:  "Peer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPeer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogEntryValidationError{
				field:  "Peer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	{
		sorted_keys := make([]string, len(m.GetMetadata()))
		i := 0
		for key := range m.GetMetadata() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetMetadata()[key]
			_ = val

			// no validation rules for Metadata[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, AccessLogEntryValidationError{
							field:  fmt.Sprintf("Metadata[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, AccessLogEntryValidationError{
							field:  fmt.Sprintf("Metadata[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return AccessLogEntryValidationError{
						field:  fmt.Sprintf("Metadata[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for Method

	// no validation rules for StatusCode

	if len(errors) > 0 {
		return AccessLogEntryMultiError(errors)
	}

	return nil
}

// AccessLogEntryMultiError is an error wrapping multiple validation errors
// returned by AccessLogEntry.ValidateAll() if the designated constraints
// aren't met.
type AccessLogEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccessLogEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccessLogEntryMultiError) AllErrors() []error { return m }

// AccessLogEntryValidationError is the validation error returned by
// AccessLogEntry.Validate if the designated constraints aren't met.
type AccessLogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessLogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessLogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessLogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessLogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessLogEntryValidationError) ErrorName() string { return "AccessLogEntryValidationError" }

// Error satisfies the builtin error interface
func (e AccessLogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessLogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessLogEntryValidationError{}

// Validate checks the field values on DecisionLogEntry with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DecisionLogEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DecisionLogEntry with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DecisionLogEntryMultiError, or nil if none found.
func (m *DecisionLogEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *DecisionLogEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CallId

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DecisionLogEntryValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DecisionLogEntryValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecisionLogEntryValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPeer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DecisionLogEntryValidationError{
					field:  "Peer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DecisionLogEntryValidationError{
					field:  "Peer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPeer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecisionLogEntryValidationError{
				field:  "Peer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetInputs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DecisionLogEntryValidationError{
						field:  fmt.Sprintf("Inputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DecisionLogEntryValidationError{
						field:  fmt.Sprintf("Inputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecisionLogEntryValidationError{
					field:  fmt.Sprintf("Inputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetOutputs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DecisionLogEntryValidationError{
						field:  fmt.Sprintf("Outputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DecisionLogEntryValidationError{
						field:  fmt.Sprintf("Outputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecisionLogEntryValidationError{
					field:  fmt.Sprintf("Outputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Error

	{
		sorted_keys := make([]string, len(m.GetMetadata()))
		i := 0
		for key := range m.GetMetadata() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetMetadata()[key]
			_ = val

			// no validation rules for Metadata[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, DecisionLogEntryValidationError{
							field:  fmt.Sprintf("Metadata[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, DecisionLogEntryValidationError{
							field:  fmt.Sprintf("Metadata[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return DecisionLogEntryValidationError{
						field:  fmt.Sprintf("Metadata[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	switch v := m.Method.(type) {
	case *DecisionLogEntry_CheckResources_:
		if v == nil {
			err := DecisionLogEntryValidationError{
				field:  "Method",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetCheckResources()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DecisionLogEntryValidationError{
						field:  "CheckResources",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DecisionLogEntryValidationError{
						field:  "CheckResources",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCheckResources()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecisionLogEntryValidationError{
					field:  "CheckResources",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DecisionLogEntry_PlanResources_:
		if v == nil {
			err := DecisionLogEntryValidationError{
				field:  "Method",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetPlanResources()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DecisionLogEntryValidationError{
						field:  "PlanResources",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DecisionLogEntryValidationError{
						field:  "PlanResources",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPlanResources()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecisionLogEntryValidationError{
					field:  "PlanResources",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return DecisionLogEntryMultiError(errors)
	}

	return nil
}

// DecisionLogEntryMultiError is an error wrapping multiple validation errors
// returned by DecisionLogEntry.ValidateAll() if the designated constraints
// aren't met.
type DecisionLogEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DecisionLogEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DecisionLogEntryMultiError) AllErrors() []error { return m }

// DecisionLogEntryValidationError is the validation error returned by
// DecisionLogEntry.Validate if the designated constraints aren't met.
type DecisionLogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecisionLogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecisionLogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecisionLogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecisionLogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecisionLogEntryValidationError) ErrorName() string { return "DecisionLogEntryValidationError" }

// Error satisfies the builtin error interface
func (e DecisionLogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecisionLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecisionLogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecisionLogEntryValidationError{}

// Validate checks the field values on MetaValues with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MetaValues) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetaValues with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetaValuesMultiError, or
// nil if none found.
func (m *MetaValues) ValidateAll() error {
	return m.validate(true)
}

func (m *MetaValues) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MetaValuesMultiError(errors)
	}

	return nil
}

// MetaValuesMultiError is an error wrapping multiple validation errors
// returned by MetaValues.ValidateAll() if the designated constraints aren't met.
type MetaValuesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetaValuesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetaValuesMultiError) AllErrors() []error { return m }

// MetaValuesValidationError is the validation error returned by
// MetaValues.Validate if the designated constraints aren't met.
type MetaValuesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetaValuesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetaValuesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetaValuesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetaValuesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetaValuesValidationError) ErrorName() string { return "MetaValuesValidationError" }

// Error satisfies the builtin error interface
func (e MetaValuesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetaValues.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetaValuesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetaValuesValidationError{}

// Validate checks the field values on Peer with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Peer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Peer with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PeerMultiError, or nil if none found.
func (m *Peer) ValidateAll() error {
	return m.validate(true)
}

func (m *Peer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for AuthInfo

	// no validation rules for UserAgent

	// no validation rules for ForwardedFor

	if len(errors) > 0 {
		return PeerMultiError(errors)
	}

	return nil
}

// PeerMultiError is an error wrapping multiple validation errors returned by
// Peer.ValidateAll() if the designated constraints aren't met.
type PeerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PeerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PeerMultiError) AllErrors() []error { return m }

// PeerValidationError is the validation error returned by Peer.Validate if the
// designated constraints aren't met.
type PeerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PeerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PeerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PeerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PeerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PeerValidationError) ErrorName() string { return "PeerValidationError" }

// Error satisfies the builtin error interface
func (e PeerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPeer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PeerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PeerValidationError{}

// Validate checks the field values on DecisionLogEntry_CheckResources with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DecisionLogEntry_CheckResources) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DecisionLogEntry_CheckResources with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DecisionLogEntry_CheckResourcesMultiError, or nil if none found.
func (m *DecisionLogEntry_CheckResources) ValidateAll() error {
	return m.validate(true)
}

func (m *DecisionLogEntry_CheckResources) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetInputs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DecisionLogEntry_CheckResourcesValidationError{
						field:  fmt.Sprintf("Inputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DecisionLogEntry_CheckResourcesValidationError{
						field:  fmt.Sprintf("Inputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecisionLogEntry_CheckResourcesValidationError{
					field:  fmt.Sprintf("Inputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetOutputs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DecisionLogEntry_CheckResourcesValidationError{
						field:  fmt.Sprintf("Outputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DecisionLogEntry_CheckResourcesValidationError{
						field:  fmt.Sprintf("Outputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecisionLogEntry_CheckResourcesValidationError{
					field:  fmt.Sprintf("Outputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Error

	if len(errors) > 0 {
		return DecisionLogEntry_CheckResourcesMultiError(errors)
	}

	return nil
}

// DecisionLogEntry_CheckResourcesMultiError is an error wrapping multiple
// validation errors returned by DecisionLogEntry_CheckResources.ValidateAll()
// if the designated constraints aren't met.
type DecisionLogEntry_CheckResourcesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DecisionLogEntry_CheckResourcesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DecisionLogEntry_CheckResourcesMultiError) AllErrors() []error { return m }

// DecisionLogEntry_CheckResourcesValidationError is the validation error
// returned by DecisionLogEntry_CheckResources.Validate if the designated
// constraints aren't met.
type DecisionLogEntry_CheckResourcesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecisionLogEntry_CheckResourcesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecisionLogEntry_CheckResourcesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecisionLogEntry_CheckResourcesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecisionLogEntry_CheckResourcesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecisionLogEntry_CheckResourcesValidationError) ErrorName() string {
	return "DecisionLogEntry_CheckResourcesValidationError"
}

// Error satisfies the builtin error interface
func (e DecisionLogEntry_CheckResourcesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecisionLogEntry_CheckResources.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecisionLogEntry_CheckResourcesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecisionLogEntry_CheckResourcesValidationError{}

// Validate checks the field values on DecisionLogEntry_PlanResources with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DecisionLogEntry_PlanResources) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DecisionLogEntry_PlanResources with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DecisionLogEntry_PlanResourcesMultiError, or nil if none found.
func (m *DecisionLogEntry_PlanResources) ValidateAll() error {
	return m.validate(true)
}

func (m *DecisionLogEntry_PlanResources) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInput()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DecisionLogEntry_PlanResourcesValidationError{
					field:  "Input",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DecisionLogEntry_PlanResourcesValidationError{
					field:  "Input",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecisionLogEntry_PlanResourcesValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOutput()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DecisionLogEntry_PlanResourcesValidationError{
					field:  "Output",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DecisionLogEntry_PlanResourcesValidationError{
					field:  "Output",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOutput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecisionLogEntry_PlanResourcesValidationError{
				field:  "Output",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Error

	if len(errors) > 0 {
		return DecisionLogEntry_PlanResourcesMultiError(errors)
	}

	return nil
}

// DecisionLogEntry_PlanResourcesMultiError is an error wrapping multiple
// validation errors returned by DecisionLogEntry_PlanResources.ValidateAll()
// if the designated constraints aren't met.
type DecisionLogEntry_PlanResourcesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DecisionLogEntry_PlanResourcesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DecisionLogEntry_PlanResourcesMultiError) AllErrors() []error { return m }

// DecisionLogEntry_PlanResourcesValidationError is the validation error
// returned by DecisionLogEntry_PlanResources.Validate if the designated
// constraints aren't met.
type DecisionLogEntry_PlanResourcesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecisionLogEntry_PlanResourcesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecisionLogEntry_PlanResourcesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecisionLogEntry_PlanResourcesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecisionLogEntry_PlanResourcesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecisionLogEntry_PlanResourcesValidationError) ErrorName() string {
	return "DecisionLogEntry_PlanResourcesValidationError"
}

// Error satisfies the builtin error interface
func (e DecisionLogEntry_PlanResourcesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecisionLogEntry_PlanResources.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecisionLogEntry_PlanResourcesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecisionLogEntry_PlanResourcesValidationError{}

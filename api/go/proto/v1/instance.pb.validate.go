// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/v1/instance.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Instance with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Instance) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Instance with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InstanceMultiError, or nil
// if none found.
func (m *Instance) ValidateAll() error {
	return m.validate(true)
}

func (m *Instance) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetKey()); l < 2 || l > 253 {
		err := InstanceValidationError{
			field:  "Key",
			reason: "value length must be between 2 and 253 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return InstanceMultiError(errors)
	}
	return nil
}

// InstanceMultiError is an error wrapping multiple validation errors returned
// by Instance.ValidateAll() if the designated constraints aren't met.
type InstanceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InstanceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InstanceMultiError) AllErrors() []error { return m }

// InstanceValidationError is the validation error returned by
// Instance.Validate if the designated constraints aren't met.
type InstanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InstanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InstanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InstanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InstanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InstanceValidationError) ErrorName() string { return "InstanceValidationError" }

// Error satisfies the builtin error interface
func (e InstanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInstance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InstanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InstanceValidationError{}
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/common/matcher/action/v3/skip_action.proto

package envoy_extensions_filters_common_matcher_action_v3

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

// Validate checks the field values on SkipFilter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SkipFilter) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SkipFilterValidationError is the validation error returned by
// SkipFilter.Validate if the designated constraints aren't met.
type SkipFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SkipFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SkipFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SkipFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SkipFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SkipFilterValidationError) ErrorName() string { return "SkipFilterValidationError" }

// Error satisfies the builtin error interface
func (e SkipFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSkipFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SkipFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SkipFilterValidationError{}

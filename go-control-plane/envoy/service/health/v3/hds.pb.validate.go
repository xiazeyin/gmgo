// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/health/v3/hds.proto

package envoy_service_health_v3

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

	v3 "github.com/xiazeyin/gmgo/go-control-plane/envoy/config/core/v3"
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

	_ = v3.HealthStatus(0)
)

// Validate checks the field values on Capability with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Capability) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CapabilityValidationError is the validation error returned by
// Capability.Validate if the designated constraints aren't met.
type CapabilityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CapabilityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CapabilityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CapabilityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CapabilityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CapabilityValidationError) ErrorName() string { return "CapabilityValidationError" }

// Error satisfies the builtin error interface
func (e CapabilityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCapability.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CapabilityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CapabilityValidationError{}

// Validate checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckRequestValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCapability()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckRequestValidationError{
				field:  "Capability",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HealthCheckRequestValidationError is the validation error returned by
// HealthCheckRequest.Validate if the designated constraints aren't met.
type HealthCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckRequestValidationError) ErrorName() string {
	return "HealthCheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckRequestValidationError{}

// Validate checks the field values on EndpointHealth with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EndpointHealth) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EndpointHealthValidationError{
				field:  "Endpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HealthStatus

	return nil
}

// EndpointHealthValidationError is the validation error returned by
// EndpointHealth.Validate if the designated constraints aren't met.
type EndpointHealthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointHealthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointHealthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointHealthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointHealthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointHealthValidationError) ErrorName() string { return "EndpointHealthValidationError" }

// Error satisfies the builtin error interface
func (e EndpointHealthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointHealth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointHealthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointHealthValidationError{}

// Validate checks the field values on LocalityEndpointsHealth with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LocalityEndpointsHealth) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityEndpointsHealthValidationError{
				field:  "Locality",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetEndpointsHealth() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityEndpointsHealthValidationError{
					field:  fmt.Sprintf("EndpointsHealth[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LocalityEndpointsHealthValidationError is the validation error returned by
// LocalityEndpointsHealth.Validate if the designated constraints aren't met.
type LocalityEndpointsHealthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityEndpointsHealthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityEndpointsHealthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityEndpointsHealthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityEndpointsHealthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityEndpointsHealthValidationError) ErrorName() string {
	return "LocalityEndpointsHealthValidationError"
}

// Error satisfies the builtin error interface
func (e LocalityEndpointsHealthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityEndpointsHealth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityEndpointsHealthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityEndpointsHealthValidationError{}

// Validate checks the field values on ClusterEndpointsHealth with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterEndpointsHealth) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ClusterName

	for idx, item := range m.GetLocalityEndpointsHealth() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterEndpointsHealthValidationError{
					field:  fmt.Sprintf("LocalityEndpointsHealth[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ClusterEndpointsHealthValidationError is the validation error returned by
// ClusterEndpointsHealth.Validate if the designated constraints aren't met.
type ClusterEndpointsHealthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterEndpointsHealthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterEndpointsHealthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterEndpointsHealthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterEndpointsHealthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterEndpointsHealthValidationError) ErrorName() string {
	return "ClusterEndpointsHealthValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterEndpointsHealthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterEndpointsHealth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterEndpointsHealthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterEndpointsHealthValidationError{}

// Validate checks the field values on EndpointHealthResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EndpointHealthResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEndpointsHealth() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EndpointHealthResponseValidationError{
					field:  fmt.Sprintf("EndpointsHealth[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetClusterEndpointsHealth() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EndpointHealthResponseValidationError{
					field:  fmt.Sprintf("ClusterEndpointsHealth[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EndpointHealthResponseValidationError is the validation error returned by
// EndpointHealthResponse.Validate if the designated constraints aren't met.
type EndpointHealthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointHealthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointHealthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointHealthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointHealthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointHealthResponseValidationError) ErrorName() string {
	return "EndpointHealthResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EndpointHealthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointHealthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointHealthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointHealthResponseValidationError{}

// Validate checks the field values on
// HealthCheckRequestOrEndpointHealthResponse with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HealthCheckRequestOrEndpointHealthResponse) Validate() error {
	if m == nil {
		return nil
	}

	switch m.RequestType.(type) {

	case *HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest:

		if v, ok := interface{}(m.GetHealthCheckRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckRequestOrEndpointHealthResponseValidationError{
					field:  "HealthCheckRequest",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse:

		if v, ok := interface{}(m.GetEndpointHealthResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckRequestOrEndpointHealthResponseValidationError{
					field:  "EndpointHealthResponse",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HealthCheckRequestOrEndpointHealthResponseValidationError is the validation
// error returned by HealthCheckRequestOrEndpointHealthResponse.Validate if
// the designated constraints aren't met.
type HealthCheckRequestOrEndpointHealthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) ErrorName() string {
	return "HealthCheckRequestOrEndpointHealthResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequestOrEndpointHealthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckRequestOrEndpointHealthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckRequestOrEndpointHealthResponseValidationError{}

// Validate checks the field values on LocalityEndpoints with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LocalityEndpoints) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityEndpointsValidationError{
				field:  "Locality",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityEndpointsValidationError{
					field:  fmt.Sprintf("Endpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LocalityEndpointsValidationError is the validation error returned by
// LocalityEndpoints.Validate if the designated constraints aren't met.
type LocalityEndpointsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityEndpointsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityEndpointsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityEndpointsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityEndpointsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityEndpointsValidationError) ErrorName() string {
	return "LocalityEndpointsValidationError"
}

// Error satisfies the builtin error interface
func (e LocalityEndpointsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityEndpoints.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityEndpointsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityEndpointsValidationError{}

// Validate checks the field values on ClusterHealthCheck with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ClusterName

	for idx, item := range m.GetHealthChecks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterHealthCheckValidationError{
					field:  fmt.Sprintf("HealthChecks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetLocalityEndpoints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterHealthCheckValidationError{
					field:  fmt.Sprintf("LocalityEndpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetTransportSocketMatches() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterHealthCheckValidationError{
					field:  fmt.Sprintf("TransportSocketMatches[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ClusterHealthCheckValidationError is the validation error returned by
// ClusterHealthCheck.Validate if the designated constraints aren't met.
type ClusterHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterHealthCheckValidationError) ErrorName() string {
	return "ClusterHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterHealthCheckValidationError{}

// Validate checks the field values on HealthCheckSpecifier with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheckSpecifier) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetClusterHealthChecks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckSpecifierValidationError{
					field:  fmt.Sprintf("ClusterHealthChecks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckSpecifierValidationError{
				field:  "Interval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HealthCheckSpecifierValidationError is the validation error returned by
// HealthCheckSpecifier.Validate if the designated constraints aren't met.
type HealthCheckSpecifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckSpecifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckSpecifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckSpecifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckSpecifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckSpecifierValidationError) ErrorName() string {
	return "HealthCheckSpecifierValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckSpecifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckSpecifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckSpecifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckSpecifierValidationError{}

// Validate checks the field values on HdsDummy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HdsDummy) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// HdsDummyValidationError is the validation error returned by
// HdsDummy.Validate if the designated constraints aren't met.
type HdsDummyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HdsDummyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HdsDummyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HdsDummyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HdsDummyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HdsDummyValidationError) ErrorName() string { return "HdsDummyValidationError" }

// Error satisfies the builtin error interface
func (e HdsDummyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHdsDummy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HdsDummyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HdsDummyValidationError{}

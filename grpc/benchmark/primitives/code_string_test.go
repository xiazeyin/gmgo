/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package primitives_test

import (
	"strconv"
	"testing"

	"github.com/xiazeyin/gmgo/grpc/codes"
)

type codeBench uint32

const (
	OK codeBench = iota
	Canceled
	Unknown
	InvalidArgument
	DeadlineExceeded
	NotFound
	AlreadyExists
	PermissionDenied
	ResourceExhausted
	FailedPrecondition
	Aborted
	OutOfRange
	Unimplemented
	Internal
	Unavailable
	DataLoss
	Unauthenticated
)

// The following String() function was generated by stringer.
const _Code_name = "OKCanceledUnknownInvalidArgumentDeadlineExceededNotFoundAlreadyExistsPermissionDeniedResourceExhaustedFailedPreconditionAbortedOutOfRangeUnimplementedInternalUnavailableDataLossUnauthenticated"

var _Code_index = [...]uint8{0, 2, 10, 17, 32, 48, 56, 69, 85, 102, 120, 127, 137, 150, 158, 169, 177, 192}

func (i codeBench) String() string {
	if i >= codeBench(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}

var nameMap = map[codeBench]string{
	OK:                 "OK",
	Canceled:           "Canceled",
	Unknown:            "Unknown",
	InvalidArgument:    "InvalidArgument",
	DeadlineExceeded:   "DeadlineExceeded",
	NotFound:           "NotFound",
	AlreadyExists:      "AlreadyExists",
	PermissionDenied:   "PermissionDenied",
	ResourceExhausted:  "ResourceExhausted",
	FailedPrecondition: "FailedPrecondition",
	Aborted:            "Aborted",
	OutOfRange:         "OutOfRange",
	Unimplemented:      "Unimplemented",
	Internal:           "Internal",
	Unavailable:        "Unavailable",
	DataLoss:           "DataLoss",
	Unauthenticated:    "Unauthenticated",
}

func (i codeBench) StringUsingMap() string {
	if s, ok := nameMap[i]; ok {
		return s
	}
	return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
}

func BenchmarkCodeStringStringer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := codeBench(uint32(i % 17))
		_ = c.String()
	}
	b.StopTimer()
}

func BenchmarkCodeStringMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := codeBench(uint32(i % 17))
		_ = c.StringUsingMap()
	}
	b.StopTimer()
}

// codes.Code.String() does a switch.
func BenchmarkCodeStringSwitch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := codes.Code(uint32(i % 17))
		_ = c.String()
	}
	b.StopTimer()
}

// Testing all codes (0<=c<=16) and also one overflow (17).
func BenchmarkCodeStringStringerWithOverflow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := codeBench(uint32(i % 18))
		_ = c.String()
	}
	b.StopTimer()
}

// Testing all codes (0<=c<=16) and also one overflow (17).
func BenchmarkCodeStringSwitchWithOverflow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := codes.Code(uint32(i % 18))
		_ = c.String()
	}
	b.StopTimer()
}

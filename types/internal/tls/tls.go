// Copyright 2016 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package tls implements functionality for dealing with TLS-encoded data,
// as defined in RFC 5246.  This includes parsing and generation of TLS-encoded
// data, together with utility functions for dealing with the DigitallySigned
// TLS type.
// N.B. This is copied from https://github.com/google/certificate-transparency-go/tree/master/tls
// - DO NOT MAKE CHANGES TO THIS FILE except to sync to the latest from ct-go.
package tls

import (
	"bytes"
	"reflect"
)

// This file holds utility functions for TLS encoding/decoding data
// as per RFC 5246 section 4.

// A structuralError suggests that the TLS data is valid, but the Go type
// which is receiving it doesn't match.
type structuralError struct {
	field string
	msg   string
}

func (e structuralError) Error() string { _ = "STUB: not implemented"; return "" }

// A syntaxError suggests that the TLS data is invalid.
type syntaxError struct {
	field string
	msg   string
}

func (e syntaxError) Error() string { _ = "STUB: not implemented"; return "" }

// Uint24 is an unsigned 3-byte integer.
type Uint24 uint32

// Enum is an unsigned integer.
type Enum uint64

var (
	uint8Type  = reflect.TypeOf(uint8(0))
	uint16Type = reflect.TypeOf(uint16(0))
	uint24Type = reflect.TypeOf(Uint24(0))
	uint32Type = reflect.TypeOf(uint32(0))
	uint64Type = reflect.TypeOf(uint64(0))
	enumType   = reflect.TypeOf(Enum(0))
)

// Unmarshal parses the TLS-encoded data in b and uses the reflect package to
// fill in an arbitrary value pointed at by val.  Because Unmarshal uses the
// reflect package, the structs being written to must use exported fields
// (upper case names).
//
// The mappings between TLS types and Go types is as follows; some fields
// must have tags (to indicate their encoded size).
//
//	TLS		Go		Required Tags
//	opaque		byte / uint8
//	uint8		byte / uint8
//	uint16		uint16
//	uint24		tls.Uint24
//	uint32		uint32
//	uint64		uint64
//	enum		tls.Enum	size:S or maxval:N
//	Type<N,M>	[]Type		minlen:N,maxlen:M
//	opaque[N]	[N]byte / [N]uint8
//	uint8[N]	[N]byte / [N]uint8
//	struct { }	struct { }
//	select(T) {
//	 case e1: Type	*T		selector:Field,val:e1
//	}
//
// TLS variants (RFC 5246 s4.6.1) are only supported when the value of the
// associated enumeration type is available earlier in the same enclosing
// struct, and each possible variant is marked with a selector tag (to
// indicate which field selects the variants) and a val tag (to indicate
// what value of the selector picks this particular field).
//
// For example, a TLS structure:
//
//	enum { e1(1), e2(2) } EnumType;
//	struct {
//	   EnumType sel;
//	   select(sel) {
//	      case e1: uint16
//	      case e2: uint32
//	   } data;
//	} VariantItem;
//
// would have a corresponding Go type:
//
//	type VariantItem struct {
//	   Sel    tls.Enum  `tls:"maxval:2"`
//	   Data16 *uint16   `tls:"selector:Sel,val:1"`
//	   Data32 *uint32   `tls:"selector:Sel,val:2"`
//	 }
//
// TLS fixed-length vectors of types other than opaque or uint8 are not supported.
//
// For TLS variable-length vectors that are themselves used in other vectors,
// create a single-field structure to represent the inner type. For example, for:
//
//	opaque InnerType<1..65535>;
//	struct {
//	  InnerType inners<1,65535>;
//	} Something;
//
// convert to:
//
//	type InnerType struct {
//	   Val    []byte       `tls:"minlen:1,maxlen:65535"`
//	}
//	type Something struct {
//	   Inners []InnerType  `tls:"minlen:1,maxlen:65535"`
//	}
//
// If the encoded value does not fit in the Go type, Unmarshal returns a parse error.
func Unmarshal(b []byte, val interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
func UnmarshalWithParams(b []byte, val interface{}, params string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The passed in interface{} is a pointer (to allow the value to be written
// to); extract the pointed-to object as a reflect.Value, so parseField
// can do various introspection things.

// Return the number of bytes needed to encode values up to (and including) x.
func byteCount(x uint64) uint { _ = "STUB: not implemented"; return 0 }

type fieldInfo struct {
	count    uint // Number of bytes
	countSet bool
	minlen   uint64 // Only relevant for slices
	maxlen   uint64 // Only relevant for slices
	selector string // Only relevant for select sub-values
	val      uint64 // Only relevant for select sub-values
	name     string // Used for better error messages
}

func (i *fieldInfo) fieldName() string { _ = "STUB: not implemented"; return "" }

// Given a tag string, return a fieldInfo describing the field.
func fieldTagToFieldInfo(str string, name string) (*fieldInfo, error) {
	_ = "STUB: not implemented"

	// Iterate over clauses in the tag, ignoring any that don't parse properly.
	return nil, nil
}

// Check that a value fits into a field described by a fieldInfo structure.
func (i fieldInfo) check(val uint64, fldName string) error { _ = "STUB: not implemented"; return nil }

// readVarUint reads an big-endian unsigned integer of the given size in
// bytes.
func readVarUint(data []byte, info *fieldInfo) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// parseField is the main parsing function. Given a byte slice and an offset
// (in bytes) into the data, it will try to parse a suitable ASN.1 value out
// and store it in the given Value.
func parseField(v reflect.Value, data []byte, initOffset int, info *fieldInfo) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// First look for known fixed types.

// Now deal with user-defined types.

// Assume that anything of the same kind as Enum is an Enum, so that
// users can alias types of their own to Enum.

// TLS includes a select(Enum) {..} construct, where the value of an enum
// indicates which variant field is present (like a C union). We require
// that the enum value be an earlier field in the same structure (the selector),
// and that each of the possible variant destination fields be pointers.
// So the Go mapping looks like:
//     type variantType struct {
//         Which  tls.Enum  `tls:"size:1"`                // this is the selector
//         Val1   *type1    `tls:"selector:Which,val:1"`  // this is a destination
//         Val2   *type2    `tls:"selector:Which,val:1"`  // this is a destination
//     }

// To deal with this, we track any enum-like fields and their values...

// .. and we track which selector names we've seen (in the destination field tags),
// and whether a destination for that selector has been chosen.

// Find information about this field.

// This is a possible select(Enum) destination, so first check that the referenced
// selector field has already been seen earlier in the struct.

// Is this the first mention of the selector field name?  If so, remember it.

// This destination field was not the chosen one, so make it nil (we checked
// it was a pointer above).

// We already saw a different destination field receive the value for this
// selector value, which indicates a badly annotated structure.

// Make an object of the pointed-to type and parse into that.

// Remember any possible tls.Enum values encountered in case they are selectors.

// Now we have seen all fields in the structure, check that all select(Enum) {..} selector
// fields found a destination to put their data in.

// Only byte/uint8 arrays are supported

// Slices represent variable-length vectors, which are prefixed by a length field.
// The fieldInfo indicates the size of that length field.

// Fast version for []byte

// Marshal returns the TLS encoding of val.
func Marshal(val interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalWithParams returns the TLS encoding of val, and allows field
// parameters to be specified for the top-level element.  The form
// of the params is the same as the field tags.
func MarshalWithParams(val interface{}, params string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalField(out *bytes.Buffer, v reflect.Value, info *fieldInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// First look for known fixed types.

// Now deal with user-defined types.

// Values of any Enum fields
// The comment parseField() describes the mapping of the TLS select(Enum) {..} construct;
// here we have selector and source (rather than destination) fields.

// Track which selector names we've seen (in the source field tags), and whether a source
// value for that selector has been processed.

// Find information about this field.

// This field is a possible source for a select(Enum) {..}.  First check
// the selector field name has been seen.

// Is this the first mention of the selector field name? If so, remember it.

// This source was not chosen; police that it should be nil.

// We already saw a different source field generate the value for this
// selector value, which indicates a badly annotated structure.

// Marshal from the pointed-to source object.

// Remember any tls.Enum values encountered in case they are selectors.

// Now we have seen all fields in the structure, check that all select(Enum) {..} selector
// fields found a source field get get their data from.

// Only byte/uint8 arrays are supported

// Fast version for []byte: first write the length as info.count bytes.

// Then just write the data.

// General version: use a separate Buffer to write the slice entries into.

// Now insert (and check) the size.

// Then copy the data.

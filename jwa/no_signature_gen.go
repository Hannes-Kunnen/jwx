// Code generated by tools/cmd/genjwa/main.go. DO NOT EDIT.

package jwa

import (
	"fmt"
)

// NoSignatureAlgorithm represents the "none" algorithm that indicates no signature is used.
// This is described in https://tools.ietf.org/html/rfc7518#section-3.1.
type NoSignatureAlgorithm string

// Only supported value for NoSignatureAlgorithm
const NoSignature NoSignatureAlgorithm = "none"

// Accept is used when conversion from values given by
// outside sources (such as JSON payloads) is required
func (v *NoSignatureAlgorithm) Accept(value interface{}) error {
	var tmp NoSignatureAlgorithm
	if x, ok := value.(NoSignatureAlgorithm); ok {
		tmp = x
	} else {
		var s string
		switch x := value.(type) {
		case fmt.Stringer:
			s = x.String()
		case string:
			s = x
		default:
			return fmt.Errorf(`invalid type for jwa.NoSignatureAlgorithm: %T`, value)
		}
		tmp = NoSignatureAlgorithm(s)
	}
	if tmp != NoSignature {
		return fmt.Errorf(`invalid jwa.NoSignatureAlgorithm value`)
	}

	*v = tmp
	return nil
}

// String returns the string representation of a NoSignatureAlgorithm
func (v NoSignatureAlgorithm) String() string {
	return string(v)
}

// isSignatureAlgorithm is just added so it would qualify as a SignatureAlgorithm.
func (v NoSignatureAlgorithm) isSignatureAlgorithm() {}

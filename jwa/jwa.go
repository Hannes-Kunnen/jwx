//go:generate ../tools/cmd/genjwa.sh

// Package jwa defines the various algorithm described in https://tools.ietf.org/html/rfc7518
package jwa

import "fmt"

// KeyAlgorithm is a workaround for jwk.Key being able to contain different
// types of algorithms in its `alg` field. It can only contain a SigningAlgorithm or a
// KeyEncryptionAlgorithm.
//
// Previously the storage for the `alg` field was represented as a string,
// but this caused some users to wonder why the field was not typed appropriately
// like other fields.
//
// Ideally we would like to keep track of SigningAlgorithms and KeyEncryptionAlgorithms separately,
// and force the APIs to type-check at compile time, but this allows users to pass a value from a
// jwk.Key directly
type KeyAlgorithm interface {
	String() string
	IsSymmetric() bool
	keyAlgorithm()
}

// InvalidKeyAlgorithm represents an algorithm that the library is not aware of.
type InvalidKeyAlgorithm string

func (s InvalidKeyAlgorithm) String() string {
	return string(s)
}

func (s InvalidKeyAlgorithm) IsSymmetric() bool {
	return false
}

func (InvalidKeyAlgorithm) keyAlgorithm() {}

func (InvalidKeyAlgorithm) Accept(_ interface{}) error {
	return fmt.Errorf(`jwa.InvalidKeyAlgorithm does not support Accept() method calls`)
}

// KeyAlgorithmFrom takes either a string, `jwa.SigningAlgorithm` or `jwa.KeyEncryptionAlgorithm`
// and returns a `jwa.KeyAlgorithm`.
//
// If the value cannot be handled, it returns an `jwa.InvalidKeyAlgorithm`
// object instead of returning an error. This design choice was made to allow
// users to directly pass the return value to functions such as `jws.Sign()`
func KeyAlgorithmFrom(v interface{}) KeyAlgorithm {
	switch v := v.(type) {
	case SigningAlgorithm:
		return v
	case KeyEncryptionAlgorithm:
		return v
	case string:
		var salg SigningAlgorithm
		if err := salg.Accept(v); err == nil {
			return salg
		}

		var kealg KeyEncryptionAlgorithm
		if err := kealg.Accept(v); err == nil {
			return kealg
		}

		return InvalidKeyAlgorithm(v)
	default:
		return InvalidKeyAlgorithm(fmt.Sprintf("%s", v))
	}
}

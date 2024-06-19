// Code generated by tools/cmd/genjwa/main.go. DO NOT EDIT.

package jwa

import (
	"fmt"
	"sort"
	"sync"
)

// SignatureAlgorithm represents the various signature algorithms as described in https://tools.ietf.org/html/rfc7518#section-3.1
type SignatureAlgorithm string

// Supported values for SignatureAlgorithm
const (
	ES256       SignatureAlgorithm = "ES256"  // ECDSA using P-256 and SHA-256
	ES256K      SignatureAlgorithm = "ES256K" // ECDSA using secp256k1 and SHA-256
	ES384       SignatureAlgorithm = "ES384"  // ECDSA using P-384 and SHA-384
	ES512       SignatureAlgorithm = "ES512"  // ECDSA using P-521 and SHA-512
	EdDSA       SignatureAlgorithm = "EdDSA"  // EdDSA signature algorithms
	HS256       SignatureAlgorithm = "HS256"  // HMAC using SHA-256
	HS384       SignatureAlgorithm = "HS384"  // HMAC using SHA-384
	HS512       SignatureAlgorithm = "HS512"  // HMAC using SHA-512
	NoSignature SignatureAlgorithm = "none"
	PS256       SignatureAlgorithm = "PS256" // RSASSA-PSS using SHA256 and MGF1-SHA256
	PS384       SignatureAlgorithm = "PS384" // RSASSA-PSS using SHA384 and MGF1-SHA384
	PS512       SignatureAlgorithm = "PS512" // RSASSA-PSS using SHA512 and MGF1-SHA512
	RS256       SignatureAlgorithm = "RS256" // RSASSA-PKCS-v1.5 using SHA-256
	RS384       SignatureAlgorithm = "RS384" // RSASSA-PKCS-v1.5 using SHA-384
	RS512       SignatureAlgorithm = "RS512" // RSASSA-PKCS-v1.5 using SHA-512
)

var muSignatureAlgorithms sync.RWMutex
var allSignatureAlgorithms map[SignatureAlgorithm]struct{}
var listSignatureAlgorithm []SignatureAlgorithm

func init() {
	muSignatureAlgorithms.Lock()
	defer muSignatureAlgorithms.Unlock()
	allSignatureAlgorithms = make(map[SignatureAlgorithm]struct{})
	allSignatureAlgorithms[ES256] = struct{}{}
	allSignatureAlgorithms[ES256K] = struct{}{}
	allSignatureAlgorithms[ES384] = struct{}{}
	allSignatureAlgorithms[ES512] = struct{}{}
	allSignatureAlgorithms[EdDSA] = struct{}{}
	allSignatureAlgorithms[HS256] = struct{}{}
	allSignatureAlgorithms[HS384] = struct{}{}
	allSignatureAlgorithms[HS512] = struct{}{}
	allSignatureAlgorithms[NoSignature] = struct{}{}
	allSignatureAlgorithms[PS256] = struct{}{}
	allSignatureAlgorithms[PS384] = struct{}{}
	allSignatureAlgorithms[PS512] = struct{}{}
	allSignatureAlgorithms[RS256] = struct{}{}
	allSignatureAlgorithms[RS384] = struct{}{}
	allSignatureAlgorithms[RS512] = struct{}{}
	rebuildSignatureAlgorithm()
}

// RegisterSignatureAlgorithm registers a new SignatureAlgorithm so that the jwx can properly handle the new value.
// Duplicates will silently be ignored
func RegisterSignatureAlgorithm(v SignatureAlgorithm) {
	muSignatureAlgorithms.Lock()
	defer muSignatureAlgorithms.Unlock()
	if _, ok := allSignatureAlgorithms[v]; !ok {
		allSignatureAlgorithms[v] = struct{}{}
		rebuildSignatureAlgorithm()
	}
}

// UnregisterSignatureAlgorithm unregisters a SignatureAlgorithm from its known database.
// Non-existentn entries will silently be ignored
func UnregisterSignatureAlgorithm(v SignatureAlgorithm) {
	muSignatureAlgorithms.Lock()
	defer muSignatureAlgorithms.Unlock()
	if _, ok := allSignatureAlgorithms[v]; ok {
		delete(allSignatureAlgorithms, v)
		rebuildSignatureAlgorithm()
	}
}

func rebuildSignatureAlgorithm() {
	listSignatureAlgorithm = make([]SignatureAlgorithm, 0, len(allSignatureAlgorithms))
	for v := range allSignatureAlgorithms {
		listSignatureAlgorithm = append(listSignatureAlgorithm, v)
	}
	sort.Slice(listSignatureAlgorithm, func(i, j int) bool {
		return string(listSignatureAlgorithm[i]) < string(listSignatureAlgorithm[j])
	})
}

// SignatureAlgorithms returns a list of all available values for SignatureAlgorithm
func SignatureAlgorithms() []SignatureAlgorithm {
	muSignatureAlgorithms.RLock()
	defer muSignatureAlgorithms.RUnlock()
	return listSignatureAlgorithm
}

// Accept is used when conversion from values given by
// outside sources (such as JSON payloads) is required
func (v *SignatureAlgorithm) Accept(value interface{}) error {
	var tmp SignatureAlgorithm
	if x, ok := value.(SignatureAlgorithm); ok {
		tmp = x
	} else {
		var s string
		switch x := value.(type) {
		case fmt.Stringer:
			s = x.String()
		case string:
			s = x
		default:
			return fmt.Errorf(`invalid type for jwa.SignatureAlgorithm: %T`, value)
		}
		tmp = SignatureAlgorithm(s)
	}
	if _, ok := allSignatureAlgorithms[tmp]; !ok {
		return fmt.Errorf(`invalid jwa.SignatureAlgorithm value`)
	}

	*v = tmp
	return nil
}

// String returns the string representation of a SignatureAlgorithm
func (v SignatureAlgorithm) String() string {
	return string(v)
}

// IsSymmetric returns true if the algorithm is a symmetric type
func (v SignatureAlgorithm) IsSymmetric() bool {
	switch v {
	case HS256, HS384, HS512:
		return true
	}
	return false
}

package jwa

// ToDo: This file should be generated

import (
	"sync"
)

// SignatureAlgorithm represents the various algorithms a payload can use to indicate its signature.
// This encapsulates all algorithms that can be used to sign payloads (SigningAlgorithm) and
// NoSignature.
// These are described in https://tools.ietf.org/html/rfc7518#section-3.1.
type SignatureAlgorithm interface {
	String() string
	isSignatureAlgorithm()
}

var muSignatureAlgorithms sync.RWMutex
var allSignatureAlgorithms map[SignatureAlgorithm]struct{}
var listSignatureAlgorithm []SignatureAlgorithm

func init() {
	muSignatureAlgorithms.Lock()
	defer muSignatureAlgorithms.Unlock()
	muSigningAlgorithms.Lock()
	defer muSigningAlgorithms.Unlock()
	allSignatureAlgorithms = make(map[SignatureAlgorithm]struct{})
	allSignatureAlgorithms[ES256] = struct{}{}
	allSignatureAlgorithms[ES256K] = struct{}{}
	allSignatureAlgorithms[ES384] = struct{}{}
	allSignatureAlgorithms[ES512] = struct{}{}
	allSignatureAlgorithms[EdDSA] = struct{}{}
	allSignatureAlgorithms[HS256] = struct{}{}
	allSignatureAlgorithms[HS384] = struct{}{}
	allSignatureAlgorithms[HS512] = struct{}{}
	allSignatureAlgorithms[PS256] = struct{}{}
	allSignatureAlgorithms[PS384] = struct{}{}
	allSignatureAlgorithms[PS512] = struct{}{}
	allSignatureAlgorithms[RS256] = struct{}{}
	allSignatureAlgorithms[RS384] = struct{}{}
	allSignatureAlgorithms[RS512] = struct{}{}
	allSignatureAlgorithms[NoSignature] = struct{}{}
	rebuildSignatureAlgorithm()
}

func rebuildSignatureAlgorithm() {
	listSignatureAlgorithm = make([]SignatureAlgorithm, 0, len(allSigningAlgorithms)+1)
	listSignatureAlgorithm = append(listSignatureAlgorithm, NoSignature)
	for v := range allSigningAlgorithms {
		listSignatureAlgorithm = append(listSignatureAlgorithm, v)
	}
}

// SignatureAlgorithms returns a list of all available values for SignatureAlgorithm
func SignatureAlgorithms() []SignatureAlgorithm {
	muSignatureAlgorithms.RLock()
	defer muSignatureAlgorithms.RUnlock()
	return listSignatureAlgorithm
}

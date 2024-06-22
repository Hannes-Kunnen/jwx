package jwa

// ToDo: This file should be generated

import (
	"fmt"
	"sort"
	"sync"
)

// SignatureAlgorithm represents the various algorithms a payload can use to indicate its signature.
// This encapsulates all algorithms that can be used to sign payloads (SigningAlgorithm) and
// NoSignature.
// These are described in https://tools.ietf.org/html/rfc7518#section-3.1.
type SignatureAlgorithm interface {
	String() string
	signatureAlgorithm()
}

var muSignatureAlgorithms sync.RWMutex
var listSignatureAlgorithm []SignatureAlgorithm

func rebuildSignatureAlgorithm() {
	listSignatureAlgorithm = make([]SignatureAlgorithm, 0, len(allSigningAlgorithms)+1)
	listSignatureAlgorithm = append(listSignatureAlgorithm, NoSignature)
	for v := range allSigningAlgorithms {
		listSignatureAlgorithm = append(listSignatureAlgorithm, v)
	}
	sort.Slice(listSignatureAlgorithm, func(i, j int) bool {
		return listSignatureAlgorithm[i].String() < listSignatureAlgorithm[j].String()
	})
}

// SignatureAlgorithms returns a list of all available values for SignatureAlgorithm
func SignatureAlgorithms() []SignatureAlgorithm {
	muSignatureAlgorithms.RLock()
	defer muSignatureAlgorithms.RUnlock()
	return listSignatureAlgorithm
}

// AcceptSignatureAlgorithm is used when conversion from values given by
// outside sources (such as JSON payloads) is required
func AcceptSignatureAlgorithm(value interface{}) (SignatureAlgorithm, error) {
	var signingAlg SigningAlgorithm
	if err := signingAlg.Accept(value); err == nil {
		return signingAlg, nil
	}

	var noSignatureAlg NoSignatureAlgorithm
	if err := noSignatureAlg.Accept(value); err == nil {
		return noSignatureAlg, nil
	}

	return nil, fmt.Errorf(`invalid jwa.SignatureAlgorithm value`)
}

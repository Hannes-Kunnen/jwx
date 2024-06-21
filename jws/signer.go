package jws

import (
	"fmt"
	"sync"

	"github.com/lestrrat-go/jwx/v2/jwa"
)

type SignerFactory interface {
	Create() (Signer, error)
}
type SignerFactoryFn func() (Signer, error)

func (fn SignerFactoryFn) Create() (Signer, error) {
	return fn()
}

var muSignerDB sync.RWMutex
var signerDB map[jwa.SigningAlgorithm]SignerFactory

// RegisterSigner is used to register a factory object that creates
// Signer objects based on the given algorithm. Previous object instantiated
// by the factory is discarded.
//
// For example, if you would like to provide a custom signer for
// jwa.EdDSA, use this function to register a `SignerFactory`
// (probably in your `init()`)
//
// Unlike the `UnregisterSigner` function, this function automatically
// calls `jwa.RegisterSigningAlgorithm` to register the algorithm
// in the known-algorithms database.
func RegisterSigner(alg jwa.SigningAlgorithm, f SignerFactory) {
	jwa.RegisterSigningAlgorithm(alg)
	muSignerDB.Lock()
	signerDB[alg] = f
	muSignerDB.Unlock()

	// Remove previous signer, if there was one
	removeSigner(alg)
}

// UnregisterSigner removes the signer factory associated with
// the given algorithm, as well as the signer instance created
// by the factory.
//
// Note that when you call this function, the algorithm itself is
// not automatically unregistered from the known algorithms database.
// This is because the algorithm may still be required for verification or
// some other operation (however unlikely, it is still possible).
// Therefore, in order to completely remove the algorithm, you must
// call `jwa.UnregisterSigningAlgorithm` yourself.
func UnregisterSigner(alg jwa.SigningAlgorithm) {
	muSignerDB.Lock()
	delete(signerDB, alg)
	muSignerDB.Unlock()
	// Remove previous signer
	removeSigner(alg)
}

func init() {
	signerDB = make(map[jwa.SigningAlgorithm]SignerFactory)

	for _, alg := range []jwa.SigningAlgorithm{jwa.RS256, jwa.RS384, jwa.RS512, jwa.PS256, jwa.PS384, jwa.PS512} {
		RegisterSigner(alg, func(alg jwa.SigningAlgorithm) SignerFactory {
			return SignerFactoryFn(func() (Signer, error) {
				return newRSASigner(alg), nil
			})
		}(alg))
	}

	for _, alg := range []jwa.SigningAlgorithm{jwa.ES256, jwa.ES384, jwa.ES512, jwa.ES256K} {
		RegisterSigner(alg, func(alg jwa.SigningAlgorithm) SignerFactory {
			return SignerFactoryFn(func() (Signer, error) {
				return newECDSASigner(alg), nil
			})
		}(alg))
	}

	for _, alg := range []jwa.SigningAlgorithm{jwa.HS256, jwa.HS384, jwa.HS512} {
		RegisterSigner(alg, func(alg jwa.SigningAlgorithm) SignerFactory {
			return SignerFactoryFn(func() (Signer, error) {
				return newHMACSigner(alg), nil
			})
		}(alg))
	}

	RegisterSigner(jwa.EdDSA, SignerFactoryFn(func() (Signer, error) {
		return newEdDSASigner(), nil
	}))
}

// NewSigner creates a signer that signs payloads using the given signing algorithm.
func NewSigner(alg jwa.SigningAlgorithm) (Signer, error) {
	muSignerDB.RLock()
	f, ok := signerDB[alg]
	muSignerDB.RUnlock()

	if ok {
		return f.Create()
	}
	return nil, fmt.Errorf(`unsupported signing algorithm "%s"`, alg)
}

type noneSigner struct{}

func (noneSigner) Algorithm() jwa.SignatureAlgorithm {
	return jwa.NoSignature
}

func (noneSigner) Sign([]byte, interface{}) ([]byte, error) {
	return nil, nil
}

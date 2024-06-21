package jwa_test

import (
	"fmt"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/stretchr/testify/assert"
)

type stringer struct {
	src string
}

func (s stringer) String() string {
	return s.src
}

func TestSanity(t *testing.T) {
	var k1 jwa.KeyAlgorithm = jwa.RS256
	if _, ok := k1.(jwa.SignatureAlgorithm); !assert.True(t, ok, `converting k1 to jws.SignatureAlgorithm should succeed`) {
		return
	}
	if _, ok := k1.(jwa.SigningAlgorithm); !assert.True(t, ok, `converting k1 to jws.SigningAlgorithm should succeed`) {
		return
	}
	if _, ok := k1.(jwa.KeyEncryptionAlgorithm); !assert.False(t, ok, `converting k1 to jws.KeyEncryptionAlgorithm should fail`) {
		return
	}
	var k2 jwa.KeyAlgorithm = jwa.DIRECT
	if _, ok := k2.(jwa.SignatureAlgorithm); !assert.False(t, ok, `converting k2 to jws.SignatureAlgorithm should fail`) {
		return
	}
	if _, ok := k2.(jwa.SigningAlgorithm); !assert.False(t, ok, `converting k2 to jws.SigningAlgorithm should fail`) {
		return
	}
	if _, ok := k2.(jwa.KeyEncryptionAlgorithm); !assert.True(t, ok, `converting k2 to jws.KeyEncryptionAlgorithm should succeed`) {
		return
	}
}

// ToDo Maybe also generate this? (Maybe entire file)
func TestSignatureAlgorithmToKeyAlgorithm(t *testing.T) {
	valid := []jwa.SignatureAlgorithm{
		jwa.ES256,
		jwa.ES256K,
		jwa.ES384,
		jwa.ES512,
		jwa.EdDSA,
		jwa.HS256,
		jwa.HS384,
		jwa.HS512,
		jwa.PS256,
		jwa.PS384,
		jwa.PS512,
		jwa.RS256,
		jwa.RS384,
		jwa.RS512,
	}
	invalid := []jwa.SignatureAlgorithm{
		jwa.NoSignature,
	}
	for _, alg := range valid {
		if _, ok := alg.(jwa.KeyAlgorithm); !assert.True(t, ok, `converting to jws.KeyAlgorithm should succeed`) {
			return
		}
	}
	for _, alg := range invalid {
		if _, ok := alg.(jwa.KeyAlgorithm); !assert.False(t, ok, `converting to jws.KeyAlgorithm should fail`) {
			return
		}
	}
}

func TestKeyAlgorithmFrom(t *testing.T) {
	testcases := []struct {
		Input interface{}
		Error bool
	}{
		{
			Input: jwa.RS256,
		},
		{
			Input: jwa.DIRECT,
		},
		{
			Input: jwa.A128CBC_HS256,
			Error: true,
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprintf("%T", tc.Input), func(t *testing.T) {
			alg := jwa.KeyAlgorithmFrom(tc.Input)
			if tc.Error {
				if !assert.IsType(t, alg, jwa.InvalidKeyAlgorithm(""), `key should be invalid`) {
					return
				}
			} else {
				if !assert.Equal(t, alg, tc.Input, `key should be valid`) {
					return
				}
			}
		})
	}
}

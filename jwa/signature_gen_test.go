package jwa_test

// ToDo: This file should actually be generated
import (
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/stretchr/testify/assert"
)

func TestSignatureAlgorithm(t *testing.T) {
	t.Run(`check list of elements`, func(t *testing.T) {
		t.Parallel()
		var expected = map[jwa.SignatureAlgorithm]struct{}{
			jwa.ES256:       {},
			jwa.ES256K:      {},
			jwa.ES384:       {},
			jwa.ES512:       {},
			jwa.EdDSA:       {},
			jwa.HS256:       {},
			jwa.HS384:       {},
			jwa.HS512:       {},
			jwa.NoSignature: {},
			jwa.PS256:       {},
			jwa.PS384:       {},
			jwa.PS512:       {},
			jwa.RS256:       {},
			jwa.RS384:       {},
			jwa.RS512:       {},
		}
		for _, v := range jwa.SignatureAlgorithms() {
			if _, ok := expected[v]; !assert.True(t, ok, `%s should be in the expected list`, v) {
				return
			}
			delete(expected, v)
		}
		if !assert.Len(t, expected, 0) {
			return
		}
	})
}

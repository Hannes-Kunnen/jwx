package jwa_test

// ToDo: This file should actually be generated
import (
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/stretchr/testify/assert"
)

func TestSignatureAlgorithm(t *testing.T) {
	t.Parallel()
	t.Run(`accept jwa constant ES256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.ES256)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string ES256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("ES256")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for ES256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "ES256"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant ES256K`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.ES256K)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES256K, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string ES256K`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("ES256K")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES256K, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for ES256K`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "ES256K"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES256K, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant ES384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.ES384)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string ES384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("ES384")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for ES384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "ES384"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant ES512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.ES512)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string ES512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("ES512")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for ES512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "ES512"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.ES512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant EdDSA`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.EdDSA)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.EdDSA, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string EdDSA`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("EdDSA")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.EdDSA, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for EdDSA`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "EdDSA"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.EdDSA, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant HS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.HS256)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string HS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("HS256")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for HS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "HS256"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant HS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.HS384)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string HS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("HS384")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for HS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "HS384"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant HS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.HS512)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string HS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("HS512")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for HS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "HS512"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.HS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant NoSignature`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.NoSignature)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.NoSignature, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string none`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("none")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.NoSignature, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for none`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "none"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.NoSignature, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant PS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.PS256)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string PS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("PS256")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for PS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "PS256"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant PS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.PS384)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string PS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("PS384")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for PS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "PS384"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant PS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.PS512)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string PS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("PS512")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for PS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "PS512"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.PS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant RS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.RS256)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string RS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("RS256")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for RS256`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "RS256"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant RS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.RS384)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string RS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("RS384")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for RS384`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "RS384"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant RS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(jwa.RS512)
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string RS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm("RS512")
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for RS512`, func(t *testing.T) {
		t.Parallel()
		dst, err := jwa.AcceptSignatureAlgorithm(stringer{src: "RS512"})
		if !assert.NoError(t, err, `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.RS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`bail out on random integer value`, func(t *testing.T) {
		t.Parallel()
		_, err := jwa.AcceptSignatureAlgorithm(1)
		if !assert.Error(t, err, `accept should fail`) {
			return
		}
	})
	t.Run(`do not accept invalid (totally made up) string value`, func(t *testing.T) {
		t.Parallel()
		_, err := jwa.AcceptSignatureAlgorithm(`totallyInvfalidValue`)
		if !assert.Error(t, err, `accept should fail`) {
			return
		}
	})
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

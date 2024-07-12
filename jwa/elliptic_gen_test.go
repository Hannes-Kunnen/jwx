// Code generated by tools/cmd/genjwa/main.go. DO NOT EDIT

package jwa_test

import (
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/stretchr/testify/assert"
)

func TestEllipticCurveAlgorithm(t *testing.T) {
	t.Parallel()
	t.Run(`accept jwa constant Ed25519`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.Ed25519), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.Ed25519, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string Ed25519`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept("Ed25519"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.Ed25519, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for Ed25519`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "Ed25519"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.Ed25519, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`stringification for Ed25519`, func(t *testing.T) {
		t.Parallel()
		if !assert.Equal(t, "Ed25519", jwa.Ed25519.String(), `stringified value matches`) {
			return
		}
	})
	t.Run(`accept jwa constant Ed448`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.Ed448), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.Ed448, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string Ed448`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept("Ed448"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.Ed448, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for Ed448`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "Ed448"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.Ed448, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`stringification for Ed448`, func(t *testing.T) {
		t.Parallel()
		if !assert.Equal(t, "Ed448", jwa.Ed448.String(), `stringified value matches`) {
			return
		}
	})
	t.Run(`accept jwa constant P256`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.P256), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string P-256`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept("P-256"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for P-256`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "P-256"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`stringification for P-256`, func(t *testing.T) {
		t.Parallel()
		if !assert.Equal(t, "P-256", jwa.P256.String(), `stringified value matches`) {
			return
		}
	})
	t.Run(`accept jwa constant P384`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.P384), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string P-384`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept("P-384"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for P-384`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "P-384"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`stringification for P-384`, func(t *testing.T) {
		t.Parallel()
		if !assert.Equal(t, "P-384", jwa.P384.String(), `stringified value matches`) {
			return
		}
	})
	t.Run(`accept jwa constant P521`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.P521), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P521, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string P-521`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept("P-521"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P521, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for P-521`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "P-521"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.P521, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`stringification for P-521`, func(t *testing.T) {
		t.Parallel()
		if !assert.Equal(t, "P-521", jwa.P521.String(), `stringified value matches`) {
			return
		}
	})
	t.Run(`accept jwa constant X25519`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.X25519), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.X25519, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string X25519`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept("X25519"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.X25519, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for X25519`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "X25519"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.X25519, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`stringification for X25519`, func(t *testing.T) {
		t.Parallel()
		if !assert.Equal(t, "X25519", jwa.X25519.String(), `stringified value matches`) {
			return
		}
	})
	t.Run(`accept jwa constant X448`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.X448), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.X448, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string X448`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept("X448"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.X448, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for X448`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "X448"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.X448, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`stringification for X448`, func(t *testing.T) {
		t.Parallel()
		if !assert.Equal(t, "X448", jwa.X448.String(), `stringified value matches`) {
			return
		}
	})
	t.Run(`do not accept invalid constant InvalidEllipticCurve`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.Error(t, dst.Accept(jwa.InvalidEllipticCurve), `accept should fail`) {
			return
		}
	})
	t.Run(`bail out on random integer value`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.Error(t, dst.Accept(1), `accept should fail`) {
			return
		}
	})
	t.Run(`do not accept invalid (totally made up) string value`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.EllipticCurveAlgorithm
		if !assert.Error(t, dst.Accept(`totallyInvfalidValue`), `accept should fail`) {
			return
		}
	})
	t.Run(`check list of elements`, func(t *testing.T) {
		t.Parallel()
		var expected = map[jwa.EllipticCurveAlgorithm]struct{}{
			jwa.Ed25519: {},
			jwa.Ed448:   {},
			jwa.P256:    {},
			jwa.P384:    {},
			jwa.P521:    {},
			jwa.X25519:  {},
			jwa.X448:    {},
		}
		for _, v := range jwa.EllipticCurveAlgorithms() {
			// There is no good way to detect from a test if es256k (secp256k1)
			// is supported, so just allow it
			if v.String() == `secp256k1` {
				continue
			}
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

// Note: this test can NOT be run in parallel as it uses options with global effect.
func TestEllipticCurveAlgorithmCustomAlgorithm(t *testing.T) {
	// These subtests can NOT be run in parallel as options with global effect change.
	customAlgorithm := jwa.EllipticCurveAlgorithm("custom-algorithm")
	// Unregister the custom algorithm, in case tests fail.
	t.Cleanup(func() {
		jwa.UnregisterEllipticCurveAlgorithm(customAlgorithm)
	})
	t.Run(`with custom algorithm registered`, func(t *testing.T) {
		jwa.RegisterEllipticCurveAlgorithm(customAlgorithm)
		t.Run(`accept variable used to register custom algorithm`, func(t *testing.T) {
			t.Parallel()
			var dst jwa.EllipticCurveAlgorithm
			if !assert.NoError(t, dst.Accept(customAlgorithm), `accept is successful`) {
				return
			}
			assert.Equal(t, customAlgorithm, dst, `accepted value should be equal to variable`)
		})
		t.Run(`accept the string custom-algorithm`, func(t *testing.T) {
			t.Parallel()
			var dst jwa.EllipticCurveAlgorithm
			if !assert.NoError(t, dst.Accept(`custom-algorithm`), `accept is successful`) {
				return
			}
			assert.Equal(t, customAlgorithm, dst, `accepted value should be equal to variable`)
		})
		t.Run(`accept fmt.Stringer for custom-algorithm`, func(t *testing.T) {
			t.Parallel()
			var dst jwa.EllipticCurveAlgorithm
			if !assert.NoError(t, dst.Accept(stringer{src: `custom-algorithm`}), `accept is successful`) {
				return
			}
			assert.Equal(t, customAlgorithm, dst, `accepted value should be equal to variable`)
		})
	})
	t.Run(`with custom algorithm deregistered`, func(t *testing.T) {
		jwa.UnregisterEllipticCurveAlgorithm(customAlgorithm)
		t.Run(`reject variable used to register custom algorithm`, func(t *testing.T) {
			t.Parallel()
			var dst jwa.EllipticCurveAlgorithm
			assert.Error(t, dst.Accept(customAlgorithm), `accept failed`)
		})
		t.Run(`reject the string custom-algorithm`, func(t *testing.T) {
			t.Parallel()
			var dst jwa.EllipticCurveAlgorithm
			assert.Error(t, dst.Accept(`custom-algorithm`), `accept failed`)
		})
		t.Run(`reject fmt.Stringer for custom-algorithm`, func(t *testing.T) {
			t.Parallel()
			var dst jwa.EllipticCurveAlgorithm
			assert.Error(t, dst.Accept(stringer{src: `custom-algorithm`}), `accept failed`)
		})
	})
}

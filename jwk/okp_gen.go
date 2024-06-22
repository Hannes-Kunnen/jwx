// Code generated by tools/cmd/genjwk/main.go. DO NOT EDIT.

package jwk

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/iter/mapiter"
	"github.com/lestrrat-go/jwx/v2/cert"
	"github.com/lestrrat-go/jwx/v2/internal/base64"
	"github.com/lestrrat-go/jwx/v2/internal/iter"
	"github.com/lestrrat-go/jwx/v2/internal/json"
	"github.com/lestrrat-go/jwx/v2/internal/pool"
	"github.com/lestrrat-go/jwx/v2/jwa"
)

const (
	OKPCrvKey = "crv"
	OKPDKey   = "d"
	OKPXKey   = "x"
)

type OKPPublicKey interface {
	Key
	FromRaw(interface{}) error
	Crv() jwa.EllipticCurveAlgorithm
	X() []byte
}

type okpPublicKey struct {
	algorithm              *jwa.KeyAlgorithm // https://tools.ietf.org/html/rfc7517#section-4.4
	crv                    *jwa.EllipticCurveAlgorithm
	keyID                  *string           // https://tools.ietf.org/html/rfc7515#section-4.1.4
	keyOps                 *KeyOperationList // https://tools.ietf.org/html/rfc7517#section-4.3
	keyUsage               *string           // https://tools.ietf.org/html/rfc7517#section-4.2
	x                      []byte
	x509CertChain          *cert.Chain // https://tools.ietf.org/html/rfc7515#section-4.1.6
	x509CertThumbprint     *string     // https://tools.ietf.org/html/rfc7515#section-4.1.7
	x509CertThumbprintS256 *string     // https://tools.ietf.org/html/rfc7515#section-4.1.8
	x509URL                *string     // https://tools.ietf.org/html/rfc7515#section-4.1.5
	privateParams          map[string]interface{}
	mu                     *sync.RWMutex
	dc                     json.DecodeCtx
}

var _ OKPPublicKey = &okpPublicKey{}
var _ Key = &okpPublicKey{}

func newOKPPublicKey() *okpPublicKey {
	return &okpPublicKey{
		mu:            &sync.RWMutex{},
		privateParams: make(map[string]interface{}),
	}
}

func (h okpPublicKey) KeyType() jwa.KeyType {
	return jwa.OKP
}

func (h okpPublicKey) IsPrivate() bool {
	return false
}

func (h *okpPublicKey) Algorithm() jwa.KeyAlgorithm {
	if h.algorithm != nil {
		return *(h.algorithm)
	}
	return nil
}

func (h *okpPublicKey) Crv() jwa.EllipticCurveAlgorithm {
	if h.crv != nil {
		return *(h.crv)
	}
	return jwa.InvalidEllipticCurve
}

func (h *okpPublicKey) KeyID() string {
	if h.keyID != nil {
		return *(h.keyID)
	}
	return ""
}

func (h *okpPublicKey) KeyOps() KeyOperationList {
	if h.keyOps != nil {
		return *(h.keyOps)
	}
	return nil
}

func (h *okpPublicKey) KeyUsage() string {
	if h.keyUsage != nil {
		return *(h.keyUsage)
	}
	return ""
}

func (h *okpPublicKey) X() []byte {
	return h.x
}

func (h *okpPublicKey) X509CertChain() *cert.Chain {
	return h.x509CertChain
}

func (h *okpPublicKey) X509CertThumbprint() string {
	if h.x509CertThumbprint != nil {
		return *(h.x509CertThumbprint)
	}
	return ""
}

func (h *okpPublicKey) X509CertThumbprintS256() string {
	if h.x509CertThumbprintS256 != nil {
		return *(h.x509CertThumbprintS256)
	}
	return ""
}

func (h *okpPublicKey) X509URL() string {
	if h.x509URL != nil {
		return *(h.x509URL)
	}
	return ""
}

func (h *okpPublicKey) makePairs() []*HeaderPair {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var pairs []*HeaderPair
	pairs = append(pairs, &HeaderPair{Key: "kty", Value: jwa.OKP})
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.crv != nil {
		pairs = append(pairs, &HeaderPair{Key: OKPCrvKey, Value: *(h.crv)})
	}
	if h.keyID != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyIDKey, Value: *(h.keyID)})
	}
	if h.keyOps != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyOpsKey, Value: *(h.keyOps)})
	}
	if h.keyUsage != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyUsageKey, Value: *(h.keyUsage)})
	}
	if h.x != nil {
		pairs = append(pairs, &HeaderPair{Key: OKPXKey, Value: h.x})
	}
	if h.x509CertChain != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertChainKey, Value: h.x509CertChain})
	}
	if h.x509CertThumbprint != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintKey, Value: *(h.x509CertThumbprint)})
	}
	if h.x509CertThumbprintS256 != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintS256Key, Value: *(h.x509CertThumbprintS256)})
	}
	if h.x509URL != nil {
		pairs = append(pairs, &HeaderPair{Key: X509URLKey, Value: *(h.x509URL)})
	}
	for k, v := range h.privateParams {
		pairs = append(pairs, &HeaderPair{Key: k, Value: v})
	}
	return pairs
}

func (h *okpPublicKey) PrivateParams() map[string]interface{} {
	return h.privateParams
}

func (h *okpPublicKey) Get(name string) (interface{}, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch name {
	case KeyTypeKey:
		return h.KeyType(), true
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case OKPCrvKey:
		if h.crv == nil {
			return nil, false
		}
		return *(h.crv), true
	case KeyIDKey:
		if h.keyID == nil {
			return nil, false
		}
		return *(h.keyID), true
	case KeyOpsKey:
		if h.keyOps == nil {
			return nil, false
		}
		return *(h.keyOps), true
	case KeyUsageKey:
		if h.keyUsage == nil {
			return nil, false
		}
		return *(h.keyUsage), true
	case OKPXKey:
		if h.x == nil {
			return nil, false
		}
		return h.x, true
	case X509CertChainKey:
		if h.x509CertChain == nil {
			return nil, false
		}
		return h.x509CertChain, true
	case X509CertThumbprintKey:
		if h.x509CertThumbprint == nil {
			return nil, false
		}
		return *(h.x509CertThumbprint), true
	case X509CertThumbprintS256Key:
		if h.x509CertThumbprintS256 == nil {
			return nil, false
		}
		return *(h.x509CertThumbprintS256), true
	case X509URLKey:
		if h.x509URL == nil {
			return nil, false
		}
		return *(h.x509URL), true
	default:
		v, ok := h.privateParams[name]
		return v, ok
	}
}

func (h *okpPublicKey) Set(name string, value interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.setNoLock(name, value)
}

func (h *okpPublicKey) setNoLock(name string, value interface{}) error {
	switch name {
	case "kty":
		return nil
	case AlgorithmKey:
		switch v := value.(type) {
		case string, jwa.SigningAlgorithm, jwa.ContentEncryptionAlgorithm:
			var tmp = jwa.KeyAlgorithmFrom(v)
			h.algorithm = &tmp
		case fmt.Stringer:
			s := v.String()
			var tmp = jwa.KeyAlgorithmFrom(s)
			h.algorithm = &tmp
		default:
			return fmt.Errorf(`invalid type for %s key: %T`, AlgorithmKey, value)
		}
		return nil
	case OKPCrvKey:
		if v, ok := value.(jwa.EllipticCurveAlgorithm); ok {
			h.crv = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, OKPCrvKey, value)
	case KeyIDKey:
		if v, ok := value.(string); ok {
			h.keyID = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, KeyIDKey, value)
	case KeyOpsKey:
		var acceptor KeyOperationList
		if err := acceptor.Accept(value); err != nil {
			return fmt.Errorf(`invalid value for %s key: %w`, KeyOpsKey, err)
		}
		h.keyOps = &acceptor
		return nil
	case KeyUsageKey:
		switch v := value.(type) {
		case KeyUsageType:
			switch v {
			case ForSignature, ForEncryption:
				tmp := v.String()
				h.keyUsage = &tmp
			default:
				return fmt.Errorf(`invalid key usage type %s`, v)
			}
		case string:
			h.keyUsage = &v
		default:
			return fmt.Errorf(`invalid key usage type %s`, v)
		}
	case OKPXKey:
		if v, ok := value.([]byte); ok {
			h.x = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, OKPXKey, value)
	case X509CertChainKey:
		if v, ok := value.(*cert.Chain); ok {
			h.x509CertChain = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertChainKey, value)
	case X509CertThumbprintKey:
		if v, ok := value.(string); ok {
			h.x509CertThumbprint = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertThumbprintKey, value)
	case X509CertThumbprintS256Key:
		if v, ok := value.(string); ok {
			h.x509CertThumbprintS256 = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertThumbprintS256Key, value)
	case X509URLKey:
		if v, ok := value.(string); ok {
			h.x509URL = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509URLKey, value)
	default:
		if h.privateParams == nil {
			h.privateParams = map[string]interface{}{}
		}
		h.privateParams[name] = value
	}
	return nil
}

func (k *okpPublicKey) Remove(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	switch key {
	case AlgorithmKey:
		k.algorithm = nil
	case OKPCrvKey:
		k.crv = nil
	case KeyIDKey:
		k.keyID = nil
	case KeyOpsKey:
		k.keyOps = nil
	case KeyUsageKey:
		k.keyUsage = nil
	case OKPXKey:
		k.x = nil
	case X509CertChainKey:
		k.x509CertChain = nil
	case X509CertThumbprintKey:
		k.x509CertThumbprint = nil
	case X509CertThumbprintS256Key:
		k.x509CertThumbprintS256 = nil
	case X509URLKey:
		k.x509URL = nil
	default:
		delete(k.privateParams, key)
	}
	return nil
}

func (k *okpPublicKey) Clone() (Key, error) {
	return cloneKey(k)
}

func (k *okpPublicKey) DecodeCtx() json.DecodeCtx {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.dc
}

func (k *okpPublicKey) SetDecodeCtx(dc json.DecodeCtx) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.dc = dc
}

func (h *okpPublicKey) UnmarshalJSON(buf []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.algorithm = nil
	h.crv = nil
	h.keyID = nil
	h.keyOps = nil
	h.keyUsage = nil
	h.x = nil
	h.x509CertChain = nil
	h.x509CertThumbprint = nil
	h.x509CertThumbprintS256 = nil
	h.x509URL = nil
	dec := json.NewDecoder(bytes.NewReader(buf))
LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf(`error reading token: %w`, err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			// Assuming we're doing everything correctly, we should ONLY
			// get either '{' or '}' here.
			if tok == '}' { // End of object
				break LOOP
			} else if tok != '{' {
				return fmt.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string: // Objects can only have string keys
			switch tok {
			case KeyTypeKey:
				val, err := json.ReadNextStringToken(dec)
				if err != nil {
					return fmt.Errorf(`error reading token: %w`, err)
				}
				if val != jwa.OKP.String() {
					return fmt.Errorf(`invalid kty value for RSAPublicKey (%s)`, val)
				}
			case AlgorithmKey:
				var s string
				if err := dec.Decode(&s); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, AlgorithmKey, err)
				}
				alg := jwa.KeyAlgorithmFrom(s)
				h.algorithm = &alg
			case OKPCrvKey:
				var decoded jwa.EllipticCurveAlgorithm
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, OKPCrvKey, err)
				}
				h.crv = &decoded
			case KeyIDKey:
				if err := json.AssignNextStringToken(&h.keyID, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, KeyIDKey, err)
				}
			case KeyOpsKey:
				var decoded KeyOperationList
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, KeyOpsKey, err)
				}
				h.keyOps = &decoded
			case KeyUsageKey:
				if err := json.AssignNextStringToken(&h.keyUsage, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, KeyUsageKey, err)
				}
			case OKPXKey:
				if err := json.AssignNextBytesToken(&h.x, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, OKPXKey, err)
				}
			case X509CertChainKey:
				var decoded cert.Chain
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertChainKey, err)
				}
				h.x509CertChain = &decoded
			case X509CertThumbprintKey:
				if err := json.AssignNextStringToken(&h.x509CertThumbprint, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertThumbprintKey, err)
				}
			case X509CertThumbprintS256Key:
				if err := json.AssignNextStringToken(&h.x509CertThumbprintS256, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertThumbprintS256Key, err)
				}
			case X509URLKey:
				if err := json.AssignNextStringToken(&h.x509URL, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509URLKey, err)
				}
			default:
				if dc := h.dc; dc != nil {
					if localReg := dc.Registry(); localReg != nil {
						decoded, err := localReg.Decode(dec, tok)
						if err == nil {
							h.setNoLock(tok, decoded)
							continue
						}
					}
				}
				decoded, err := registry.Decode(dec, tok)
				if err == nil {
					h.setNoLock(tok, decoded)
					continue
				}
				return fmt.Errorf(`could not decode field %s: %w`, tok, err)
			}
		default:
			return fmt.Errorf(`invalid token %T`, tok)
		}
	}
	if h.crv == nil {
		return fmt.Errorf(`required field crv is missing`)
	}
	if h.x == nil {
		return fmt.Errorf(`required field x is missing`)
	}
	return nil
}

func (h okpPublicKey) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	fields := make([]string, 0, 10)
	for _, pair := range h.makePairs() {
		fields = append(fields, pair.Key.(string))
		data[pair.Key.(string)] = pair.Value
	}

	sort.Strings(fields)
	buf := pool.GetBytesBuffer()
	defer pool.ReleaseBytesBuffer(buf)
	buf.WriteByte('{')
	enc := json.NewEncoder(buf)
	for i, f := range fields {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('"')
		buf.WriteString(f)
		buf.WriteString(`":`)
		v := data[f]
		switch v := v.(type) {
		case []byte:
			buf.WriteRune('"')
			buf.WriteString(base64.EncodeToString(v))
			buf.WriteRune('"')
		default:
			if err := enc.Encode(v); err != nil {
				return nil, fmt.Errorf(`failed to encode value for field %s: %w`, f, err)
			}
			buf.Truncate(buf.Len() - 1)
		}
	}
	buf.WriteByte('}')
	ret := make([]byte, buf.Len())
	copy(ret, buf.Bytes())
	return ret, nil
}

func (h *okpPublicKey) Iterate(ctx context.Context) HeaderIterator {
	pairs := h.makePairs()
	ch := make(chan *HeaderPair, len(pairs))
	go func(ctx context.Context, ch chan *HeaderPair, pairs []*HeaderPair) {
		defer close(ch)
		for _, pair := range pairs {
			select {
			case <-ctx.Done():
				return
			case ch <- pair:
			}
		}
	}(ctx, ch, pairs)
	return mapiter.New(ch)
}

func (h *okpPublicKey) Walk(ctx context.Context, visitor HeaderVisitor) error {
	return iter.WalkMap(ctx, h, visitor)
}

func (h *okpPublicKey) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, h)
}

type OKPPrivateKey interface {
	Key
	FromRaw(interface{}) error
	Crv() jwa.EllipticCurveAlgorithm
	D() []byte
	X() []byte
}

type okpPrivateKey struct {
	algorithm              *jwa.KeyAlgorithm // https://tools.ietf.org/html/rfc7517#section-4.4
	crv                    *jwa.EllipticCurveAlgorithm
	d                      []byte
	keyID                  *string           // https://tools.ietf.org/html/rfc7515#section-4.1.4
	keyOps                 *KeyOperationList // https://tools.ietf.org/html/rfc7517#section-4.3
	keyUsage               *string           // https://tools.ietf.org/html/rfc7517#section-4.2
	x                      []byte
	x509CertChain          *cert.Chain // https://tools.ietf.org/html/rfc7515#section-4.1.6
	x509CertThumbprint     *string     // https://tools.ietf.org/html/rfc7515#section-4.1.7
	x509CertThumbprintS256 *string     // https://tools.ietf.org/html/rfc7515#section-4.1.8
	x509URL                *string     // https://tools.ietf.org/html/rfc7515#section-4.1.5
	privateParams          map[string]interface{}
	mu                     *sync.RWMutex
	dc                     json.DecodeCtx
}

var _ OKPPrivateKey = &okpPrivateKey{}
var _ Key = &okpPrivateKey{}

func newOKPPrivateKey() *okpPrivateKey {
	return &okpPrivateKey{
		mu:            &sync.RWMutex{},
		privateParams: make(map[string]interface{}),
	}
}

func (h okpPrivateKey) KeyType() jwa.KeyType {
	return jwa.OKP
}

func (h okpPrivateKey) IsPrivate() bool {
	return true
}

func (h *okpPrivateKey) Algorithm() jwa.KeyAlgorithm {
	if h.algorithm != nil {
		return *(h.algorithm)
	}
	return nil
}

func (h *okpPrivateKey) Crv() jwa.EllipticCurveAlgorithm {
	if h.crv != nil {
		return *(h.crv)
	}
	return jwa.InvalidEllipticCurve
}

func (h *okpPrivateKey) D() []byte {
	return h.d
}

func (h *okpPrivateKey) KeyID() string {
	if h.keyID != nil {
		return *(h.keyID)
	}
	return ""
}

func (h *okpPrivateKey) KeyOps() KeyOperationList {
	if h.keyOps != nil {
		return *(h.keyOps)
	}
	return nil
}

func (h *okpPrivateKey) KeyUsage() string {
	if h.keyUsage != nil {
		return *(h.keyUsage)
	}
	return ""
}

func (h *okpPrivateKey) X() []byte {
	return h.x
}

func (h *okpPrivateKey) X509CertChain() *cert.Chain {
	return h.x509CertChain
}

func (h *okpPrivateKey) X509CertThumbprint() string {
	if h.x509CertThumbprint != nil {
		return *(h.x509CertThumbprint)
	}
	return ""
}

func (h *okpPrivateKey) X509CertThumbprintS256() string {
	if h.x509CertThumbprintS256 != nil {
		return *(h.x509CertThumbprintS256)
	}
	return ""
}

func (h *okpPrivateKey) X509URL() string {
	if h.x509URL != nil {
		return *(h.x509URL)
	}
	return ""
}

func (h *okpPrivateKey) makePairs() []*HeaderPair {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var pairs []*HeaderPair
	pairs = append(pairs, &HeaderPair{Key: "kty", Value: jwa.OKP})
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.crv != nil {
		pairs = append(pairs, &HeaderPair{Key: OKPCrvKey, Value: *(h.crv)})
	}
	if h.d != nil {
		pairs = append(pairs, &HeaderPair{Key: OKPDKey, Value: h.d})
	}
	if h.keyID != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyIDKey, Value: *(h.keyID)})
	}
	if h.keyOps != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyOpsKey, Value: *(h.keyOps)})
	}
	if h.keyUsage != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyUsageKey, Value: *(h.keyUsage)})
	}
	if h.x != nil {
		pairs = append(pairs, &HeaderPair{Key: OKPXKey, Value: h.x})
	}
	if h.x509CertChain != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertChainKey, Value: h.x509CertChain})
	}
	if h.x509CertThumbprint != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintKey, Value: *(h.x509CertThumbprint)})
	}
	if h.x509CertThumbprintS256 != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintS256Key, Value: *(h.x509CertThumbprintS256)})
	}
	if h.x509URL != nil {
		pairs = append(pairs, &HeaderPair{Key: X509URLKey, Value: *(h.x509URL)})
	}
	for k, v := range h.privateParams {
		pairs = append(pairs, &HeaderPair{Key: k, Value: v})
	}
	return pairs
}

func (h *okpPrivateKey) PrivateParams() map[string]interface{} {
	return h.privateParams
}

func (h *okpPrivateKey) Get(name string) (interface{}, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch name {
	case KeyTypeKey:
		return h.KeyType(), true
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case OKPCrvKey:
		if h.crv == nil {
			return nil, false
		}
		return *(h.crv), true
	case OKPDKey:
		if h.d == nil {
			return nil, false
		}
		return h.d, true
	case KeyIDKey:
		if h.keyID == nil {
			return nil, false
		}
		return *(h.keyID), true
	case KeyOpsKey:
		if h.keyOps == nil {
			return nil, false
		}
		return *(h.keyOps), true
	case KeyUsageKey:
		if h.keyUsage == nil {
			return nil, false
		}
		return *(h.keyUsage), true
	case OKPXKey:
		if h.x == nil {
			return nil, false
		}
		return h.x, true
	case X509CertChainKey:
		if h.x509CertChain == nil {
			return nil, false
		}
		return h.x509CertChain, true
	case X509CertThumbprintKey:
		if h.x509CertThumbprint == nil {
			return nil, false
		}
		return *(h.x509CertThumbprint), true
	case X509CertThumbprintS256Key:
		if h.x509CertThumbprintS256 == nil {
			return nil, false
		}
		return *(h.x509CertThumbprintS256), true
	case X509URLKey:
		if h.x509URL == nil {
			return nil, false
		}
		return *(h.x509URL), true
	default:
		v, ok := h.privateParams[name]
		return v, ok
	}
}

func (h *okpPrivateKey) Set(name string, value interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.setNoLock(name, value)
}

func (h *okpPrivateKey) setNoLock(name string, value interface{}) error {
	switch name {
	case "kty":
		return nil
	case AlgorithmKey:
		switch v := value.(type) {
		case string, jwa.SigningAlgorithm, jwa.ContentEncryptionAlgorithm:
			var tmp = jwa.KeyAlgorithmFrom(v)
			h.algorithm = &tmp
		case fmt.Stringer:
			s := v.String()
			var tmp = jwa.KeyAlgorithmFrom(s)
			h.algorithm = &tmp
		default:
			return fmt.Errorf(`invalid type for %s key: %T`, AlgorithmKey, value)
		}
		return nil
	case OKPCrvKey:
		if v, ok := value.(jwa.EllipticCurveAlgorithm); ok {
			h.crv = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, OKPCrvKey, value)
	case OKPDKey:
		if v, ok := value.([]byte); ok {
			h.d = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, OKPDKey, value)
	case KeyIDKey:
		if v, ok := value.(string); ok {
			h.keyID = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, KeyIDKey, value)
	case KeyOpsKey:
		var acceptor KeyOperationList
		if err := acceptor.Accept(value); err != nil {
			return fmt.Errorf(`invalid value for %s key: %w`, KeyOpsKey, err)
		}
		h.keyOps = &acceptor
		return nil
	case KeyUsageKey:
		switch v := value.(type) {
		case KeyUsageType:
			switch v {
			case ForSignature, ForEncryption:
				tmp := v.String()
				h.keyUsage = &tmp
			default:
				return fmt.Errorf(`invalid key usage type %s`, v)
			}
		case string:
			h.keyUsage = &v
		default:
			return fmt.Errorf(`invalid key usage type %s`, v)
		}
	case OKPXKey:
		if v, ok := value.([]byte); ok {
			h.x = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, OKPXKey, value)
	case X509CertChainKey:
		if v, ok := value.(*cert.Chain); ok {
			h.x509CertChain = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertChainKey, value)
	case X509CertThumbprintKey:
		if v, ok := value.(string); ok {
			h.x509CertThumbprint = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertThumbprintKey, value)
	case X509CertThumbprintS256Key:
		if v, ok := value.(string); ok {
			h.x509CertThumbprintS256 = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertThumbprintS256Key, value)
	case X509URLKey:
		if v, ok := value.(string); ok {
			h.x509URL = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509URLKey, value)
	default:
		if h.privateParams == nil {
			h.privateParams = map[string]interface{}{}
		}
		h.privateParams[name] = value
	}
	return nil
}

func (k *okpPrivateKey) Remove(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	switch key {
	case AlgorithmKey:
		k.algorithm = nil
	case OKPCrvKey:
		k.crv = nil
	case OKPDKey:
		k.d = nil
	case KeyIDKey:
		k.keyID = nil
	case KeyOpsKey:
		k.keyOps = nil
	case KeyUsageKey:
		k.keyUsage = nil
	case OKPXKey:
		k.x = nil
	case X509CertChainKey:
		k.x509CertChain = nil
	case X509CertThumbprintKey:
		k.x509CertThumbprint = nil
	case X509CertThumbprintS256Key:
		k.x509CertThumbprintS256 = nil
	case X509URLKey:
		k.x509URL = nil
	default:
		delete(k.privateParams, key)
	}
	return nil
}

func (k *okpPrivateKey) Clone() (Key, error) {
	return cloneKey(k)
}

func (k *okpPrivateKey) DecodeCtx() json.DecodeCtx {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.dc
}

func (k *okpPrivateKey) SetDecodeCtx(dc json.DecodeCtx) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.dc = dc
}

func (h *okpPrivateKey) UnmarshalJSON(buf []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.algorithm = nil
	h.crv = nil
	h.d = nil
	h.keyID = nil
	h.keyOps = nil
	h.keyUsage = nil
	h.x = nil
	h.x509CertChain = nil
	h.x509CertThumbprint = nil
	h.x509CertThumbprintS256 = nil
	h.x509URL = nil
	dec := json.NewDecoder(bytes.NewReader(buf))
LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf(`error reading token: %w`, err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			// Assuming we're doing everything correctly, we should ONLY
			// get either '{' or '}' here.
			if tok == '}' { // End of object
				break LOOP
			} else if tok != '{' {
				return fmt.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string: // Objects can only have string keys
			switch tok {
			case KeyTypeKey:
				val, err := json.ReadNextStringToken(dec)
				if err != nil {
					return fmt.Errorf(`error reading token: %w`, err)
				}
				if val != jwa.OKP.String() {
					return fmt.Errorf(`invalid kty value for RSAPublicKey (%s)`, val)
				}
			case AlgorithmKey:
				var s string
				if err := dec.Decode(&s); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, AlgorithmKey, err)
				}
				alg := jwa.KeyAlgorithmFrom(s)
				h.algorithm = &alg
			case OKPCrvKey:
				var decoded jwa.EllipticCurveAlgorithm
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, OKPCrvKey, err)
				}
				h.crv = &decoded
			case OKPDKey:
				if err := json.AssignNextBytesToken(&h.d, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, OKPDKey, err)
				}
			case KeyIDKey:
				if err := json.AssignNextStringToken(&h.keyID, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, KeyIDKey, err)
				}
			case KeyOpsKey:
				var decoded KeyOperationList
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, KeyOpsKey, err)
				}
				h.keyOps = &decoded
			case KeyUsageKey:
				if err := json.AssignNextStringToken(&h.keyUsage, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, KeyUsageKey, err)
				}
			case OKPXKey:
				if err := json.AssignNextBytesToken(&h.x, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, OKPXKey, err)
				}
			case X509CertChainKey:
				var decoded cert.Chain
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertChainKey, err)
				}
				h.x509CertChain = &decoded
			case X509CertThumbprintKey:
				if err := json.AssignNextStringToken(&h.x509CertThumbprint, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertThumbprintKey, err)
				}
			case X509CertThumbprintS256Key:
				if err := json.AssignNextStringToken(&h.x509CertThumbprintS256, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertThumbprintS256Key, err)
				}
			case X509URLKey:
				if err := json.AssignNextStringToken(&h.x509URL, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509URLKey, err)
				}
			default:
				if dc := h.dc; dc != nil {
					if localReg := dc.Registry(); localReg != nil {
						decoded, err := localReg.Decode(dec, tok)
						if err == nil {
							h.setNoLock(tok, decoded)
							continue
						}
					}
				}
				decoded, err := registry.Decode(dec, tok)
				if err == nil {
					h.setNoLock(tok, decoded)
					continue
				}
				return fmt.Errorf(`could not decode field %s: %w`, tok, err)
			}
		default:
			return fmt.Errorf(`invalid token %T`, tok)
		}
	}
	if h.crv == nil {
		return fmt.Errorf(`required field crv is missing`)
	}
	if h.d == nil {
		return fmt.Errorf(`required field d is missing`)
	}
	if h.x == nil {
		return fmt.Errorf(`required field x is missing`)
	}
	return nil
}

func (h okpPrivateKey) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	fields := make([]string, 0, 11)
	for _, pair := range h.makePairs() {
		fields = append(fields, pair.Key.(string))
		data[pair.Key.(string)] = pair.Value
	}

	sort.Strings(fields)
	buf := pool.GetBytesBuffer()
	defer pool.ReleaseBytesBuffer(buf)
	buf.WriteByte('{')
	enc := json.NewEncoder(buf)
	for i, f := range fields {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('"')
		buf.WriteString(f)
		buf.WriteString(`":`)
		v := data[f]
		switch v := v.(type) {
		case []byte:
			buf.WriteRune('"')
			buf.WriteString(base64.EncodeToString(v))
			buf.WriteRune('"')
		default:
			if err := enc.Encode(v); err != nil {
				return nil, fmt.Errorf(`failed to encode value for field %s: %w`, f, err)
			}
			buf.Truncate(buf.Len() - 1)
		}
	}
	buf.WriteByte('}')
	ret := make([]byte, buf.Len())
	copy(ret, buf.Bytes())
	return ret, nil
}

func (h *okpPrivateKey) Iterate(ctx context.Context) HeaderIterator {
	pairs := h.makePairs()
	ch := make(chan *HeaderPair, len(pairs))
	go func(ctx context.Context, ch chan *HeaderPair, pairs []*HeaderPair) {
		defer close(ch)
		for _, pair := range pairs {
			select {
			case <-ctx.Done():
				return
			case ch <- pair:
			}
		}
	}(ctx, ch, pairs)
	return mapiter.New(ch)
}

func (h *okpPrivateKey) Walk(ctx context.Context, visitor HeaderVisitor) error {
	return iter.WalkMap(ctx, h, visitor)
}

func (h *okpPrivateKey) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, h)
}

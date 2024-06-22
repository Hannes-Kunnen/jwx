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
	SymmetricOctetsKey = "k"
)

type SymmetricKey interface {
	Key
	FromRaw([]byte) error
	Octets() []byte
}

type symmetricKey struct {
	algorithm              *jwa.KeyAlgorithm // https://tools.ietf.org/html/rfc7517#section-4.4
	keyID                  *string           // https://tools.ietf.org/html/rfc7515#section-4.1.4
	keyOps                 *KeyOperationList // https://tools.ietf.org/html/rfc7517#section-4.3
	keyUsage               *string           // https://tools.ietf.org/html/rfc7517#section-4.2
	octets                 []byte
	x509CertChain          *cert.Chain // https://tools.ietf.org/html/rfc7515#section-4.1.6
	x509CertThumbprint     *string     // https://tools.ietf.org/html/rfc7515#section-4.1.7
	x509CertThumbprintS256 *string     // https://tools.ietf.org/html/rfc7515#section-4.1.8
	x509URL                *string     // https://tools.ietf.org/html/rfc7515#section-4.1.5
	privateParams          map[string]interface{}
	mu                     *sync.RWMutex
	dc                     json.DecodeCtx
}

var _ SymmetricKey = &symmetricKey{}
var _ Key = &symmetricKey{}

func newSymmetricKey() *symmetricKey {
	return &symmetricKey{
		mu:            &sync.RWMutex{},
		privateParams: make(map[string]interface{}),
	}
}

func (h symmetricKey) KeyType() jwa.KeyType {
	return jwa.OctetSeq
}

func (h *symmetricKey) Algorithm() jwa.KeyAlgorithm {
	if h.algorithm != nil {
		return *(h.algorithm)
	}
	return nil
}

func (h *symmetricKey) KeyID() string {
	if h.keyID != nil {
		return *(h.keyID)
	}
	return ""
}

func (h *symmetricKey) KeyOps() KeyOperationList {
	if h.keyOps != nil {
		return *(h.keyOps)
	}
	return nil
}

func (h *symmetricKey) KeyUsage() string {
	if h.keyUsage != nil {
		return *(h.keyUsage)
	}
	return ""
}

func (h *symmetricKey) Octets() []byte {
	return h.octets
}

func (h *symmetricKey) X509CertChain() *cert.Chain {
	return h.x509CertChain
}

func (h *symmetricKey) X509CertThumbprint() string {
	if h.x509CertThumbprint != nil {
		return *(h.x509CertThumbprint)
	}
	return ""
}

func (h *symmetricKey) X509CertThumbprintS256() string {
	if h.x509CertThumbprintS256 != nil {
		return *(h.x509CertThumbprintS256)
	}
	return ""
}

func (h *symmetricKey) X509URL() string {
	if h.x509URL != nil {
		return *(h.x509URL)
	}
	return ""
}

func (h *symmetricKey) makePairs() []*HeaderPair {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var pairs []*HeaderPair
	pairs = append(pairs, &HeaderPair{Key: "kty", Value: jwa.OctetSeq})
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
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
	if h.octets != nil {
		pairs = append(pairs, &HeaderPair{Key: SymmetricOctetsKey, Value: h.octets})
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

func (h *symmetricKey) PrivateParams() map[string]interface{} {
	return h.privateParams
}

func (h *symmetricKey) Get(name string) (interface{}, bool) {
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
	case SymmetricOctetsKey:
		if h.octets == nil {
			return nil, false
		}
		return h.octets, true
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

func (h *symmetricKey) Set(name string, value interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.setNoLock(name, value)
}

func (h *symmetricKey) setNoLock(name string, value interface{}) error {
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
	case SymmetricOctetsKey:
		if v, ok := value.([]byte); ok {
			h.octets = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, SymmetricOctetsKey, value)
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

func (k *symmetricKey) Remove(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	switch key {
	case AlgorithmKey:
		k.algorithm = nil
	case KeyIDKey:
		k.keyID = nil
	case KeyOpsKey:
		k.keyOps = nil
	case KeyUsageKey:
		k.keyUsage = nil
	case SymmetricOctetsKey:
		k.octets = nil
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

func (k *symmetricKey) Clone() (Key, error) {
	return cloneKey(k)
}

func (k *symmetricKey) DecodeCtx() json.DecodeCtx {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.dc
}

func (k *symmetricKey) SetDecodeCtx(dc json.DecodeCtx) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.dc = dc
}

func (h *symmetricKey) UnmarshalJSON(buf []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.algorithm = nil
	h.keyID = nil
	h.keyOps = nil
	h.keyUsage = nil
	h.octets = nil
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
				if val != jwa.OctetSeq.String() {
					return fmt.Errorf(`invalid kty value for RSAPublicKey (%s)`, val)
				}
			case AlgorithmKey:
				var s string
				if err := dec.Decode(&s); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, AlgorithmKey, err)
				}
				alg := jwa.KeyAlgorithmFrom(s)
				h.algorithm = &alg
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
			case SymmetricOctetsKey:
				if err := json.AssignNextBytesToken(&h.octets, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, SymmetricOctetsKey, err)
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
	if h.octets == nil {
		return fmt.Errorf(`required field k is missing`)
	}
	return nil
}

func (h symmetricKey) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	fields := make([]string, 0, 9)
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

func (h *symmetricKey) Iterate(ctx context.Context) HeaderIterator {
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

func (h *symmetricKey) Walk(ctx context.Context, visitor HeaderVisitor) error {
	return iter.WalkMap(ctx, h, visitor)
}

func (h *symmetricKey) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, h)
}

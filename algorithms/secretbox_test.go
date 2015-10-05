package algorithms

import "testing"
import (
	"crypto/rand"
	"encoding/base64"

	"github.com/stretchr/testify/assert"
)

func TestSecretboxSymmetry(t *testing.T) {
	var key [32]byte
	_, err := rand.Read(key[:])
	assert.NoError(t, err)
	box := newSecretBox()

	nonces := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		var message [4096]byte
		rand.Read(message[:])
		ciphertext, err := box.Encrypt(&key, message[:])

		// Test that nonces seem to be unique
		nonce := base64.StdEncoding.EncodeToString(ciphertext[:24])
		_, seen := nonces[nonce]
		assert.False(t, seen)
		nonces[nonce] = struct{}{}

		assert.NoError(t, err)
		plaintext, err := box.Decrypt(&key, ciphertext)
		assert.NoError(t, err)
		assert.Equal(t, message[:], plaintext)
	}
}

package ec2meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLookup(t *testing.T) {
	assert := assert.New(t)

	if !CanLookup() {
		return
	}

	tests := map[string]struct {
		output bool
	}{
		"check lookup": {output: true},
	}

	for _, t := range tests {
		assert.Equal(t.output, CanLookup())
	}
}

func TestHostname(t *testing.T) {
	assert := assert.New(t)
	if !CanLookup() {
		return
	}

	result, err := Hostname()
	assert.NoError(err)
	assert.NotEqual(0, len(result))
}

func TestLocalHostname(t *testing.T) {
	assert := assert.New(t)
	if !CanLookup() {
		return
	}

	result, err := LocalHostname()
	assert.NoError(err)
	assert.NotEqual(0, len(result))
}

func TestLocalIPV4(t *testing.T) {
	assert := assert.New(t)
	if !CanLookup() {
		return
	}

	result, err := LocalIPV4()
	assert.NoError(err)
	assert.NotEqual(0, len(result))
}

func TestMac(t *testing.T) {
	assert := assert.New(t)
	if !CanLookup() {
		return
	}

	result, err := Mac()
	assert.NoError(err)
	assert.NotEqual(0, len(result))
}

func TestPublicHostname(t *testing.T) {
	assert := assert.New(t)
	if !CanLookup() {
		return
	}

	result, err := PublicHostname()
	assert.NoError(err)
	assert.NotEqual(0, len(result))
}

func TestPublicIPV4(t *testing.T) {
	assert := assert.New(t)
	if !CanLookup() {
		return
	}

	result, err := PublicIPV4()
	assert.NoError(err)
	assert.NotEqual(0, len(result))
}

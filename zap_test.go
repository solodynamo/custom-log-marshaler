package main

import (
	"testing"

	"github.com/solodynamo/pii-marshaler/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestZapGenerate(t *testing.T) {
	generate("./fixtures/zapexample.go", &UberZap{})
	assert.True(t, fixtures.ContainsText("\t\tenc.AddString(\"name\", l.Name)", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddReflected(\"user\", l.User)", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddString(\"request_id\", l.RequestID)", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddBool(\"from_cache\", l.FromCache)", "./fixtures/zapexample.go"))
}

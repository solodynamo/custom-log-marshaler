package main

import (
	"testing"

	"github.com/solodynamo/custom-log-marshaler/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestZerologGenerate(t *testing.T) {
	generate("./fixtures/zerologexample.go", &ZeroLog{})
	assert.True(t, fixtures.ContainsText("\t\tenc.Object(\"user\", l.User)", "./fixtures/zerologexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.Str(\"request_id\", l.RequestID)", "./fixtures/zerologexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.Bool(\"from_cache\", l.FromCache)", "./fixtures/zerologexample.go"))
}

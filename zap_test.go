package main

import (
	"testing"

	"github.com/solodynamo/custom-log-marshaler/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestZapGenerate(t *testing.T) {
	generate("./fixtures/zapexample.go", &UberZap{})
	assert.True(t, fixtures.ContainsText("\t\tif l.Name != nil { enc.AddString(\"name\", *l.Name) }", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddObject(\"user\", l.User)", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddBool(\"from_cache\", l.FromCache)", "./fixtures/zapexample.go"))

	assert.True(t, fixtures.ContainsText("\t\tenc.AddString(\"language\", l.Language)", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddString(\"translation\", l.Translation)", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddReflected(\"translations\", l.Translations)", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddReflected(\"metadata\", l.Metadata)", "./fixtures/zapexample.go"))
	assert.True(t, fixtures.ContainsText("\t\tenc.AddReflected(\"no\", l.No)", "./fixtures/zapexample.go"))
}

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Pair struct {
	a, b interface{}
}

var spaceToUnderscore []Pair
var underscoreToSpace []Pair

func init() {
	spaceToUnderscore = []Pair{
		Pair{"thishasnospaces", "thishasnospaces"},
		Pair{"this has spaces", "this_has_spaces"},
		Pair{"this_has_underscores", "this_has_underscores"},
		Pair{"5     spaces", "5_____spaces"},
	}

	underscoreToSpace = []Pair{
		Pair{"thishasnous", "thishasnous"},
		Pair{"this_has_us", "this has us"},
		Pair{"this has spaces", "this has spaces"},
		Pair{"5_____us", "5     us"},
	}
}

func TestReplaceSpaces(t *testing.T) {
	for i := range spaceToUnderscore {
		in := spaceToUnderscore[i].a.(string)
		expected := spaceToUnderscore[i].b.(string)
		result := ReplaceSpaces(in)
		assert.Equal(t, expected, result)
	}
}

func TestReplaceUnderscores(t *testing.T) {
	for i := range underscoreToSpace {
		in := underscoreToSpace[i].a.(string)
		expected := underscoreToSpace[i].b.(string)
		result := ReplaceUnderscores(in)
		assert.Equal(t, expected, result)
	}
}

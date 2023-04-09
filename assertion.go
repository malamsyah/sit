package sit

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func (s *SIT) AssertStringEqualInJSONPath(path, json, expected string) {
	value := gjson.Get(json, path)
	assert.Equal(s.T, expected, value.String())
}

func (s *SIT) AssertIntEqualInJSONPath(path, json string, expected int64) {
	value := gjson.Get(json, path)
	assert.Equal(s.T, expected, value.Int())
}

func (s *SIT) AssertFloatEqualInJSONPath(path, json string, expected float64) {
	value := gjson.Get(json, path)
	assert.Equal(s.T, expected, value.Float())
}

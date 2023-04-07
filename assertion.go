package sit

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func (s *SIT) AssertStringEqualInJSONPath(jsonPath, json, expected string) {
	value := gjson.Get(json, jsonPath)
	assert.Equal(s.T, expected, value.String())
}

func (s *SIT) AssertIntEqualInJSONPath(jsonPath, json string, expected int) {
	value := gjson.Get(json, jsonPath)
	assert.Equal(s.T, expected, value.Int())
}

func (s *SIT) AssertNumEqualInJSONPath(jsonPath, json string, expected float64) {
	value := gjson.Get(json, jsonPath)
	assert.Equal(s.T, expected, value.Float())
}

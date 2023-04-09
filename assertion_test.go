package sit

import (
	"bytes"
	"io"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (suite *AssertionTestSuite) TestAssertStringEqualInJSONPath() {
	setupStub(suite.sit)

	data := post("http://localhost:8080/test?action=create")

	assert.NotNil(suite.T(), data)
	suite.sit.AssertStringEqualInJSONPath("data.name", data, "john")
}

func (suite *AssertionTestSuite) TestAssertIntEqualInJSONPath() {
	setupStub(suite.sit)

	data := post("http://localhost:8080/test?action=create")

	assert.NotNil(suite.T(), data)
	suite.sit.AssertIntEqualInJSONPath("data.age", data, int64(20))
}

func (suite *AssertionTestSuite) TestAssertFloatEqualInJSONPath() {
	setupStub(suite.sit)

	data := post("http://localhost:8080/test?action=create")

	assert.NotNil(suite.T(), data)
	suite.sit.AssertFloatEqualInJSONPath("data.height", data, float64(178.5))
}

func setupStub(sit *SIT) {
	sit.StubFor(http.MethodPost, "/test", NewStubOption().
		WithEqualQueryParam("action", "create").
		WithEqualHeader("Content-Type", "application/json").
		WithEquaJSONlBody(`{"name":"john"}`).
		WillReturnJSON(map[string]interface{}{
			"data": map[string]interface{}{
				"name":   "john",
				"age":    20,
				"height": 178.5,
			},
		}, nil, http.StatusOK))
}

func post(url string) string {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(`{"name":"john"}`)))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

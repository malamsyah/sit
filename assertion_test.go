package sit

import (
	"io/ioutil"
	"net/http"
)

func (suite *AssertionTestSuite) TestAssertStringEqualInJSONPath() {
	stub := suite.sit.Post("/test")

	stub.WillReturnJSON(map[string]interface{}{
		"data": map[string]string{
			"foo": "bar",
		},
	}, nil, 200)
	stub.AtPriority(1)
	suite.sit.StubFor(stub)

	data := post("http://localhost:8080/test")
	suite.sit.AssertStringEqualInJSONPath("data.foo", data, "bar")
}

func post(url string) string {
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

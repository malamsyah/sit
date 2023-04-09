package sit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/walkerus/go-wiremock"
)

type SIT struct {
	wiremock *wiremock.Client
	T        *testing.T
}

func NewSIT(t *testing.T, url string) *SIT {
	return &SIT{
		wiremock: wiremock.NewClient(url),
		T:        t,
	}
}

func (s *SIT) StubFor(method, path string, option *StubOption) {
	req := wiremock.NewStubRule(method, wiremock.URLPathEqualTo(path))

	if option.requestJSONBody != "" {
		req.WithBodyPattern(wiremock.EqualToJson(option.requestJSONBody))
	}

	if option.requestQueryParams != nil {
		for key, value := range option.requestQueryParams {
			req.WithQueryParam(key, wiremock.EqualTo(value))
		}
	}

	if option.requestHeader != nil {
		for key, value := range option.requestHeader {
			req.WithHeader(key, wiremock.EqualTo(value))
		}
	}

	req.WillReturnJSON(option.responseValue, option.responseHeaders, option.responseStatus)

	err := s.wiremock.StubFor(req)
	assert.NoError(s.T, err)
}

func (s *SIT) Reset() error {
	return s.wiremock.Reset()
}

package sit

import "github.com/walkerus/go-wiremock"

type StubRule struct {
	*wiremock.StubRule
}

func (s *StubRule) WithEqualQueryParam(key, value string) *StubRule {
	s.StubRule = s.WithQueryParam(key, wiremock.EqualTo(value))
	return s
}

func (s *StubRule) WithEqualHeader(key, value string) *StubRule {
	s.StubRule = s.WithHeader(key, wiremock.EqualTo(value))
	return s
}

func (s *StubRule) WithEquaJSONlBody(value string) *StubRule {
	s.StubRule = s.WithBodyPattern(wiremock.EqualToJson(value))
	return s
}

func (s *StubRule) WillReturnJSON(value interface{}, headers map[string]string, status int64) *StubRule {
	s.StubRule = s.StubRule.WillReturnJSON(value, headers, status)
	return s
}

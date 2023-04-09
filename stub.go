package sit

type StubOption struct {
	//Request
	requestJSONBody    string
	requestQueryParams map[string]string
	requestHeader      map[string]string

	//Response
	responseStatus  int64
	responseHeaders map[string]string
	responseValue   interface{}
}

func NewStubOption() *StubOption {
	return &StubOption{
		requestQueryParams: make(map[string]string),
		requestHeader:      make(map[string]string),
	}
}

func (s *StubOption) WithEqualQueryParam(key, value string) *StubOption {
	s.requestQueryParams[key] = value
	return s
}

func (s *StubOption) WithEqualHeader(key, value string) *StubOption {
	s.requestHeader[key] = value
	return s
}

func (s *StubOption) WithEquaJSONlBody(value string) *StubOption {
	s.requestJSONBody = value
	return s
}

func (s *StubOption) WillReturnJSON(value interface{}, headers map[string]string, status int64) *StubOption {
	s.responseHeaders = headers
	s.responseStatus = status
	s.responseValue = value

	return s
}

package sit

import (
	"testing"

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

func (s *SIT) StubFor(stub StubRule) {
	s.wiremock.StubFor(stub.StubRule)
}

func (s *SIT) Post(path string) StubRule {
	return StubRule{
		wiremock.Post(wiremock.URLPathEqualTo(path)),
	}
}

func (s *SIT) Get(path string) StubRule {
	return StubRule{
		wiremock.Get(wiremock.URLPathEqualTo(path)),
	}
}

func (s *SIT) Put(path string) StubRule {
	return StubRule{
		wiremock.Put(wiremock.URLPathEqualTo(path)),
	}
}

func (s *SIT) Delete(path string) StubRule {
	return StubRule{
		wiremock.Delete(wiremock.URLPathEqualTo(path)),
	}
}

func (s *SIT) Patch(path string) StubRule {
	return StubRule{
		wiremock.Patch(wiremock.URLPathEqualTo(path)),
	}
}

func (s *SIT) Reset() error {
	return s.wiremock.Clear()
}

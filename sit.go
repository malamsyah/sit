package sit

import "github.com/walkerus/go-wiremock"

type SIT struct {
	wiremock *wiremock.Client
}

func NewSIT(url string) *SIT {
	return &SIT{
		wiremock: wiremock.NewClient(url),
	}
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

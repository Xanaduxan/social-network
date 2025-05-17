//go:build http_v2

package test

import (
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/httpclientv2"
)

type ProfileClient = httpclientv2.Client

func BuildProfile(s *Suite) {
	var err error
	s.profile, err = httpclientv2.New("http://localhost:8080/okarpova/my-app/api/v2")
	s.NoError(err)
}

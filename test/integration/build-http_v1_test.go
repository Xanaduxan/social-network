//go:build http_v1

package test

import "gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/httpclient"

type ProfileClient = httpclient.Client

func BuildProfile(s *Suite) {
	s.profile = httpclient.New("localhost:8080")
}

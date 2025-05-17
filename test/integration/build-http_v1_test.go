//go:build http_v1

package test

import "github.com/okarpova/my-app/pkg/httpclient"

type ProfileClient = httpclient.Client

func BuildProfile(s *Suite) {
	s.profile = httpclient.New("localhost:8080")
}

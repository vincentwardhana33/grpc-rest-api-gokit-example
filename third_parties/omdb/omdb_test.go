package omdb

import (
	"testing"

	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/config"
)

var (
	getMovies string = "True"
)

func TestGetMovies(t *testing.T) {
	cases := []struct {
		desc       string
		apikey     string
		searchword string
		pagination string
		exp        string
	}{{
		desc:       "Success Case",
		apikey:     config.OmdbAPIKey,
		searchword: "Spiderman",
		pagination: "1",
		exp:        `True`,
	}, {
		desc:       "Other Searchword",
		apikey:     config.OmdbAPIKey,
		searchword: "Batman",
		pagination: "1",
		exp:        `True`,
	}, {
		desc:       "No Searchword",
		apikey:     config.OmdbAPIKey,
		searchword: "",
		pagination: "1",
		exp:        `False`,
	}, {
		desc:       "No Pagination",
		apikey:     config.OmdbAPIKey,
		searchword: "Spiderman",
		pagination: "",
		exp:        `True`,
	}}

	for _, c := range cases {
		result := GetMovies(c.searchword, c.pagination)

		t.Logf("%s -> Got: %s", c.desc, result.Response)

		if result.Response != c.exp {
			t.Errorf("--FAILED -> Expected: %s", c.exp)
		}
	}
}

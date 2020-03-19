package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/config"
	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/pb"
)

const (
	ApiURL = "http://www.omdbapi.com"
)

func GetMovies(keyword string, pagination string) pb.MovieResponse {
	apikey := config.OmdbAPIKey
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%s", ApiURL, apikey, keyword, pagination)
	result := httpGetPublic(url)

	var movies *pb.MovieResponse = &pb.MovieResponse{}
	err := json.Unmarshal(result, movies)

	if err != nil {
		log.Fatal(err)
	}

	return *movies
}

func httpGetPublic(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

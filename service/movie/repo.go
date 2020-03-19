package movie

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/pb"
	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/third_parties/omdb"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	logger log.Logger
}

func NewRepo(logger log.Logger) Repository {
	return &repo{
		logger: log.With(logger, "repo"),
	}
}

func (repo *repo) SearchMovies(ctx context.Context, keyword string, pagination string) pb.MovieResponse {
	movies := omdb.GetMovies(keyword, pagination)
	return movies
}

package movie

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/pb"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s service) SearchMovies(ctx context.Context, keyword string, pagination string) pb.MovieResponse {
	logger := log.With(s.logger, "method", "SearchMovies")

	movies := s.repostory.SearchMovies(ctx, keyword, pagination)

	logger.Log("Search Movies", keyword, pagination)

	return movies
}

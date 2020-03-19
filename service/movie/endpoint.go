package movie

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/pb"
)

type Endpoints struct {
	SearchMovies endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		SearchMovies: makeSearchMoviesEndpoint(s),
	}
}

func makeSearchMoviesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.MovieRequest)
		movies := s.SearchMovies(ctx, req.Keyword, req.Pagination)

		return movies, nil
	}
}

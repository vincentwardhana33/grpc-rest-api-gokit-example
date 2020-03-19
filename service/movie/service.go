package movie

import (
	"context"

	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/pb"
)

type Service interface {
	SearchMovies(ctx context.Context, keyword string, pagination string) pb.MovieResponse
}

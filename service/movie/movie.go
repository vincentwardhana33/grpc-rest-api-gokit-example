package movie

import (
	"context"

	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/pb"
)

type Repository interface {
	SearchMovies(ctx context.Context, keyword string, pagination string) pb.MovieResponse
}

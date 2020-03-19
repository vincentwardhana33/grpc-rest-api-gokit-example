package movie

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vincentwardhana33/grpc-rest-api-gokit-example/pb"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeMovieRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	// vars := mux.Vars(r)

	// searchword := r.URL.Query().Get("searchword")
	// pagination := r.URL.Query().Get("pagination")

	req := pb.MovieRequest{
		Keyword:    r.URL.Query().Get("searchword"),
		Pagination: r.URL.Query().Get("pagination"),
	}
	return req, nil
}

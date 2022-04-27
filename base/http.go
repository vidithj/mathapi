package base

import (
	"context"
	"encoding/json"
	"mathapi/model"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(s Service, version string, basePath string) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	baseRoute := "/" + basePath + "/" + version

	r.Methods(http.MethodPost).Path(baseRoute + "/min").Handler(httptransport.NewServer(
		e.Min,
		decodeRequest,
		EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	))

	r.Methods(http.MethodPost).Path(baseRoute + "/max").Handler(httptransport.NewServer(
		e.Max,
		decodeRequest,
		EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	))
	r.Methods(http.MethodPost).Path(baseRoute + "/avg").Handler(httptransport.NewServer(
		e.Avg,
		decodeRequest,
		EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	))
	r.Methods(http.MethodPost).Path(baseRoute + "/median").Handler(httptransport.NewServer(
		e.Median,
		decodeRequest,
		EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	))
	r.Methods(http.MethodPost).Path(baseRoute + "/percentile").Handler(httptransport.NewServer(
		e.Percentile,
		decodeRequest,
		EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	))
	return r
}
func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func decodeRequest(ctx context.Context, req *http.Request) (r interface{}, err error) {
	var request = model.MathRequest{}
	if err = json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

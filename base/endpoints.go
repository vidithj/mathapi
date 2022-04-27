package base

import (
	"context"
	"errors"

	"mathapi/model"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Min        endpoint.Endpoint
	Max        endpoint.Endpoint
	Avg        endpoint.Endpoint
	Median     endpoint.Endpoint
	Percentile endpoint.Endpoint
}

//MakeServerEndpoints ...
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		Min:        makeMin(s),
		Max:        makeMax(s),
		Avg:        makeAvg(s),
		Median:     makeMedian(s),
		Percentile: makePercentile(s),
	}
}

func makeMin(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.MathRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}
		return s.CalcMin(ctx, req)
	}
}

func makeAvg(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.MathRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}
		return s.CalcAvg(ctx, req)
	}
}

func makeMedian(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.MathRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}
		return s.CalcMedian(ctx, req)
	}
}

func makePercentile(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.MathRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}
		return s.CalcPercentile(ctx, req)
	}
}

func makeMax(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.MathRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}
		return s.CalcMax(ctx, req)
	}
}

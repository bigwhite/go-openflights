package openflights

import (
	"time"

	"go.pedge.io/google-protobuf"
	"go.pedge.io/proto/rpclog"
	"golang.org/x/net/context"
)

type logAPIServer struct {
	protorpclog.Logger
	delegate APIServer
}

func newLogAPIServer(delegate APIServer) *logAPIServer {
	return &logAPIServer{protorpclog.NewLogger("openflights.API"), delegate}
}

func (a *logAPIServer) GetAllAirports(ctx context.Context, request *google_protobuf.Empty) (response *Airports, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetAllAirports(ctx, request)
}

func (a *logAPIServer) GetAllAirlines(ctx context.Context, request *google_protobuf.Empty) (response *Airlines, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetAllAirlines(ctx, request)
}

func (a *logAPIServer) GetAllRoutes(ctx context.Context, request *google_protobuf.Empty) (response *Routes, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetAllRoutes(ctx, request)
}

func (a *logAPIServer) GetAirportByID(ctx context.Context, request *GetAirportByIDRequest) (response *Airport, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetAirportByID(ctx, request)
}

func (a *logAPIServer) GetAirlineByID(ctx context.Context, request *GetAirlineByIDRequest) (response *Airline, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetAirlineByID(ctx, request)
}

func (a *logAPIServer) GetRouteByID(ctx context.Context, request *GetRouteByIDRequest) (response *Route, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetRouteByID(ctx, request)
}

func (a *logAPIServer) GetDistanceByID(ctx context.Context, request *GetDistanceByIDRequest) (response *google_protobuf.UInt32Value, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetDistanceByID(ctx, request)
}

func (a *logAPIServer) GetAirportByCode(ctx context.Context, request *GetAirportByCodeRequest) (response *Airport, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetAirportByCode(ctx, request)
}

func (a *logAPIServer) GetAirlineByCode(ctx context.Context, request *GetAirlineByCodeRequest) (response *Airline, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetAirlineByCode(ctx, request)
}

func (a *logAPIServer) GetRoutesByCode(ctx context.Context, request *GetRoutesByCodeRequest) (response *Routes, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetRoutesByCode(ctx, request)
}

func (a *logAPIServer) GetDistanceByCode(ctx context.Context, request *GetDistanceByCodeRequest) (response *google_protobuf.UInt32Value, err error) {
	defer func(start time.Time) { a.Log(request, response, err, time.Since(start)) }(time.Now())
	return a.delegate.GetDistanceByCode(ctx, request)
}

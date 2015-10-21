package openflights

import (
	"go.pedge.io/google-protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type localAPIClient struct {
	apiServer APIServer
}

func newLocalAPIClient(apiServer APIServer) *localAPIClient {
	return &localAPIClient{apiServer}
}

func (a *localAPIClient) GetAllAirports(ctx context.Context, request *google_protobuf.Empty, _ ...grpc.CallOption) (*Airports, error) {
	return a.apiServer.GetAllAirports(ctx, request)
}

func (a *localAPIClient) GetAllAirlines(ctx context.Context, request *google_protobuf.Empty, _ ...grpc.CallOption) (*Airlines, error) {
	return a.apiServer.GetAllAirlines(ctx, request)
}

func (a *localAPIClient) GetAllRoutes(ctx context.Context, request *google_protobuf.Empty, _ ...grpc.CallOption) (*Routes, error) {
	return a.apiServer.GetAllRoutes(ctx, request)
}

func (a *localAPIClient) GetAirportByID(ctx context.Context, request *GetAirportByIDRequest, _ ...grpc.CallOption) (*Airport, error) {
	return a.apiServer.GetAirportByID(ctx, request)
}

func (a *localAPIClient) GetAirlineByID(ctx context.Context, request *GetAirlineByIDRequest, _ ...grpc.CallOption) (*Airline, error) {
	return a.apiServer.GetAirlineByID(ctx, request)
}

func (a *localAPIClient) GetRouteByID(ctx context.Context, request *GetRouteByIDRequest, _ ...grpc.CallOption) (*Route, error) {
	return a.apiServer.GetRouteByID(ctx, request)
}

func (a *localAPIClient) GetDistanceByID(ctx context.Context, request *GetDistanceByIDRequest, _ ...grpc.CallOption) (*google_protobuf.UInt32Value, error) {
	return a.apiServer.GetDistanceByID(ctx, request)
}

func (a *localAPIClient) GetAirportByCode(ctx context.Context, request *GetAirportByCodeRequest, _ ...grpc.CallOption) (*Airport, error) {
	return a.apiServer.GetAirportByCode(ctx, request)
}

func (a *localAPIClient) GetAirlineByCode(ctx context.Context, request *GetAirlineByCodeRequest, _ ...grpc.CallOption) (*Airline, error) {
	return a.apiServer.GetAirlineByCode(ctx, request)
}

func (a *localAPIClient) GetRoutesByCode(ctx context.Context, request *GetRoutesByCodeRequest, _ ...grpc.CallOption) (*Routes, error) {
	return a.apiServer.GetRoutesByCode(ctx, request)
}

func (a *localAPIClient) GetDistanceByCode(ctx context.Context, request *GetDistanceByCodeRequest, _ ...grpc.CallOption) (*google_protobuf.UInt32Value, error) {
	return a.apiServer.GetDistanceByCode(ctx, request)
}

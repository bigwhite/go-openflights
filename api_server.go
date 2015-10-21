package openflights

import (
	"go.pedge.io/google-protobuf"
	"golang.org/x/net/context"
)

var (
	emptyInstance = &google_protobuf.Empty{}
)

type apiServer struct {
	client Client
}

func newAPIServer(client Client) *apiServer {
	return &apiServer{client}
}

func (a *apiServer) GetAllAirports(_ context.Context, request *google_protobuf.Empty) (response *Airports, err error) {
	airports, err := a.client.GetAllAirports()
	if err != nil {
		return nil, err
	}
	return &Airports{
		Airport: airports,
	}, nil
}

func (a *apiServer) GetAllAirlines(_ context.Context, request *google_protobuf.Empty) (response *Airlines, err error) {
	airlines, err := a.client.GetAllAirlines()
	if err != nil {
		return nil, err
	}
	return &Airlines{
		Airline: airlines,
	}, nil
}

func (a *apiServer) GetAllRoutes(_ context.Context, request *google_protobuf.Empty) (response *Routes, err error) {
	routes, err := a.client.GetAllRoutes()
	if err != nil {
		return nil, err
	}
	return &Routes{
		Route: routes,
	}, nil
}

func (a *apiServer) GetAirportByID(_ context.Context, request *GetAirportByIDRequest) (response *Airport, err error) {
	return a.client.GetAirportByID(request.Id)
}

func (a *apiServer) GetAirlineByID(_ context.Context, request *GetAirlineByIDRequest) (response *Airline, err error) {
	return a.client.GetAirlineByID(request.Id)
}

func (a *apiServer) GetRouteByID(_ context.Context, request *GetRouteByIDRequest) (response *Route, err error) {
	route, err := a.client.GetRouteByID(request.Id)
	if err != nil {
		return nil, err
	}
	return route, nil
}

func (a *apiServer) GetDistanceByID(_ context.Context, request *GetDistanceByIDRequest) (response *google_protobuf.UInt32Value, err error) {
	distance, err := a.client.GetDistanceByID(request.SourceAirportId, request.DestinationAirportId)
	if err != nil {
		return nil, err
	}
	return &google_protobuf.UInt32Value{
		Value: distance,
	}, nil
}

func (a *apiServer) GetAirportByCode(_ context.Context, request *GetAirportByCodeRequest) (response *Airport, err error) {
	return a.client.GetAirportByCode(request.Code)
}

func (a *apiServer) GetAirlineByCode(_ context.Context, request *GetAirlineByCodeRequest) (response *Airline, err error) {
	return a.client.GetAirlineByCode(request.Code)
}

func (a *apiServer) GetRoutesByCode(_ context.Context, request *GetRoutesByCodeRequest) (response *Routes, err error) {
	routes, err := a.client.GetRoutesByCode(request.AirlineCode, request.SourceAirportCode, request.DestinationAirportCode)
	if err != nil {
		return nil, err
	}
	return &Routes{
		Route: routes,
	}, nil
}

func (a *apiServer) GetDistanceByCode(_ context.Context, request *GetDistanceByCodeRequest) (response *google_protobuf.UInt32Value, err error) {
	distance, err := a.client.GetDistanceByCode(request.SourceAirportCode, request.DestinationAirportCode)
	if err != nil {
		return nil, err
	}
	return &google_protobuf.UInt32Value{
		Value: distance,
	}, nil
}

package flights

import (
	"time"

	"github.com/gogo/protobuf/proto"
	"go.pedge.io/google-protobuf"
	"go.pedge.io/proto/rpclog"
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
	start := time.Now()
	defer func() { a.log("GetAllAirports", request, response, err, start) }()
	airports, err := a.client.GetAllAirports()
	if err != nil {
		return nil, err
	}
	return &Airports{
		Airport: airports,
	}, nil
}

func (a *apiServer) GetAllAirlines(_ context.Context, request *google_protobuf.Empty) (response *Airlines, err error) {
	start := time.Now()
	defer func() { a.log("GetAllAirlines", request, response, err, start) }()
	airlines, err := a.client.GetAllAirlines()
	if err != nil {
		return nil, err
	}
	return &Airlines{
		Airline: airlines,
	}, nil
}

func (a *apiServer) GetAllRoutes(_ context.Context, request *google_protobuf.Empty) (response *Routes, err error) {
	start := time.Now()
	defer func() { a.log("GetAllRoutes", request, response, err, start) }()
	routes, err := a.client.GetAllRoutes()
	if err != nil {
		return nil, err
	}
	return &Routes{
		Route: routes,
	}, nil
}

func (a *apiServer) GetAirportByID(_ context.Context, request *GetAirportByIDRequest) (response *Airport, err error) {
	start := time.Now()
	defer func() { a.log("GetAirportByID", request, response, err, start) }()
	return a.client.GetAirportByID(request.Id)
}

func (a *apiServer) GetAirlineByID(_ context.Context, request *GetAirlineByIDRequest) (response *Airline, err error) {
	start := time.Now()
	defer func() { a.log("GetAirlineByID", request, response, err, start) }()
	return a.client.GetAirlineByID(request.Id)
}

func (a *apiServer) GetRouteByID(_ context.Context, request *GetRouteByIDRequest) (response *Route, err error) {
	start := time.Now()
	defer func() { a.log("GetRouteByID", request, response, err, start) }()
	route, err := a.client.GetRouteByID(request.Id)
	if err != nil {
		return nil, err
	}
	return route, nil
}

func (a *apiServer) GetDistanceByID(_ context.Context, request *GetDistanceByIDRequest) (response *google_protobuf.UInt32Value, err error) {
	start := time.Now()
	defer func() { a.log("GetDistanceByID", request, response, err, start) }()
	distance, err := a.client.GetDistanceByID(request.SourceAirportId, request.DestinationAirportId)
	if err != nil {
		return nil, err
	}
	return &google_protobuf.UInt32Value{
		Value: distance,
	}, nil
}

func (a *apiServer) GetAirportByCode(_ context.Context, request *GetAirportByCodeRequest) (response *Airport, err error) {
	start := time.Now()
	defer func() { a.log("GetAirportByCode", request, response, err, start) }()
	return a.client.GetAirportByCode(request.Code)
}

func (a *apiServer) GetAirlineByCode(_ context.Context, request *GetAirlineByCodeRequest) (response *Airline, err error) {
	start := time.Now()
	defer func() { a.log("GetAirlineByCode", request, response, err, start) }()
	return a.client.GetAirlineByCode(request.Code)
}

func (a *apiServer) GetRoutesByCode(_ context.Context, request *GetRoutesByCodeRequest) (response *Routes, err error) {
	start := time.Now()
	defer func() { a.log("GetRoutesByCode", request, response, err, start) }()
	routes, err := a.client.GetRoutesByCode(request.AirlineCode, request.SourceAirportCode, request.DestinationAirportCode)
	if err != nil {
		return nil, err
	}
	return &Routes{
		Route: routes,
	}, nil
}

func (a *apiServer) GetDistanceByCode(_ context.Context, request *GetDistanceByCodeRequest) (response *google_protobuf.UInt32Value, err error) {
	start := time.Now()
	defer func() { a.log("GetDistanceByCode", request, response, err, start) }()
	distance, err := a.client.GetDistanceByCode(request.SourceAirportCode, request.DestinationAirportCode)
	if err != nil {
		return nil, err
	}
	return &google_protobuf.UInt32Value{
		Value: distance,
	}, nil
}

func (a *apiServer) log(methodName string, request proto.Message, response proto.Message, err error, start time.Time) {
	protorpclog.Info("flights.API", methodName, request, response, err, time.Since(start))
}

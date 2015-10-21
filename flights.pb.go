// Code generated by protoc-gen-go.
// source: flights.proto
// DO NOT EDIT!

package openflights

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import google_api1 "google/api"
import google_protobuf1 "go.pedge.io/google-protobuf"
import google_protobuf2 "go.pedge.io/google-protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DST represents the DST value for a timezone.
type DST int32

const (
	DST_DST_NONE DST = 0
	DST_DST_A    DST = 1
	DST_DST_E    DST = 2
	DST_DST_N    DST = 3
	DST_DST_O    DST = 4
	DST_DST_S    DST = 5
	DST_DST_U    DST = 6
	DST_DST_Z    DST = 7
)

var DST_name = map[int32]string{
	0: "DST_NONE",
	1: "DST_A",
	2: "DST_E",
	3: "DST_N",
	4: "DST_O",
	5: "DST_S",
	6: "DST_U",
	7: "DST_Z",
}
var DST_value = map[string]int32{
	"DST_NONE": 0,
	"DST_A":    1,
	"DST_E":    2,
	"DST_N":    3,
	"DST_O":    4,
	"DST_S":    5,
	"DST_U":    6,
	"DST_Z":    7,
}

func (x DST) String() string {
	return proto.EnumName(DST_name, int32(x))
}

// CSVStore stores information on flights in CSV format.
type CSVStore struct {
	Airports []byte `protobuf:"bytes,1,opt,name=airports,proto3" json:"airports,omitempty"`
	Airlines []byte `protobuf:"bytes,2,opt,name=airlines,proto3" json:"airlines,omitempty"`
	Routes   []byte `protobuf:"bytes,3,opt,name=routes,proto3" json:"routes,omitempty"`
}

func (m *CSVStore) Reset()         { *m = CSVStore{} }
func (m *CSVStore) String() string { return proto.CompactTextString(m) }
func (*CSVStore) ProtoMessage()    {}

// IDStore stores maps from id to object.
type IDStore struct {
	IdToAirport map[string]*Airport `protobuf:"bytes,1,rep,name=id_to_airport" json:"id_to_airport,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IdToAirline map[string]*Airline `protobuf:"bytes,2,rep,name=id_to_airline" json:"id_to_airline,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IdToRoute   map[string]*Route   `protobuf:"bytes,3,rep,name=id_to_route" json:"id_to_route,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *IDStore) Reset()         { *m = IDStore{} }
func (m *IDStore) String() string { return proto.CompactTextString(m) }
func (*IDStore) ProtoMessage()    {}

func (m *IDStore) GetIdToAirport() map[string]*Airport {
	if m != nil {
		return m.IdToAirport
	}
	return nil
}

func (m *IDStore) GetIdToAirline() map[string]*Airline {
	if m != nil {
		return m.IdToAirline
	}
	return nil
}

func (m *IDStore) GetIdToRoute() map[string]*Route {
	if m != nil {
		return m.IdToRoute
	}
	return nil
}

// Airport represents an airport.
type Airport struct {
	Id                    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	City                  string `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	Country               string `protobuf:"bytes,4,opt,name=country" json:"country,omitempty"`
	IataFaa               string `protobuf:"bytes,5,opt,name=iata_faa" json:"iata_faa,omitempty"`
	Icao                  string `protobuf:"bytes,6,opt,name=icao" json:"icao,omitempty"`
	LatitudeMicros        int32  `protobuf:"zigzag32,7,opt,name=latitude_micros" json:"latitude_micros,omitempty"`
	LongitudeMicros       int32  `protobuf:"zigzag32,8,opt,name=longitude_micros" json:"longitude_micros,omitempty"`
	AltitudeFeet          uint32 `protobuf:"varint,9,opt,name=altitude_feet" json:"altitude_feet,omitempty"`
	TimezoneOffsetMinutes int32  `protobuf:"zigzag32,10,opt,name=timezone_offset_minutes" json:"timezone_offset_minutes,omitempty"`
	Dst                   DST    `protobuf:"varint,11,opt,name=dst,enum=openflights.DST" json:"dst,omitempty"`
	Timezone              string `protobuf:"bytes,12,opt,name=timezone" json:"timezone,omitempty"`
}

func (m *Airport) Reset()         { *m = Airport{} }
func (m *Airport) String() string { return proto.CompactTextString(m) }
func (*Airport) ProtoMessage()    {}

// Airline represents an airline.
type Airline struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Alias    string `protobuf:"bytes,3,opt,name=alias" json:"alias,omitempty"`
	Iata     string `protobuf:"bytes,4,opt,name=iata" json:"iata,omitempty"`
	Icao     string `protobuf:"bytes,5,opt,name=icao" json:"icao,omitempty"`
	Callsign string `protobuf:"bytes,6,opt,name=callsign" json:"callsign,omitempty"`
	Country  string `protobuf:"bytes,7,opt,name=country" json:"country,omitempty"`
	Active   bool   `protobuf:"varint,8,opt,name=active" json:"active,omitempty"`
}

func (m *Airline) Reset()         { *m = Airline{} }
func (m *Airline) String() string { return proto.CompactTextString(m) }
func (*Airline) ProtoMessage()    {}

// Route represents a route.
type Route struct {
	Id                 string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Airline            *Airline `protobuf:"bytes,2,opt,name=airline" json:"airline,omitempty"`
	SourceAirport      *Airport `protobuf:"bytes,3,opt,name=source_airport" json:"source_airport,omitempty"`
	DestinationAirport *Airport `protobuf:"bytes,4,opt,name=destination_airport" json:"destination_airport,omitempty"`
	Codeshare          bool     `protobuf:"varint,5,opt,name=codeshare" json:"codeshare,omitempty"`
	Stops              uint32   `protobuf:"varint,6,opt,name=stops" json:"stops,omitempty"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}

func (m *Route) GetAirline() *Airline {
	if m != nil {
		return m.Airline
	}
	return nil
}

func (m *Route) GetSourceAirport() *Airport {
	if m != nil {
		return m.SourceAirport
	}
	return nil
}

func (m *Route) GetDestinationAirport() *Airport {
	if m != nil {
		return m.DestinationAirport
	}
	return nil
}

// Airports is the protobuf plural for Airport.
type Airports struct {
	Airport []*Airport `protobuf:"bytes,1,rep,name=airport" json:"airport,omitempty"`
}

func (m *Airports) Reset()         { *m = Airports{} }
func (m *Airports) String() string { return proto.CompactTextString(m) }
func (*Airports) ProtoMessage()    {}

func (m *Airports) GetAirport() []*Airport {
	if m != nil {
		return m.Airport
	}
	return nil
}

// Airlines is the protobuf plural for Airline.
type Airlines struct {
	Airline []*Airline `protobuf:"bytes,1,rep,name=airline" json:"airline,omitempty"`
}

func (m *Airlines) Reset()         { *m = Airlines{} }
func (m *Airlines) String() string { return proto.CompactTextString(m) }
func (*Airlines) ProtoMessage()    {}

func (m *Airlines) GetAirline() []*Airline {
	if m != nil {
		return m.Airline
	}
	return nil
}

// Routes is the protobuf plural for Airport.
type Routes struct {
	Route []*Route `protobuf:"bytes,1,rep,name=route" json:"route,omitempty"`
}

func (m *Routes) Reset()         { *m = Routes{} }
func (m *Routes) String() string { return proto.CompactTextString(m) }
func (*Routes) ProtoMessage()    {}

func (m *Routes) GetRoute() []*Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type GetAirportByIDRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetAirportByIDRequest) Reset()         { *m = GetAirportByIDRequest{} }
func (m *GetAirportByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetAirportByIDRequest) ProtoMessage()    {}

type GetAirlineByIDRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetAirlineByIDRequest) Reset()         { *m = GetAirlineByIDRequest{} }
func (m *GetAirlineByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetAirlineByIDRequest) ProtoMessage()    {}

type GetRouteByIDRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRouteByIDRequest) Reset()         { *m = GetRouteByIDRequest{} }
func (m *GetRouteByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetRouteByIDRequest) ProtoMessage()    {}

type GetDistanceByIDRequest struct {
	SourceAirportId      string `protobuf:"bytes,1,opt,name=source_airport_id" json:"source_airport_id,omitempty"`
	DestinationAirportId string `protobuf:"bytes,2,opt,name=destination_airport_id" json:"destination_airport_id,omitempty"`
}

func (m *GetDistanceByIDRequest) Reset()         { *m = GetDistanceByIDRequest{} }
func (m *GetDistanceByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetDistanceByIDRequest) ProtoMessage()    {}

type GetAirportByCodeRequest struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *GetAirportByCodeRequest) Reset()         { *m = GetAirportByCodeRequest{} }
func (m *GetAirportByCodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetAirportByCodeRequest) ProtoMessage()    {}

type GetAirlineByCodeRequest struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *GetAirlineByCodeRequest) Reset()         { *m = GetAirlineByCodeRequest{} }
func (m *GetAirlineByCodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetAirlineByCodeRequest) ProtoMessage()    {}

type GetRoutesByCodeRequest struct {
	AirlineCode            string `protobuf:"bytes,1,opt,name=airline_code" json:"airline_code,omitempty"`
	SourceAirportCode      string `protobuf:"bytes,2,opt,name=source_airport_code" json:"source_airport_code,omitempty"`
	DestinationAirportCode string `protobuf:"bytes,3,opt,name=destination_airport_code" json:"destination_airport_code,omitempty"`
}

func (m *GetRoutesByCodeRequest) Reset()         { *m = GetRoutesByCodeRequest{} }
func (m *GetRoutesByCodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetRoutesByCodeRequest) ProtoMessage()    {}

type GetDistanceByCodeRequest struct {
	SourceAirportCode      string `protobuf:"bytes,1,opt,name=source_airport_code" json:"source_airport_code,omitempty"`
	DestinationAirportCode string `protobuf:"bytes,2,opt,name=destination_airport_code" json:"destination_airport_code,omitempty"`
}

func (m *GetDistanceByCodeRequest) Reset()         { *m = GetDistanceByCodeRequest{} }
func (m *GetDistanceByCodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetDistanceByCodeRequest) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("openflights.DST", DST_name, DST_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	GetAllAirports(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Airports, error)
	GetAllAirlines(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Airlines, error)
	GetAllRoutes(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Routes, error)
	GetAirportByID(ctx context.Context, in *GetAirportByIDRequest, opts ...grpc.CallOption) (*Airport, error)
	GetAirlineByID(ctx context.Context, in *GetAirlineByIDRequest, opts ...grpc.CallOption) (*Airline, error)
	GetRouteByID(ctx context.Context, in *GetRouteByIDRequest, opts ...grpc.CallOption) (*Route, error)
	GetDistanceByID(ctx context.Context, in *GetDistanceByIDRequest, opts ...grpc.CallOption) (*google_protobuf2.UInt32Value, error)
	GetAirportByCode(ctx context.Context, in *GetAirportByCodeRequest, opts ...grpc.CallOption) (*Airport, error)
	GetAirlineByCode(ctx context.Context, in *GetAirlineByCodeRequest, opts ...grpc.CallOption) (*Airline, error)
	GetRoutesByCode(ctx context.Context, in *GetRoutesByCodeRequest, opts ...grpc.CallOption) (*Routes, error)
	GetDistanceByCode(ctx context.Context, in *GetDistanceByCodeRequest, opts ...grpc.CallOption) (*google_protobuf2.UInt32Value, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetAllAirports(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Airports, error) {
	out := new(Airports)
	err := grpc.Invoke(ctx, "/openflights.API/GetAllAirports", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAllAirlines(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Airlines, error) {
	out := new(Airlines)
	err := grpc.Invoke(ctx, "/openflights.API/GetAllAirlines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAllRoutes(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Routes, error) {
	out := new(Routes)
	err := grpc.Invoke(ctx, "/openflights.API/GetAllRoutes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAirportByID(ctx context.Context, in *GetAirportByIDRequest, opts ...grpc.CallOption) (*Airport, error) {
	out := new(Airport)
	err := grpc.Invoke(ctx, "/openflights.API/GetAirportByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAirlineByID(ctx context.Context, in *GetAirlineByIDRequest, opts ...grpc.CallOption) (*Airline, error) {
	out := new(Airline)
	err := grpc.Invoke(ctx, "/openflights.API/GetAirlineByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetRouteByID(ctx context.Context, in *GetRouteByIDRequest, opts ...grpc.CallOption) (*Route, error) {
	out := new(Route)
	err := grpc.Invoke(ctx, "/openflights.API/GetRouteByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetDistanceByID(ctx context.Context, in *GetDistanceByIDRequest, opts ...grpc.CallOption) (*google_protobuf2.UInt32Value, error) {
	out := new(google_protobuf2.UInt32Value)
	err := grpc.Invoke(ctx, "/openflights.API/GetDistanceByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAirportByCode(ctx context.Context, in *GetAirportByCodeRequest, opts ...grpc.CallOption) (*Airport, error) {
	out := new(Airport)
	err := grpc.Invoke(ctx, "/openflights.API/GetAirportByCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAirlineByCode(ctx context.Context, in *GetAirlineByCodeRequest, opts ...grpc.CallOption) (*Airline, error) {
	out := new(Airline)
	err := grpc.Invoke(ctx, "/openflights.API/GetAirlineByCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetRoutesByCode(ctx context.Context, in *GetRoutesByCodeRequest, opts ...grpc.CallOption) (*Routes, error) {
	out := new(Routes)
	err := grpc.Invoke(ctx, "/openflights.API/GetRoutesByCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetDistanceByCode(ctx context.Context, in *GetDistanceByCodeRequest, opts ...grpc.CallOption) (*google_protobuf2.UInt32Value, error) {
	out := new(google_protobuf2.UInt32Value)
	err := grpc.Invoke(ctx, "/openflights.API/GetDistanceByCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	GetAllAirports(context.Context, *google_protobuf1.Empty) (*Airports, error)
	GetAllAirlines(context.Context, *google_protobuf1.Empty) (*Airlines, error)
	GetAllRoutes(context.Context, *google_protobuf1.Empty) (*Routes, error)
	GetAirportByID(context.Context, *GetAirportByIDRequest) (*Airport, error)
	GetAirlineByID(context.Context, *GetAirlineByIDRequest) (*Airline, error)
	GetRouteByID(context.Context, *GetRouteByIDRequest) (*Route, error)
	GetDistanceByID(context.Context, *GetDistanceByIDRequest) (*google_protobuf2.UInt32Value, error)
	GetAirportByCode(context.Context, *GetAirportByCodeRequest) (*Airport, error)
	GetAirlineByCode(context.Context, *GetAirlineByCodeRequest) (*Airline, error)
	GetRoutesByCode(context.Context, *GetRoutesByCodeRequest) (*Routes, error)
	GetDistanceByCode(context.Context, *GetDistanceByCodeRequest) (*google_protobuf2.UInt32Value, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetAllAirports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetAllAirports(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetAllAirlines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetAllAirlines(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetAllRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetAllRoutes(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetAirportByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetAirportByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetAirportByID(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetAirlineByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetAirlineByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetAirlineByID(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetRouteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetRouteByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetRouteByID(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetDistanceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetDistanceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetDistanceByID(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetAirportByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetAirportByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetAirportByCode(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetAirlineByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetAirlineByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetAirlineByCode(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetRoutesByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetRoutesByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetRoutesByCode(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetDistanceByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetDistanceByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetDistanceByCode(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openflights.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAirports",
			Handler:    _API_GetAllAirports_Handler,
		},
		{
			MethodName: "GetAllAirlines",
			Handler:    _API_GetAllAirlines_Handler,
		},
		{
			MethodName: "GetAllRoutes",
			Handler:    _API_GetAllRoutes_Handler,
		},
		{
			MethodName: "GetAirportByID",
			Handler:    _API_GetAirportByID_Handler,
		},
		{
			MethodName: "GetAirlineByID",
			Handler:    _API_GetAirlineByID_Handler,
		},
		{
			MethodName: "GetRouteByID",
			Handler:    _API_GetRouteByID_Handler,
		},
		{
			MethodName: "GetDistanceByID",
			Handler:    _API_GetDistanceByID_Handler,
		},
		{
			MethodName: "GetAirportByCode",
			Handler:    _API_GetAirportByCode_Handler,
		},
		{
			MethodName: "GetAirlineByCode",
			Handler:    _API_GetAirlineByCode_Handler,
		},
		{
			MethodName: "GetRoutesByCode",
			Handler:    _API_GetRoutesByCode_Handler,
		},
		{
			MethodName: "GetDistanceByCode",
			Handler:    _API_GetDistanceByCode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

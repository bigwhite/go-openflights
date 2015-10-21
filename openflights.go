/*
Package openflights exposes various flight data from OpenFlights.org.


If you do use this package, I ask you to donate to OpenFlights, the source for all the data
in here as of now, at http://openflights.org/donate. Seriously, if you can afford it, the OpenFlights
team is responsible for putting all this data together and maintaining it, and we owe it to them
to support their work.
*/
package openflights // import "go.pedge.io/openflights"

// GetCSVStore gets the CSVStore from GitHub.
func GetCSVStore() (*CSVStore, error) {
	return getCSVStore()
}

// NewIDStore creates a new IDStore from a CSVStore.
func NewIDStore(csvStore *CSVStore) (*IDStore, error) {
	return newIDStore(csvStore)
}

// CodeStore is a mapping for airline/airport codes (ICAO or IATA/FAA) to object.
//
// Duplicates may be filtered, ie there may be airlines/airports in a CSVStore or IDStore
// that are not present in this structure.
type CodeStore struct {
	CodeToAirport                                                  map[string]*Airport
	CodeToAirline                                                  map[string]*Airline
	AirlineCodeToSourceAirportCodeToDestinationAirportCodeToRoutes map[string]map[string]map[string][]*Route
}

// NewCodeStore creates a new CodeStore from an IDStore.
func NewCodeStore(idStore *IDStore) (*CodeStore, error) {
	return newCodeStore(idStore)
}

// IDClient is the client to interface with flights data by ID.
type IDClient interface {
	GetAllAirports() ([]*Airport, error)
	GetAllAirlines() ([]*Airline, error)
	GetAllRoutes() ([]*Route, error)
	GetAirportByID(id string) (*Airport, error)
	GetAirlineByID(id string) (*Airline, error)
	GetRouteByID(id string) (*Route, error)
	GetDistanceByID(sourceAirportID string, destinationAirportID string) (uint32, error)
}

// CodeClient is the client to interface with flights data by ICAO/IATA/FAA code.
type CodeClient interface {
	GetAirportByCode(code string) (*Airport, error)
	GetAirlineByCode(code string) (*Airline, error)
	GetRoutesByCode(airlineCode string, sourceAirportCode string, destinationAirportCode string) ([]*Route, error)
	GetDistanceByCode(sourceAirportCode string, destinationAirportCode string) (uint32, error)
}

// Client is the client to interface with flights data.
type Client interface {
	IDClient
	CodeClient
}

// NewClient creates a new Client that calles the given APIClient.
func NewClient(apiClient APIClient) Client {
	return newClient(apiClient)
}

// NewLocalAPIClient creates a new APIClient using the given APIServer.
func NewLocalAPIClient(apiServer APIServer) APIClient {
	return newLocalAPIClient(apiServer)
}

// NewAPIServer creates a new APIServer using the given Client.
func NewAPIServer(client Client) APIServer {
	return newLogAPIServer(newAPIServer(client))
}

// NewServerClient creates a new server-side Client.
func NewServerClient(idStore *IDStore) (Client, error) {
	codeStore, err := newCodeStore(idStore)
	if err != nil {
		return nil, err
	}
	return newServerClient(idStore, codeStore)
}

// NewDefaultServerClient creates a new server-side Client from the generated CSVStore.
func NewDefaultServerClient() (Client, error) {
	idStore, err := newIDStore(_GlobalCSVStore)
	if err != nil {
		return nil, err
	}
	return NewServerClient(idStore)
}

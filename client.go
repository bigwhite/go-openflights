package flights

import "golang.org/x/net/context"

type client struct {
	apiClient APIClient
}

func newClient(apiClient APIClient) *client {
	return &client{
		apiClient,
	}
}

func (c *client) GetAllAirports() ([]*Airport, error) {
	airports, err := c.apiClient.GetAllAirports(context.Background(), emptyInstance)
	if err != nil {
		return nil, err
	}
	return airports.Airport, nil
}

func (c *client) GetAllAirlines() ([]*Airline, error) {
	airlines, err := c.apiClient.GetAllAirlines(context.Background(), emptyInstance)
	if err != nil {
		return nil, err
	}
	return airlines.Airline, nil
}

func (c *client) GetAllRoutes() ([]*Route, error) {
	routes, err := c.apiClient.GetAllRoutes(context.Background(), emptyInstance)
	if err != nil {
		return nil, err
	}
	return routes.Route, nil
}

func (c *client) GetAirportByID(id string) (*Airport, error) {
	return c.apiClient.GetAirportByID(
		context.Background(),
		&GetAirportByIDRequest{
			Id: id,
		},
	)
}

func (c *client) GetAirlineByID(id string) (*Airline, error) {
	return c.apiClient.GetAirlineByID(
		context.Background(),
		&GetAirlineByIDRequest{
			Id: id,
		},
	)
}

func (c *client) GetRouteByID(id string) (*Route, error) {
	route, err := c.apiClient.GetRouteByID(
		context.Background(),
		&GetRouteByIDRequest{
			Id: id,
		},
	)
	if err != nil {
		return nil, err
	}
	return route, nil
}

func (c *client) GetDistanceByID(sourceAirportID string, destinationAirportID string) (uint32, error) {
	uint32Value, err := c.apiClient.GetDistanceByID(
		context.Background(),
		&GetDistanceByIDRequest{
			SourceAirportId:      sourceAirportID,
			DestinationAirportId: destinationAirportID,
		},
	)
	if err != nil {
		return 0, err
	}
	return uint32Value.Value, nil
}

func (c *client) GetAirportByCode(code string) (*Airport, error) {
	return c.apiClient.GetAirportByCode(
		context.Background(),
		&GetAirportByCodeRequest{
			Code: code,
		},
	)
}

func (c *client) GetAirlineByCode(code string) (*Airline, error) {
	return c.apiClient.GetAirlineByCode(
		context.Background(),
		&GetAirlineByCodeRequest{
			Code: code,
		},
	)
}

func (c *client) GetRoutesByCode(airlineCode string, sourceAirportCode string, destinationAirportCode string) ([]*Route, error) {
	routes, err := c.apiClient.GetRoutesByCode(
		context.Background(),
		&GetRoutesByCodeRequest{
			AirlineCode:            airlineCode,
			SourceAirportCode:      sourceAirportCode,
			DestinationAirportCode: destinationAirportCode,
		},
	)
	if err != nil {
		return nil, err
	}
	return routes.Route, nil
}

func (c *client) GetDistanceByCode(sourceAirportCode string, destinationAirportCode string) (uint32, error) {
	uint32Value, err := c.apiClient.GetDistanceByCode(
		context.Background(),
		&GetDistanceByCodeRequest{
			SourceAirportCode:      sourceAirportCode,
			DestinationAirportCode: destinationAirportCode,
		},
	)
	if err != nil {
		return 0, err
	}
	return uint32Value.Value, nil
}

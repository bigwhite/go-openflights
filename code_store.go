package openflights

import (
	"fmt"
	"strings"

	"go.pedge.io/protolog"
)

func newCodeStore(idStore *IDStore, options CodeStoreOptions) (*CodeStore, error) {
	codeToAirport, err := getAirportMap(idStore, options)
	if err != nil {
		return nil, err
	}
	codeToAirline, err := getAirlineMap(idStore, options)
	if err != nil {
		return nil, err
	}
	airlineToSourceToDestToRoutes, err := getRoutesMap(idStore, options)
	if err != nil {
		return nil, err
	}
	return &CodeStore{
		codeToAirport,
		codeToAirline,
		airlineToSourceToDestToRoutes,
	}, nil
}

// GetAirportByCode returns the Airport for the given ICAO/IATA/FAA code, or error if it does not exist.
func (c *CodeStore) GetAirportByCode(code string) (*Airport, error) {
	return getAirportByKey(c.CodeToAirport, code)
}

// GetAirlineByCode returns the Airline for the given ICAO/IATA/FAA code, or error if it does not exist.
func (c *CodeStore) GetAirlineByCode(code string) (*Airline, error) {
	return getAirlineByKey(c.CodeToAirline, code)
}

// GetRoutesByCode returns the Routes for the given ICAO/IATA/FAA codes.
func (c *CodeStore) GetRoutesByCode(airline string, source string, dest string) ([]*Route, error) {
	return getRoutesByKeys(c.AirlineCodeToSourceAirportCodeToDestinationAirportCodeToRoutes, airline, source, dest)
}

// GetDistanceByCode returns the distance in miles between the two Airports with the given ICAO/IATA/FAA codes.
func (c *CodeStore) GetDistanceByCode(sourceAirportCode string, destinationAirportCode string) (uint32, error) {
	sourceAirport, err := c.GetAirportByCode(sourceAirportCode)
	if err != nil {
		return 0, err
	}
	destinationAirport, err := c.GetAirportByCode(destinationAirportCode)
	if err != nil {
		return 0, err
	}
	return getDistanceForAirports(sourceAirport, destinationAirport), nil
}

func getAirportMap(idStore *IDStore, options CodeStoreOptions) (map[string]*Airport, error) {
	m := make(map[string]*Airport)
	for _, airport := range idStore.IdToAirport {
		if !options.NoFilterDuplicates {
			include, err := includeAirport(airport)
			if err != nil {
				return nil, err
			}
			if !include {
				continue
			}
		}
		for _, s := range airport.Codes() {
			if _, ok := m[strings.ToLower(s)]; ok {
				err := fmt.Errorf("openflights: duplicate airport key: %s", s)
				if options.NoFilterDuplicates || options.NoErrorOnDuplicates {
					protolog.Warnln(err.Error())
				} else {
					return nil, err
				}
			}
			m[strings.ToLower(s)] = airport
		}
	}
	return m, nil
}

func getAirlineMap(idStore *IDStore, options CodeStoreOptions) (map[string]*Airline, error) {
	airlineCodeToAirlineIDToRouteIDs := getAirlineCodeToAirlineIDToRouteIDs(idStore.IdToRoute)
	m := make(map[string]*Airline)
	for _, airline := range idStore.IdToAirline {
		if !options.NoFilterDuplicates {
			include, err := includeAirline(airline, airlineCodeToAirlineIDToRouteIDs)
			if err != nil {
				return nil, err
			}
			if !include {
				continue
			}
		}
		for _, s := range airline.Codes() {
			if _, ok := m[strings.ToLower(s)]; ok {
				err := fmt.Errorf("openflights: duplicate airline key: %s", s)
				if options.NoFilterDuplicates || options.NoErrorOnDuplicates {
					protolog.Warnln(err.Error())
				} else {
					return nil, err
				}
			}
			m[strings.ToLower(s)] = airline
		}
	}
	return m, nil
}

func getRoutesMap(idStore *IDStore, options CodeStoreOptions) (map[string]map[string]map[string][]*Route, error) {
	m := make(map[string]map[string]map[string][]*Route)
	for _, route := range idStore.IdToRoute {
		if !options.NoFilterDuplicates {
			include, err := includeRoute(route)
			if err != nil {
				return nil, err
			}
			if !include {
				continue
			}
		}
		for _, airline := range route.Airline.Codes() {
			for _, source := range route.SourceAirport.Codes() {
				for _, dest := range route.DestinationAirport.Codes() {
					airline = strings.ToLower(airline)
					source = strings.ToLower(source)
					dest = strings.ToLower(dest)
					if _, ok := m[airline]; !ok {
						m[airline] = make(map[string]map[string][]*Route)
					}
					if _, ok := m[airline][source]; !ok {
						m[airline][source] = make(map[string][]*Route)
					}
					if _, ok := m[airline][source][dest]; !ok {
						m[airline][source][dest] = []*Route{route}
					} else {
						if !containsRoute(m[airline][source][dest], route) {
							m[airline][source][dest] = append(m[airline][source][dest], route)
						}
					}
				}
			}
		}
	}
	return m, nil
}

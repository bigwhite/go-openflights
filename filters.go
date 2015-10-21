package openflights

import (
	"fmt"
	"math"

	"go.pedge.io/protolog"
)

var (
	filterAirportIDToAirportCode = map[string]string{
		"2431": "RPVM",
		"2919": "UATE",
		"3769": "BFT",
		"4330": "OIIE",
		"7708": "ZYA",
		"8420": "EDDB",
	}
)

func includeAirport(airport *Airport) (bool, error) {
	airportCode, ok := filterAirportIDToAirportCode[airport.Id]
	if ok {
		if airport.IataFaa != airportCode && airport.Icao != airportCode {
			return false, fmt.Errorf("openflights: expected airport %v to have code %s", airport, airportCode)
		}
		return false, nil
	}
	return true, nil
}

func includeAirline(airline *Airline, airlineCodeToAirlineIDToRouteIDs map[string]map[string]map[string]bool) (bool, error) {
	if !airline.Active {
		return false, nil
	}
	for _, code := range airline.Codes() {
		airlineIDToRouteIDs, ok := airlineCodeToAirlineIDToRouteIDs[code]
		if !ok {
			continue
		}
		airlineRouteIDs, ok := airlineIDToRouteIDs[airline.Id]
		numAirlineRouteIDs := 0
		if ok {
			numAirlineRouteIDs = len(airlineRouteIDs)
		}
		max := math.MinInt32
		maxAirlineID := ""
		for airlineID, routeIDs := range airlineIDToRouteIDs {
			if len(routeIDs) > max {
				max = len(routeIDs)
				maxAirlineID = airlineID
			}
		}
		if maxAirlineID != airline.Id {
			protolog.Debugf("filtering airline %v because it only had %d routes when max was airline id %s with %d", airline, numAirlineRouteIDs, maxAirlineID, max)
			return false, nil
		}
	}
	return true, nil
}

func includeRoute(route *Route) (bool, error) {
	return true, nil
}

func getAirlineCodeToAirlineIDToRouteIDs(idToRoute map[string]*Route) map[string]map[string]map[string]bool {
	m := make(map[string]map[string]map[string]bool)
	for _, route := range idToRoute {
		airline := route.Airline
		for _, code := range airline.Codes() {
			if _, ok := m[code]; !ok {
				m[code] = make(map[string]map[string]bool)
			}
			if _, ok := m[code][airline.Id]; !ok {
				m[code][airline.Id] = make(map[string]bool)
			}
			m[code][airline.Id][route.Id] = true
		}
	}
	return m
}

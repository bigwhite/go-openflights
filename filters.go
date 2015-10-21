package openflights

import "fmt"

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

func includeAirline(airline *Airline) (bool, error) {
	if !airline.Active {
		return false, nil
	}
	return true, nil
}

func includeRoute(route *Route) (bool, error) {
	return true, nil
}

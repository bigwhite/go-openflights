package openflights

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAirportByCode(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	for _, s := range []string{"VIE", "vie", "Vie"} {
		vie, err := client.GetAirportByCode(s)
		require.NoError(t, err)
		require.Equal(
			t,
			&Airport{
				Id:                    "1613",
				Name:                  "Schwechat",
				City:                  "Vienna",
				Country:               "Austria",
				IataFaa:               "VIE",
				Icao:                  "LOWW",
				LatitudeMicros:        48110278,
				LongitudeMicros:       16569722,
				AltitudeFeet:          600,
				TimezoneOffsetMinutes: 60,
				Dst:      DST_DST_E,
				Timezone: "Europe/Vienna",
			},
			vie,
		)
	}
}

func TestGetRoutesByCode(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	for _, code := range []string{"LH", "DLH", "lh", "dlh"} {
		routes, err := client.GetRoutesByCode(code, "SFO", "MUC")
		require.NoError(t, err)
		dlh, err := client.GetAirlineByCode("DLH")
		require.NoError(t, err)
		sfo, err := client.GetAirportByCode("SFO")
		require.NoError(t, err)
		muc, err := client.GetAirportByCode("MUC")
		require.NoError(t, err)
		expected := []*Route{
			&Route{
				Id:                 "38756",
				Airline:            dlh,
				SourceAirport:      sfo,
				DestinationAirport: muc,
				Codeshare:          false,
				Stops:              0,
			},
		}
		require.Equal(t, len(expected), len(routes))
		for i, route := range routes {
			require.Equal(t, expected[i], route)
		}
	}
}

func TestGetDistanceByCode(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	for _, s := range []string{"SFO", "sfo", "Sfo"} {
		for _, s2 := range []string{"FRA", "fra", "Fra"} {
			distanceMiles, err := client.GetDistanceByCode(s, s2)
			require.NoError(t, err)
			require.Equal(t, uint32(5684), distanceMiles)
		}
	}
}

func getTestClient(t *testing.T) Client {
	idStore, err := NewIDStore(_GlobalCSVStore)
	require.NoError(t, err)
	serverClient, err := NewServerClient(idStore, CodeStoreOptions{NoFilterDuplicates: true})
	require.NoError(t, err)
	// normally in Go code you would directly call the Client,
	// but for testing I want to go through the whole chain.
	return NewClient(NewLocalAPIClient(NewAPIServer(serverClient)))
}

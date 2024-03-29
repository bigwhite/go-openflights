package main

import (
	"fmt"
	"os"
	"text/template"

	"go.pedge.io/openflights"
	"go.pedge.io/openflights/cmd/internal/common"
)

var (
	idTmpl = template.Must(template.New("idTmpl").Parse(`
{{$import := .Import}}
{{$package := .Package}}
{{$packagePrefix := .PackagePrefix}}
{{$private := .Private}}
{{$data := .Data}}
// Code generated by gen-openflights-id-store
// DO NOT EDIT!

package {{$package}}
{{$import}}

var (
	{{if not $private}}// GlobalIDStore is the generated *{{$packagePrefix}}IDStore for all flights information.
	{{end}}{{if $private}_{{end}}IDStore = &{{$packagePrefix}}IDStore{
		IdToAirport: map[string]*{{$packagePrefix}}Airport{
			{{range $id, $airport := $data.IdToAirport}} "{{$id}}": airport{{$id}},
		{{end}}
		},
		IdToAirline: map[string]*{{$packagePrefix}}Airline{
			{{range $id, $airline := $data.IdToAirline}} "{{$id}}": airline{{$id}},
		{{end}}
		},
		IdToRoute: map[string]*{{$packagePrefix}}Route{
			{{range $id, $route := $data.IdToRoute}} "{{$id}}": route{{$id}},
		{{end}}
		},
	}
	{{range $id, $airport := $data.IdToAirport}}airport{{$id}} = &{{$packagePrefix}}Airport{
		Id: "{{$airport.Id}}",
		Name: "{{$airport.Name}}",
		City: "{{$airport.City}}",
		Country: "{{$airport.Country}}",
		IataFaa: "{{$airport.IataFaa}}",
		Icao: "{{$airport.Icao}}",
		LatitudeMicros: {{$airport.LatitudeMicros}},
		LongitudeMicros: {{$airport.LongitudeMicros}},
		AltitudeFeet: {{$airport.AltitudeFeet}},
		TimezoneOffsetMinutes: {{$airport.TimezoneOffsetMinutes}},
		Dst: {{$packagePrefix}}DST_{{$airport.Dst}},
		Timezone: "{{$airport.Timezone}}",
	}
	{{end}}
	{{range $id, $airline := $data.IdToAirline}}airline{{$id}} = &{{$packagePrefix}}Airline{
		Id: "{{$airline.Id}}",
		Name: "{{$airline.Name}}",
		Alias: "{{$airline.Alias}}",
		Iata: "{{$airline.Iata}}",
		Icao: "{{$airline.Icao}}",
		Callsign: "{{$airline.Callsign}}",
		Country: "{{$airline.Country}}",
		Active: {{$airline.Active}},
	}
	{{end}}
	{{range $id, $route := $data.IdToRoute}}route{{$id}} = &{{$packagePrefix}}Route{
		Id: "{{$route.Id}}",
		Airline: airline{{$route.Airline.Id}},
		SourceAirport: airport{{$route.SourceAirport.Id}},
		DestinationAirport: airport{{$route.DestinationAirport.Id}},
		Codeshare: {{$route.Codeshare}},
		Stops: {{$route.Stops}},
	}
	{{end}}
)
`))
)

func main() {
	if err := do(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("usage: %s package path/to/out.go", os.Args[0])
	}
	pkg := os.Args[1]
	outFilePath := os.Args[2]
	csvStore, err := openflights.GetCSVStore()
	if err != nil {
		return err
	}
	idStore, err := openflights.NewIDStore(csvStore)
	if err != nil {
		return nil
	}
	return common.WriteData(pkg, outFilePath, idTmpl, idStore)
}

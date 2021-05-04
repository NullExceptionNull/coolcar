package main

import (
	trippb "coolcar/proto/gen/go"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start:       "ABC",
		End:         "EDF",
		DurationSec: 1000,
		FeeCent:     10,
		StartPos: &trippb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		EndPos: &trippb.Location{
			Latitude:  35,
			Longitude: 115,
		},
		PathLocations: []*trippb.Location{
			{
				Latitude:  31,
				Longitude: 119,
			}, {
				Latitude:  32,
				Longitude: 118,
			},
		},
	}
	fmt.Println(&trip)
	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X \n", b)

	var tripCopy trippb.Trip

	err = proto.Unmarshal(b, &tripCopy)

	if err != nil {
		panic(err)
	}
	fmt.Println("-----", &tripCopy)

	jsonStr, _ := json.Marshal(&trip)

	fmt.Printf("%s \n", jsonStr)

}

package main

import "fmt"

type newSatellite struct {
	name  string
	line1 string
	line2 string
}

type TLEData struct {
	lineNumber             string
	satelliteNumber        string
	classification         string
	intDesLaunchYear       string
	intDesLaunchYearNumber string
	intDesLaunchPiece      string
	epochYear              string
	epochDay               string
	meanMotionFirst        string
	meanMotionSecond       string
	BSTAR                  string
	emphermerisType        string
	elementNumber          string
	checksum               string
}

func main() {

	satelliteISS := newSatellite{
		name:  "ISS",
		line1: "1 25544U 98067A   20358.54791667  .00016717  00000-0  10270-3 0  9002",
		line2: "2 25544  51.6432  21.5264 0002184  90.4728 285.4598 15.49212921247662",
	}

	fmt.Println(satelliteISS)
}

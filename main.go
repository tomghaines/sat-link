package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type newSatellite struct {
	name  string
	line1 string
	line2 string
}

type lineOneTLE struct {
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

type lineTwoTLE struct {
	lineNumber                     string
	satelliteNumber                string
	inclination                    string
	rightAscensionsOfAscendingNode string
	eccentricity                   string
	argumentOfPerigee              string
	meanAnomoly                    string
	meanMotion                     string
	revolutionNumAtEpoch           string
	checksum                       string
}

type KeplarianElements struct {
	inclination       string
	RAAN              string
	eccentricity      string
	argumentOfPerigee string
	meanAnomoly       string
	meanMotion        string
}

type SatellitePosition struct {
	X float64
	Y float64
	Z float64
}

func fmtLineOneTLE(sat newSatellite) lineOneTLE {
	lineOneArr := strings.Split(sat.line1, "")

	tle := lineOneTLE{
		lineNumber:             strings.Join(lineOneArr[0:1], ""),
		satelliteNumber:        strings.Join(lineOneArr[2:7], ""),
		classification:         strings.Join(lineOneArr[7:8], ""),
		intDesLaunchYear:       strings.Join(lineOneArr[9:11], ""),
		intDesLaunchYearNumber: strings.Join(lineOneArr[11:14], ""),
		intDesLaunchPiece:      strings.Join(lineOneArr[14:17], ""),
		epochYear:              strings.Join(lineOneArr[18:20], ""),
		epochDay:               strings.Join(lineOneArr[20:32], ""),
		meanMotionFirst:        strings.Join(lineOneArr[33:43], ""),
		meanMotionSecond:       strings.Join(lineOneArr[44:52], ""),
		BSTAR:                  strings.Join(lineOneArr[53:61], ""),
		emphermerisType:        strings.Join(lineOneArr[62:63], ""),
		elementNumber:          strings.Join(lineOneArr[64:68], ""),
		checksum:               strings.Join(lineOneArr[68:69], ""),
	}

	return tle
}

func fmtLineTwoTLE(sat newSatellite) lineTwoTLE {
	lineTwoArr := strings.Split(sat.line2, "")

	tle := lineTwoTLE{
		lineNumber:                     strings.Join(lineTwoArr[0:1], ""),
		satelliteNumber:                strings.Join(lineTwoArr[2:7], ""),
		inclination:                    strings.Join(lineTwoArr[8:16], ""),
		rightAscensionsOfAscendingNode: strings.Join(lineTwoArr[17:25], ""),
		eccentricity:                   strings.Join(lineTwoArr[26:33], ""),
		argumentOfPerigee:              strings.Join(lineTwoArr[34:42], ""),
		meanAnomoly:                    strings.Join(lineTwoArr[43:51], ""),
		meanMotion:                     strings.Join(lineTwoArr[52:63], ""),
		revolutionNumAtEpoch:           strings.Join(lineTwoArr[63:68], ""),
		checksum:                       strings.Join(lineTwoArr[68:69], ""),
	}
	return tle
}

func getKeplearians(l2 lineTwoTLE) KeplarianElements {
	kepEls := KeplarianElements{
		inclination:       l2.inclination,
		RAAN:              l2.rightAscensionsOfAscendingNode,
		eccentricity:      l2.eccentricity,
		argumentOfPerigee: l2.argumentOfPerigee,
		meanAnomoly:       l2.meanAnomoly,
		meanMotion:        l2.meanMotion,
	}

	return kepEls
}

const mu = 398600.4418 // Earth's gravitational parameter, km^3/s^2

func deriveSemiMajorAxis(meanMotion string) float64 {
	mm, err := strconv.ParseFloat(meanMotion, 64)
	if err != nil {
		fmt.Println("Error parsing mean motion:", err)
	}
	n := mm * 2 * math.Pi / 86400
	a := math.Cbrt(mu / (n * n))
	return a
}

func meanToEccentricAnomaly(M, e float64) float64 {
	E := M
	for i := 0; i < 100; i++ {
		dE := (M - E + e*math.Sin(E)) / (1 - e*math.Cos(E))
		E += dE
		if math.Abs(dE) < 1e-6 {
			break
		}
	}
	return E
}

func eccentricToTrueAnomaly(E, e float64) float64 {
	sinNu := math.Sqrt(1-e*e) * math.Sin(E) / (1 - e*math.Cos(E))
	cosNu := (math.Cos(E) - e) / (1 - e*math.Cos(E))
	return math.Atan2(sinNu, cosNu)
}

func positionInOrbitalPlane(a, e, nu float64) (float64, float64) {
	r := a * (1 - e*e) / (1 + e*math.Cos(nu))
	x := r * math.Cos(nu)
	y := r * math.Sin(nu)
	return x, y
}

func transformToECI(x, y float64, inclination, RAAN, argumentOfPerigee string) (float64, float64, float64) {
	inc, err := strconv.ParseFloat(strings.TrimSpace(inclination), 64)
	if err != nil {
		fmt.Println("Error parsing inclination:", err)
	}
	raan, err := strconv.ParseFloat(strings.TrimSpace(RAAN), 64)
	if err != nil {
		fmt.Println("Error parsing RAAN:", err)
	}
	argPerigee, err := strconv.ParseFloat(strings.TrimSpace(argumentOfPerigee), 64)
	if err != nil {
		fmt.Println("Error parsing argument of perigee:", err)
	}

	// Convert degrees to radians
	inc = inc * math.Pi / 180
	raan = raan * math.Pi / 180
	argPerigee = argPerigee * math.Pi / 180

	cosw := math.Cos(argPerigee)
	sinw := math.Sin(argPerigee)
	cosi := math.Cos(inc)
	sini := math.Sin(inc)
	cosO := math.Cos(raan)
	sinO := math.Sin(raan)

	r := [3]float64{x, y, 0}

	rp := [3]float64{
		cosw*r[0] - sinw*r[1],
		sinw*r[0] + cosw*r[1],
		r[2],
	}

	ri := [3]float64{
		rp[0],
		cosi*rp[1] - sini*rp[2],
		sini*rp[1] + cosi*rp[2],
	}

	rf := [3]float64{
		cosO*ri[0] - sinO*ri[1],
		sinO*ri[0] + cosO*ri[1],
		ri[2],
	}

	return rf[0], rf[1], rf[2]
}

func calculatePosition(keplarian KeplarianElements) SatellitePosition {
	a := deriveSemiMajorAxis(strings.TrimSpace(keplarian.meanMotion))
	e, err := strconv.ParseFloat("0."+strings.TrimSpace(keplarian.eccentricity), 64)
	if err != nil {
		fmt.Println("Error parsing eccentricity:", err)
	}
	M, err := strconv.ParseFloat(strings.TrimSpace(keplarian.meanAnomoly), 64)
	if err != nil {
		fmt.Println("Error parsing mean anomaly:", err)
	}
	M = M * math.Pi / 180 // convert to radians
	E := meanToEccentricAnomaly(M, e)
	nu := eccentricToTrueAnomaly(E, e)
	x, y := positionInOrbitalPlane(a, e, nu)
	ecix, eciy, eciz := transformToECI(x, y, strings.TrimSpace(keplarian.inclination), strings.TrimSpace(keplarian.RAAN), strings.TrimSpace(keplarian.argumentOfPerigee))

	return SatellitePosition{
		X: ecix,
		Y: eciy,
		Z: eciz,
	}
}

func printSatData(sat newSatellite) {
	keps := getKeplearians(fmtLineTwoTLE(sat))
	pos := calculatePosition(keps)

	fmt.Printf("Satellite: %v\nTLE Data:\nLine One:\n%v\nLine Two:\n%v\nKeplarians:\n%v\nPosition (ECI): X=%.2f km, Y=%.2f km, Z=%.2f km\n\n",
		sat.name, fmtLineOneTLE(sat), fmtLineTwoTLE(sat), keps, pos.X, pos.Y, pos.Z)
}

func main() {
	satelliteISS := newSatellite{
		name:  "ISS (ZARYA)",
		line1: "1 25544U 98067A   20358.54791667  .00016717  00000-0  10270-3 0  9002",
		line2: "2 25544  51.6432  21.5264 0002184  90.4728 285.4598 15.49212921247662",
	}

	printSatData(satelliteISS)
}

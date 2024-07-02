package main

import (
	"fmt"
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

func fmtLineOneTLE(sat newSatellite) lineOneTLE {

	lineOneArr := strings.Split(sat.line1, "")

	tle := lineOneTLE{
		lineNumber:             "\nLine Number: " + strings.Join(lineOneArr[0:1], ""),
		satelliteNumber:        "\nSatellite Number: " + strings.Join(lineOneArr[3:7], ""),
		classification:         "\nClassifcation (U=Unclassified): " + strings.Join(lineOneArr[7:8], ""),
		intDesLaunchYear:       "\nInternational Designator (Last two digits of launch year): " + strings.Join(lineOneArr[10:11], ""),
		intDesLaunchYearNumber: "\nInternational Designator (Launch number of the year): " + strings.Join(lineOneArr[12:14], ""),
		intDesLaunchPiece:      "\nInternational Designator (Piece of the launch): " + strings.Join(lineOneArr[15:17], ""),
		epochYear:              "\nEpoch Year (Last two digits of year): " + strings.Join(lineOneArr[19:20], ""),
		epochDay:               "\nEpoch (Day of the year and fractional portion of the day): " + strings.Join(lineOneArr[21:32], ""),
		meanMotionFirst:        "\nFirst Time Derivative of the Mean Motion: " + strings.Join(lineOneArr[34:43], ""),
		meanMotionSecond:       "\nSecond Time Derivative of Mean Motion (Leading decimal point assumed): " + strings.Join(lineOneArr[45:52], ""),
		BSTAR:                  "\nBSTAR drag term (Leading decimal point assumed): " + strings.Join(lineOneArr[54:61], ""),
		emphermerisType:        "\nEphemeris type: " + strings.Join(lineOneArr[62:63], ""),
		elementNumber:          "\nElement Number: " + strings.Join(lineOneArr[65:68], ""),
		checksum:               "\nChecksum (Modulo 10) - (Letters, blanks, periods, plus signs = 0; minus signs = 1): " + strings.Join(lineOneArr[68:69], ""),
	}

	return tle
}

func fmtLineTwoTLE(sat newSatellite) lineTwoTLE {

	lineTwoArr := strings.Split(sat.line2, "")

	tle := lineTwoTLE{
		lineNumber:                     "\nLine Number: " + strings.Join(lineTwoArr[0:1], ""),
		satelliteNumber:                "\nSatellite Number: " + strings.Join(lineTwoArr[3:7], ""),
		inclination:                    "\nInclination [Degrees]: " + strings.Join(lineTwoArr[9:16], ""),
		rightAscensionsOfAscendingNode: "\nRight Ascension of the Ascending Node [Degrees]: " + strings.Join(lineTwoArr[18:25], ""),
		eccentricity:                   "\nEccentricity (Leading decimal point assumed): " + strings.Join(lineTwoArr[27:33], ""),
		argumentOfPerigee:              "\nArgument of Perigee [Degrees]: " + strings.Join(lineTwoArr[35:42], ""),
		meanAnomoly:                    "\nMean Anomaly [Degrees]: " + strings.Join(lineTwoArr[44:51], ""),
		meanMotion:                     "\nMean Motion [Revs per day]: " + strings.Join(lineTwoArr[53:63], ""),
		revolutionNumAtEpoch:           "\nRevolution number at epoch [Revs]: " + strings.Join(lineTwoArr[64:68], ""),
		checksum:                       "\nChecksum: " + strings.Join(lineTwoArr[68:69], ""),
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

func printSatData(sat newSatellite) {
	fmt.Printf("TLE Data:\n%v\nLine One:\n%v\nLine Two:\n%v\nKeplarians:\n%v", sat.name, fmtLineOneTLE(sat), fmtLineTwoTLE(sat), getKeplearians(fmtLineTwoTLE(sat)))
}

func main() {

	satelliteISS := newSatellite{
		name:  "ISS (ZARYA)",
		line1: "1 25544U 98067A   20358.54791667  .00016717  00000-0  10270-3 0  9002",
		line2: "2 25544  51.6432  21.5264 0002184  90.4728 285.4598 15.49212921247662",
	}

	satelliteNOAA := newSatellite{
		name:  "NOAA",
		line1: "1 23455U 94089A   97320.90946019  .00000140  00000-0  10191-3 0  2621",
		line2: "2 23455  99.0090 272.6745 0008546 223.1686 136.8816 14.11711747148495",
	}

	printSatData(satelliteNOAA)
	printSatData(satelliteISS)
}

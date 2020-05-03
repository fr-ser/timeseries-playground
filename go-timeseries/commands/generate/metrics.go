package generate

import (
	"encoding/csv"
	"math"
	"math/rand"
	"strconv"
	"time"

	"timeseries/tools"
)

// createAndSaveReading generates a reading for each metrics
// and writes it to the Writer
func createAndSaveReading(writer *csv.Writer, readingTime time.Time, machineID int) {
	stringTime := readingTime.Format(time.RFC3339)
	stringMachine := strconv.Itoa(machineID)
	turnedOff := isTurnedOff(readingTime, machineID)

	readings := [][]string{
		{stringTime, "1", stringMachine, getEngineTemp(turnedOff)},
		{stringTime, "2", stringMachine, getOilTemp(turnedOff)},
		{stringTime, "3", stringMachine, getOilPressure(turnedOff)},
		{stringTime, "4", stringMachine, getRunningHours(readingTime, machineID)},
		{stringTime, "5", stringMachine, getEngineLoad(turnedOff)},
	}
	for _, reading := range readings {
		err := writer.Write(reading)
		tools.CheckError("Cannot write to file", err)
	}
}

const (
	engineTempOn  = 80
	engineStdOn   = 20
	engineTempOff = 12
	engineStdOff  = 2
)

func getEngineTemp(turnedOff bool) string {
	var temp float64
	volatility := 1.0
	if rand.Intn(100) > 95 {
		volatility = 3.0
	}

	if turnedOff {
		temp = rand.NormFloat64()*engineStdOff*volatility + engineTempOff
	} else {
		temp = rand.NormFloat64()*engineStdOn*volatility + engineTempOn
	}

	return strconv.FormatFloat(temp, 'f', 5, 32)
}

const (
	oilTempOn     = 120
	oilTempStdOn  = 30
	oilTempOff    = 20
	oilTempStdOff = 3
)

func getOilTemp(turnedOff bool) string {
	var temp float64
	volatility := 1.0
	if rand.Intn(100) > 95 {
		volatility = 3.0
	}

	if turnedOff {
		temp = rand.NormFloat64()*oilTempStdOff*volatility + oilTempOff
	} else {
		temp = rand.NormFloat64()*oilTempStdOn*volatility + oilTempOn
	}

	return strconv.FormatFloat(temp, 'f', 5, 32)
}

const (
	oilPressureOn     = 560
	oilPressureStdOn  = 200
	oilPressureOff    = 1
	oilPressureStdOff = 0.2
)

func getOilPressure(turnedOff bool) string {
	var pressure float64
	volatility := 1.0
	if rand.Intn(100) > 95 {
		volatility = 2.0
	}
	if turnedOff {
		pressure = rand.NormFloat64()*oilPressureStdOff*volatility + oilPressureOff
	} else {
		pressure = rand.NormFloat64()*oilPressureStdOn*volatility + oilPressureOn
	}

	return strconv.FormatFloat(pressure, 'f', 5, 32)
}

const (
	loadOn    = 80
	loadStdOn = 10
	loadOff   = 0
)

func getEngineLoad(turnedOff bool) string {
	var load float64
	volatility := 1.0
	if rand.Intn(100) > 95 {
		volatility = 6.0
	}

	if turnedOff {
		load = 0
	} else {
		load = rand.NormFloat64()*loadStdOn*volatility + loadOn
	}

	return strconv.FormatFloat(load, 'f', 5, 32)
}

var engineStart, _ = time.Parse(time.RFC3339, "2012-12-12T00:00:00Z")

func getRunningHours(readingTime time.Time, machineID int) string {
	var hours = (float64(readingTime.Unix()-engineStart.Unix()) / 60)
	return strconv.FormatFloat(hours, 'f', 2, 32)
}

func isTurnedOff(someTime time.Time, machineID int) bool {
	// vary time to turn on depending on day and machineID
	var varianceMin float64 = 180
	var machineHash = float64(machineID+12) * 13 * 11 * 7
	offset := math.Mod(float64(someTime.Day())*machineHash, varianceMin) - varianceMin/2
	realMinutes := someTime.Hour()*60 + someTime.Minute()
	offsetMinutes := realMinutes + int(offset)

	return (offsetMinutes > 20*60 || offsetMinutes < 8*60)
}

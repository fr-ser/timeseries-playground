package generate

import (
	"encoding/csv"
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

	readings := [][]string{
		{stringTime, "1", stringMachine, getEngineTemp(readingTime, machineID)},
		{stringTime, "2", stringMachine, getOilTemp(readingTime, machineID)},
		{stringTime, "3", stringMachine, getOilPressure(readingTime, machineID)},
		{stringTime, "4", stringMachine, getRunningHours(readingTime, machineID)},
		{stringTime, "5", stringMachine, getEngineLoad(readingTime, machineID)},
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

func getEngineTemp(readingTime time.Time, machineID int) string {
	var temp float64

	if isNight(readingTime) {
		temp = rand.NormFloat64()*engineStdOff + engineTempOff
	} else {
		temp = rand.NormFloat64()*engineStdOn + engineTempOn
	}

	return strconv.FormatFloat(temp, 'f', 5, 32)
}

const (
	oilTempOn     = 120
	oilTempStdOn  = 30
	oilTempOff    = 20
	oilTempStdOff = 3
)

func getOilTemp(readingTime time.Time, machineID int) string {
	var temp float64

	if isNight(readingTime) {
		temp = rand.NormFloat64()*oilTempStdOff + oilTempOff
	} else {
		temp = rand.NormFloat64()*oilTempStdOn + oilTempOn
	}

	return strconv.FormatFloat(temp, 'f', 5, 32)
}

const (
	oilPressureOn     = 560
	oilPressureStdOn  = 200
	oilPressureOff    = 1
	oilPressureStdOff = 0.2
)

func getOilPressure(readingTime time.Time, machineID int) string {
	var pressure float64

	if isNight(readingTime) {
		pressure = rand.NormFloat64()*oilPressureStdOff + oilPressureOff
	} else {
		pressure = rand.NormFloat64()*oilPressureStdOn + oilPressureOn
	}

	return strconv.FormatFloat(pressure, 'f', 5, 32)
}

const (
	loadOn    = 80
	loadStdOn = 10
	loadOff   = 0
)

func getEngineLoad(readingTime time.Time, machineID int) string {
	var load float64

	if isNight(readingTime) {
		load = 0
	} else {
		load = rand.NormFloat64()*loadStdOn + loadOn
	}

	return strconv.FormatFloat(load, 'f', 5, 32)
}

var engineStart, _ = time.Parse(time.RFC3339, "1992-12-12T00:00:00Z")

func getRunningHours(readingTime time.Time, machineID int) string {
	var hours = (float64(readingTime.Unix()-engineStart.Unix()) / 60)
	return strconv.FormatFloat(hours, 'f', 2, 32)
}

func isNight(someTime time.Time) bool {
	hour := someTime.Hour()
	return (hour > 20 || hour < 8)
}

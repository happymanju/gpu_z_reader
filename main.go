package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type sensorReading struct {
	date              string
	gpuClock          float64
	memClock          float64
	gpuTemp           float64
	hotSpot           float64
	fan1per           int64
	fan1RPM           int64
	fan2per           int64
	fan2RPM           int64
	memUsed           int64
	gpuLoad           int64
	memControllerLoad int64
	videoEngineLoad   int64
	busInterfaceLoad  float64
	boardPowerDraw    float64
	gpuChipPowerDraw  float64
	pcieSlotPower     float64
	pcieSlotVoltage   float64
	eightPinPower     float64
	eightPinVoltage   float64
	powerConsumption  float64
	perfCapReason     int64
	gpuVoltage        float64
	cpuTemp           float64
	sysMemory         int64
}

type dataFrame []sensorReading

func calcAverage(arr dataFrame, prop string) float64 {
	var sum int
	length := len(arr)

	for _, sr := range arr {
		//sum += sr[prop]
	}
}

func parseFloats(s string) float64{
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("error: couldn't parse value as float; ", err)
		return 0.0
	}
	return v
}

func parseInts(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("error: couldn't parse value as int; ", err)
		return 0
	}

	return v
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No filename passed in")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f)
	_, err = reader.Read()
	if err != nil {
		println("error reading csv line: ", err)
	}
	
	var df dataFrame

	for {

		r, err := reader.Read()
		if err != nil {
			log.Println("end of file")
			break
		}
		
		newSensorReading := sensorReading {
			date: r[0],
			gpuClock: parseFloats(r[1]),
			memClock: parseFloats(r[2]),
			gpuTemp: parseFloats(r[3]),
			hotSpot: parseFloats(r[4]),
			fan1per: parseInts(r[5]),
			fan1RPM: parseInts(r[6]),
			fan2per: parseInts(r[7]),
			fan2RPM: parseInts(r[8]),
			memUsed: parseInts(r[9]),
			gpuLoad: parseInts(r[10]),
			memControllerLoad: parseInts(r[11]),
			videoEngineLoad: parseInts(r[12]),
			busInterfaceLoad: parseFloats(r[14]),
			boardPowerDraw: parseFloats(r[15]),
			gpuChipPowerDraw: parseFloats(r[16]),
			pcieSlotPower: parseFloats(r[17]),
			pcieSlotVoltage: parseFloats(r[18]),
			eightPinPower: parseFloats(r[19]),
			eightPinVoltage: parseFloats(r[20]),
			powerConsumption: parseFloats(r[21]),
			perfCapReason: parseInts(r[22]),
			gpuVoltage: parseFloats(r[23]),
			cpuTemp: parseFloats(r[24]),
			sysMemory: parseInts(r[25]),
		}

		df = append(df,newSensorReading)
	}





}

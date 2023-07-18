package main

import (
	"encoding/csv"
	"log"
	"os"
)

type sensorReading struct {
	date              string
	gpuClock          float64
	memClock          float64
	gpuTemp           float64
	hotSpot           float64
	fan1per           int
	fan1RPM           int
	fan2per           int
	fan2RPM           int
	memUsed           int
	gpuLoad           int
	memControllerLoad int
	videoEngineLoad   int
	busInterfaceLoad  float64
	boardPowerDraw    float64
	gpuChipPowerDraw  float64
	pcieSlotPower     float64
	pcieSlotVoltage   float64
	eightPinPower     float64
	eightPinVoltage   float64
	powerConsumption  float64
	perfCapReason     int
	gpuVoltage        float64
	cpuTemp           float64
	sysMemory         int
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

}

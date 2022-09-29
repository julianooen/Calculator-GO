package services

import (
	"calc/entities"
)

var memoryCalc []entities.CalcResult

func AppendCalcResult(resStruct entities.CalcResult) {
	memoryCalc = append(memoryCalc, resStruct)
}

func ShowHistory() (history []entities.CalcResult) {
	return memoryCalc
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	result := []map[string]any{}
	for _, value := range data {
		// fmt.Println(value)

		//	Membuat map untuk menampung isi data
		currentResult := make(map[string]any)
		isiData := strings.Split(value, ";")

		//	Buat variabel di tiap datanya
		name := isiData[0]
		age, _ := strconv.Atoi(isiData[1])
		address := isiData[2]
		height, _ := strconv.ParseFloat(isiData[3], 64)
		isMarried, _ := strconv.ParseBool(isiData[4])

		//	masukkan data tadi kedalam map current result
		currentResult["name"] = name
		currentResult["age"] = age
		currentResult["address"] = address

		if height != 0.0 {
			currentResult["height"] = height
		}
		if isiData[4] != "" {
			currentResult["isMarried"] = isMarried
		}

		fmt.Println("current result", currentResult)
		result = append(result, currentResult)
	}
	return result // TODO: replace this
}

func main() {
	biodata := []string{
		"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;",
	}
	PopulationData(biodata)
}

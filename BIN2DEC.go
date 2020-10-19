package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isBinValidator(binData string) bool{
	binDataset := strings.Split(binData, "")
	for _, bin := range binDataset {
		if bin != "0" && bin != "1" {
			return false
		}
	}
	return true
}

func stringReverser(message string) string {
	runes := []rune(message)
	for i, j := 0, len(runes)-1; i<j;i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func converter(binPayload string) float64 {
	var isBin bool = isBinValidator(binPayload)
	if isBin {
		revertedPayload := stringReverser(binPayload)
		binDataset := strings.Split(revertedPayload, "")
		var decVal float64
		for i, bin := range binDataset {
			intBin, _ := strconv.Atoi(bin)
			decVal += float64(intBin) * math.Pow(2, float64(i))
		}
		return decVal
	}
	return 0
}

func main (){
	var binPayload string
	fmt.Print("Binary\t:")
	fmt.Scanf("%s", &binPayload)
	fmt.Print("Decimal\t:")
	fmt.Println(converter(binPayload))
}
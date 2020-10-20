package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func usdConverter(value int) (finalValue int) {
	exchangeValue := map[string]int{}
	response, err := http.Get("https://free.currconv.com/api/v7/convert?q=USD_IDR&compact=ultra&apiKey=7ea48091941cc8b3badb")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	byteResponse, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(byteResponse, &exchangeValue)
	finalValue = value * exchangeValue["USD_IDR"]
	return
}

func main() {
	var valuePayload int
	fmt.Print("USD\t:")
	fmt.Scanf("%d", &valuePayload)
	fmt.Print("IDR\t:")
	fmt.Println(usdConverter(valuePayload))
}

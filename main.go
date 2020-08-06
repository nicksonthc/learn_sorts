package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// TestData struct for json test data
type TestData struct {
	Data []int `json:"data"`
}

type sortFunction func(arr []int) []int

func countElapse(sf sortFunction, arr []int) {
	start := time.Now()
	fmt.Println(sf(arr))
	elapsed := time.Since(start)
	log.Printf("Insertion Sort took %s", elapsed)
}

func openTestData() []int {
	jsonFile, err := os.Open("./testdata.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var testData TestData

	json.Unmarshal([]byte(byteValue), &testData)

	return testData.Data
}

func insertionSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		currentValue := arr[i]
		j := i - 1

		for j >= 0 && currentValue < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = currentValue
		i++
	}
	return arr
}

func main() {
	testArray := openTestData()
	countElapse(insertionSort, testArray)
}

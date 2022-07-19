package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//CountSum Function calculates the sum from a text file that contains a number
func CountSum(arrayData []int, ch chan int) {

	var sum int = 0

	for i := 0; i < len(arrayData); i++ {
		sum += arrayData[i]
	}

	ch <- sum

}

func main() {
	var path string
	routines := flag.Int("routines", 2, "an int")
	fmt.Println("routine:", *routines)
	flag.StringVar(&path, "path", "", "")
	flag.Parse()
	fmt.Println(path, "The File Contain Numbers")
	data, err := ioutil.ReadFile(path)
	//reading a file
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fieldData := strings.Fields(string(data))
	var arrayData []int

	for _, value := range fieldData {
		number, _ := strconv.Atoi(value)
		arrayData = append(arrayData, number)
	}
	chunksize := 0
	chunksize = len(arrayData) / *routines
	fmt.Println("The chunck size will be =", chunksize)
	start := 0
	end := chunksize
	ch := make(chan int)
	for i := 0; i < *routines; i++ {
		chunkData := arrayData[start:end]
		start = start + chunksize
		end = end + chunksize
		go CountSum(chunkData, ch)

	}
	for i := 0; i < *routines; i++ {
		chunkData := <-ch
		fmt.Println("The sum of each chunk is =", chunkData)

	}

}

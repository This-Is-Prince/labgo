package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type City struct {
	Name       string
	Population int
}

func main() {
	log.Println("Tutorial 5 - Fan-In-Fan-Out")

	start := time.Now()
	f, err := os.Open("./cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cities := genRows(f)

	// fan out pattern
	// more than one worker competing to consume the same channel
	totalWorkers := 2
	workers := []<-chan City{}
	for i := 0; i < totalWorkers; i++ {
		workers = append(workers, upperCityName(cities))
	}

	// wg := &sync.WaitGroup{}
	// wg.Add(totalWorkers)
	// for index, upperCities := range workers {
	// 	go func(index int, upperCities <-chan City) {
	// 		for c := range upperCities {
	// 			log.Printf("UpperCities%d:- %v", index, c)
	// 		}
	// 		wg.Done()
	// 	}(index, upperCities)
	// }
	// wg.Wait()

	// fan in pattern consolidates the outputs from multiple channels into one
	for c := range fanIn(workers...) {
		log.Println(c)
	}
	log.Println("Total Time:", time.Since(start))
}

func fanIn(cityChannels ...<-chan City) <-chan City {
	cityChannelOut := make(chan City)
	wg := &sync.WaitGroup{}
	wg.Add(len(cityChannels))
	for index, cityChannelIn := range cityChannels {
		go func(index int, cityChannelIn <-chan City) {
			for city := range cityChannelIn {
				cityChannelOut <- city
			}
			wg.Done()
		}(index, cityChannelIn)
	}
	go func() {
		wg.Wait()
		close(cityChannelOut)
	}()
	return cityChannelOut
}

func upperCityName(cities <-chan City) <-chan City {
	upperCities := make(chan City)
	go func() {
		for c := range cities {
			upperCities <- City{Name: strings.ToUpper(c.Name), Population: c.Population}
		}
		close(upperCities)
	}()
	return upperCities
}

func genRows(r io.Reader) chan City {
	cities := make(chan City)
	go func() {
		reader := csv.NewReader(r)
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			populationInt, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			cities <- City{Name: row[1], Population: populationInt}
		}
		close(cities)
	}()
	return cities
}

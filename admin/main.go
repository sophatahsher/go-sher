package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	/*
		var aString string = "This is Go!"
		fmt.Println(aString)
		fmt.Println("The variable's type is %T", aString)

		var anInteger int = 49
		fmt.Println(anInteger)

		var defaultInt int
		fmt.Println(defaultInt)

		n := time.Now()
		fmt.Println("I record this video at ", n)

		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Println("Go launched at ", t)
		fmt.Println(t.Format(time.ANSIC))

		parseTime, _ := time.Parse(time.ANSIC, " Tue Nov 10 23:00:00 2009")
		fmt.Printf("The type of parseTime is %T\n", parseTime)
	*/

	content := readTextFromUrl("http://services.explorecalifornia.org/json/tours.php")
	//fmt.Print(content)

	tours := toursFromJson(content)

	for _, tour := range tours {
		fmt.Println(tour.Name)
	}
}

func readTextFromUrl(url string) string {

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()
	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(byte)
	//fmt.Print(content)

	return content
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()

	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)

		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}

	return tours
}

type Tour struct {
	Name, Price string
}

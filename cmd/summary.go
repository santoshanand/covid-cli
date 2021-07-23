/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type Response struct {
	ID          string    `json:"ID"`
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Province    string    `json:"Province"`
	City        string    `json:"City"`
	CityCode    string    `json:"CityCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Confirmed   int       `json:"Confirmed"`
	Deaths      int       `json:"Deaths"`
	Recovered   int       `json:"Recovered"`
	Active      int       `json:"Active"`
	Date        time.Time `json:"Date"`
}

type Summary struct {
	Deaths    int    `json:"Deaths"`
	Recovered int    `json:"Recovered"`
	Active    int    `json:"Active"`
	Country   string `json:"Country"`
}

// curl --location --request GET 'https://api.covid19api.com/live/country/south-africa/status/confirmed'
func getSummary(country string) {
	url := fmt.Sprintf("https://api.covid19api.com/live/country/%v/status/confirmed", country)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println("Unable to get summary")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Unable to get global summary data")
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to read data")
	}
	resArray := []Response{}
	err = json.Unmarshal(resData, &resArray)
	if err != nil {
		fmt.Println("Unable to parse data")
	}

	var sumDeath = 0
	var sumActive = 0
	var sumRecoverd = 0
	for _, item := range resArray {
		sumDeath += item.Deaths
		sumActive += item.Active
		sumRecoverd += item.Recovered
	}
	fmt.Printf("----------------------\n")
	fmt.Printf("   Summary of %s\n", country)
	fmt.Printf("----------------------\n")
	fmt.Printf("Active Cases   : %v\n", sumActive)
	fmt.Printf("Recoverd Cases : %v\n", sumRecoverd)
	fmt.Printf("Death Cases    : %v\n", sumDeath)

}

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A summary of new and total cases(active, recovered, death) per country name 'covid-cli summary south-africa'",
	Long:  `A summary of new and total cases per country updated daily`,
	Run: func(cmd *cobra.Command, args []string) {
		country := "india"
		if len(args) > 0 {
			country = args[0]
		}
		getSummary(country)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

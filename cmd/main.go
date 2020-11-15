package main

import (
	etsy "etsgo"
	"etsgo/cmd/generator"
	"os"
)

var etsyApiKey string

func init() {
	etsyApiKey = os.Getenv("ETSY_API_KEY")
}

func main() {
	apiClient := etsy.BuildApiClient(etsyApiKey)
	if err := generator.GenerateTypes(apiClient); err != nil {
		panic(err)
	}
}

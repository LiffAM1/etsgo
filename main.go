package main

import (
	etsgo "etsgo/etsgo"
	"fmt"
)

func main() {
	apiClient := etsgo.BuildApiClient("<apiKey>")
	shopId := "<shopId>"

	pathParams := map[string]string{"shop_id": shopId}

	result, statusCode, err := apiClient.MakeRequest("findAllShopListingsActive", pathParams, nil)
	fmt.Printf("%d", statusCode)
	fmt.Printf(result)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

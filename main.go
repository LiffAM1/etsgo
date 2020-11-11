package main

import (
	etsgo "etsgo/etsgo"
	"fmt"
)

func main() {
	apiClient := etsgo.BuildApiClient("vuj1qgnarw7oqljekwaf5s96")
	shopId := "22695679"

	pathParams := map[string]string{"shop_id": shopId}

	result, statusCode, err := apiClient.MakeRequest("findAllShopListingsActive", pathParams, nil)
	fmt.Printf("%d", statusCode)
	fmt.Printf(result)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

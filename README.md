# EtsGo

EtsGo is a basic Etsy API wrapper for GoLang.
I am an Etsy seller and created this wrapper to use for an upcoming inventory management tool I will create for Google Sheets, but I figured other people might want to use it in their Go projects.

NOTE: Currently only supports GET Etsy API requests

## Get Started

Get it!

go get github.com/LiffAM1/etsgo

Import it!

```go
import etsgo "github.com/LiffAM1/etsgo"
```

## Usage

First, create an API client. The API client will store your Etsy API key and a map of API methods that etsgo supports:

```go
package main

import etsgo "github.com/LiffAM1/etso"
import "fmt"

func main() {
	apiClient := etsgo.BuildApiClient("myCoolApiKey")
	fmt.Println(apiClient.SupportedMethods[" 	findAllShopListingsActive"])
}
```
Now, make a request! The MakeRequest method takes a method name, a list of path parameters (ones that are part of the URI), and a list of querystring parameters:

```go
func main() {
	apiClient := etsgo.BuildApiClient("myCoolApiKey")
	shopId := "myCoolShopId"

	pathParams := map[string]string{"shop_id": shopId}

	result, statusCode, err := apiClient.MakeRequest("findAllShopListingsActive", pathParams, nil)
	fmt.Printf("%d", statusCode)
	fmt.Printf(result)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
```
## Contributing
Pull requests are welcome :)

## License
[MIT](https://choosealicense.com/licenses/mit/)
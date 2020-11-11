package etsgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// API Client
var baseUrl = "https://openapi.etsy.com/v2/"

type Endpoint struct {
	Name        string
	Description string
	Uri         string
	Params      map[string]string
	Defaults    []string
	Http_method string
	Type        string
}

type Endpoints struct {
	Count   int
	Results []Endpoint
}

type EtsyAPIClient struct {
	ApiKey           string
	SupportedMethods map[string]Endpoint
}

func (client *EtsyAPIClient) buildAPIUrl(url string) string {
	return baseUrl + url + "?api_key=" + client.ApiKey
}

func BuildApiClient(apiKey string) *EtsyAPIClient {
	client := EtsyAPIClient{
		ApiKey:           apiKey,
		SupportedMethods: make(map[string]Endpoint),
	}

	url := client.buildAPIUrl("")

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	result, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var endpoints Endpoints
	json.Unmarshal([]byte(result), &endpoints)
	for _, element := range endpoints.Results {
		// TODO: Add support for other HTTP methods
		if element.Http_method == "GET" {
			client.SupportedMethods[element.Name] = element
		}
	}

	return &client
}

func (client *EtsyAPIClient) MakeRequest(methodName string, pathParams map[string]string, params map[string]string) (string, int, error) {
	var endpoint Endpoint
	if v, ok := client.SupportedMethods[methodName]; ok {
		endpoint = v
	} else {
		return "", 405, errors.New(methodName + " not supported")
	}

	url := endpoint.Uri
	url = client.buildAPIUrl(url)

	if pathParams != nil {
		for k, v := range pathParams {
			url = strings.ReplaceAll(url, ":"+k, v)
		}
	}

	if params != nil {
		for k, v := range params {
			url = url + fmt.Sprintf("?%s=%s", k, v)
		}
	}

	// TODO Add support for other HTTP methods
	if endpoint.Http_method == "GET" {
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return "", response.StatusCode, err
		}

		result, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			log.Fatal(err)
			return "", response.StatusCode, err
		}
		return string([]byte(result)), response.StatusCode, nil
	} else {
		return "", 405, errors.New(endpoint.Http_method + " not yet supported")
	}
}

package generator

import (
	"encoding/json"
	etsy "etsgo"
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"github.com/iancoleman/strcase"
	"io"
	"io/ioutil"
)

const GENTYPES_DIR = "gentypes"

// RetryableError implements the Error interface
type EmptyError struct {
	msg string
}

// Error implements the error interface for RetryableError
func (e EmptyError) Error() string {
	return e.msg
}

func newEmptyError(msg string) *EmptyError {
	return &EmptyError{msg: msg}
}

type TypeGenParams struct {
	Method     string
	PathParams map[string]string
	UrlParams  map[string]string
}

var SupportedTypes = map[string]TypeGenParams{
	"Country": {
		Method: "getCountry",
		PathParams: map[string]string{
			"country_id": "10",
		},
	},
	"User": {
		Method: "getUser",
		PathParams: map[string]string{
			"user_id": "bntwbx4eiyc5g774",
		},
	},
	"Shop": {
		Method: "getShop",
		PathParams: map[string]string{
			"shop_id": "26078993",
		},
	},
	"Listing": {
		Method: "getListing",
		PathParams: map[string]string{
			"listing_id": "890221638",
		},
	},
}

func GenerateTypes(client *etsy.Client) error {
	methods := client.GetMethodTable()
	generatedTypes := map[string]bool{
		"ApiMethod": true,
	}

	for _, method := range methods {
		if _, found := generatedTypes[method.Type]; found {
			continue
		}

		supType, found := SupportedTypes[method.Type]
		if !found {
			continue
		}
		if supType.Method != method.Name {
			// TODO: generate functions for supported types
			continue
		}

		url := client.BuildUrl(method.Uri, supType.PathParams, supType.UrlParams)
		fmt.Printf("%s: %s (%s)\n", method.Name, method.Type, method.Uri)
		b, err := client.Get(url, false)
		if err != nil {
			return err
		}

		filename := GENTYPES_DIR + "/" + strcase.ToLowerCamel(method.Type) + ".go"
		fmt.Printf("Generating file %s...\n", filename)
		generatedFile, err := gojson.Generate(b.(io.ReadCloser), ParseEtsyJson, method.Type, "etsy", []string{"json"}, false, true)
		if err != nil {
			switch err.(type) {
			case *EmptyError:
				continue
			default:
				return err
			}
		}
		if err := ioutil.WriteFile(filename, generatedFile, 0664); err != nil {
			return err
		}
		fmt.Printf("%s generated successfully\n", filename)
		generatedTypes[method.Type] = true
	}
	return nil
}

// ParseEtsyJson is a custom parser that accounts for etsy's standard response format https://www.etsy.com/developers/documentation/getting_started/api_basics#section_standard_response_format
func ParseEtsyJson(input io.Reader) (interface{}, error) {
	var result []interface{}
	var sResp etsy.StandardResponseFormat
	if err := json.NewDecoder(input).Decode(&sResp); err != nil {
		return nil, err
	}
	if err := etsy.UnmarshalSlice(sResp.Results, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, newEmptyError("parser found no content to generate the struct")
	}
	// Since we already account for Results being a slice, generated structs
	// don't need to be slices too, so just get the first and generate
	return result[0], nil
}

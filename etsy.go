package etsy

import (
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// API Client
const BASE_URL = "https://openapi.etsy.com/v2/"

// RetryableError implements the Error interface
type RetryableError struct {
	msg string
}

// Error implements the error interface for RetryableError
func (e RetryableError) Error() string {
	return e.msg
}

func newRetryableError(msg string) *RetryableError {
	return &RetryableError{msg: msg}
}

type StandardResponseFormat struct {
	Count   int
	Results []json.RawMessage
	Params  interface{}
	Type    string
}

type Client struct {
	ApiKey string
}

func (client *Client) decodeJSON(resp *http.Response) (interface{}, error) {
	defer resp.Body.Close()
	var sResp StandardResponseFormat
	if err := json.NewDecoder(resp.Body).Decode(&sResp); err != nil {
		return nil, err
	}
	switch sResp.Type {
	case "ApiMethod":
		var decResults []ApiMethod
		if err := UnmarshalSlice(sResp.Results, &decResults); err != nil {
			return nil, nil
		}
		return decResults, nil
	default:
		panic("unsupported response type")
	}

	return nil, nil
}

// https://stackoverflow.com/questions/24777603/create-slice-of-unknown-type
func UnmarshalSlice(docs []json.RawMessage, rv interface{}) error {
	slice := reflect.ValueOf(rv).Elem()
	// Allocate slice with desired capacity
	slice.Set(reflect.MakeSlice(slice.Type(), 0, len(docs)))

	v := reflect.New(slice.Type().Elem())
	for _, doc := range docs {
		if err := json.Unmarshal(doc, v.Interface()); err != nil {
			return err
		}
		slice.Set(reflect.Append(slice, v.Elem()))
	}
	return nil
}

func (client *Client) Get(url string, unmarshal bool) (interface{}, error) {
	var body interface{}
	err := retry.Do(
		func() error {
			resp, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("error: %s", err)
			}
			if resp.StatusCode >= 500 {
				return newRetryableError(fmt.Sprintf("retryable http error: (%d) %s", resp.StatusCode, resp.Header.Get("X-Error-Detail")))
			} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				return fmt.Errorf("http error: (%d) %s", resp.StatusCode, resp.Header.Get("X-Error-Detail"))
			}
			if unmarshal {
				//fmt.Printf("%s\n", resp.Header.Get("X-RateLimit-Remaining"))
				body, err = client.decodeJSON(resp)
				return err
			}
			body = resp.Body
			return nil
		},
		retry.RetryIf(func(err error) bool {
			switch err.(type) {
			case *RetryableError:
				return true
			default:
				return false
			}
		}),
		retry.Attempts(4),
		retry.DelayType(retry.BackOffDelay),
		retry.Delay(time.Duration(200)*time.Millisecond),
		retry.LastErrorOnly(true),
		retry.OnRetry(func(n uint, err error) {
			fmt.Printf("retry #%d, %s\n", n, err)
		}),
	)
	return body, err
}

func (client *Client) BuildUrl(uri string, pathParams map[string]string, params map[string]string) string {
	if pathParams != nil {
		for k, v := range pathParams {
			uri = strings.ReplaceAll(uri, ":"+k, v)
		}
	}

	if params != nil {
		for k, v := range params {
			uri = uri + fmt.Sprintf("?%s=%s", k, v)
		}
	}
	return BASE_URL + uri + "?api_key=" + client.ApiKey
}

func BuildApiClient(apiKey string) *Client {
	client := Client{
		ApiKey: apiKey,
	}
	return &client
}

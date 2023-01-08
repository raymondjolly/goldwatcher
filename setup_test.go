package main

import (
	"bytes"
	"fyne-gold/repository"
	"fyne.io/fyne/v2/test"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	log.Println("setting client to test client")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `
{
  "ts": 1672424658476,
  "tsj": 1672424656077,
  "date": "Dec 30th 2022, 01:24:16 pm NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 1821.745,
      "xagPrice": 23.8367,
      "chgXau": 6.61,
      "chgXag": -0.0383,
      "pcXau": 0.3642,
      "pcXag": -0.1604,
      "xauClose": 1815.135,
      "xagClose": 23.875
    }
  ]
}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})

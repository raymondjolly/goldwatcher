package main

import (
	"bytes"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image"
	"image/png"
	"io"
	"os"
	"strings"
)

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	chartContainer := container.NewVBox(chart)
	app.PriceChartContainer = chartContainer
	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))
	var img *canvas.Image

	err := app.downloadFile(apiURL, "gold.png")
	if err != nil {
		//use bundled image
		img = canvas.NewImageFromResource(resourceUnreachablePng)
	} else {
		img = canvas.NewImageFromFile("gold.png")
	}

	img.SetMinSize(fyne.Size{
		Width:  770,
		Height: 410,
	})

	img.FillMode = canvas.ImageFillOriginal

	return img
}

func (app *Config) downloadFile(URL, filename string) error {
	//get the response bytes from calling a URL
	resp, err := app.HTTPClient.Get(URL)

	errCheck(err)
	if resp.StatusCode != 200 {
		return errors.New("received incorrect response code when downloading image")
	}

	b, err := io.ReadAll(resp.Body)
	errCheck(err)
	defer resp.Body.Close()
	img, _, err := image.Decode(bytes.NewReader(b))
	errCheck(err)

	out, err := os.Create(fmt.Sprintf("./%s", filename))
	errCheck(err)
	png.Encode(out, img)
	return nil
}

func errCheck(err error) error {
	if err != nil {
		return err
	}
	return nil
}

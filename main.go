package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"os"
)

type Config struct {
	App      fyne.App
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var myApp Config

func main() {
	//create a fyne application
	fyneApp := app.NewWithID("me.raymondjolly.goldwatcher.preferences")
	myApp.App = fyneApp

	//create a logger
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//open a connection to the database

	//create a database repository

	//create and size a fyne window
	win := fyneApp.NewWindow("GoldWatcher")

	//show and run the application
	win.ShowAndRun()
}

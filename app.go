package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/reujab/wallpaper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const URL = "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US"
const DOMAIN = "https://www.bing.com"

// App struct
type App struct {
	ctx context.Context
}

type BingInfo struct {
	Images []struct {
		StartDate     string        `json:"startdate"`
		FullStartDate string        `json:"fullstartdate"`
		EndDate       string        `json:"enddate"`
		URL           string        `json:"url"`
		UrlBase       string        `json:"urlbase"`
		CopyRight     string        `json:"copyright"`
		CopyRightLink string        `json:"copyrightlink"`
		Title         string        `json:"title"`
		Quiz          string        `json:"quiz"`
		Wp            bool          `json:"wp"`
		Hsh           string        `json:"hsh"`
		Drk           int           `json:"drk"`
		Top           int           `json:"top"`
		Bot           int           `json:"bot"`
		Hs            []interface{} `json:"hs"`
	} `json:"images"`
	Tooltips struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltips"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (b *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func getBingInfo() BingInfo {
	client := &http.Client{}
	resp, err := client.Get(URL)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println("error")
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var bingInfo BingInfo
	err = json.Unmarshal(body, &bingInfo)
	if err != nil {
		println("json convert error")
	}
	return bingInfo
}

type ImageInfo struct {
	Url   string `json:"url"`
	Date  string `json:"date"`
	Title string `json:"title"`
}

func (b *App) GetImageInfo() ImageInfo {
	bingInfo := getBingInfo()
	imageInfo := bingInfo.Images[0]
	url := DOMAIN + strings.Replace(imageInfo.URL, "1920x1080", "UHD", 1)
	return ImageInfo{url, imageInfo.StartDate, imageInfo.Title}
}

func (b App) SetWallpaper(url string) {
	err := wallpaper.SetFromURL(url)
	if err != nil {
		log.Panicln("can't set wallpaper")
	}
}

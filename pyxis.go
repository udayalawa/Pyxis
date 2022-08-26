package main

import (
	"fmt"
)

var (
	purpose       int
	attackSurface string
	threads       int
	help          bool
	verbose       bool
	version       bool
	headless      bool
	cookie        string
	rawRequest    string
)

func main() {
	fmt.Println("(!) Legal disclaimer: Usage of Pyxis for attacking targets without prior mutual consent is illegal. It is the end user's responsibility to obey all applicable local, state and federal laws. Developers assume no liability and are not responsible for any misuse or damage caused by this program.")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println("Hello, it's Pyxis, How's the name?")
	fmt.Println("Pyxis is a Robust Web Crawler, Dorks Scanner, Links finder, Sensitive files Scanner, JS Scraper/Downloader, Attack Surface Visualizer, Capable of creating fake acount using temp mail api and Going through the whole applicateion js and inner frameworks")

	purpose = 1
	fmt.Println(purpose)
	attackSurface = "www.google.com"
	fmt.Printf("target is : %v", attackSurface)
}
func concent() {

}

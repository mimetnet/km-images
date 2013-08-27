package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Scrape(url string) {
	// Load the HTML document (in real use, the type would be *goquery.Document)
	var doc *goquery.Document
	var e error

	if doc, e = goquery.NewDocument(url); e != nil {
		panic(e.Error())
	}

	doc.Find("h2.headingTypeB01 a").Each(func(i int, s *goquery.Selection) {
		var cnt int
		var title string
		var parts []string

		// For each item found, get the band, title and score, and print it
		title = s.Text()

		if "" == title || true != strings.HasPrefix(title, "bizhub ") {
			return
		}

		title = strings.TrimPrefix(title, "bizhub ")
		parts = strings.Split(title, " ")
		cnt = len(parts)

		switch cnt {
		case 0:
			return
		case 1:
			d := NewDevice(parts[0])
			fmt.Printf("'%s' (%s)\n", d, parts)
			break
		default:
			var i = 0

			for i = 0; i < cnt; i++ {
				if "PRO" != parts[i] && "PRESS" != parts[i] {
					break
				}
			}

			d := NewDevice(parts[i])
			fmt.Printf("'%s' (%s)\n", d, parts)
			break
		}

	})
}

func main() {
	Scrape("http://www.konicaminolta-images.eu/images/list/category/Copier%20Print%20Systems__Multifunctional%20Systems%20Black%20and%20White/?tx_kmmediapool_pi1[itemsperpage]=1000")
	Scrape("http://www.konicaminolta-images.eu/images/list/category/Copier%20Print%20Systems__Multifunctional%20Systems%20Colour/?tx_kmmediapool_pi1[itemsperpage]=1000")
	Scrape("http://www.konicaminolta-images.eu/images/list/category/Production%20Printing%20Systems__Production%20Printing%20Black%20and%20White/?tx_kmmediapool_pi1[itemsperpage]=1000")
	Scrape("http://www.konicaminolta-images.eu/images/list/category/Production%20Printing%20Systems__Production%20Printing%20Colour/?tx_kmmediapool_pi1[itemsperpage]=1000")
}

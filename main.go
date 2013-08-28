package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func MapSelection(i int, s *goquery.Selection) string {
	var title string

	// For each item found, get the band, title and score, and print it
	if title = s.Text(); "" == title || !strings.HasPrefix(title, "bizhub ") {
		return ""
	}

	return strings.TrimPrefix(title, "bizhub ")
}

func ConvertTitle(title string) <-chan string {
	parts := strings.Split(title, " ")
	cnt := len(parts)
	ret := make(chan string)

	go func() {
		switch cnt {
		case 0:
			break
		case 1:
			//d := NewDevice(parts[0])
			//fmt.Println(d)
			ret <- parts[0]
			break
		default:
			var i = 0

			for i = 0; i < cnt; i++ {
				if "PRO" != parts[i] && "PRESS" != parts[i] {
					break
				}
			}

			for ; i < cnt; i++ {
				if "P" != parts[i] && "DS" != parts[i] {
					//d := NewDevice(parts[i])
					//fmt.Println(d)
					ret <- parts[i]
				}
			}
			break
		}

		close(ret)
	}()

	return ret
}

func Scrape(url string) {
	var e error
	var doc *goquery.Document

	if doc, e = goquery.NewDocument(url); e != nil {
		panic(e.Error())
	}

	titles := doc.Find("h2.headingTypeB01 a").Map(MapSelection)
	//devices := make([]Device, len(titles))

	for _, title := range titles {
		if "" != title {
			for val := range ConvertTitle(title) {
				fmt.Println(NewDevice(val))
			}
		}
	}
}

func main() {
	Scrape("http://www.konicaminolta-images.eu/images/list/category/Copier%20Print%20Systems__Multifunctional%20Systems%20Black%20and%20White/?tx_kmmediapool_pi1[itemsperpage]=1000")
	Scrape("http://www.konicaminolta-images.eu/images/list/category/Copier%20Print%20Systems__Multifunctional%20Systems%20Colour/?tx_kmmediapool_pi1[itemsperpage]=1000")
	Scrape("http://www.konicaminolta-images.eu/images/list/category/Production%20Printing%20Systems__Production%20Printing%20Black%20and%20White/?tx_kmmediapool_pi1[itemsperpage]=1000")
	Scrape("http://www.konicaminolta-images.eu/images/list/category/Production%20Printing%20Systems__Production%20Printing%20Colour/?tx_kmmediapool_pi1[itemsperpage]=1000")
}

package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
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
					ret <- parts[i]
				}
			}
			break
		}

		close(ret)
	}()

	return ret
}

func Scrape(url string, out chan *Device, done chan bool) {
	var e error
	var doc *goquery.Document

	if doc, e = goquery.NewDocument(url); e != nil {
		panic(e.Error())
	}

	cnt := 0
	titles := doc.Find("h2.headingTypeB01 a").Map(MapSelection)

	for _, title := range titles {
		if "" != title {
			for val := range ConvertTitle(title) {
				cnt++
				out <- NewDevice(val)
			}
		}
	}

	done <- true
	//close(out)
}

func main() {
	black := flag.Bool("bw", false, "Find B&W Devices")
	color := flag.Bool("color", false, "Find Color Devices")
	office := flag.Bool("office", false, "Find Office MFPs")
	pro := flag.Bool("pro", false, "Find Pro/Press MFPs")
	press := flag.Bool("press", false, "Find Pro/Press MFPs")

	flag.Parse()

	if 1 == len(os.Args) {
		*black, *color, *office, *pro, *press = true, true, true, true, true
	}

	cnt := 0
	outChan := make(chan *Device)
	doneChan := make(chan bool)

	if *black || *office {
		go Scrape("http://www.konicaminolta-images.eu/images/list/category/Copier%20Print%20Systems__Multifunctional%20Systems%20Black%20and%20White/?tx_kmmediapool_pi1[itemsperpage]=1000", outChan, doneChan)
		cnt++
	}

	if *color || *office {
		go Scrape("http://www.konicaminolta-images.eu/images/list/category/Copier%20Print%20Systems__Multifunctional%20Systems%20Colour/?tx_kmmediapool_pi1[itemsperpage]=1000", outChan, doneChan)
		cnt++
	}

	if *black || *pro || *press {
		go Scrape("http://www.konicaminolta-images.eu/images/list/category/Production%20Printing%20Systems__Production%20Printing%20Black%20and%20White/?tx_kmmediapool_pi1[itemsperpage]=1000", outChan, doneChan)
		cnt++
	}

	if *color || *pro || *press {
		go Scrape("http://www.konicaminolta-images.eu/images/list/category/Production%20Printing%20Systems__Production%20Printing%20Colour/?tx_kmmediapool_pi1[itemsperpage]=1000", outChan, doneChan)
		cnt++
	}

	for {
		select {
		case device := <-outChan:
			fmt.Println(device)
			break
		case <-doneChan:
			cnt--
			break
		}

		if 0 == cnt {
			break
		}
	}
}

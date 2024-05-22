package main

import (
   "encoding/json"
   "fmt"
   "log"
   "os"

   "strconv"

   "github.com/gocolly/colly"
)

type Product struct {
   Name     string
   Image    string
   Price    string
}

const targetURL = "https://books.toscrape.com/"

func main() {
   currentPage := 1


   c := colly.NewCollector()
   products := make([]Product, 0)


   // Callback Functions

   // Execute before visit - 1
   c.OnRequest(func(r *colly.Request) {
       fmt.Println("Visiting", r.URL)
   })

   // scrape logic - 3
   c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
	   item := Product{}
	   item.Name = e.ChildAttr(".thumbnail", "alt")
	   src := e.ChildAttr(".thumbnail", "src")
	   item.Image = targetURL + src
	   item.Price = e.ChildText(".price_color")
	   products = append(products, item)
   })

   // if err - #
   c.OnError(func(r *colly.Response, e error) {
       fmt.Println("Got this error:", e)
   })

   // after scrape done - 4
   c.OnScraped(func(r *colly.Response) {

   	   if currentPage == 10 {

	       fmt.Println("Finished", r.Request.URL)
	       js, err := json.MarshalIndent(products, "", "    ")
	       if err != nil {
	           log.Fatal(err)
	       }
	       fmt.Println("Writing data to file")
	       if err := os.WriteFile("products.json", js, 0664); err == nil {
	           fmt.Println("Data written to file successfully")
	       }

       } else {
       	   currentPage += 1
           cnvtd := strconv.Itoa(currentPage)
           c.Visit(targetURL + "catalogue/page-" + cnvtd + ".html")
       }



   })
   // visit - 2
   c.Visit(targetURL)
}

// task took avg 7s

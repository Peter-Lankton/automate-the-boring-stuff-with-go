package main

import (
	"encoding/csv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"os"
)

// define a custom data type for the scraped data
type Product struct {
	name, price string
}

func main() {
	// where to store the scraped data
	var products []Product

	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver/chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// visit the target page
	err = driver.Get("https://scrapingclub.com/exercise/list_infinite_scroll/")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// select the product elements
	productElements, err := driver.FindElements(selenium.ByCSSSelector, ".post")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// iterate over the product elements
	// and extract data from them
	for _, productElement := range productElements {
		// select the name and price nodes
		nameElement, err := productElement.FindElement(selenium.ByCSSSelector, "h4")
		priceElement, err := productElement.FindElement(selenium.ByCSSSelector, "h5")
		// extract the data of interest
		name, err := nameElement.Text()
		price, err := priceElement.Text()
		if err != nil {
			log.Fatal("Error:", err)
		}

		// add the scraped data to the list
		product := Product{}
		product.name = name
		product.price = price
		products = append(products, product)
	}

	// export the scraped data to CSV
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer file.Close()

	// initialize a file writer
	writer := csv.NewWriter(file)

	// define the CSV headers
	headers := []string{
		"name",
		"price",
	}

	// write the column headers
	writer.Write(headers)

	// adding each product to the CSV output file
	for _, product := range products {

		// converting a Product to an array of strings
		record := []string{
			product.name,
			product.price,
		}

		// writing a new CSV record
		writer.Write(record)
	}

	defer writer.Flush()
}

package main

import (
	"fmt"

	"github.com/gentil-eilison/tax-calculator/filemanager"
	"github.com/gentil-eilison/tax-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		// cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <- errorChans[index]:
			fmt.Println(err)
		
		case <- doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
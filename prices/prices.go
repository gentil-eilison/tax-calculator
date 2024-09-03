package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Couldn't open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Converting string to float failed.")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]string)
	job.LoadData()
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		key := fmt.Sprintf("%.2f", price)
		result[key] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}
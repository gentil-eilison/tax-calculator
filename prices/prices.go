package prices

import (
	"fmt"

	"github.com/gentil-eilison/tax-calculator/conversion"
	"github.com/gentil-eilison/tax-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager filemanager.FileManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
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

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
		IOManager: fm,
	}
}
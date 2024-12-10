package helper

import (
	"encoding/csv"
	"fmt"
	"os"
	"tugas/model"

	"github.com/xuri/excelize/v2"
)

// Fungsi untuk membuat laporan dalam format Excel
func GenerateExcelReport(orders []model.Order, filePath string) error {
	f := excelize.NewFile()
	sheetName := "Orders"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}

	// Membuat header
	f.SetCellValue(sheetName, "A1", "OrderID")
	f.SetCellValue(sheetName, "B1", "Customer")
	f.SetCellValue(sheetName, "C1", "Amount")
	f.SetCellValue(sheetName, "D1", "OrderDate")

	// Membuat style untuk header
	headerStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4CAF50"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold:   true,
			Color:  "#FFFFFF",
			Family: "Calibri",
			Size:   11,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return err
	}

	// Menerapkan style ke header
	f.SetCellStyle(sheetName, "A1", "D1", headerStyle)

	// Mengisi data
	for i, order := range orders {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), order.OrderID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), order.Customer)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), order.Amount)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), order.OrderDate)
	}

	//set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Menyimpan file
	if err := f.SaveAs(filePath); err != nil {
		return err
	}

	return nil
}

// Fungsi untuk membuat laporan dalam format CSV
func GenerateCSVReport(orders []model.Order, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Menulis header
	writer.Write([]string{"OrderID", "Customer", "Amount", "OrderDate"})

	// Menulis data
	for _, order := range orders {
		record := []string{
			order.OrderID,
			order.Customer,
			fmt.Sprintf("%.2f", order.Amount),
			order.OrderDate,
		}
		writer.Write(record)
	}

	return nil
}

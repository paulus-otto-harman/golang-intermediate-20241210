package service

import (
	"fmt"
	"path/filepath"
	"time"
	"tugas/helper"
	"tugas/repository"
)

func ReportOrderExcel() {
	orders := repository.GetOrders() // Dapatkan data order (di sini hanya data dummy)

	// Tentukan folder penyimpanan
	reportFolder := "report"
	if err := helper.EnsureFolder(reportFolder); err != nil {
		fmt.Println("Gagal membuat folder laporan:", err)
		return
	}

	// Buat laporan Excel
	excelPath := filepath.Join(reportFolder, fmt.Sprintf("orders_report_%s.xlsx", time.Now().Format("20060102")))
	if err := helper.GenerateExcelReport(orders, excelPath); err != nil {
		fmt.Println("Gagal membuat laporan Excel:", err)
	}
}

func ReportOrderCSV() {
	orders := repository.GetOrders() // Dapatkan data order (di sini hanya data dummy)

	// Tentukan folder penyimpanan
	reportFolder := "report"
	if err := helper.EnsureFolder(reportFolder); err != nil {
		fmt.Println("Gagal membuat folder laporan:", err)
		return
	}

	// Buat laporan CSV
	csvPath := filepath.Join(reportFolder, fmt.Sprintf("orders_report_%s.csv", time.Now().Format("20060102")))
	if err := helper.GenerateCSVReport(orders, csvPath); err != nil {
		fmt.Println("Gagal membuat laporan CSV:", err)
	} else {
		fmt.Println("Laporan CSV berhasil dibuat:", csvPath)
	}
}

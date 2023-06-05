package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	FilePath      string   `env:"FILE_PATH" envDefault:"E:\\BaiduDiskDownload\\企业年报15-22\\"`
	OutPutName    string   `env:"OUTPUT_NAME" envDefault:"财报数据.xlsx"`
	UnitList      []string `env:"UNIT_LIST" envDefault:"元,千元,万元"`
	CostCondition []string `env:"COST_CONDITION" envDefault:"研发费用加计扣除,研究开发费加成扣除的纳税影响,可加计扣除费用的影响,研发加计扣除影响,额外可扣除费用的影响"`
}

var cfg = config{}

var newExcel ExcelFile

func init() {
	flag.Parse()
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err)
	}
	newExcel.f = NewExcelizeFile()
}

func main() {
	for row, rowData := range TxtHelper() {
		newExcel.f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row+2), rowData.Info[0])
		newExcel.f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row+2), rowData.Info[1])
		newExcel.f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row+2), rowData.Info[2])
		newExcel.f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row+2), rowData.Info[3])
	}
	if err := newExcel.f.SaveAs(cfg.OutPutName); err != nil {
		panic(err)
	}
	fmt.Println("done")
}

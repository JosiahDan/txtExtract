package main

import (
	"github.com/xuri/excelize/v2"
)

type ExcelFile struct {
	f *excelize.File
}

func NewExcelizeFile() *excelize.File {
	return excelize.NewFile()
}

func (e *ExcelFile) SetCell(sheet string, cell string, value string) {
	if err := e.f.SetCellValue(sheet, cell, value); err != nil {
		panic(err)
	}
}

func (e *ExcelFile) CreateNewSheet(name string) int {
	index, err := e.f.NewSheet(name)
	if err != nil {
		panic(err)
	}

	return index
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func TxtHelper() []BusInfo {
	var fileList []string
	var busInfoList []BusInfo

	DateList := GetDateList(cfg.FilePath)

	for _, date := range DateList {
		fmt.Println("开始采集" + date + ".....")
		fileList, error := GetAllFile(cfg.FilePath+"\\"+date+"\\"+"TXT", fileList)
		if error != nil {
			panic(error)
		}
		bar := progressbar.Default(int64(len(fileList)))
		for _, file := range fileList {
			if strings.Contains(file, "已取消") {
				continue
			}
			f, err := os.Open(file)
			if err != nil {
				panic(err)
			}
			var newBusInfo BusInfo
			reader := bufio.NewReader(f)

			// 获取股票代码
			newBusInfo.GetStockCode(file)
			newBusInfo.GetReportDate(date)
			newBusInfo.GetUnit(reader)
			newBusInfo.GetDeductions(reader)

			if newBusInfo.Info[3] != "暂无数据" {
				busInfoList = append(busInfoList, newBusInfo)
			}

			f.Close()
			bar.Add(1)
		}
	}

	return busInfoList
}

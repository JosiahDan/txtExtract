package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/duke-git/lancet/v2/strutil"
)

type BusInfo struct {
	Info [4]string
}

// 获取股票代码
func (b *BusInfo) GetStockCode(file string) {
	fileName := strutil.After(file, "/")
	stockCode := fileName[:6]
	b.Info[0] = stockCode
}

// 获取财报年份
func (b *BusInfo) GetReportDate(date string) {

	b.Info[1] = date
}

// 获取会计利润与所得税费用调整过程的单位
func (b *BusInfo) GetUnit(reader *bufio.Reader) {

	// 逐行读取文件内容
	parseFlag := -1

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}

		if strings.Contains(line, "会计利润与所得税费用调整过程") {
			parseFlag = 1
		}
		if parseFlag > 0 && parseFlag <= 7 {
			if strings.Contains(line, "单位") {
				for _, unit := range cfg.UnitList {
					if strings.Contains(line, unit) {
						b.Info[2] = unit
						return
					}
				}
				parseFlag++
			}
		} else if parseFlag > 7 {
			b.Info[2] = "元"
			return
		}
	}
	b.Info[2] = "元"
}

func (b *BusInfo) GetDeductions(reader *bufio.Reader) {
	// 逐行读取文件内容

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}

		for _, condition := range cfg.CostCondition {
			if strings.Contains(line, condition) {
				tempDeductions := strutil.After(line, "-")
				re := regexp.MustCompile(`[\d,\.]+`) // 正则表达式，匹配数字、逗号和句点

				matches := re.FindAllString(tempDeductions, -1) // 提取匹配的内容
				var res string
				for i := 0; i < len(matches); i++ {
					res += matches[i]
				}

				b.Info[3] = res
				return

			}
		}
	}
	b.Info[3] = "暂无数据"
}

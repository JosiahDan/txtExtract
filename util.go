package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

func GetDateList(dirPath string) []string {
	var dirList []string
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		// 排除当前文件夹路径
		if path != dirPath {
			// 获取文件或文件夹名称
			name := info.Name()

			// 判断是否为文件夹
			if info.IsDir() {
				if name != "TXT" {
					dirList = append(dirList, name)
				}
			}
		}
		return nil
	})
	return dirList
}

package utils

import (
	"os"
)

func Mkdir(path string) {
	// 判断文件夹是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		// 其他错误
		panic(err)
	}
}

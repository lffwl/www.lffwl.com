package utils

import (
	"os"
)

// CreateMultiDir 调用os.MkdirAll递归创建文件夹
func CreateMultiDir(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, 0755)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// isExist 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

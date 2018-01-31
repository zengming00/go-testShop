package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

type DateDir struct {
	Dir      string
	Fullpath string
}

/**
 * 获取日期文件夹（不存在则自动创建）
 */
func GetDateDir() (*DateDir, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	var date = GetCnDate()
	var dir = fmt.Sprintf("/public/uploads/%.4d/%.2d/%.2d/", date.Year(), int(date.Month()), date.Day())
	var path = filepath.Join(baseDir, dir)

	err = os.MkdirAll(path, 0666)
	if err != nil {
		return nil, err
	}
	return &DateDir{Dir: dir, Fullpath: path}, nil
}

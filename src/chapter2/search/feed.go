package search

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const dataFile = "src/chapter2/data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	//open file
	filePath, _ := filepath.Abs(dataFile)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	//when func return close file
	defer file.Close()

	//将文件解码到一个slice里
	//这个slice的每一项是一个指向一个Feed类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	//这个函数不需要检查错误， 调用者会做
	return feeds, err
}

package pinyin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type WordData struct {
	Word string `json:"word"`
	Rate int    `json:"rate"`
	Type string `json:"type"`
}

type RecognizeResult struct {
	Pinyin string
	Result []WordData
}

func getData() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory:", err)
		return
	}
	absolutePath := filepath.Join(currentDir, "data.json")
	// 读取 JSON 文件
	file, err := os.Open(absolutePath)
	if err != nil {
		fmt.Println("Failed to open JSON file:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read JSON file:", err)
		return
	}

	// 解析 JSON 数据到 map 中
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}
}

var jsonData map[string][]WordData

func Recognize(pinyinList []string) []RecognizeResult {
	if jsonData == nil {
		getData()
	}
	result := make([]RecognizeResult, len(pinyinList))
	for index, pinyin := range pinyinList {
		if jsonData[pinyin] != nil {
			result[index] = RecognizeResult{pinyin, jsonData[pinyin]}
		} else {
			result[index] = RecognizeResult{pinyin, make([]WordData, 0)}
		}
	}
	return result
}

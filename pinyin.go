package pinyin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
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

func GetCurrentFilePath() (string, error) {
	// 获取调用方的文件路径
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("Failed to get caller information")
	}

	// 提取文件所在的目录路径
	dirPath := filepath.Dir(callerFile)
	return dirPath, nil
}

func getData() {
	dirPath, err := GetCurrentFilePath()
	if err != nil {
		fmt.Println("Failed to get current file path:", err)
		return
	}
	absolutePath := filepath.Join(dirPath, "data.json")
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

func RecognizeDetail(pinyinList []string) []RecognizeResult {
	if jsonData == nil {
		getData()
	}
	result := make([]RecognizeResult, 0)
	for _, pinyin := range pinyinList {
		if jsonData[pinyin] != nil {
			result = append(result, RecognizeResult{pinyin, jsonData[pinyin]})
		}
	}
	return result
}

func RecognizePinyin(pinyinList []string) []string {
	if jsonData == nil {
		getData()
	}
	result := make([]string, 0)
	for _, pinyin := range pinyinList {
		if jsonData[pinyin] != nil {
			result = append(result, pinyin)
		}
	}
	return result
}

func RecognizeNonPinyin(pinyinList []string) []string {
	if jsonData == nil {
		getData()
	}
	result := make([]string, 0)
	for _, pinyin := range pinyinList {
		if jsonData[pinyin] == nil {
			result = append(result, pinyin)
		}
	}
	return result
}

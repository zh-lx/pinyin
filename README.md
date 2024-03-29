# pinyin

识别输入的英文字符串是否为拼音

## 安装

```shell
go get -u github.com/zh-lx/pinyin
```

## 使用

```go
package main

import (
	"fmt"
	"github.com/zh-lx/pinyin"
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

func main() {
	// 待识别的拼音数组
	pinyinList := []string{"yisi", "baohupinghe", "notpinyinaaa"}
	var result RecognizeResult[]
	result= pinyin.Recognize(pinyinList)
}
```
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


func main() {
	// 待识别的拼音数组
	pinyinList := []string{"yisi", "baohupinghe", "notpinyinaaa"}
	var result []pinyin.RecognizeResult
	result= pinyin.Recognize(pinyinList)
	
	fmt.Println(result)
	// [{Pinyin:yisi Result:[{Word:意思 Rate:6089 Type:n} {Word:一丝 Rate:1186 Type:m} {Word:疑似 Rate:36 Type:v} {Word:伊斯 Rate:15 Type:ns} {Word:乙巳 Rate:8 Type:m} {四 Rate:5 Type:m} {Word:一似 Rate:3 Type:a} {Word:一俟 Rate:3 Type:nrt} {Word:缢死 Rate:3 Type:n}]} {Pinyin:baohupinghe Result:[{Word:暴虎冯河 Rate:3 Type:nr} {Woe:3 Type:nr}]} {Pinyin:notpinyinaaa Result:[]}]419 <nil>
}
```
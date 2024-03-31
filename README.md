# pinyin

识别输入的英文字符串是否为拼音

## 安装

```shell
go get -u github.com/zh-lx/pinyin
```

## 使用

### 返回全部结果
```go
package main

import (
	"fmt"
	"github.com/zh-lx/pinyin"
)


func main() {
	// 待识别的拼音数组
	pinyinList := []string{"yisi", "notpinyinaaa"}
	var result []pinyin.RecognizeResult
	result= pinyin.Recognize(pinyinList)
	
	fmt.Println(result)
	// [{Pinyin:yisi Result:[{Word:意思 Rate:6089 Type:n} {Word:一丝 Rate:1186 Type:m}]} {Pinyin:notpinyinaaa Result:[]}]118 <nil>
}
```

### 返回是拼音的结果
```go
package main

import (
	"fmt"
	"github.com/zh-lx/pinyin"
)


func main() {
	// 待识别的拼音数组
	pinyinList := []string{"yisi", "notpinyinaaa"}
	var result []pinyin.RecognizeResult
	result= pinyin.Recognize(pinyinList)
	
	fmt.Println(result)
    // [{Pinyin:yisi Result:[{Word:意思 Rate:6089 Type:n} {Word:一丝 Rate:1186 Type:m}]}]86 <nil>
}
```

### 返回是拼音的结果(仅英文字母)
```go
package main

import (
	"fmt"
	"github.com/zh-lx/pinyin"
)


func main() {
	// 待识别的拼音数组
	pinyinList := []string{"yisi", "notpinyinaaa"}
	var result []pinyin.RecognizeResult
	result= pinyin.RecognizeDetail(pinyinList)
	
	fmt.Println(result)
	// [yisi]6 <nil>
}
```

### 返回是非拼音的结果(仅英文字母)
```go
package main

import (
	"fmt"
	"github.com/zh-lx/pinyin"
)


func main() {
	// 待识别的拼音数组
	pinyinList := []string{"yisi", "notpinyinaaa"}
	var result []pinyin.RecognizeResult
	result= pinyin.Recognize(pinyinList)
	
	fmt.Println(result)
	// [notpinyinaaa]14 <nil>
}
```
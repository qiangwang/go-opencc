# go-opencc

封装 libopencc 库，提供简单的 golang 繁简转换接口

## 使用说明

1. 依赖 libopencc，详见 Dockerfile
2. 接口调用方法详见测试用例
3. 欢迎提 Issue 或 PR

## Example

```go
package main

import (
	"fmt"

	"github.com/qiangwang/go-opencc"
)

func main() {
	converter := opencc.Converter{}
	defer converter.Close()

	text, err := converter.Convert("開放中文轉換", "t2s")
	fmt.Println(text, err)
}
```
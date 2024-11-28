// File Name: recovery
// Description: recovery
//
// Created on: 2024/11/28
// Author: Tiger

package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(message string) string {
	// 获取调用栈信息 uintptr 是一个无符号整数，它的长度是32位或者64位，取决于操作系统
	var pcs [32]uintptr
	// 这里的 3 表示跳过前三个调用栈，因为前三个调用栈都是一些内部函数，我们不关心
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc) // 获取对应的函数
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}

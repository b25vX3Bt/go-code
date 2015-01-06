package main

/*
type Flag struct {
    Name     string // コマンドラインで使われる名前
    Usage    string // ヘルプメッセージ
    Value    Value  // セットされた値
    DefValue string // デフォルト値(文字列で)。これはヘルプメッセージで使用される
}
*/

import (
	"flag"
	"fmt"
)

func main() {

	var (
		//name, default-value, Usage
		//return: pointer
		debug = flag.String("debug", "off", "Usage: debug on/off")
	)

	var port int
	flag.IntVar(&port, "port", 8080, "Listen Port") //
	flag.IntVar(&port, "p", 8080, "Listen Port")    // 省略

	flag.Parse()

	fmt.Printf("arg count[%d]\n", flag.NArg())

	fmt.Printf("debug:[%s]\n", *debug)
	fmt.Printf("port:[%d]\n", port)
}

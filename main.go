package main

import (
	doudou "doudouCheckin/src"
	"flag"
	"strings"
)

func main() {
	email := flag.String("u", "", "your email")
	pwd := flag.String("p", "", "pwd")
	flag.Parse()
	if !strings.Contains(*email, "@") {
		panic("请输入正确的邮箱地址")
	}
	if *pwd == "" {
		panic("请输入密码")
	}
	cookies, err := doudou.Login(*email, *pwd)
	if err != nil {
		panic(err)
	}
	err = doudou.CheckIn(cookies)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"city/boss/parse"
	"city/engine"
)

func main() {

	engine.Run(engine.Request{
		Url:       "https://www.zhipin.com/",
		ParseFunc: parse.ParseCityList,
	})

}

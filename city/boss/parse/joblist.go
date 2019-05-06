package parse

import (
	"city/engine"
	"regexp"
)

const cityList = `<a ka=".+" href="(/.+/)">([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityList)
	match := re.FindAllSubmatch(contents, -1)
	request := engine.ParseResult{}
	for _, m := range match {
		//fmt.Printf("%s", m[1]) //wangzhi
		name := string(m[2])
		request.Items = append(request.Items, "joblist:"+name) //job
		request.Requests = append(request.Requests, engine.Request{
			Url: "https://www.zhipin.com" + string(m[1]), //url
			ParseFunc: func(c []byte) engine.ParseResult { //job
				return ParseJob(c)
			}, //job
			//                         //joblist
		})
		//fmt.Printf("job:%s ,URL:https://www.zhipin.com%s\n", m[2], m[1])
	}
	return request
}

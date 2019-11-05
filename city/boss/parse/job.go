package parse

import (
	"city/engine"
	"regexp"
)

const urlall = `<a href="(/job_detail/[\w|~]+[\w\W]*?)data-jid="[\w|~]+[\w\W]*?data-itemid="\d+"[\w\W]*?lid="[\w|.]+[\w\W]*?jobid="\d+"[\w\W]*?class="job-title">(.*)<[\w\W]*?class="red">([\w|-]+)</span>`

// const urlall = `<p>(.*?)<em class="vline"></em>(.*?)<em class="vline"></em>(.*?)</p>`
func ParseJob(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(urlall)
	match := re.FindAllSubmatch(contents, -1)
	request := engine.ParseResult{}
	for _, m := range match {
		//fmt.Printf("is urls:%s", m[1]) //url
		name := string(m[2])
		request.Items = append(request.Items, "job:"+name+"url:"+"https://www.zhipin.com"+string(m[1])+"money:"+string(m[3])) //jobname
		request.Requests = append(request.Requests, engine.Request{
			Url: "https://www.zhipin.com" + string(m[1]), //url
			ParseFunc: func(c []byte) engine.ParseResult { //jobs
				return ParseJobs(c)
			}, //
		})
		//fmt.Printf("jobname:%s ,nextURL:https://www.zhipin.com%s\n", m[2], m[1])
	}
	return requests

}

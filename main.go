package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	gjson "github.com/bitly/go-simplejson"
)

func main() {
	StartServer()
}

func StartServer() {

	for {

		tasks := LoadTasks("task")

		for _, task := range tasks {

			//	for each tasks
			go func(task *gjson.Json) {

				taskname, err := task.Get("name").String()
				Assert(err)
				log.Printf("Task[%v]...\n", taskname)

				//	Assert Array
				v, err := task.Get("list").Array()
				Assert(err)

				length := len(v)

				//	Last record
				resp := new(http.Response)

				//	for each request in per task
				for i := 0; i < length; i++ {

					jo := task.Get("list").GetIndex(i)

					url, _ := jo.Get("url").String()
					method, _ := jo.Get("method").String()
					data, _ := jo.Get("data").Map()

					//	url.values
					value := make(map[string][]string)

					for k, v := range data {
						vs := v.([]interface{})
						for i := range vs {
							value[k] = append(value[k], vs[i].(string))
						}
					}

					//	request
					req, err := http.NewRequest(method, url, nil)
					Assert(err)

					//	copy the last cookies
					for _, v := range resp.Cookies() {
						req.AddCookie(v)
					}

					//	temp resp
					var re *http.Response
					clt := new(http.Client)

					if method == "POST" {
						re, err = clt.PostForm(url, value)
					}

					if method == "GET" {
						re, err = clt.Do(req)
					}

					Assert(err)

					b, err := ioutil.ReadAll(re.Body)
					Assert(err)

					log.Printf("%s\n", b)

					resp = re

				}
			}(task)

		}
		time.Sleep(1 * time.Hour)
	}

}

func LoadTasks(pathname string) (ret []*gjson.Json) {

	fi, err := ioutil.ReadDir(pathname)
	Assert(err)

	for i := range fi {

		b, err := ioutil.ReadFile(path.Join(pathname, fi[i].Name()))
		Assert(err)

		jo, err := gjson.NewJson(b)
		Assert(err)

		ret = append(ret, jo)

	}

	return
}

func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

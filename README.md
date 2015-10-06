#Gosigner

auto sign in everyday.

每日签到

## Usage

- Download

    go get coding.net/kzzhr/gosigner
    
- Run It

    cd $GOPATH/src/gosigner
    nohup ./gosigner &
    

## Config

Read below according to [task/foo.com_demo.json](task/foo.com_demo.json)

请参考[task/foo.com_demo.json](task/foo.com_demo.json)阅读以下说明：


- `name` field is just use in logs
- `list` contains several `request`, these `request` will running in order, last response's `cookie` will insert as next request's `cookie`
- `data` design as a array is considering the golang programing


- `name`仅作为日志打印
- `list`中包含多个`request`，每个`request`依次执行，上一次执行得到的`cookie`会作为下次请求的`cookie`
- `data`之所以使用数组，只是为了在`go`中解析


[task/foo.com_demo.json](task/foo.com_demo.json)

```
{
	"name": "foo.com_demouser",
	"list": [
		{
			"url": "http://foo.com/login",
			"method": "POST",
			"data": {
				"username": ["your_username"],
				"password": ["your_password"]
			}
		}, 
		{
			"url": "http://foo.com/check_in.html",
			"method": "GET"
		}
	]
}

```

## About

Official Website: https://coding.net/u/kzzhr/p/gosigner/git

Author: https://coding.net/u/kzzhr

package main

import (
	"api/apis"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// 路由定义
type routeInfo struct {
	pattern string                                       // 正则表达式
	f       func(w http.ResponseWriter, r *http.Request) //Controller函数
}

// 路由添加
var routePath = []routeInfo{
	routeInfo{`^/trains?from=.*&to=.*$`, apis.GetTrains},
}

// 使用正则路由转发
func Route(w http.ResponseWriter, r *http.Request) {
	isFound := false
	for _, p := range routePath {
		// 这里循环匹配Path，先添加的先匹配
		reg, err := regexp.Compile(p.pattern)
		if err != nil {
			continue
		}
		if reg.MatchString(r.URL.Path) {
			isFound = true
			p.f(w, r)
		}
	}
	if !isFound {
		// 未匹配到路由
		fmt.Fprint(w, "404 Page Not Found!")
	}
}

func main() {
	defer apis.Shutdown()
	sever := http.NewServeMux()
	sever.HandleFunc("/fromscode", apis.GetPathFromStationCode)
	log.Fatalln(http.ListenAndServe(":8080", sever))
	// if api, err := apis.NewTrains(); err != nil {
	// 	log.Panicln(err.Error())
	// } else {
	// 	if p, err := api.GetTrainsFromStationCode("D1"); err != nil {
	// 		fmt.Println(err.Error())
	// 	} else {
	// 		fmt.Println(apis.Trains2AmapPathSimplifier(p))
	// 	}
	// }

}

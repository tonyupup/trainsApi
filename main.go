package main

import (
	"api/apis"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

// 路由定义
type routeInfo struct {
	pattern string                                       // 正则表达式
	f       func(w http.ResponseWriter, r *http.Request) //Controller函数
}

// 路由添加
var routePath = []routeInfo{
	routeInfo{`^/trains`, apis.GetTrains},
	routeInfo{`^/fromscode`, apis.GetPathFromStationCode},
}

// 使用正则路由转发
func Route(w http.ResponseWriter, r *http.Request) {
	isFound := false
	for _, p := range routePath {
		// 这里循环匹配Path，先添加的先匹配
		reg := regexp.MustCompile(p.pattern)
		// if err != nil {
		// 	continue
		// }
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
	sever.HandleFunc("/", Route)
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Start server :", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), sever))
	// if api, err := apis.NewTrains(); err != nil {
	// 	log.Panicln(err.Error())
	// } else {
	// 	if p, err := api.GetTrainsFromStationCode("D1"); err != nil {
	// 		fmt.Println(err.Error())
	// 	} else {
	// 		fmt.Println(apis.Trains2AmapPathSimplifier(p))
	// 	}
	// }
	// if p, err := apis.Tains.GetTrainsFromAddress("大同", "清河"); err != nil {
	// 	print(err.Error())
	// } else {
	// 	print(p)
	// }

}

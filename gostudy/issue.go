package main

import (
	"fmt"
	"net/url"
)

func main() {
    u := "http://www.baidu.com"
	urlo, _ := url.Parse(u)
	//urlo.Path,_ = url.QueryUnescape(urlo.Path)
	tt := fmt.Sprintf("%s://%s%s?%s", urlo.Scheme, urlo.Host, urlo.Path, urlo.RawQuery)
	fmt.Printf("url:%s",tt)

	//object := "php-oss-sdk/tests/129715.csv"
	//object = url.QueryEscape(object)
	//fmt.Printf("%s\n",object)
	//s,_:= url.QueryUnescape(object)
	//fmt.Printf("%s", s)
	//result, err := github.SearchIssues(os.Args[1:])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n",
	//		item.Number, item.User.Login, item.Title)
	//}
}

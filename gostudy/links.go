package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	_, err = html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	// visitNode := func(n *html.Node) {
	// 	if n.Type == html.ElementNode && n.Data == "a" {
	// 		for _, a := range n.Attr {
	// 			if a.Key != "href" {
	// 				continue
	// 			}
	// 			link, err := resp.Request.URL.Parse(a.Val)
	// 			if err != nil {
	// 				continue
	// 			}
	// 			links = append(links, link.String())
	// 		}
	// 	}
	// }
	// forEachNode(doc, visitNode, nil)
	return links, nil
}

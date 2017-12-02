package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/html"
)

//Extract xxx
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	var wg sync.WaitGroup
	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("=======开启 %d 协程\n", i)
			for linkx := range unseenLinks { //阻塞
				fmt.Printf("协程 %d 未阻塞正在执行！\n", i)
				foundLinks := crawl(linkx)
				go func() { fmt.Println("xxxxx"); worklist <- foundLinks }()
			}
			//fmt.Println("=======协程 %d 关闭\n", i)
		}(i)
	}

	go func() {
		wg.Wait()
		close(unseenLinks)
	}()

	seen := make(map[string]bool)
	for list := range worklist { //阻塞
		if len(list) == 0 {
			wg.Done()
			continue
		}
		fmt.Println("========================")
		for _, linkx := range list {
			if !seen[linkx] {
				seen[linkx] = true
				fmt.Printf("href------:%s\n", linkx)

				unseenLinks <- linkx
			}
		}
		fmt.Println("list loop finished.")
	}
}

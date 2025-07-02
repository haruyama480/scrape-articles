// kubernetes-doc/main.go
// This file is part of the scrape-articles project.
//
// run:
//
//	go run main.go > kubernetes-doc.md
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// headless ブラウザを起動
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	// タイムアウト設定
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// ファイルパスを取得
	url := "https://kubernetes.io/docs/home/"
	xpath := "//nav[@id=\"td-section-nav\"]//a"

	// タスク実行
	var linksAndDepths []map[string]interface{}
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Evaluate(fmt.Sprintf(`
            let results = [];
            let nodes = document.evaluate('%s', document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null);

            for (let i = 0; i < nodes.snapshotLength; i++) {
                const link = nodes.snapshotItem(i);
                let depth = 0;
                let parent = link.parentNode;
                while (parent) {
                    if (parent.tagName === 'UL') {
                         depth++;
                    }
                    parent = parent.parentNode;
                }
                results.push({
                    href: link.href,
                    text: link.innerText,
                    depth: depth
                });
            }
            results;
		`, xpath), &linksAndDepths),
	)

	if err != nil {
		log.Fatal(err)
	}

	// // fmt.Println("Extracted Links and Depths:")
	// for _, item := range linksAndDepths {
	// 	// fmt.Printf("  Href: %s, Text: %s, Depth: %d\n", item["href"], item["text"], int(item["depth"].(float64)))
	// 	fmt.Printf("%s, %s\n", item["href"], item["text"])
	// }

	minIndent := math.MaxInt
	for _, item := range linksAndDepths {
		minIndent = min(minIndent, int(item["depth"].(float64)))
	}

	var buffer strings.Builder
	buffer.WriteString("# Kubernetes Documentation Links\n\n")
	for _, item := range linksAndDepths {
		indent := strings.Repeat("  ", int(item["depth"].(float64))-minIndent)
		buffer.WriteString(fmt.Sprintf("%s- [%s](%s)\n", indent, item["text"], item["href"]))
	}
	fmt.Println(buffer.String())
}

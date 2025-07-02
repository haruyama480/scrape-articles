// run:
//
//	go run main.go > go-blog.md
package main

import (
	"context"
	"fmt"
	"log"
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
	url := "https://go.dev/blog/all"
	xpath1 := `//div[@id="blogindex"]//p[contains(@class,"blogtitle")]`
	xpath2 := `//div[@id="blogindex"]//p[contains(@class,"blogsummary")]`

	// タスク実行
	var linksAndDepths []map[string]interface{}
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Evaluate(fmt.Sprintf(`
            let results = [];
            let posts = document.evaluate('%s', document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null);
            let summarys = document.evaluate('%s', document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null);

            for (let i = 0; i < posts.snapshotLength; i++) {
                const item = posts.snapshotItem(i);
				const summary = summarys.snapshotItem(i);
                results.push({
					title: item.querySelector("a").innerText,
                    href: item.querySelector("a").href,
					date: item.querySelector("span.date").innerText,
					author: item.querySelector("span.author").innerText.trim(),
					summary: summary.innerText,
				});
            }
            results;
		`, xpath1, xpath2), &linksAndDepths),
	)

	if err != nil {
		log.Fatal(err)
	}

	// for _, item := range linksAndDepths {
	// 	fmt.Printf("%s, %s, %s, %s, %s\n", item["href"], item["title"], item["date"], item["author"], item["summary"])
	// }

	var buffer strings.Builder
	buffer.WriteString("# Go Blog Posts\n\n")
	for _, item := range linksAndDepths {
		date, err := dateConverter(item["date"].(string))
		if err != nil {
			log.Printf("Error converting date: %s, %s, %v", item["title"], item["date"], err)
			panic("")
		}
		buffer.WriteString(fmt.Sprintf("- [%s](%s), %s, %s\n", item["title"], item["href"], date, item["author"]))
		buffer.WriteString(fmt.Sprintf("  > %s\n", item["summary"]))
		buffer.WriteString("\n")
	}
	fmt.Println(buffer.String())
}

func dateConverter(dateStr string) (string, error) {
	t, err := time.Parse("2 January 2006", dateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse date: %w", err)
	}
	return t.Format("2006-01-02"), nil
}

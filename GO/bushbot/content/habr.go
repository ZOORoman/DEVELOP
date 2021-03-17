package content

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"news-bot/db"

	"github.com/PuerkitoBio/goquery"
)

func NabrGo(tag string) {
	URL := fmt.Sprintf("https://habr.com/ru/hub/%s/all/", tag)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Ошибка кода %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	linkAll := doc.Find(".post_preview").Find(".post_title")
	link, _ := linkAll.Find("a").Attr("href")
	linkText, _ := linkAll.Find("a").Html()
	fmt.Println(linkText.link)
	linkMD5Sum := md5.Sum([]byte(link))

	text := fmt.Sprintf(`<b>Habr - %s </b>: <a href\=\"%s\">%s</a>`, tag, link, linkText)
	fmt.Println(text)
	db.CheckSiteNewsBot(URL, link, text, fmt.Sprintf("%x", linkMD5Sum))
}

package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

// type Fact stcuct {
// 	ID int `json:"id"`
// 	Description string `json:"description"`
// }

func main() {
	c := colly.NewCollector()
	// c.OnHTML("#merItemThumbnail fluid__ac0ce62c radius__ac0ce62c", func(e *colly.HTMLElement) {
	// 	fmt.Println(e)
	// })

	c.OnHTML("li", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println(r.StatusCode)
	// })
	c.Visit("https://jp.mercari.com/")
}

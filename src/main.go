package extract

import (
	"appengine"
	"appengine/urlfetch"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"fmt"
)

type Apart struct {
	Reserved []byte `json:",string"`
	Data     []byte `json:",string"`
	UnitId   string
	Name     string
}

func init() {
	http.HandleFunc("/proxy", serveDistant)
	http.HandleFunc("/go/extract", process)
}

func extractFrom(doc *goquery.Document) string {
	result := ""

	// 1) Title of the product
	item1 := doc.Find(".product-name h1").First().Text()
	result += "Title: " + item1 + "\n"

	// 2) Price of the product
	item2 := doc.Find(".product-shop .price").First().Text()
	result += "Price: " + item2 + "\n"
	
	// 3) The brand
	// Although there is no distinct field "brand" in the page,
	// we can observe that the title always begins with what looks
	// like the Brand
	title := item1
	if title != "" {
		item3 := strings.Split(title, " ")[0]
		result += "Brand: " + item3 + "\n"
	}
	
	// 4) The manufacturers and the Etronics part numbers
	item41 := doc.Find("#product-attribute-specs-table>tbody>tr>th:contains(Manufacturer)").Next().Text()
	result += "Manufacturer: " + item41 + "\n"
	item42 := doc.Find("#product-attribute-specs-table>tbody>tr>th:contains(Model\\ Number)").Next().Text()
	result += "Number: " + item42 + "\n"
	
	// 5) Short description
	item5 := doc.Find(".short-description .std").First().Text()
	result += "Short description: " + item5 + "\n"
	
	// 6) The category structure
	item6 := doc.Find(".breadcrumbs li")
	cats := []string{}
	breadcrumb: for _, li := range item6.Nodes{
		for _, attr := range li.Attr {
			if attr.Key=="class" && !strings.Contains(attr.Val, "ategor") {
				continue breadcrumb
			}
		}
		child := li.FirstChild
		for child.Type != 3 && child.NextSibling != nil {
			child = child.NextSibling
		}
		txt := child.FirstChild 
		cat := txt.Data
		cats = append(cats, fmt.Sprintf("%v",cat) )
	}
	result += "Categories: " + strings.Join(cats, ", ") + "\n"

	return result
}

func process(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	url := r.FormValue("targetURL")

	resp, err := urlfetch.Client(c).Get(url)
	if err != nil {
		c.Errorf("%v", err)
		return
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		c.Errorf("%v", err)
		return
	}

	result := extractFrom(doc)
	w.Write([]byte(result))
}

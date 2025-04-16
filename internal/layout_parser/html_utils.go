package layout_parser

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func FindElementInHtml(htmlBody string, attributeKey string, attributeValue string) string {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return ""
	}
	result := ""
	var findElement func(*html.Node)
	findElement = func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, attr := range node.Attr {
				if attr.Key == attributeKey && strings.Contains(attr.Val, attributeValue) {
					result += findTextElement(node)
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			findElement(c)
		}
	}
	findElement(doc)
	return html.UnescapeString(result)
}

func findTextElement(node *html.Node) string {
	result := ""
	for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
		if nextNode.Type == html.TextNode {
			return nextNode.Data
		}
		result += findTextElement(nextNode)
	}
	return result
}

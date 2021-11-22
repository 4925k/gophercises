package parseWebPage

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link struct hold the urls and texts
type Link struct {
	Href string
	Text string
}

// Parse will an html page as io reader and return a list of links with error
func Parse(r io.Reader) ([]Link, error) {
	// parse reader into a html node
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	// get all links in the html page
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

// buildLink returns a link struct from the given link node
func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

// linkNodes takes and html node and
// returns a list of link nodes
// recursive function
func linkNodes(n *html.Node) []*html.Node {
	// if node is a link node return the node
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	//loop over nodes and take out all link nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

// text returns the text value from html node
func text(n *html.Node) string {
	// return node data if it a text node
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	// if nested nodes, connect all the strings and return
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}

	return strings.Join(strings.Fields(ret), "")
}

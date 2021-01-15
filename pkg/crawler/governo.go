package crawler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/leogr/the-colors-of-italy/pkg/types"
	"golang.org/x/net/html"
)

const GOVERNO_SOURCE = "http://www.governo.it/it/node/15638"

func Governo() (types.Regions, error) {
	out := make(types.Regions, 0, 21)

	resp, err := http.Get(GOVERNO_SOURCE)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "path" {
			reg := types.Region{}
			for _, a := range n.Attr {
				switch a.Key {
				case "id":
					reg.ID = a.Val
				case "fill":
					reg.Color = a.Val
				case "onclick":
					parts := strings.Split(a.Val, "'")
					if len(parts) < 3 {
						return
					}
					reg.Status = parts[1]
				}
			}

			if name, ok := types.RegionNamingMap[reg.ID]; ok {
				reg.Name = name
				reg.SourceURL = GOVERNO_SOURCE
				out = append(out, reg)
			}
			return
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	return out, nil
}

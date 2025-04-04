package services

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExpressVPN() (VpnStatus, error) {
	html, err := getHttpBytes(
		"https://www.expressvpn.com/what-is-my-ip",
	)
	if err != nil {
		return VpnStatus{}, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return VpnStatus{}, err
	}

	ipEl := doc.Find(".ip-address > span")

	ip := strings.TrimSpace(ipEl.Text())
	if ip == "" {
		ip = "Unknown"
	}

	status := ipEl.HasClass("green")

	locationHeadingEl := doc.Find("h6").FilterFunction(
		func(_ int, s *goquery.Selection) bool {
			return strings.TrimSpace(s.Text()) == "Location"
		},
	)

	locationEl := locationHeadingEl.Parent().Find("h4").First()

	location := strings.TrimSpace(
		strings.ReplaceAll(locationEl.Text(), "\n", " "),
	)
	if location == "" {
		location = "Unknown"
	}

	return VpnStatus{
		IP:       ip,
		Status:   status,
		Location: location,
		Name:     "ExpressVPN",
	}, nil
}

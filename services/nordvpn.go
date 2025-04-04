package services

import "fmt"

type NordVPNData struct {
	IP        string `json:"ip"`
	Country   string `json:"country"`
	Region    string `json:"region"`
	City      string `json:"city"`
	Protected bool   `json:"protected"`
}

func NordVPN() (VpnStatus, error) {
	var data NordVPNData
	err := getHttpJson(
		"https://web-api.nordvpn.com/v1/ips/info",
		&data,
	)
	if err != nil {
		return VpnStatus{}, err
	}

	return VpnStatus{
		IP:       data.IP,
		Status:   data.Protected,
		Location: fmt.Sprintf("%s, %s, %s", data.Country, data.Region, data.City),
		Name:     "NordVPN",
	}, nil
}

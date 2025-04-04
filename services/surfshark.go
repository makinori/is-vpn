package services

import "fmt"

type SurfsharkData struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
	Secured bool   `json:"secured"`
}

func Surfshark() (VpnStatus, error) {
	var data SurfsharkData
	err := getHttpJson(
		"https://surfshark.com/api/v1/server/user",
		&data,
	)
	if err != nil {
		return VpnStatus{}, err
	}

	return VpnStatus{
		IP:       data.IP,
		Status:   data.Secured,
		Location: fmt.Sprintf("%s, %s, %s", data.Country, data.Region, data.City),
		Name:     "Surfshark",
	}, nil
}

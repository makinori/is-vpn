package services

import "fmt"

type PrivateInternetAccessData struct {
	IP      string `json:"ip"`
	Country string `json:"cn"`
	City    string `json:"cty"`
	Region  string `json:"rgn"`
}

type PrivateInternetAccessExposed struct {
	Status bool `json:"status"`
}

func PrivateInternetAccess() (VpnStatus, error) {
	var data PrivateInternetAccessData
	err := getHttpJson(
		"https://www.privateinternetaccess.com/site-api/get-location-info",
		&data,
	)
	if err != nil {
		return VpnStatus{}, err
	}

	var exposed PrivateInternetAccessExposed
	err = postHttpJson(
		"https://www.privateinternetaccess.com/site-api/exposed-check",
		map[string]string{
			"ipAddress": data.IP,
		},
		&exposed,
	)
	if err != nil {
		return VpnStatus{}, err
	}

	return VpnStatus{
		IP:     data.IP,
		Status: !exposed.Status,
		Location: fmt.Sprintf(
			"%s, %s, %s", data.Country, data.Region, data.City,
		),
		Name: "Private Internet Access",
	}, nil
}

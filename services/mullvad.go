package services

import "fmt"

type MullvadData struct {
	IP             string `json:"ip"`
	Country        string `json:"country"`
	City           string `json:"city"`
	ExitIP         bool   `json:"mullvad_exit_ip"`
	ExitIPHostname string `json:"mullvad_exit_ip_hostname,omitempty"`
}

func Mullvad() (VpnStatus, error) {
	var data MullvadData
	err := getHttpJson(
		"https://ipv4.am.i.mullvad.net/json",
		&data,
	)
	if err != nil {
		return VpnStatus{}, err
	}

	location := fmt.Sprintf("%s, %s", data.Country, data.City)

	if data.ExitIPHostname != "" {
		location += " (" + data.ExitIPHostname + ")"
	}

	return VpnStatus{
		IP:       data.IP,
		Status:   data.ExitIP,
		Location: location,
		Name:     "Mullvad",
	}, nil
}

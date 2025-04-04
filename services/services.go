package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"maps"
	"net/http"
	"slices"
)

type VpnStatus struct {
	IP       string `json:"ip"`
	Status   bool   `json:"status"`
	Location string `json:"location"`
	Name     string `json:"name"`
}

func getHttpJson[T any](url string, out T) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return err
	}

	return nil
}

func getHttpBytes(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []byte{}, err
	}

	out, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

func postHttpJson[T any](url string, data map[string]string, out T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return err
	}

	return nil
}

var SERVICE_RESOLVE_MAP = map[string](func() (VpnStatus, error)){
	"expressvpn":            ExpressVPN,
	"mullvad":               Mullvad,
	"nordvpn":               NordVPN,
	"pia":                   PrivateInternetAccess,
	"privateinternetaccess": PrivateInternetAccess,
	"surfshark":             Surfshark,
}

var SERVICE_LIST = slices.Sorted(maps.Keys(SERVICE_RESOLVE_MAP))

func GetStatusResolveFunc(service string) (func() (VpnStatus, error), error) {
	serviceFunc, exists := SERVICE_RESOLVE_MAP[service]
	if !exists {
		return nil, errors.New("service not found")
	}

	return serviceFunc, nil
}

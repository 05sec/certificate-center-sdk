package license

import (
	"encoding/json"
	"errors"
)

func Decode(licenseStr string) (*Info, error) {
	if licenseStr == "" {
		return nil, errors.New("empty license")
	}

	var decodedLicense Info

	if json.Unmarshal([]byte(licenseStr), &decodedLicense) != nil {
		return nil, errors.New("failed to decode license")
	}

	return &decodedLicense, nil
}

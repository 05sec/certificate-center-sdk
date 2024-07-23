package license

import (
	"encoding/json"
	"errors"
	"github.com/05sec/certificate-center-sdk/pkg/types"
)

func Decode(licenseStr string) (*types.License, error) {
	if licenseStr == "" {
		return nil, errors.New("empty license")
	}

	var decodedLicense types.License

	if json.Unmarshal([]byte(licenseStr), &decodedLicense) != nil {
		return nil, errors.New("failed to decode license")
	}

	return &decodedLicense, nil
}

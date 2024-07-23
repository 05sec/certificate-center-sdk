package certificate_center_sdk

import (
	licenseV1 "github.com/05sec/certificate-center-sdk/pkg/proto/license/v1"
	"log"
	"testing"
)

func TestGetLicense(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	clinet := NewClient(&Config{
		BaseURL: "http://localhost:7716",
		ApiKey:  "123",
	})
	resp, err := clinet.GetLicense(&licenseV1.GetLicenseRequest{Code: "123"})
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("resp: %v", resp)
}

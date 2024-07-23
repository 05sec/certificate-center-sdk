package test

import (
	licenseV1 "github.com/05sec/certificate-center-sdk/gen/proto/license/v1"
	"github.com/05sec/certificate-center-sdk/pkg/client"
	"github.com/05sec/certificate-center-sdk/pkg/license"
	"log"
	"testing"
)

func TestGetLicense(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	clinet := client.NewClient(&client.Config{
		BaseURL: "http://localhost:7716",
		ApiKey:  "123",
	})
	//resp, err := clinet.GetLicense(&licenseV1.GetLicenseRequest{Code: "123"})
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//log.Printf("resp: %v", resp)

	resp2, err := clinet.ReadLicense(&licenseV1.ReadLicenseRequest{Code: "123"})
	if err != nil {
		t.Error(err)
		return
	}

	ddd, _ := license.Decode(resp2.LicenseInfo)

	log.Printf("resp: %v", resp2)
	log.Printf("ddd: %v", ddd)
}

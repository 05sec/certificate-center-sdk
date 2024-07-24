package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	licenseV1 "github.com/05sec/certificate-center-sdk/gen/proto/license/v1"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLicense(code string) (*licenseV1.GetLicenseResponse, error) {
	url := c.config.BaseURL + "/gapi/product/v1/license/get"

	request := &licenseV1.GetLicenseRequest{
		Code: code,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %v, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var license struct {
		License string `json:"license"`
	}
	err = json.Unmarshal(body, &license)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &licenseV1.GetLicenseResponse{
		License: license.License,
	}, nil
}

func (c *Client) ReadLicense(code string) (*licenseV1.ReadLicenseResponse, error) {
	url := c.config.BaseURL + "/gapi/product/v1/license/info"

	request := &licenseV1.ReadLicenseRequest{
		Code: code,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %v, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return &licenseV1.ReadLicenseResponse{LicenseInfo: string(body)}, nil
}

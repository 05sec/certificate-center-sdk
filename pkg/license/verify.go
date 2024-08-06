package license

import (
	"errors"
	"github.com/asjdf/machineid"
	"github.com/wuhan005/gadget"
	"os"
	"time"
)

// VerifyLicense verifies the license string.
// AppID is the application ID.
// licenseStr is the license string.
// pubKeyStr is the public key string.
// TODO: give client the pubKey
func VerifyLicense(AppID, licenseStr string, pubKeyStr string) error {
	machineID := GetMachineID(AppID)

	info, err := Read(licenseStr, pubKeyStr)
	if err != nil {
		return errors.New("failed to decode license")
	}

	if info.ExpiresAt.Before(time.Now()) || info.StartsAt.After(time.Now()) || info.MachineID != machineID {
		return errors.New("invalid license")
	}

	return nil
}

func GetMachineID(AppID string) string {
	machineID, err := machineid.ProtectedID(AppID)
	if err != nil {
		hostName, err := os.ReadFile("/etc/resolv.conf")
		if err != nil {
			return "none"
		}
		return gadget.Sha1(string(hostName))
	}
	return machineID
}

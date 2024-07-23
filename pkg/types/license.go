package types

import "time"

type License struct {
	LicenseInfo string    `json:"licenseInfo"`
	MachineCode string    `json:"m"`
	Tags        []string  `json:"t"`
	StartsAt    time.Time `json:"s"`
	ExpiresAt   time.Time `json:"e"`
}

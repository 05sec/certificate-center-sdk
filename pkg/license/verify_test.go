package license

import "testing"

func TestGetMachineID(t *testing.T) {
	machineID := GetMachineID("test")
	t.Logf("machineID: %s", machineID)
}

func TestVerifyLicense(t *testing.T) {
	licenseStr := "eyJzaWciOnsiRm9ybWF0Ijoic3NoLWVkMjU1MTkiLCJCbG9iIjoiVy9na25hSVBrSTBoakcxR1hWRVR1SkxnQ0xwQnQ4MkdKZGtiZ2NkM1hmaE15UnY3Q2NWVWUvVW5NQTh2ZkNZMXlPc2h5ZEFqWHpkaUZTNnZkWlorREE9PSIsIlJlc3QiOm51bGx9LCJpbmZvIjoiZXlKMklqb3hMQ0p1SWpwYk5ERXNNVFExTERJeU55d3hNU3czTml3eE5UQXNNVEUxTERRMFhTd2liU0k2SWprMll6QTJaVEZtWm1OaE1tWmtZekpqTXpFeE5XWmhPR014TXpNNFlqTTBOMkkyWmpRd1pHRTJaRE01TkRobVkyVXpObVZsWVRWbU9EVTJZMkpsWXpNaUxDSm1JanBiWFN3aWRDSTZXeUlpWFN3aWN5STZJakU1TnpBdE1ERXRNREZVTURBNk1EQTZNREFyTURnNk1EQWlMQ0psSWpvaU1qQXlOQzB3T0MweU9GUXdNRG93TURvd01Dc3dPRG93TUNKOSJ9"
	pubKeyStr := "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDOaK7ymIYlb6Yo3PWf/aGeez2ehmVUaUjMDU9NBtfpK"

	err := VerifyLicense("test", licenseStr, pubKeyStr)
	if err != nil {
		t.Error(err)
	}
}

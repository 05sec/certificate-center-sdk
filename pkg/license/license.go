package license

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zalando/go-keyring"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

// Info represents license information
type Info struct {
	MachineID string    `json:"m"`
	Features  []string  `json:"f"`
	Tags      []string  `json:"t"`
	StartsAt  time.Time `json:"s"`
	ExpiresAt time.Time `json:"e"`
}

// Nonce is a random number used to generate different
// signatures for the same license information
type encodeInfo struct {
	Version int     `json:"v"`
	Nonce   [8]byte `json:"n"`
	Info
}

type signedKey struct {
	Signature   *ssh.Signature `json:"sig"`
	EncodedInfo []byte         `json:"info"`
}

const formatVersion = 1

func Read(encodedLicense string, pubKey string) (*Info, error) {
	signedKeyData, err := base64.RawURLEncoding.DecodeString(encodedLicense)
	if err != nil {
		return nil, errors.Wrap(err, "decode license key")
	}

	var signedKey signedKey
	if err := json.Unmarshal(signedKeyData, &signedKey); err != nil {
		return nil, errors.Wrap(err, "unmarshal signed license key")
	}

	parsedPubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(pubKey))
	if err != nil {
		return nil, fmt.Errorf("parse public key: %w", err)
	}

	sshPubKey, ok := parsedPubKey.(ssh.PublicKey)
	if !ok {
		return nil, errors.New("invalid public key type")
	}

	if err := sshPubKey.Verify(signedKey.EncodedInfo, signedKey.Signature); err != nil {
		return nil, errors.Wrap(err, "verify license key signature")
	}

	var encodedInfo encodeInfo
	if err := json.Unmarshal(signedKey.EncodedInfo, &encodedInfo); err != nil {
		return nil, errors.Wrap(err, "unmarshal license info")
	}

	return &encodedInfo.Info, nil
}

// StoreLicense 将编码的许可证存储在密钥环中
func StoreLicense(service, username, encodedLicense string) error {
	return keyring.Set(service, username, encodedLicense)
}

// RetrieveLicense 从密钥环中检索编码的许可证
func RetrieveLicense(service, username string) (string, error) {
	encodedLicense, err := keyring.Get(service, username)
	if err != nil {
		return "", errors.Wrap(err, "retrieve license from keyring")
	}
	return encodedLicense, nil
}

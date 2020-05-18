package vendorconsent

import (
	"fmt"

	"github.com/prebid/go-gdpr/api"
	tcf1 "github.com/prebid/go-gdpr/vendorconsent/tcf1"
	tcf2 "github.com/prebid/go-gdpr/vendorconsent/tcf2"
)

var (
	errEmptyDecodedConsent = fmt.Errorf("decoded consent cannot be empty")
)

// ParseString parses a Raw (unpadded) base64 URL encoded string.
func ParseString(consent string) (api.VendorConsents, error) {
	if consent == "" {
		return nil, errEmptyDecodedConsent
	}

	if tcf2.IsConsentV2(consent) {
		return tcf2.ParseString(consent)
	}

	return tcf1.ParseString(consent)
}

// ParseVersion parses version from base64-decoded consent string
func ParseVersion(decodedConsent []byte) (uint8, error) {
	if len(decodedConsent) == 0 {
		return 0, errEmptyDecodedConsent
	}
	// read version from first 6 bits
	return decodedConsent[0] >> 2, nil
}

// Backwards compatibility

// VendorConsents old interface
type VendorConsents interface {
	api.VendorConsents
}

// Parse method is the equivalent of call tcf1.Parse(...)
func Parse(data []byte) (api.VendorConsents, error) {
	return tcf1.Parse(data)
}

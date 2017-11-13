package msisdn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func trim(s string) string {
	notWordRE := regexp.MustCompile(`[\W_]`)
	return notWordRE.ReplaceAllString(s, "")
}

func isMSISDN(s string) bool {
	if strings.HasPrefix(s, "1800") {
		return true
	}
	for _, country := range countries {
		if strings.HasPrefix(s, country.areaCode) {
			_, err := ParseMSISDN(s)
			if err != nil {
				return false
			}
			return true
		}
	}
	return false
}

func ParseMSISDN(msisdn string) (*MSISDN, error) {
	m := &MSISDN{}
	m.msisdn = trim(msisdn)
	m.countryCode = m.getCountryCode()
	if !m.validate() {
		return nil, fmt.Errorf("MSISDN is invalid")
	}
	m.provider = m.getProvider()
	m.landline = m.isLandLine()
	return m, nil
}

// phone: can be MSISDN or local number
// countryCode: needed only if `phone` is a local number, support "MY", "ID", "SG"
func Parse(phone string, countryCode string) (*MSISDN, error) {
	phone = trim(phone)
	if isMSISDN(phone) {
		return ParseMSISDN(phone)
	}
	country, ok := countries[countryCode]
	if !ok {
		return nil, fmt.Errorf("unsupported countryCode=%#v", countryCode)
	}
	if country.localPrefix != "" && strings.HasPrefix(phone, country.localPrefix) {
		phone = phone[len(country.localPrefix):]
	}
	msisdn := country.areaCode + phone
	return ParseMSISDN(msisdn)
}

//
type MSISDN struct {
	msisdn      string
	countryCode string
	landline    bool
	provider    string
}

//
func (m *MSISDN) GetMSISDN() string {
	return m.msisdn
}

//
func (m *MSISDN) GetProvider() string {
	return m.provider
}

//
func (m *MSISDN) GetCountryCode() string {
	return m.countryCode
}

func (m *MSISDN) GetAreaCode() string {
	country, ok := countries[m.countryCode]
	if !ok {
		return ""
	}
	return country.areaCode
}

//
func (m *MSISDN) IsLandLine() bool {
	return m.landline
}

//
func (m *MSISDN) GetLocal() string {
	msisdn := m.msisdn
	if strings.HasPrefix(msisdn, "1800") {
		return msisdn
	}
	country, ok := countries[m.countryCode]
	if !ok {
		return ""
	}
	return country.localPrefix + msisdn[len(country.areaCode):]
}

func (m *MSISDN) GetLocalFormatted() string {
	country, ok := countries[m.countryCode]
	if !ok {
		return ""
	}
	local := m.GetLocal()
	if local == "" {
		return ""
	}
	for _, format := range country.localFormats {
		local = format.find.ReplaceAllString(local, format.replace)
	}
	return local
}

func (m *MSISDN) getCountryCode() string {
	prefix := m.msisdn[:2]
	for cc, country := range countries {
		if prefix == country.areaCode {
			return cc
		}
	}
	return ""
}

func (m *MSISDN) getProvider() string {
	if m.countryCode == "" {
		return ""
	}

	subscriberNumber, _ := strconv.Atoi(m.msisdn)
	sn := int64(subscriberNumber)

	for _, p := range countries[m.countryCode].providers {
		if sn >= p.startRange && sn <= p.endRange {
			return p.name
		}
	}
	return ""
}

func (m *MSISDN) isLandLine() bool {
	if strings.HasPrefix(m.msisdn, "1800") {
		return true
	}
	// TODO: move this logic to `countries`
	switch m.countryCode {
	case "MY":
		if !strings.HasPrefix(m.msisdn, "601") {
			return true
		}
	case "SG":
		return strings.HasPrefix(m.msisdn, "656")
	default:
		return false
	}
	return false
}

func (m *MSISDN) validate() bool {
	// must be number and not empty
	_, err := strconv.Atoi(m.msisdn)
	if err != nil {
		return false
	}
	// must not exceed 15 digits
	if len(m.msisdn) > 15 {
		return false
	}
	country := countries[m.countryCode]
	if len(m.msisdn) < country.minLength || len(m.msisdn) > country.maxLength {
		return false
	}
	return true
}

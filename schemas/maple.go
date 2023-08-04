package schemas

import (
	"fmt"
	"strings"
)

type ErrEmptyFields struct {
	Fields []string `json:"fields"`
}

func (ef ErrEmptyFields) Error() string {
	return fmt.Sprintf("Empty fields not allowed: %s", strings.Join(ef.Fields, ", "))
}

type MetaData struct {
	SourceUID       string `json:"sourceUID,omitempty"`
	SourcePlatform  string `json:"sourcePlatform,omitempty"`
	EventOriginType string `json:"eventOriginType,omitempty"`
	SchemaVersion   string `json:"schemaVersion,omitempty"`
}
type SessionLookupData struct {
	CartToken         string   `json:"cartToken,omitempty"`
	SessionLookupKeys []string `json:"sessionLookupKeys,omitempty"`
	SessionIDFound    bool     `json:"sessionIdFound,omitempty"`
	SessionSyncMethod string   `json:"sessionSyncMethod,omitempty"`
}
type Page struct {
	ReferrerURL string `json:"referrerUrl,omitempty"`
	URLPath     string `json:"urlPath,omitempty"`
	URLTitle    string `json:"urlTitle,omitempty"`
	Host        string `json:"host,omitempty"`
	URL         string `json:"url,omitempty"`
}
type ContextData struct {
	SessionLookupData SessionLookupData `json:"sessionLookupData,omitempty"`
	Page              Page              `json:"page,omitempty"`
}
type EventTimeData struct {
	OriginalTimestamp  int `json:"originalTimestamp,omitempty"`
	SourceProccessedAt int `json:"sourceProccessedAt,omitempty"`
}
type Shopify struct {
	OrderID     int64  `json:"order_id,omitempty"`
	AppID       int    `json:"app_id,omitempty"`
	ShopifyCart string `json:"shopifyCart,omitempty"`
	ShopifyTags string `json:"shopifyTags,omitempty"`
}
type SourceCustomData struct {
	Shopify Shopify `json:"shopify,omitempty"`
}

type Address struct {
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Region      string `json:"region,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Street      string `json:"street,omitempty"`
}
type CustomerData struct {
	ID          string  `json:"id,omitempty"`
	Email       string  `json:"email,omitempty"`
	FirstName   string  `json:"firstName,omitempty"`
	LastName    string  `json:"lastName,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	DateOfBirth string  `json:"dateOfBirth,omitempty"`
	Address     Address `json:"address,omitempty"`
}
type Price struct {
	TotalPrice float64 `json:"totalPrice,omitempty"`
}
type Products struct {
	ID           string `json:"id,omitempty"`
	Brand        string `json:"brand,omitempty"`
	Category     string `json:"category,omitempty"`
	Name         string `json:"name,omitempty"`
	VariantID    string `json:"variantId,omitempty"`
	VariantTitle string `json:"variantTitle,omitempty"`
	Sku          string `json:"sku,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Price        Price  `json:"price,omitempty"`
}
type CartData struct {
	Currency   string     `json:"currency,omitempty"`
	Products   []Products `json:"products,omitempty"`
	TotalPrice float64    `json:"total_price,omitempty"`
}

type SessionData struct {
	SessionID      string         `json:"sessionId,omitempty"`
	AnonymousID    string         `json:"anonymousId,omitempty"`
	GdprData       GdprData       `json:"gdprData,omitempty"`
	BrowserData    BrowserData    `json:"browserData,omitempty"`
	CustomerData   CustomerData   `json:"customerData,omitempty"`
	CartData       CartData       `json:"cartData,omitempty"`
	IPLocationData IPLocationData `json:"ipLocationData,omitempty"`
}
type Etag struct {
	ID       int64 `json:"id,omitempty"`
	LastSeen int64 `json:"lastSeen,omitempty"`
	Session  int   `json:"session,omitempty"`
}
type Utm struct {
	UtmCampaign string `json:"utmCampaign,omitempty"`
	UtmContent  string `json:"utmContent,omitempty"`
	UtmMedium   string `json:"utmMedium,omitempty"`
	UtmSource   string `json:"utmSource,omitempty"`
}
type Cookies struct {
	Fbp   string `json:"_fbp,omitempty"`
	Ttp   string `json:"_ttp,omitempty"`
	Fbc   string `json:"_fbc,omitempty"`
	Ga    string `json:"_ga,omitempty"`
	Gid   string `json:"_gid,omitempty"`
	GclAu string `json:"_gcl_au,omitempty"`
	Clck  string `json:"_clck,omitempty"`
}
type Clid struct {
	Fbclid string `json:"fbclid,omitempty"`
	Gclid  string `json:"gclid,omitempty"`
	Ttclid string `json:"ttclid,omitempty"`
}
type BrowserData struct {
	UserAgent string  `json:"userAgent,omitempty"`
	IP        string  `json:"ip,omitempty"`
	Etag      Etag    `json:"etag,omitempty"`
	Utm       Utm     `json:"utm,omitempty"`
	Cookies   Cookies `json:"cookies,omitempty"`
	Clid      Clid    `json:"clid,omitempty"`
}
type GdprData struct {
	CustomerConsent string `json:"customerConsent,omitempty"`
	GdprAllowed     bool   `json:"gdprAllowed,omitempty"`
	GdprPlugin      string `json:"gdprPlugin,omitempty"`
	GdprConsent     bool   `json:"gdprConsent,omitempty"`
	GdprMode        string `json:"gdprMode,omitempty"`
}
type IPLocationData struct {
	IP            string `json:"ip,omitempty"`
	As            string `json:"as,omitempty"`
	City          string `json:"city,omitempty"`
	Continent     string `json:"continent,omitempty"`
	ContinentCode string `json:"continentCode,omitempty"`
	Country       string `json:"country,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
	Currency      string `json:"currency,omitempty"`
	District      string `json:"district,omitempty"`
	Isp           string `json:"isp,omitempty"`
	Proxy         bool   `json:"proxy,omitempty"`
	Region        string `json:"region,omitempty"`
	RegionName    string `json:"regionName,omitempty"`
	Timezone      string `json:"timezone,omitempty"`
	Zip           string `json:"zip,omitempty"`
	Org           string `json:"org,omitempty"`
}

type SessionData0 struct {
	SessionID      string         `json:"sessionId,omitempty"`
	AnonymousID    string         `json:"anonymousId,omitempty"`
	BrowserData    BrowserData    `json:"browserData,omitempty"`
	GdprData       GdprData       `json:"gdprData,omitempty"`
	CustomerData   CustomerData   `json:"customerData,omitempty"`
	IPLocationData IPLocationData `json:"ipLocationData,omitempty"`
}
type GaParams struct {
	Ga    string `json:"ga,omitempty"`
	Gid   string `json:"gid,omitempty"`
	GclAu string `json:"gclAu,omitempty"`
}
type MetaData0 struct {
	SourceID        int    `json:"sourceID,omitempty"`
	SchemaVersion   string `json:"schemaVersion,omitempty"`
	UserID          string `json:"userID,omitempty"`
	SourcePlatform  string `json:"sourcePlatform,omitempty"`
	EventOriginType string `json:"eventOriginType,omitempty"`
	SourceUID       string `json:"sourceUID,omitempty"`
}
type EventTimeData0 struct {
	OriginalTimestamp int `json:"originalTimestamp,omitempty"`
}
type SourceCustomData0 struct {
}
type Mclid struct {
	Ad        string `json:"ad,omitempty"`
	Adset     string `json:"adset,omitempty"`
	Campaign  string `json:"campaign,omitempty"`
	Raw       string `json:"raw,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

type CustomData struct {
	P                int     `json:"_p,omitempty"`
	S                string  `json:"_s,omitempty"`
	AdBlockerEnabled bool    `json:"ad_blocker_enabled,omitempty"`
	Clid             Clid    `json:"clid,omitempty"`
	Cookies          Cookies `json:"cookies,omitempty"`
	CustomerConsent  string  `json:"customer_consent,omitempty"`
	Dl               string  `json:"dl,omitempty"`
	Dr               string  `json:"dr,omitempty"`
	Dt               string  `json:"dt,omitempty"`
	GdprAllowed      bool    `json:"gdpr_allowed,omitempty"`
	GdprConsent      string  `json:"gdpr_consent,omitempty"`
	GdprMode         string  `json:"gdpr_mode,omitempty"`
	GdprPlugin       string  `json:"gdpr_plugin,omitempty"`
	Host             string  `json:"host,omitempty"`
	Name             string  `json:"name,omitempty"`
	Path             string  `json:"path,omitempty"`
	Referrer         string  `json:"referrer,omitempty"`
	Sct              string  `json:"sct,omitempty"`
	Search           string  `json:"search,omitempty"`
	Seg              string  `json:"seg,omitempty"`
	Sr               string  `json:"sr,omitempty"`
	Title            string  `json:"title,omitempty"`
	Uaa              string  `json:"uaa,omitempty"`
	Uab              string  `json:"uab,omitempty"`
	Uafvl            string  `json:"uafvl,omitempty"`
	Uamb             string  `json:"uamb,omitempty"`
	Ul               string  `json:"ul,omitempty"`
	URL              string  `json:"url,omitempty"`
	Utm              Utm     `json:"utm,omitempty"`
}
type Event struct {
	EventID          string           `json:"eventId,omitempty"`
	EventName        string           `json:"eventName,omitempty"`
	EventType        string           `json:"eventType,omitempty"`
	MetaData         MetaData         `json:"metaData,omitempty"`
	ContextData      ContextData      `json:"contextData,omitempty"`
	EventTimeData    EventTimeData    `json:"eventTimeData,omitempty"`
	SourceCustomData SourceCustomData `json:"SourceCustomData,omitempty"`
	SessionData      SessionData      `json:"sessionData,omitempty"`
	GaParams         GaParams         `json:"gaParams,omitempty"`
	CustomData       CustomData       `json:"customData,omitempty"`
}

type DestinationSettings struct {
	PixelID       string `json:"pixelId"`
	TestEventCode string `json:"testEventCode"`
	AccessKey     string `json:"accessKey"`
}

type FacebookEventPayload struct {
	DestinationSettings DestinationSettings `json:"destinationSettings"`
	Events              []Event             `json:"events"`
	//Events []map[string]string `json:"events"`
}

func (f FacebookEventPayload) Validate() error {
	// Check for DestinationSettings fields and make sure they are not empty
	if f.DestinationSettings.PixelID == "" || f.DestinationSettings.AccessKey == "" {
		return &ErrEmptyFields{
			Fields: []string{"pixelId", "accessKey"},
		}
	}

	return nil
}

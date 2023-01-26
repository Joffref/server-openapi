/*
 * 3gpp-nidd
 *
 * API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NIDD

import (
	"time"
)

// NiddConfiguration - Represents the configuration for NIDD.
type NiddConfiguration struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// Identifies the MTC Service Provider and/or MTC Application.
	MtcProviderId string `json:"mtcProviderId,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn string `json:"msisdn,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId string `json:"externalGroupId,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	Duration time.Time `json:"duration,omitempty"`

	// Indicates whether the reliable data service (as defined in clause 4.5.14.3 of 3GPP TS  23.682) acknowledgement is requested (true) or not (false). Default value is false. 
	ReliableDataService bool `json:"reliableDataService,omitempty"`

	// Indicates the static port configuration that is used for reliable data transfer between specific applications using RDS (as defined in clause 5.2.4 and 5.2.5 of 3GPP TS 24.250).
	RdsPorts []RdsPort `json:"rdsPorts,omitempty"`

	PdnEstablishmentOption PdnEstablishmentOptions `json:"pdnEstablishmentOption,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination"`

	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`

	// The Maximum Packet Size is the maximum NIDD packet size that was transferred to the UE by the SCEF in the PCO, see clause 4.5.14.1 of 3GPP TS 23.682. If no maximum packet size was provided to the UE by the SCEF, the SCEF sends a default configured max packet size to SCS/AS. Unit  bit.
	MaximumPacketSize int32 `json:"maximumPacketSize,omitempty"`

	// The downlink data deliveries that needed to be executed by the SCEF. The cardinality of the property shall be 0..1 in the request and 0..N in the response (i.e. response may contain multiple buffered MT NIDD).
	NiddDownlinkDataTransfers []NiddDownlinkDataTransfer `json:"niddDownlinkDataTransfers,omitempty"`

	Status NiddStatus `json:"status,omitempty"`
}
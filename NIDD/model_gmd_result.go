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

// GmdResult - Represents the group message delivery result.
type GmdResult struct {

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn string `json:"msisdn,omitempty"`

	DeliveryStatus DeliveryStatus `json:"deliveryStatus"`

	// string with format \"date-time\" as defined in OpenAPI.
	RequestedRetransmissionTime time.Time `json:"requestedRetransmissionTime,omitempty"`
}
/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// TrafficInfluDataNotif - Represents traffic influence data for notification.
type TrafficInfluDataNotif struct {

	// String providing an URI formatted according to RFC 3986.
	ResUri string `json:"resUri"`

	TrafficInfluData TrafficInfluData `json:"trafficInfluData,omitempty"`
}

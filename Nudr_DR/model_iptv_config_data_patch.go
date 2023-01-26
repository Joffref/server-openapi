/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// IptvConfigDataPatch - Represents the parameters to request the modification of an IPTV Configuration resource. 
type IptvConfigDataPatch struct {

	// Identifies a list of multicast address access control information. Any string value can be used as a key of the map. 
	MultiAccCtrls map[string]MulticastAccessControl `json:"multiAccCtrls,omitempty"`
}

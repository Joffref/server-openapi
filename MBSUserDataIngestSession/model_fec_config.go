/*
 * 3gpp-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSUserDataIngestSession

// FecConfig - Represents FEC configuration information.
type FecConfig struct {

	// String providing an URI formatted according to RFC 3986.
	FecScheme string `json:"fecScheme"`

	FecOverHead int32 `json:"fecOverHead"`

	AdditionalParams []AddFecParams `json:"additionalParams,omitempty"`
}
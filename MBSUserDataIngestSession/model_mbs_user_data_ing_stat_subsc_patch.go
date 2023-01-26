/*
 * 3gpp-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSUserDataIngestSession

// MbsUserDataIngStatSubscPatch - Represents the requested modifications to an MBS User Data Ingest Session Status  Subscription. 
type MbsUserDataIngStatSubscPatch struct {

	EventSubscs []SubscribedEvent `json:"eventSubscs,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri,omitempty"`
}

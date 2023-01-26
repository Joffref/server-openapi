/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

type MbsSecurityContext struct {

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbsSecurityContext
	KeyList map[string]MbsKeyInfo `json:"keyList"`
}

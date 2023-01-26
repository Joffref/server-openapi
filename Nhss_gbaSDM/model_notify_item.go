/*
 * Nhss_gbaSDM
 *
 * Nhss Subscriber Data Management Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_gbaSDM

// NotifyItem - Indicates changes on a resource.
type NotifyItem struct {

	// String providing an URI formatted according to RFC 3986.
	ResourceId string `json:"resourceId"`

	Changes []ChangeItem `json:"changes"`
}

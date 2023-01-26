/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// NetworkSliceCond - Subscription to a set of NFs, based on the slices (S-NSSAI and NSI) they support
type NetworkSliceCond struct {

	SnssaiList []Snssai `json:"snssaiList"`

	NsiList []string `json:"nsiList,omitempty"`
}

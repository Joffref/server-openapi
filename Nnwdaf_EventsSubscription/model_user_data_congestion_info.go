/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_EventsSubscription

// UserDataCongestionInfo - Represents the user data congestion information.
type UserDataCongestionInfo struct {

	NetworkArea NetworkAreaInfo `json:"networkArea"`

	CongestionInfo CongestionInfo `json:"congestionInfo"`

	Snssai Snssai `json:"snssai,omitempty"`
}

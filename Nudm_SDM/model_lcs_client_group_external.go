/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SDM

type LcsClientGroupExternal struct {

	LcsClientGroupId string `json:"lcsClientGroupId,omitempty"`

	AllowedGeographicArea []GeographicArea `json:"allowedGeographicArea,omitempty"`

	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`

	ValidTimePeriod ValidTimePeriod `json:"validTimePeriod,omitempty"`
}

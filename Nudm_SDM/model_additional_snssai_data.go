/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SDM

type AdditionalSnssaiData struct {

	RequiredAuthnAuthz bool `json:"requiredAuthnAuthz,omitempty"`

	SubscribedUeSliceMbr SliceMbrRm `json:"subscribedUeSliceMbr,omitempty"`

	SubscribedNsSrgList []string `json:"subscribedNsSrgList,omitempty"`
}

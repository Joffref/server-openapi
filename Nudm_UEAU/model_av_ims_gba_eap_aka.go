/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_UEAU

type AvImsGbaEapAka struct {

	AvType HssAvType `json:"avType"`

	Rand string `json:"rand"`

	Xres string `json:"xres"`

	Autn string `json:"autn"`

	Ck string `json:"ck"`

	Ik string `json:"ik"`
}
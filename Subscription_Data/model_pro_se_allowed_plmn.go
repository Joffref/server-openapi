/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

// ProSeAllowedPlmn - Contains the PLMN identities where the Prose services are authorised to use and the authorised Prose services on this given PLMNs.
type ProSeAllowedPlmn struct {

	VisitedPlmn PlmnId `json:"visitedPlmn"`

	ProseDirectAllowed []ProseDirectAllowed `json:"proseDirectAllowed,omitempty"`
}

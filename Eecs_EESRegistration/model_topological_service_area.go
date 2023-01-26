/*
 * ECS EES Registration_API
 *
 * API for EES Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eecs_EESRegistration

// TopologicalServiceArea - Represents topological service area information.
type TopologicalServiceArea struct {

	// A list of E-UTRA cell identities.
	Ecgis []Ecgi `json:"ecgis,omitempty"`

	// A list of NR cell identities.
	Ncgis []Ncgi `json:"ncgis,omitempty"`

	// A list of tracking area identities.
	Tais []Tai `json:"tais,omitempty"`

	// A list of PLMN identities.
	PlmnIds []PlmnId1 `json:"plmnIds,omitempty"`
}

/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// ServiceParameterDataPatch - Represents the service parameter data that can be updated.
type ServiceParameterDataPatch struct {

	// Represents configuration parameters for V2X communications over PC5 reference point. 
	ParamOverPc5 string `json:"paramOverPc5,omitempty"`

	// Represents configuration parameters for V2X communications over Uu reference point. 
	ParamOverUu string `json:"paramOverUu,omitempty"`

	// Represents the service parameters for 5G ProSe direct discovery.
	ParamForProSeDd string `json:"paramForProSeDd,omitempty"`

	// Represents the service parameters for 5G ProSe direct communications.
	ParamForProSeDc string `json:"paramForProSeDc,omitempty"`

	// Represents the service parameters for 5G ProSe UE-to-network relay UE.
	ParamForProSeU2NRelUe string `json:"paramForProSeU2NRelUe,omitempty"`

	// Represents the service parameters for 5G ProSe Remate UE.
	ParamForProSeRemUe string `json:"paramForProSeRemUe,omitempty"`

	// Contains the service parameter used to influence the URSP.
	UrspInfluence []UrspRuleRequest `json:"urspInfluence,omitempty"`

	// Contains the outcome of the UE Policy Delivery.
	DeliveryEvents []Event `json:"deliveryEvents,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	PolicDelivNotifUri string `json:"policDelivNotifUri,omitempty"`
}

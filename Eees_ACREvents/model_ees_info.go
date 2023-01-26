/*
 * Eees_ACREvents
 *
 * API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_ACREvents

// EesInfo - Represents EES information.
type EesInfo struct {

	// Identity of the EES
	EesId string `json:"eesId"`

	EndPt EndPoint `json:"endPt,omitempty"`

	// Application identities of the Edge Application Servers registered with the EES.
	EasIds []string `json:"easIds,omitempty"`

	// Represents an ECSP Information.
	EcspInfo string `json:"ecspInfo,omitempty"`

	SvcArea LocationArea5G `json:"svcArea,omitempty"`

	// Represents list of Data network access identifier.
	Dnais []string `json:"dnais,omitempty"`

	// Indicates if the EES supports service continuity or not, also indicates which ACR scenarios are supported by the EES. 
	EesSvcContSupp []AcrScenario `json:"eesSvcContSupp,omitempty"`

	// Indicates whether the EEC is required to register on the EES to use edge services or not. 
	EecRegConf bool `json:"eecRegConf"`
}

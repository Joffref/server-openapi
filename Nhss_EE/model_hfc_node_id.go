/*
 * Nhss_EE
 *
 * HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_EE

// HfcNodeId - REpresents the HFC Node Identifer received over NGAP.
type HfcNodeId struct {

	// This IE represents the identifier of the HFC node Id as specified in CableLabs WR-TR-5WWC-ARCH. It is provisioned by the wireline operator as part of wireline operations and may contain up to six characters. 
	HfcNId string `json:"hfcNId"`
}

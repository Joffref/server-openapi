/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_AMPolicyControl

// MappingOfSnssai - Contains the mapping of S-NSSAI in the serving network and the value of the home network
type MappingOfSnssai struct {

	ServingSnssai Snssai `json:"servingSnssai"`

	HomeSnssai Snssai `json:"homeSnssai"`
}

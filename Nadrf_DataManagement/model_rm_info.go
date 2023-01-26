/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

// RmInfo - Represents the registration state of a UE for an access type
type RmInfo struct {

	RmState RmState `json:"rmState"`

	AccessType AccessType `json:"accessType"`
}

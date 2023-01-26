/*
 * 3gpp-service-parameter
 *
 * API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ServiceParameter

// AppDescriptor - Represents an operation system and the corresponding applications.
type AppDescriptor struct {

	// Represents the Operating System of the served UE.
	OsId string `json:"osId"`

	// Identifies applications that are running on the UE's operating system. Any string value can be used as a key of the map. 
	AppIds map[string]string `json:"appIds"`
}

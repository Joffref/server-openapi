/*
 * Ntsctsf_ASTI Service API
 *
 * TSCTSF  Access Stratum time distribution configuration Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ntsctsf_ASTI

// StatusRequestData - Contains the parameters for retrieval of the status of the access stratum time distribution for a list of UEs. 
type StatusRequestData struct {

	Supis []string `json:"supis,omitempty"`

	Gpsis []string `json:"gpsis,omitempty"`
}
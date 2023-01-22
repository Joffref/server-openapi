/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_Communication

// NssaiMapping - Represents the mapping between a S-NSSAI in serving PLMN to a S-NSSAI in home PLMN
type NssaiMapping struct {

	MappedSnssai Snssai `json:"mappedSnssai"`

	HSnssai Snssai `json:"hSnssai"`
}
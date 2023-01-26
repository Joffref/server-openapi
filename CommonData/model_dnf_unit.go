/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// DnfUnit - During the processing of dnfUnits attribute, all the members in the array shall be  interpreted as logically concatenated with logical \"OR\". 
type DnfUnit struct {

	DnfUnit []Atom `json:"dnfUnit"`
}

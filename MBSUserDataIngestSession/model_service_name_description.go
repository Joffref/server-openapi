/*
 * 3gpp-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSUserDataIngestSession

// ServiceNameDescription - Represents a set of per language service names and/or service descriptions. 
type ServiceNameDescription struct {

	ServName string `json:"servName,omitempty"`

	ServDescrip string `json:"servDescrip,omitempty"`

	Language string `json:"language"`
}

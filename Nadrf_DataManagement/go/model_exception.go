/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

// Exception - Represents the Exception information.
type Exception struct {

	ExcepId ExceptionId `json:"excepId"`

	ExcepLevel int32 `json:"excepLevel,omitempty"`

	ExcepTrend ExceptionTrend `json:"excepTrend,omitempty"`
}
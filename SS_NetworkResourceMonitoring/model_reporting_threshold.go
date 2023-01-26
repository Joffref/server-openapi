/*
 * SS_NetworkResourceMonitoring
 *
 * API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_NetworkResourceMonitoring

// ReportingThreshold - Indicates the requested reporting termination threshold for the measurement index(es). 
type ReportingThreshold struct {

	MeasThrValues MeasurementData `json:"measThrValues"`

	ThrDirection MatchingDirection `json:"thrDirection"`
}
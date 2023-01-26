/*
 * 3gpp-time-sync-exposure
 *
 * API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_TimeSyncExposure

// ConfigForPort - Contains configuration for each port.
type ConfigForPort struct {

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi,omitempty"`

	N6Ind bool `json:"n6Ind,omitempty"`

	PtpEnable bool `json:"ptpEnable,omitempty"`

	LogSyncInter int32 `json:"logSyncInter,omitempty"`

	LogSyncInterInd bool `json:"logSyncInterInd,omitempty"`

	LogAnnouInter int32 `json:"logAnnouInter,omitempty"`

	LogAnnouInterInd bool `json:"logAnnouInterInd,omitempty"`
}
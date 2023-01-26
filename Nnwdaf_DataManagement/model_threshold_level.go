/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// ThresholdLevel - Represents a threshold level.
type ThresholdLevel struct {

	CongLevel int32 `json:"congLevel,omitempty"`

	NfLoadLevel int32 `json:"nfLoadLevel,omitempty"`

	NfCpuUsage int32 `json:"nfCpuUsage,omitempty"`

	NfMemoryUsage int32 `json:"nfMemoryUsage,omitempty"`

	NfStorageUsage int32 `json:"nfStorageUsage,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	AvgTrafficRate string `json:"avgTrafficRate,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxTrafficRate string `json:"maxTrafficRate,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	AvgPacketDelay int32 `json:"avgPacketDelay,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	MaxPacketDelay int32 `json:"maxPacketDelay,omitempty"`

	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	AvgPacketLossRate int32 `json:"avgPacketLossRate,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	SvcExpLevel float32 `json:"svcExpLevel,omitempty"`
}

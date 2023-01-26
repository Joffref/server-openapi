/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// NonDynamic5Qi - It indicates the QoS Characteristics for a standardized or pre-configured 5QI for downlink and uplink. 
type NonDynamic5Qi struct {

	// Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.  
	PriorityLevel int32 `json:"priorityLevel,omitempty"`

	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AverWindow int32 `json:"averWindow,omitempty"`

	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty"`

	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	ExtMaxDataBurstVol int32 `json:"extMaxDataBurstVol,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds. 
	CnPacketDelayBudgetDl int32 `json:"cnPacketDelayBudgetDl,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds. 
	CnPacketDelayBudgetUl int32 `json:"cnPacketDelayBudgetUl,omitempty"`
}

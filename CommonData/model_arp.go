/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// Arp - Contains Allocation and Retention Priority information.
type Arp struct {

	// nullable true shall not be used for this attribute. Unsigned integer indicating the ARP Priority Level (see clause 5.7.2.2 of 3GPP TS 23.501, within the range 1 to 15.Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 15 as the lowest priority.  
	PriorityLevel *int32 `json:"priorityLevel"`

	PreemptCap PreemptionCapability `json:"preemptCap"`

	PreemptVuln PreemptionVulnerability `json:"preemptVuln"`
}

/*
 * Nnef_SMContext
 *
 * Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_SMContext

// NiddInformation - Informaiton related to NIDD used for the SM Context.
type NiddInformation struct {

	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExtGroupId string `json:"extGroupId,omitempty"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi,omitempty"`

	// When present, this IE shall contain the AF Id used for the SM Context.
	AfId string `json:"afId,omitempty"`
}
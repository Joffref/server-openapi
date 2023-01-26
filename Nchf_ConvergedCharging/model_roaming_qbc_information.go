/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

type RoamingQbcInformation struct {

	MultipleQFIcontainer []MultipleQfIcontainer `json:"multipleQFIcontainer,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	UPFID string `json:"uPFID,omitempty"`

	RoamingChargingProfile RoamingChargingProfile `json:"roamingChargingProfile,omitempty"`
}

/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_ConvergedCharging

type MultipleUnitUsage struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	RatingGroup int32 `json:"ratingGroup"`

	RequestedUnit RequestedUnit `json:"requestedUnit,omitempty"`

	UsedUnitContainer []UsedUnitContainer `json:"usedUnitContainer,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	UPFID string `json:"uPFID,omitempty"`

	MultihomedPDUAddress PduAddress `json:"multihomedPDUAddress,omitempty"`
}
/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_MBSPolicyControl

type MediaTypeAnyOf string

// List of MediaTypeAnyOf
const (
	AUDIO MediaTypeAnyOf = "AUDIO"
	VIDEO MediaTypeAnyOf = "VIDEO"
	DATA MediaTypeAnyOf = "DATA"
	APPLICATION MediaTypeAnyOf = "APPLICATION"
	CONTROL MediaTypeAnyOf = "CONTROL"
	TEXT MediaTypeAnyOf = "TEXT"
	MESSAGE MediaTypeAnyOf = "MESSAGE"
	OTHER MediaTypeAnyOf = "OTHER"
)

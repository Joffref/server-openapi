/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

import (
	"os"
)

type UpdatePduSession400Response struct {

	JsonData HsmfUpdateError `json:"jsonData,omitempty"`

	BinaryDataN1SmInfoToUe *os.File `json:"binaryDataN1SmInfoToUe,omitempty"`
}
/*
 * EES Application Client Information_API
 *
 * API for EES Application Client Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_AppClientInformation

// AcServiceKpis - EAS details.
type AcServiceKpis struct {

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ConnBand string `json:"connBand,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ReqRate int32 `json:"reqRate,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	RespTime int32 `json:"respTime,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Avail int32 `json:"avail,omitempty"`

	// The compute resources required by the AC.
	ReqComp string `json:"reqComp,omitempty"`

	// The graphical compute resources required by the AC.
	ReqGrapComp string `json:"reqGrapComp,omitempty"`

	// The memory resources required by the AC.
	ReqMem string `json:"reqMem,omitempty"`

	// The storage resources required by the AC.
	ReqStrg string `json:"reqStrg,omitempty"`
}

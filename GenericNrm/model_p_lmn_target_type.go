/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GenericNrm

// PLmnTargetType - The PLMN for which sessions shall be selected in the Trace Session in case of management based activation when several PLMNs are supported in the RAN (this means that shared cells and not shared cells are allowed for the specified PLMN. Note that the PLMN Target might differ from the PLMN specified in the Trace Reference, as that specifies the PLMN that is containing the management system requesting the Trace Session from the NE. See 3GPP TS 32.422 clause 5.9b for additional details.
type PLmnTargetType struct {

	Mcc string `json:"mcc"`

	Mnc string `json:"mnc"`
}

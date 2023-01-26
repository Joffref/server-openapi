/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm

// NfService - NF Service is defined in TS 29.510
type NfService struct {

	ServiceInstanceId string `json:"serviceInstanceId,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	Version string `json:"version,omitempty"`

	Schema string `json:"schema,omitempty"`

	Fqdn string `json:"fqdn,omitempty"`

	InterPlmnFqdn string `json:"interPlmnFqdn,omitempty"`

	IpEndPoints []IpEndPoint `json:"ipEndPoints,omitempty"`

	ApiPrfix string `json:"apiPrfix,omitempty"`

	AllowedPlmns PlmnId `json:"allowedPlmns,omitempty"`

	AllowedNfTypes []NfType `json:"allowedNfTypes,omitempty"`

	AllowedNssais []Snssai `json:"allowedNssais,omitempty"`
}

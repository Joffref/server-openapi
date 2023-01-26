/*
 * M1_EdgeResourcesProvisioning
 *
 * 5GMS AF M1 Edge Resources Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_M1_EdgeResourcesProvisioning

type IpPacketFilterSet struct {

	SrcIp string `json:"srcIp,omitempty"`

	DstIp string `json:"dstIp,omitempty"`

	Protocol int32 `json:"protocol,omitempty"`

	SrcPort int32 `json:"srcPort,omitempty"`

	DstPort int32 `json:"dstPort,omitempty"`

	ToSTc string `json:"toSTc,omitempty"`

	FlowLabel int32 `json:"flowLabel,omitempty"`

	Spi int32 `json:"spi,omitempty"`

	Direction string `json:"direction"`
}

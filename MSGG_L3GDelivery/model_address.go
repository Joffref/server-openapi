/*
 * MSGG_L3GDelivery
 *
 * API for MSGG L3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSGG_L3GDelivery

// Address - Contains the Message type data
type Address struct {

	AddrType AddressType `json:"addrType"`

	Addr string `json:"addr"`
}

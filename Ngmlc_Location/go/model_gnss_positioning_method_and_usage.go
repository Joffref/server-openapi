/*
 * Ngmlc_Location
 *
 * GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ngmlc_Location

// GnssPositioningMethodAndUsage - Indicates the usage of a Global Navigation Satellite System (GNSS) positioning method.
type GnssPositioningMethodAndUsage struct {

	Mode PositioningMode `json:"mode"`

	Gnss GnssId `json:"gnss"`

	Usage Usage `json:"usage"`
}
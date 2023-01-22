/*
 * VAE_FileDistribution
 *
 * API for VAE File Distribution Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package VAE_FileDistribution

// GadShape - Common base type for GAD shapes.
type GadShape struct {

	Shape SupportedGadShapes `json:"shape"`
}
/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AiMlNrm

type ThresholdInfo struct {

	ThresholdDirection string `json:"thresholdDirection,omitempty"`

	ThresholdValue ThresholdInfoThresholdValue `json:"thresholdValue,omitempty"`

	Hysteresis ThresholdInfoHysteresis `json:"hysteresis,omitempty"`
}

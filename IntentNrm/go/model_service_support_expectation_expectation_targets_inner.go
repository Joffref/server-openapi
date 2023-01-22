/*
 * Intent NRM
 *
 * OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package IntentNrm

type ServiceSupportExpectationExpectationTargetsInner struct {

	TargetName string `json:"targetName,omitempty"`

	TargetCondition Condition `json:"targetCondition,omitempty"`

	TargetValueRange float32 `json:"targetValueRange,omitempty"`

	TargetAttribute string `json:"targetAttribute,omitempty"`

	TargetContexts []TargetContext `json:"targetContexts,omitempty"`
}
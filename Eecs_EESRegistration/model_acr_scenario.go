/*
 * ECS EES Registration_API
 *
 * API for EES Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eecs_EESRegistration

// AcrScenario - Possible values are: - EEC_INITIATED: Represents the EEC initiated ACR scenario. - EEC_EXECUTED_VIA_SOURCE_EES: Represents the EEC ACR scenario executed via the S-EES. - EEC_EXECUTED_VIA_TARGET_EES: Represents the EEC ACR scenario executed via the T-EES. - SOURCE_EAS_DECIDED: Represents the EEC ACR scenario where the S-EAS decides to perform ACR. - SOURCE_EES_EXECUTED: Represents the EEC ACR scenario where S-EES executes the ACR. - EEL_MANAGED_ACR: Represents the EEC ACR scenario where the ACR is managed by the Edge Enabler Layer. 
type AcrScenario struct {
}

/*
 * Ndcaf_DataReporting
 *
 * Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndcaf_DataReporting

type ApplicationSpecificRecordAllOf struct {

	// String providing an URI formatted according to RFC 3986.
	RecordType string `json:"recordType"`

	RecordContainer *interface{} `json:"recordContainer"`
}

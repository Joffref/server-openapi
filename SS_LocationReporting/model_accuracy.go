/*
 * SS_LocationReporting
 *
 * API for SEAL Location Reporting Configuration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_LocationReporting

// Accuracy - Possible values are - CGI_ECGI: The SCS/AS requests to be notified using cell level location accuracy. - ENODEB: The SCS/AS requests to be notified using eNodeB level location accuracy. - TA_RA: The SCS/AS requests to be notified using TA/RA level location accuracy. - PLMN: The SCS/AS requests to be notified using PLMN level location accuracy. - TWAN_ID: The SCS/AS requests to be notified using TWAN identifier level location accuracy. - GEO_AREA: The SCS/AS requests to be notified using the geographical area accuracy. - CIVIC_ADDR: The SCS/AS requests to be notified using the civic address accuracy. 
type Accuracy struct {
}

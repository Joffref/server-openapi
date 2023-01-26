/*
 * 3gpp-time-sync-exposure
 *
 * API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_TimeSyncExposure

// AsTimeResource - Possible values are: - ATOMIC_CLOCK: Indicates atomic clock is supported. - GNSS: Indicates Global Navigation Satellite System is supported. - TERRESTRIAL_RADIO: Indicates terrestrial radio is supported. - SERIAL_TIME_CODE: Indicates serial time code is supported. - PTP: Indicates PTP is supported. - NTP: Indicates NTP is supported. - HAND_SET: Indicates hand set is supported. - INTERNAL_OSCILLATOR: Indicates internal oscillator is supported. - OTHER: Indicates other source of time is supported. 
type AsTimeResource struct {
}

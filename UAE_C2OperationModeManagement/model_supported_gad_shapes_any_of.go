/*
 * UAE Server C2 Operation Mode Management Service
 *
 * UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_UAE_C2OperationModeManagement

type SupportedGadShapesAnyOf string

// List of SupportedGadShapesAnyOf
const (
	POINT SupportedGadShapesAnyOf = "POINT"
	POINT_UNCERTAINTY_CIRCLE SupportedGadShapesAnyOf = "POINT_UNCERTAINTY_CIRCLE"
	POINT_UNCERTAINTY_ELLIPSE SupportedGadShapesAnyOf = "POINT_UNCERTAINTY_ELLIPSE"
	POLYGON SupportedGadShapesAnyOf = "POLYGON"
	POINT_ALTITUDE SupportedGadShapesAnyOf = "POINT_ALTITUDE"
	POINT_ALTITUDE_UNCERTAINTY SupportedGadShapesAnyOf = "POINT_ALTITUDE_UNCERTAINTY"
	ELLIPSOID_ARC SupportedGadShapesAnyOf = "ELLIPSOID_ARC"
	LOCAL_2_D_POINT_UNCERTAINTY_ELLIPSE SupportedGadShapesAnyOf = "LOCAL_2D_POINT_UNCERTAINTY_ELLIPSE"
	LOCAL_3_D_POINT_UNCERTAINTY_ELLIPSOID SupportedGadShapesAnyOf = "LOCAL_3D_POINT_UNCERTAINTY_ELLIPSOID"
)

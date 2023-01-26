/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndividualApplicationDataSubscription - Delete the individual Application Data subscription
func DeleteIndividualApplicationDataSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadIndividualApplicationDataSubscription - Get an existing individual Application Data Subscription resource
func ReadIndividualApplicationDataSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReplaceIndividualApplicationDataSubscription - Modify a subscription to receive notification of application data changes
func ReplaceIndividualApplicationDataSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

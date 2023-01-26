/*
 * 3gpp-eas-deployment
 *
 * API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_EASDeployment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAnDeployment - Deletes an already existing EAS Deployment information resource
func DeleteAnDeployment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FullyUpdateAnDeployment - Fully updates/replaces an existing resource
func FullyUpdateAnDeployment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadAnDeployment - Read an active Individual EAS Deployment Information resource for the AF
func ReadAnDeployment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
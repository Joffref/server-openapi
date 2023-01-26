/*
 * 3gpp-nidd
 *
 * API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NIDD

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/3gpp-nidd/v1/",
		Index,
	},

	{
		"DeleteIndManagePortConfiguration",
		http.MethodDelete,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/rds-ports/:portId",
		DeleteIndManagePortConfiguration,
	},

	{
		"FetchIndManagePortConfiguration",
		http.MethodGet,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/rds-ports/:portId",
		FetchIndManagePortConfiguration,
	},

	{
		"UpdateIndManagePortConfiguration",
		http.MethodPut,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/rds-ports/:portId",
		UpdateIndManagePortConfiguration,
	},

	{
		"DeleteNIDDConfiguration",
		http.MethodDelete,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId",
		DeleteNIDDConfiguration,
	},

	{
		"FetchIndNIDDConfiguration",
		http.MethodGet,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId",
		FetchIndNIDDConfiguration,
	},

	{
		"ModifyNIDDConfiguration",
		http.MethodPatch,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId",
		ModifyNIDDConfiguration,
	},

	{
		"DeleteIndDownlinkDataDelivery",
		http.MethodDelete,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/downlink-data-deliveries/:downlinkDataDeliveryId",
		DeleteIndDownlinkDataDelivery,
	},

	{
		"FetchIndDownlinkDataDelivery",
		http.MethodGet,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/downlink-data-deliveries/:downlinkDataDeliveryId",
		FetchIndDownlinkDataDelivery,
	},

	{
		"ModifyIndDownlinkDataDelivery",
		http.MethodPatch,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/downlink-data-deliveries/:downlinkDataDeliveryId",
		ModifyIndDownlinkDataDelivery,
	},

	{
		"UpdateIndDownlinkDataDelivery",
		http.MethodPut,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/downlink-data-deliveries/:downlinkDataDeliveryId",
		UpdateIndDownlinkDataDelivery,
	},

	{
		"FetchAllManagePortConfigurations",
		http.MethodGet,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/rds-ports",
		FetchAllManagePortConfigurations,
	},

	{
		"CreateNIDDConfiguration",
		http.MethodPost,
		"/3gpp-nidd/v1/:scsAsId/configurations",
		CreateNIDDConfiguration,
	},

	{
		"FetchAllNIDDConfigurations",
		http.MethodGet,
		"/3gpp-nidd/v1/:scsAsId/configurations",
		FetchAllNIDDConfigurations,
	},

	{
		"CreateDownlinkDataDelivery",
		http.MethodPost,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/downlink-data-deliveries",
		CreateDownlinkDataDelivery,
	},

	{
		"FetchAllDownlinkDataDeliveries",
		http.MethodGet,
		"/3gpp-nidd/v1/:scsAsId/configurations/:configurationId/downlink-data-deliveries",
		FetchAllDownlinkDataDeliveries,
	},
}

// Package classification of Website API
//
// Documentation for Website API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/CodeMonkVJ/perusenews/server/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of websites
// swagger:response websitesResponse
type websitesResponseWrapper struct {
	// All current websites
	// in: body
	Body []data.Website
}

// Data structure representing a single website
// swagger:response websiteResponse
type websiteResponseWrapper struct {
	// Newly created website
	// in: body
	Body data.Website
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateWebsite createWebsite
type websiteParamsWrapper struct {
	// Website data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Website
}

// swagger:parameters listSingleWebsite deleteWebsite
type websiteIDParamsWrapper struct {
	// The id of the website for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

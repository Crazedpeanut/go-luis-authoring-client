// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/crazedpeanut/luis/models"
)

// NewNr5ade56a9d5b81c209ce2e5abParams creates a new Nr5ade56a9d5b81c209ce2e5abParams object
// with the default values initialized.
func NewNr5ade56a9d5b81c209ce2e5abParams() *Nr5ade56a9d5b81c209ce2e5abParams {
	var ()
	return &Nr5ade56a9d5b81c209ce2e5abParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5ade56a9d5b81c209ce2e5abParamsWithTimeout creates a new Nr5ade56a9d5b81c209ce2e5abParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5ade56a9d5b81c209ce2e5abParamsWithTimeout(timeout time.Duration) *Nr5ade56a9d5b81c209ce2e5abParams {
	var ()
	return &Nr5ade56a9d5b81c209ce2e5abParams{

		timeout: timeout,
	}
}

// NewNr5ade56a9d5b81c209ce2e5abParamsWithContext creates a new Nr5ade56a9d5b81c209ce2e5abParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr5ade56a9d5b81c209ce2e5abParamsWithContext(ctx context.Context) *Nr5ade56a9d5b81c209ce2e5abParams {
	var ()
	return &Nr5ade56a9d5b81c209ce2e5abParams{

		Context: ctx,
	}
}

// NewNr5ade56a9d5b81c209ce2e5abParamsWithHTTPClient creates a new Nr5ade56a9d5b81c209ce2e5abParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5ade56a9d5b81c209ce2e5abParamsWithHTTPClient(client *http.Client) *Nr5ade56a9d5b81c209ce2e5abParams {
	var ()
	return &Nr5ade56a9d5b81c209ce2e5abParams{
		HTTPClient: client,
	}
}

/*Nr5ade56a9d5b81c209ce2e5abParams contains all the parameters to send to the API endpoint
for the 5ade56a9d5b81c209ce2e5ab operation typically these are written to a http.Request
*/
type Nr5ade56a9d5b81c209ce2e5abParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*EntityID
	  Format - guid. The Pattern.Any entity ID.

	*/
	EntityID string
	/*HierarchicalModelUpdateObject
	  A model object containing the new Pattern.Any entity extractor name and explicit list.

	*/
	HierarchicalModelUpdateObject *models.HierarchicalModelUpdateObject
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) WithTimeout(timeout time.Duration) *Nr5ade56a9d5b81c209ce2e5abParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) WithContext(ctx context.Context) *Nr5ade56a9d5b81c209ce2e5abParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) WithHTTPClient(client *http.Client) *Nr5ade56a9d5b81c209ce2e5abParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) WithAppID(appID string) *Nr5ade56a9d5b81c209ce2e5abParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithEntityID adds the entityID to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) WithEntityID(entityID string) *Nr5ade56a9d5b81c209ce2e5abParams {
	o.SetEntityID(entityID)
	return o
}

// SetEntityID adds the entityId to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) SetEntityID(entityID string) {
	o.EntityID = entityID
}

// WithHierarchicalModelUpdateObject adds the hierarchicalModelUpdateObject to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) WithHierarchicalModelUpdateObject(hierarchicalModelUpdateObject *models.HierarchicalModelUpdateObject) *Nr5ade56a9d5b81c209ce2e5abParams {
	o.SetHierarchicalModelUpdateObject(hierarchicalModelUpdateObject)
	return o
}

// SetHierarchicalModelUpdateObject adds the hierarchicalModelUpdateObject to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) SetHierarchicalModelUpdateObject(hierarchicalModelUpdateObject *models.HierarchicalModelUpdateObject) {
	o.HierarchicalModelUpdateObject = hierarchicalModelUpdateObject
}

// WithVersionID adds the versionID to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) WithVersionID(versionID string) *Nr5ade56a9d5b81c209ce2e5abParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5ade56a9d5b81c209ce2e5ab params
func (o *Nr5ade56a9d5b81c209ce2e5abParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5ade56a9d5b81c209ce2e5abParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param entityId
	if err := r.SetPathParam("entityId", o.EntityID); err != nil {
		return err
	}

	if o.HierarchicalModelUpdateObject != nil {
		if err := r.SetBodyParam(o.HierarchicalModelUpdateObject); err != nil {
			return err
		}
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

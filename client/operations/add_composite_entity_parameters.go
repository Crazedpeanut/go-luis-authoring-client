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

	"github.com/crazedpeanut/go-luis-authoring-client/models"
)

// NewAddCompositeEntityParams creates a new AddCompositeEntityParams object
// with the default values initialized.
func NewAddCompositeEntityParams() *AddCompositeEntityParams {
	var ()
	return &AddCompositeEntityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCompositeEntityParamsWithTimeout creates a new AddCompositeEntityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCompositeEntityParamsWithTimeout(timeout time.Duration) *AddCompositeEntityParams {
	var ()
	return &AddCompositeEntityParams{

		timeout: timeout,
	}
}

// NewAddCompositeEntityParamsWithContext creates a new AddCompositeEntityParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCompositeEntityParamsWithContext(ctx context.Context) *AddCompositeEntityParams {
	var ()
	return &AddCompositeEntityParams{

		Context: ctx,
	}
}

// NewAddCompositeEntityParamsWithHTTPClient creates a new AddCompositeEntityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCompositeEntityParamsWithHTTPClient(client *http.Client) *AddCompositeEntityParams {
	var ()
	return &AddCompositeEntityParams{
		HTTPClient: client,
	}
}

/*AddCompositeEntityParams contains all the parameters to send to the API endpoint
for the add composite entity operation typically these are written to a http.Request
*/
type AddCompositeEntityParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*HierarchicalModelCreateObject
	  A model object containing the name and children of the new entity extractor.

	The name of the entity must be unique in the application and must not be used by a prebuilt entity.

	*/
	HierarchicalModelCreateObject *models.HierarchicalModelCreateObject
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add composite entity params
func (o *AddCompositeEntityParams) WithTimeout(timeout time.Duration) *AddCompositeEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add composite entity params
func (o *AddCompositeEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add composite entity params
func (o *AddCompositeEntityParams) WithContext(ctx context.Context) *AddCompositeEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add composite entity params
func (o *AddCompositeEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add composite entity params
func (o *AddCompositeEntityParams) WithHTTPClient(client *http.Client) *AddCompositeEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add composite entity params
func (o *AddCompositeEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the add composite entity params
func (o *AddCompositeEntityParams) WithAppID(appID string) *AddCompositeEntityParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the add composite entity params
func (o *AddCompositeEntityParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithHierarchicalModelCreateObject adds the hierarchicalModelCreateObject to the add composite entity params
func (o *AddCompositeEntityParams) WithHierarchicalModelCreateObject(hierarchicalModelCreateObject *models.HierarchicalModelCreateObject) *AddCompositeEntityParams {
	o.SetHierarchicalModelCreateObject(hierarchicalModelCreateObject)
	return o
}

// SetHierarchicalModelCreateObject adds the hierarchicalModelCreateObject to the add composite entity params
func (o *AddCompositeEntityParams) SetHierarchicalModelCreateObject(hierarchicalModelCreateObject *models.HierarchicalModelCreateObject) {
	o.HierarchicalModelCreateObject = hierarchicalModelCreateObject
}

// WithVersionID adds the versionID to the add composite entity params
func (o *AddCompositeEntityParams) WithVersionID(versionID string) *AddCompositeEntityParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the add composite entity params
func (o *AddCompositeEntityParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *AddCompositeEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.HierarchicalModelCreateObject != nil {
		if err := r.SetBodyParam(o.HierarchicalModelCreateObject); err != nil {
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

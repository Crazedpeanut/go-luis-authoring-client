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

// NewAddHierarchialEntityParams creates a new AddHierarchialEntityParams object
// with the default values initialized.
func NewAddHierarchialEntityParams() *AddHierarchialEntityParams {
	var ()
	return &AddHierarchialEntityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddHierarchialEntityParamsWithTimeout creates a new AddHierarchialEntityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddHierarchialEntityParamsWithTimeout(timeout time.Duration) *AddHierarchialEntityParams {
	var ()
	return &AddHierarchialEntityParams{

		timeout: timeout,
	}
}

// NewAddHierarchialEntityParamsWithContext creates a new AddHierarchialEntityParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddHierarchialEntityParamsWithContext(ctx context.Context) *AddHierarchialEntityParams {
	var ()
	return &AddHierarchialEntityParams{

		Context: ctx,
	}
}

// NewAddHierarchialEntityParamsWithHTTPClient creates a new AddHierarchialEntityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddHierarchialEntityParamsWithHTTPClient(client *http.Client) *AddHierarchialEntityParams {
	var ()
	return &AddHierarchialEntityParams{
		HTTPClient: client,
	}
}

/*AddHierarchialEntityParams contains all the parameters to send to the API endpoint
for the add hierarchial entity operation typically these are written to a http.Request
*/
type AddHierarchialEntityParams struct {

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

// WithTimeout adds the timeout to the add hierarchial entity params
func (o *AddHierarchialEntityParams) WithTimeout(timeout time.Duration) *AddHierarchialEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add hierarchial entity params
func (o *AddHierarchialEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add hierarchial entity params
func (o *AddHierarchialEntityParams) WithContext(ctx context.Context) *AddHierarchialEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add hierarchial entity params
func (o *AddHierarchialEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add hierarchial entity params
func (o *AddHierarchialEntityParams) WithHTTPClient(client *http.Client) *AddHierarchialEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add hierarchial entity params
func (o *AddHierarchialEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the add hierarchial entity params
func (o *AddHierarchialEntityParams) WithAppID(appID string) *AddHierarchialEntityParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the add hierarchial entity params
func (o *AddHierarchialEntityParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithHierarchicalModelCreateObject adds the hierarchicalModelCreateObject to the add hierarchial entity params
func (o *AddHierarchialEntityParams) WithHierarchicalModelCreateObject(hierarchicalModelCreateObject *models.HierarchicalModelCreateObject) *AddHierarchialEntityParams {
	o.SetHierarchicalModelCreateObject(hierarchicalModelCreateObject)
	return o
}

// SetHierarchicalModelCreateObject adds the hierarchicalModelCreateObject to the add hierarchial entity params
func (o *AddHierarchialEntityParams) SetHierarchicalModelCreateObject(hierarchicalModelCreateObject *models.HierarchicalModelCreateObject) {
	o.HierarchicalModelCreateObject = hierarchicalModelCreateObject
}

// WithVersionID adds the versionID to the add hierarchial entity params
func (o *AddHierarchialEntityParams) WithVersionID(versionID string) *AddHierarchialEntityParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the add hierarchial entity params
func (o *AddHierarchialEntityParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *AddHierarchialEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

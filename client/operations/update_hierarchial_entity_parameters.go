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

// NewUpdateHierarchialEntityParams creates a new UpdateHierarchialEntityParams object
// with the default values initialized.
func NewUpdateHierarchialEntityParams() *UpdateHierarchialEntityParams {
	var ()
	return &UpdateHierarchialEntityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHierarchialEntityParamsWithTimeout creates a new UpdateHierarchialEntityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateHierarchialEntityParamsWithTimeout(timeout time.Duration) *UpdateHierarchialEntityParams {
	var ()
	return &UpdateHierarchialEntityParams{

		timeout: timeout,
	}
}

// NewUpdateHierarchialEntityParamsWithContext creates a new UpdateHierarchialEntityParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateHierarchialEntityParamsWithContext(ctx context.Context) *UpdateHierarchialEntityParams {
	var ()
	return &UpdateHierarchialEntityParams{

		Context: ctx,
	}
}

// NewUpdateHierarchialEntityParamsWithHTTPClient creates a new UpdateHierarchialEntityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateHierarchialEntityParamsWithHTTPClient(client *http.Client) *UpdateHierarchialEntityParams {
	var ()
	return &UpdateHierarchialEntityParams{
		HTTPClient: client,
	}
}

/*UpdateHierarchialEntityParams contains all the parameters to send to the API endpoint
for the update hierarchial entity operation typically these are written to a http.Request
*/
type UpdateHierarchialEntityParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*HEntityID
	  Format - guid. The hierarchical entity extractor ID.

	*/
	HEntityID string
	/*HierarchicalModelUpdateObject
	  Model containing names of the children of the hierarchical entity.

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

// WithTimeout adds the timeout to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) WithTimeout(timeout time.Duration) *UpdateHierarchialEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) WithContext(ctx context.Context) *UpdateHierarchialEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) WithHTTPClient(client *http.Client) *UpdateHierarchialEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) WithAppID(appID string) *UpdateHierarchialEntityParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithHEntityID adds the hEntityID to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) WithHEntityID(hEntityID string) *UpdateHierarchialEntityParams {
	o.SetHEntityID(hEntityID)
	return o
}

// SetHEntityID adds the hEntityId to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) SetHEntityID(hEntityID string) {
	o.HEntityID = hEntityID
}

// WithHierarchicalModelUpdateObject adds the hierarchicalModelUpdateObject to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) WithHierarchicalModelUpdateObject(hierarchicalModelUpdateObject *models.HierarchicalModelUpdateObject) *UpdateHierarchialEntityParams {
	o.SetHierarchicalModelUpdateObject(hierarchicalModelUpdateObject)
	return o
}

// SetHierarchicalModelUpdateObject adds the hierarchicalModelUpdateObject to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) SetHierarchicalModelUpdateObject(hierarchicalModelUpdateObject *models.HierarchicalModelUpdateObject) {
	o.HierarchicalModelUpdateObject = hierarchicalModelUpdateObject
}

// WithVersionID adds the versionID to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) WithVersionID(versionID string) *UpdateHierarchialEntityParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the update hierarchial entity params
func (o *UpdateHierarchialEntityParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHierarchialEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param hEntityId
	if err := r.SetPathParam("hEntityId", o.HEntityID); err != nil {
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

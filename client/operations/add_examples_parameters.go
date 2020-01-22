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

// NewAddExamplesParams creates a new AddExamplesParams object
// with the default values initialized.
func NewAddExamplesParams() *AddExamplesParams {
	var ()
	return &AddExamplesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddExamplesParamsWithTimeout creates a new AddExamplesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddExamplesParamsWithTimeout(timeout time.Duration) *AddExamplesParams {
	var ()
	return &AddExamplesParams{

		timeout: timeout,
	}
}

// NewAddExamplesParamsWithContext creates a new AddExamplesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddExamplesParamsWithContext(ctx context.Context) *AddExamplesParams {
	var ()
	return &AddExamplesParams{

		Context: ctx,
	}
}

// NewAddExamplesParamsWithHTTPClient creates a new AddExamplesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddExamplesParamsWithHTTPClient(client *http.Client) *AddExamplesParams {
	var ()
	return &AddExamplesParams{
		HTTPClient: client,
	}
}

/*AddExamplesParams contains all the parameters to send to the API endpoint
for the add examples operation typically these are written to a http.Request
*/
type AddExamplesParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*ExampleLabelObjectArray
	  A JSON array containing non-duplicate example labels.

	*/
	ExampleLabelObjectArray models.ExampleLabelObjectArray
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add examples params
func (o *AddExamplesParams) WithTimeout(timeout time.Duration) *AddExamplesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add examples params
func (o *AddExamplesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add examples params
func (o *AddExamplesParams) WithContext(ctx context.Context) *AddExamplesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add examples params
func (o *AddExamplesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add examples params
func (o *AddExamplesParams) WithHTTPClient(client *http.Client) *AddExamplesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add examples params
func (o *AddExamplesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the add examples params
func (o *AddExamplesParams) WithAppID(appID string) *AddExamplesParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the add examples params
func (o *AddExamplesParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithExampleLabelObjectArray adds the exampleLabelObjectArray to the add examples params
func (o *AddExamplesParams) WithExampleLabelObjectArray(exampleLabelObjectArray models.ExampleLabelObjectArray) *AddExamplesParams {
	o.SetExampleLabelObjectArray(exampleLabelObjectArray)
	return o
}

// SetExampleLabelObjectArray adds the exampleLabelObjectArray to the add examples params
func (o *AddExamplesParams) SetExampleLabelObjectArray(exampleLabelObjectArray models.ExampleLabelObjectArray) {
	o.ExampleLabelObjectArray = exampleLabelObjectArray
}

// WithVersionID adds the versionID to the add examples params
func (o *AddExamplesParams) WithVersionID(versionID string) *AddExamplesParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the add examples params
func (o *AddExamplesParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *AddExamplesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.ExampleLabelObjectArray != nil {
		if err := r.SetBodyParam(o.ExampleLabelObjectArray); err != nil {
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
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

// NewCreatePatternFeatureParams creates a new CreatePatternFeatureParams object
// with the default values initialized.
func NewCreatePatternFeatureParams() *CreatePatternFeatureParams {
	var ()
	return &CreatePatternFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePatternFeatureParamsWithTimeout creates a new CreatePatternFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePatternFeatureParamsWithTimeout(timeout time.Duration) *CreatePatternFeatureParams {
	var ()
	return &CreatePatternFeatureParams{

		timeout: timeout,
	}
}

// NewCreatePatternFeatureParamsWithContext creates a new CreatePatternFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePatternFeatureParamsWithContext(ctx context.Context) *CreatePatternFeatureParams {
	var ()
	return &CreatePatternFeatureParams{

		Context: ctx,
	}
}

// NewCreatePatternFeatureParamsWithHTTPClient creates a new CreatePatternFeatureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePatternFeatureParamsWithHTTPClient(client *http.Client) *CreatePatternFeatureParams {
	var ()
	return &CreatePatternFeatureParams{
		HTTPClient: client,
	}
}

/*CreatePatternFeatureParams contains all the parameters to send to the API endpoint
for the create pattern feature operation typically these are written to a http.Request
*/
type CreatePatternFeatureParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*PatternCreateObject
	  A JSON object containing Name and Pattern of the feature.

	*/
	PatternCreateObject *models.PatternCreateObject
	/*VersionID
	  The version ID of the task.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create pattern feature params
func (o *CreatePatternFeatureParams) WithTimeout(timeout time.Duration) *CreatePatternFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pattern feature params
func (o *CreatePatternFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pattern feature params
func (o *CreatePatternFeatureParams) WithContext(ctx context.Context) *CreatePatternFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pattern feature params
func (o *CreatePatternFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pattern feature params
func (o *CreatePatternFeatureParams) WithHTTPClient(client *http.Client) *CreatePatternFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pattern feature params
func (o *CreatePatternFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the create pattern feature params
func (o *CreatePatternFeatureParams) WithAppID(appID string) *CreatePatternFeatureParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the create pattern feature params
func (o *CreatePatternFeatureParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithPatternCreateObject adds the patternCreateObject to the create pattern feature params
func (o *CreatePatternFeatureParams) WithPatternCreateObject(patternCreateObject *models.PatternCreateObject) *CreatePatternFeatureParams {
	o.SetPatternCreateObject(patternCreateObject)
	return o
}

// SetPatternCreateObject adds the patternCreateObject to the create pattern feature params
func (o *CreatePatternFeatureParams) SetPatternCreateObject(patternCreateObject *models.PatternCreateObject) {
	o.PatternCreateObject = patternCreateObject
}

// WithVersionID adds the versionID to the create pattern feature params
func (o *CreatePatternFeatureParams) WithVersionID(versionID string) *CreatePatternFeatureParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the create pattern feature params
func (o *CreatePatternFeatureParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePatternFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.PatternCreateObject != nil {
		if err := r.SetBodyParam(o.PatternCreateObject); err != nil {
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

// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// NewCreateWebhookPolicyOfProjectParams creates a new CreateWebhookPolicyOfProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateWebhookPolicyOfProjectParams() *CreateWebhookPolicyOfProjectParams {
	return &CreateWebhookPolicyOfProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWebhookPolicyOfProjectParamsWithTimeout creates a new CreateWebhookPolicyOfProjectParams object
// with the ability to set a timeout on a request.
func NewCreateWebhookPolicyOfProjectParamsWithTimeout(timeout time.Duration) *CreateWebhookPolicyOfProjectParams {
	return &CreateWebhookPolicyOfProjectParams{
		timeout: timeout,
	}
}

// NewCreateWebhookPolicyOfProjectParamsWithContext creates a new CreateWebhookPolicyOfProjectParams object
// with the ability to set a context for a request.
func NewCreateWebhookPolicyOfProjectParamsWithContext(ctx context.Context) *CreateWebhookPolicyOfProjectParams {
	return &CreateWebhookPolicyOfProjectParams{
		Context: ctx,
	}
}

// NewCreateWebhookPolicyOfProjectParamsWithHTTPClient creates a new CreateWebhookPolicyOfProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateWebhookPolicyOfProjectParamsWithHTTPClient(client *http.Client) *CreateWebhookPolicyOfProjectParams {
	return &CreateWebhookPolicyOfProjectParams{
		HTTPClient: client,
	}
}

/* CreateWebhookPolicyOfProjectParams contains all the parameters to send to the API endpoint
   for the create webhook policy of project operation.

   Typically these are written to a http.Request.
*/
type CreateWebhookPolicyOfProjectParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool `js:"xIsResourceName"`

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Policy.

	   Properties "targets" and "event_types" needed.
	*/
	Policy *models.WebhookPolicy `js:"policy"`

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string `js:"projectNameOrID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the create webhook policy of project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWebhookPolicyOfProjectParams) WithDefaults() *CreateWebhookPolicyOfProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create webhook policy of project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWebhookPolicyOfProjectParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := CreateWebhookPolicyOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) WithTimeout(timeout time.Duration) *CreateWebhookPolicyOfProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) WithContext(ctx context.Context) *CreateWebhookPolicyOfProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) WithHTTPClient(client *http.Client) *CreateWebhookPolicyOfProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) WithXIsResourceName(xIsResourceName *bool) *CreateWebhookPolicyOfProjectParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) WithXRequestID(xRequestID *string) *CreateWebhookPolicyOfProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPolicy adds the policy to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) WithPolicy(policy *models.WebhookPolicy) *CreateWebhookPolicyOfProjectParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) SetPolicy(policy *models.WebhookPolicy) {
	o.Policy = policy
}

// WithProjectNameOrID adds the projectNameOrID to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) WithProjectNameOrID(projectNameOrID string) *CreateWebhookPolicyOfProjectParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the create webhook policy of project params
func (o *CreateWebhookPolicyOfProjectParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWebhookPolicyOfProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
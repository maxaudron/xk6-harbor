// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// AddProjectMetadatasReader is a Reader for the AddProjectMetadatas structure.
type AddProjectMetadatasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProjectMetadatasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProjectMetadatasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddProjectMetadatasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddProjectMetadatasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddProjectMetadatasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddProjectMetadatasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddProjectMetadatasConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddProjectMetadatasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddProjectMetadatasOK creates a AddProjectMetadatasOK with default headers values
func NewAddProjectMetadatasOK() *AddProjectMetadatasOK {
	return &AddProjectMetadatasOK{}
}

/* AddProjectMetadatasOK describes a response with status code 200, with default header values.

Success
*/
type AddProjectMetadatasOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *AddProjectMetadatasOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasOK ", 200)
}

func (o *AddProjectMetadatasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewAddProjectMetadatasBadRequest creates a AddProjectMetadatasBadRequest with default headers values
func NewAddProjectMetadatasBadRequest() *AddProjectMetadatasBadRequest {
	return &AddProjectMetadatasBadRequest{}
}

/* AddProjectMetadatasBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddProjectMetadatasBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddProjectMetadatasBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasBadRequest  %+v", 400, o.Payload)
}
func (o *AddProjectMetadatasBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectMetadatasUnauthorized creates a AddProjectMetadatasUnauthorized with default headers values
func NewAddProjectMetadatasUnauthorized() *AddProjectMetadatasUnauthorized {
	return &AddProjectMetadatasUnauthorized{}
}

/* AddProjectMetadatasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddProjectMetadatasUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddProjectMetadatasUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasUnauthorized  %+v", 401, o.Payload)
}
func (o *AddProjectMetadatasUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectMetadatasForbidden creates a AddProjectMetadatasForbidden with default headers values
func NewAddProjectMetadatasForbidden() *AddProjectMetadatasForbidden {
	return &AddProjectMetadatasForbidden{}
}

/* AddProjectMetadatasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddProjectMetadatasForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddProjectMetadatasForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasForbidden  %+v", 403, o.Payload)
}
func (o *AddProjectMetadatasForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectMetadatasNotFound creates a AddProjectMetadatasNotFound with default headers values
func NewAddProjectMetadatasNotFound() *AddProjectMetadatasNotFound {
	return &AddProjectMetadatasNotFound{}
}

/* AddProjectMetadatasNotFound describes a response with status code 404, with default header values.

Not found
*/
type AddProjectMetadatasNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddProjectMetadatasNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasNotFound  %+v", 404, o.Payload)
}
func (o *AddProjectMetadatasNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectMetadatasConflict creates a AddProjectMetadatasConflict with default headers values
func NewAddProjectMetadatasConflict() *AddProjectMetadatasConflict {
	return &AddProjectMetadatasConflict{}
}

/* AddProjectMetadatasConflict describes a response with status code 409, with default header values.

Conflict
*/
type AddProjectMetadatasConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddProjectMetadatasConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasConflict  %+v", 409, o.Payload)
}
func (o *AddProjectMetadatasConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectMetadatasInternalServerError creates a AddProjectMetadatasInternalServerError with default headers values
func NewAddProjectMetadatasInternalServerError() *AddProjectMetadatasInternalServerError {
	return &AddProjectMetadatasInternalServerError{}
}

/* AddProjectMetadatasInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type AddProjectMetadatasInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddProjectMetadatasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasInternalServerError  %+v", 500, o.Payload)
}
func (o *AddProjectMetadatasInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

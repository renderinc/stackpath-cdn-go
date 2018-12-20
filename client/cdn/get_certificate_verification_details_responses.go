// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/renderinc/stackpath-cdn-go/models"
)

// GetCertificateVerificationDetailsReader is a Reader for the GetCertificateVerificationDetails structure.
type GetCertificateVerificationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCertificateVerificationDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCertificateVerificationDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetCertificateVerificationDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCertificateVerificationDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCertificateVerificationDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCertificateVerificationDetailsOK creates a GetCertificateVerificationDetailsOK with default headers values
func NewGetCertificateVerificationDetailsOK() *GetCertificateVerificationDetailsOK {
	return &GetCertificateVerificationDetailsOK{}
}

/*GetCertificateVerificationDetailsOK handles this case with default header values.

GetCertificateVerificationDetailsOK get certificate verification details o k
*/
type GetCertificateVerificationDetailsOK struct {
	Payload *models.CdnGetCertificateVerificationDetailsResponse
}

func (o *GetCertificateVerificationDetailsOK) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/verification_details][%d] getCertificateVerificationDetailsOK  %+v", 200, o.Payload)
}

func (o *GetCertificateVerificationDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnGetCertificateVerificationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateVerificationDetailsUnauthorized creates a GetCertificateVerificationDetailsUnauthorized with default headers values
func NewGetCertificateVerificationDetailsUnauthorized() *GetCertificateVerificationDetailsUnauthorized {
	return &GetCertificateVerificationDetailsUnauthorized{}
}

/*GetCertificateVerificationDetailsUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type GetCertificateVerificationDetailsUnauthorized struct {
	Payload *models.APIStatus
}

func (o *GetCertificateVerificationDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/verification_details][%d] getCertificateVerificationDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCertificateVerificationDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateVerificationDetailsInternalServerError creates a GetCertificateVerificationDetailsInternalServerError with default headers values
func NewGetCertificateVerificationDetailsInternalServerError() *GetCertificateVerificationDetailsInternalServerError {
	return &GetCertificateVerificationDetailsInternalServerError{}
}

/*GetCertificateVerificationDetailsInternalServerError handles this case with default header values.

Internal server error.
*/
type GetCertificateVerificationDetailsInternalServerError struct {
	Payload *models.APIStatus
}

func (o *GetCertificateVerificationDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/verification_details][%d] getCertificateVerificationDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCertificateVerificationDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateVerificationDetailsDefault creates a GetCertificateVerificationDetailsDefault with default headers values
func NewGetCertificateVerificationDetailsDefault(code int) *GetCertificateVerificationDetailsDefault {
	return &GetCertificateVerificationDetailsDefault{
		_statusCode: code,
	}
}

/*GetCertificateVerificationDetailsDefault handles this case with default header values.

Default error structure.
*/
type GetCertificateVerificationDetailsDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the get certificate verification details default response
func (o *GetCertificateVerificationDetailsDefault) Code() int {
	return o._statusCode
}

func (o *GetCertificateVerificationDetailsDefault) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/verification_details][%d] GetCertificateVerificationDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetCertificateVerificationDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
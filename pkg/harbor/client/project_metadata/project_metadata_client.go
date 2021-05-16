// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the project metadata client
type API interface {
	/*
	   AddProjectMetadatas adds metadata for the specific project

	   Add metadata for the specific project*/
	AddProjectMetadatas(ctx context.Context, params *AddProjectMetadatasParams) (*AddProjectMetadatasOK, error)
	/*
	   DeleteProjectMetadata deletes the specific metadata for the specific project

	   Delete the specific metadata for the specific project*/
	DeleteProjectMetadata(ctx context.Context, params *DeleteProjectMetadataParams) (*DeleteProjectMetadataOK, error)
	/*
	   GetProjectMetadata gets the specific metadata of the specific project

	   Get the specific metadata of the specific project*/
	GetProjectMetadata(ctx context.Context, params *GetProjectMetadataParams) (*GetProjectMetadataOK, error)
	/*
	   ListProjectMetadatas gets the metadata of the specific project

	   Get the metadata of the specific project*/
	ListProjectMetadatas(ctx context.Context, params *ListProjectMetadatasParams) (*ListProjectMetadatasOK, error)
	/*
	   UpdateProjectMetadata updates the specific metadata for the specific project

	   Update the specific metadata for the specific project*/
	UpdateProjectMetadata(ctx context.Context, params *UpdateProjectMetadataParams) (*UpdateProjectMetadataOK, error)
}

// New creates a new project metadata API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for project metadata API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
AddProjectMetadatas adds metadata for the specific project

Add metadata for the specific project
*/
func (a *Client) AddProjectMetadatas(ctx context.Context, params *AddProjectMetadatasParams) (*AddProjectMetadatasOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addProjectMetadatas",
		Method:             "POST",
		PathPattern:        "/projects/{project_name_or_id}/metadatas/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddProjectMetadatasReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddProjectMetadatasOK), nil

}

/*
DeleteProjectMetadata deletes the specific metadata for the specific project

Delete the specific metadata for the specific project
*/
func (a *Client) DeleteProjectMetadata(ctx context.Context, params *DeleteProjectMetadataParams) (*DeleteProjectMetadataOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProjectMetadata",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name_or_id}/metadatas/{meta_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProjectMetadataReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProjectMetadataOK), nil

}

/*
GetProjectMetadata gets the specific metadata of the specific project

Get the specific metadata of the specific project
*/
func (a *Client) GetProjectMetadata(ctx context.Context, params *GetProjectMetadataParams) (*GetProjectMetadataOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProjectMetadata",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/metadatas/{meta_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProjectMetadataReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProjectMetadataOK), nil

}

/*
ListProjectMetadatas gets the metadata of the specific project

Get the metadata of the specific project
*/
func (a *Client) ListProjectMetadatas(ctx context.Context, params *ListProjectMetadatasParams) (*ListProjectMetadatasOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProjectMetadatas",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/metadatas/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListProjectMetadatasReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListProjectMetadatasOK), nil

}

/*
UpdateProjectMetadata updates the specific metadata for the specific project

Update the specific metadata for the specific project
*/
func (a *Client) UpdateProjectMetadata(ctx context.Context, params *UpdateProjectMetadataParams) (*UpdateProjectMetadataOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProjectMetadata",
		Method:             "PUT",
		PathPattern:        "/projects/{project_name_or_id}/metadatas/{meta_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateProjectMetadataReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProjectMetadataOK), nil

}
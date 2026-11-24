# FoldersCreatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Name** | **string** | name | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ParentFolderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFoldersCreatePostRequest

`func NewFoldersCreatePostRequest(workspaceId string, name string, ) *FoldersCreatePostRequest`

NewFoldersCreatePostRequest instantiates a new FoldersCreatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersCreatePostRequestWithDefaults

`func NewFoldersCreatePostRequestWithDefaults() *FoldersCreatePostRequest`

NewFoldersCreatePostRequestWithDefaults instantiates a new FoldersCreatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *FoldersCreatePostRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *FoldersCreatePostRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *FoldersCreatePostRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *FoldersCreatePostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FoldersCreatePostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FoldersCreatePostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FoldersCreatePostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FoldersCreatePostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FoldersCreatePostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FoldersCreatePostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FoldersCreatePostRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FoldersCreatePostRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentFolderId

`func (o *FoldersCreatePostRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *FoldersCreatePostRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *FoldersCreatePostRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *FoldersCreatePostRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *FoldersCreatePostRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *FoldersCreatePostRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



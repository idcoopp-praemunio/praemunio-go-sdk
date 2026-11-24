# FoldersDetailsGet200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PublicId** | **string** |  | 
**ParentId** | **map[string]interface{}** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Path** | **string** |  | 
**DeletedAt** | **map[string]interface{}** |  | 
**DeletedBy** | **map[string]interface{}** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Workspace** | [**FoldersDetailsGet200ResponseResultWorkspace**](FoldersDetailsGet200ResponseResultWorkspace.md) |  | 
**Parent** | **map[string]interface{}** |  | 

## Methods

### NewFoldersDetailsGet200ResponseResult

`func NewFoldersDetailsGet200ResponseResult(id int32, publicId string, parentId map[string]interface{}, name string, description string, path string, deletedAt map[string]interface{}, deletedBy map[string]interface{}, createdAt string, updatedAt string, workspace FoldersDetailsGet200ResponseResultWorkspace, parent map[string]interface{}, ) *FoldersDetailsGet200ResponseResult`

NewFoldersDetailsGet200ResponseResult instantiates a new FoldersDetailsGet200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersDetailsGet200ResponseResultWithDefaults

`func NewFoldersDetailsGet200ResponseResultWithDefaults() *FoldersDetailsGet200ResponseResult`

NewFoldersDetailsGet200ResponseResultWithDefaults instantiates a new FoldersDetailsGet200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FoldersDetailsGet200ResponseResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FoldersDetailsGet200ResponseResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FoldersDetailsGet200ResponseResult) SetId(v int32)`

SetId sets Id field to given value.


### GetPublicId

`func (o *FoldersDetailsGet200ResponseResult) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *FoldersDetailsGet200ResponseResult) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *FoldersDetailsGet200ResponseResult) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetParentId

`func (o *FoldersDetailsGet200ResponseResult) GetParentId() map[string]interface{}`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FoldersDetailsGet200ResponseResult) GetParentIdOk() (*map[string]interface{}, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FoldersDetailsGet200ResponseResult) SetParentId(v map[string]interface{})`

SetParentId sets ParentId field to given value.


### GetName

`func (o *FoldersDetailsGet200ResponseResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FoldersDetailsGet200ResponseResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FoldersDetailsGet200ResponseResult) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FoldersDetailsGet200ResponseResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FoldersDetailsGet200ResponseResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FoldersDetailsGet200ResponseResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPath

`func (o *FoldersDetailsGet200ResponseResult) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FoldersDetailsGet200ResponseResult) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FoldersDetailsGet200ResponseResult) SetPath(v string)`

SetPath sets Path field to given value.


### GetDeletedAt

`func (o *FoldersDetailsGet200ResponseResult) GetDeletedAt() map[string]interface{}`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FoldersDetailsGet200ResponseResult) GetDeletedAtOk() (*map[string]interface{}, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FoldersDetailsGet200ResponseResult) SetDeletedAt(v map[string]interface{})`

SetDeletedAt sets DeletedAt field to given value.


### GetDeletedBy

`func (o *FoldersDetailsGet200ResponseResult) GetDeletedBy() map[string]interface{}`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *FoldersDetailsGet200ResponseResult) GetDeletedByOk() (*map[string]interface{}, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *FoldersDetailsGet200ResponseResult) SetDeletedBy(v map[string]interface{})`

SetDeletedBy sets DeletedBy field to given value.


### GetCreatedAt

`func (o *FoldersDetailsGet200ResponseResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FoldersDetailsGet200ResponseResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FoldersDetailsGet200ResponseResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FoldersDetailsGet200ResponseResult) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FoldersDetailsGet200ResponseResult) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FoldersDetailsGet200ResponseResult) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWorkspace

`func (o *FoldersDetailsGet200ResponseResult) GetWorkspace() FoldersDetailsGet200ResponseResultWorkspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *FoldersDetailsGet200ResponseResult) GetWorkspaceOk() (*FoldersDetailsGet200ResponseResultWorkspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *FoldersDetailsGet200ResponseResult) SetWorkspace(v FoldersDetailsGet200ResponseResultWorkspace)`

SetWorkspace sets Workspace field to given value.


### GetParent

`func (o *FoldersDetailsGet200ResponseResult) GetParent() map[string]interface{}`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FoldersDetailsGet200ResponseResult) GetParentOk() (*map[string]interface{}, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FoldersDetailsGet200ResponseResult) SetParent(v map[string]interface{})`

SetParent sets Parent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



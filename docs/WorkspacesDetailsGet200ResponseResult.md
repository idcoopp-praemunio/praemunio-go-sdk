# WorkspacesDetailsGet200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PublicId** | **string** |  | 
**Name** | **string** |  | 
**IsPersonal** | **bool** |  | 
**Description** | **string** |  | 
**Status** | **string** |  | 
**StorageLimit** | **string** |  | 
**StorageUsed** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Team** | [**WorkspacesDetailsGet200ResponseResultTeam**](WorkspacesDetailsGet200ResponseResultTeam.md) |  | 

## Methods

### NewWorkspacesDetailsGet200ResponseResult

`func NewWorkspacesDetailsGet200ResponseResult(id int32, publicId string, name string, isPersonal bool, description string, status string, storageLimit string, storageUsed string, createdAt string, updatedAt string, team WorkspacesDetailsGet200ResponseResultTeam, ) *WorkspacesDetailsGet200ResponseResult`

NewWorkspacesDetailsGet200ResponseResult instantiates a new WorkspacesDetailsGet200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspacesDetailsGet200ResponseResultWithDefaults

`func NewWorkspacesDetailsGet200ResponseResultWithDefaults() *WorkspacesDetailsGet200ResponseResult`

NewWorkspacesDetailsGet200ResponseResultWithDefaults instantiates a new WorkspacesDetailsGet200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkspacesDetailsGet200ResponseResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspacesDetailsGet200ResponseResult) SetId(v int32)`

SetId sets Id field to given value.


### GetPublicId

`func (o *WorkspacesDetailsGet200ResponseResult) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *WorkspacesDetailsGet200ResponseResult) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetName

`func (o *WorkspacesDetailsGet200ResponseResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspacesDetailsGet200ResponseResult) SetName(v string)`

SetName sets Name field to given value.


### GetIsPersonal

`func (o *WorkspacesDetailsGet200ResponseResult) GetIsPersonal() bool`

GetIsPersonal returns the IsPersonal field if non-nil, zero value otherwise.

### GetIsPersonalOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetIsPersonalOk() (*bool, bool)`

GetIsPersonalOk returns a tuple with the IsPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonal

`func (o *WorkspacesDetailsGet200ResponseResult) SetIsPersonal(v bool)`

SetIsPersonal sets IsPersonal field to given value.


### GetDescription

`func (o *WorkspacesDetailsGet200ResponseResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspacesDetailsGet200ResponseResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *WorkspacesDetailsGet200ResponseResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkspacesDetailsGet200ResponseResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStorageLimit

`func (o *WorkspacesDetailsGet200ResponseResult) GetStorageLimit() string`

GetStorageLimit returns the StorageLimit field if non-nil, zero value otherwise.

### GetStorageLimitOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetStorageLimitOk() (*string, bool)`

GetStorageLimitOk returns a tuple with the StorageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLimit

`func (o *WorkspacesDetailsGet200ResponseResult) SetStorageLimit(v string)`

SetStorageLimit sets StorageLimit field to given value.


### GetStorageUsed

`func (o *WorkspacesDetailsGet200ResponseResult) GetStorageUsed() string`

GetStorageUsed returns the StorageUsed field if non-nil, zero value otherwise.

### GetStorageUsedOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetStorageUsedOk() (*string, bool)`

GetStorageUsedOk returns a tuple with the StorageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsed

`func (o *WorkspacesDetailsGet200ResponseResult) SetStorageUsed(v string)`

SetStorageUsed sets StorageUsed field to given value.


### GetCreatedAt

`func (o *WorkspacesDetailsGet200ResponseResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspacesDetailsGet200ResponseResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WorkspacesDetailsGet200ResponseResult) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspacesDetailsGet200ResponseResult) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTeam

`func (o *WorkspacesDetailsGet200ResponseResult) GetTeam() WorkspacesDetailsGet200ResponseResultTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WorkspacesDetailsGet200ResponseResult) GetTeamOk() (*WorkspacesDetailsGet200ResponseResultTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WorkspacesDetailsGet200ResponseResult) SetTeam(v WorkspacesDetailsGet200ResponseResultTeam)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



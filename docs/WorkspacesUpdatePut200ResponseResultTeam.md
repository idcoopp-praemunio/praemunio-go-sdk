# WorkspacesUpdatePut200ResponseResultTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PublicId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Slug** | **string** |  | 
**Color** | **string** |  | 
**IsActive** | **bool** |  | 
**Type** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**TeamMembers** | [**[]WorkspacesUpdatePut200ResponseResultTeamTeamMembersInner**](WorkspacesUpdatePut200ResponseResultTeamTeamMembersInner.md) |  | 

## Methods

### NewWorkspacesUpdatePut200ResponseResultTeam

`func NewWorkspacesUpdatePut200ResponseResultTeam(id int32, publicId string, name string, description string, slug string, color string, isActive bool, type_ string, createdAt string, updatedAt string, teamMembers []WorkspacesUpdatePut200ResponseResultTeamTeamMembersInner, ) *WorkspacesUpdatePut200ResponseResultTeam`

NewWorkspacesUpdatePut200ResponseResultTeam instantiates a new WorkspacesUpdatePut200ResponseResultTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspacesUpdatePut200ResponseResultTeamWithDefaults

`func NewWorkspacesUpdatePut200ResponseResultTeamWithDefaults() *WorkspacesUpdatePut200ResponseResultTeam`

NewWorkspacesUpdatePut200ResponseResultTeamWithDefaults instantiates a new WorkspacesUpdatePut200ResponseResultTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetId(v int32)`

SetId sets Id field to given value.


### GetPublicId

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetName

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSlug

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetColor

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetColor(v string)`

SetColor sets Color field to given value.


### GetIsActive

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetType

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTeamMembers

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetTeamMembers() []WorkspacesUpdatePut200ResponseResultTeamTeamMembersInner`

GetTeamMembers returns the TeamMembers field if non-nil, zero value otherwise.

### GetTeamMembersOk

`func (o *WorkspacesUpdatePut200ResponseResultTeam) GetTeamMembersOk() (*[]WorkspacesUpdatePut200ResponseResultTeamTeamMembersInner, bool)`

GetTeamMembersOk returns a tuple with the TeamMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMembers

`func (o *WorkspacesUpdatePut200ResponseResultTeam) SetTeamMembers(v []WorkspacesUpdatePut200ResponseResultTeamTeamMembersInner)`

SetTeamMembers sets TeamMembers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



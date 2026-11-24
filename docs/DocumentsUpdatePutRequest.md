# DocumentsUpdatePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID | 
**Name** | Pointer to **NullableString** | name | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDocumentsUpdatePutRequest

`func NewDocumentsUpdatePutRequest(id string, ) *DocumentsUpdatePutRequest`

NewDocumentsUpdatePutRequest instantiates a new DocumentsUpdatePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsUpdatePutRequestWithDefaults

`func NewDocumentsUpdatePutRequestWithDefaults() *DocumentsUpdatePutRequest`

NewDocumentsUpdatePutRequestWithDefaults instantiates a new DocumentsUpdatePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentsUpdatePutRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentsUpdatePutRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentsUpdatePutRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DocumentsUpdatePutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentsUpdatePutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentsUpdatePutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentsUpdatePutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DocumentsUpdatePutRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DocumentsUpdatePutRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *DocumentsUpdatePutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentsUpdatePutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentsUpdatePutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocumentsUpdatePutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DocumentsUpdatePutRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DocumentsUpdatePutRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFolderId

`func (o *DocumentsUpdatePutRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DocumentsUpdatePutRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DocumentsUpdatePutRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DocumentsUpdatePutRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *DocumentsUpdatePutRequest) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *DocumentsUpdatePutRequest) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TeamsUpdatePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name | 
**Description** | Pointer to **NullableString** | Team description | [optional] 
**Type** | **string** | Type of your team | 
**Id** | **string** |  | 

## Methods

### NewTeamsUpdatePutRequest

`func NewTeamsUpdatePutRequest(name string, type_ string, id string, ) *TeamsUpdatePutRequest`

NewTeamsUpdatePutRequest instantiates a new TeamsUpdatePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsUpdatePutRequestWithDefaults

`func NewTeamsUpdatePutRequestWithDefaults() *TeamsUpdatePutRequest`

NewTeamsUpdatePutRequestWithDefaults instantiates a new TeamsUpdatePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamsUpdatePutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamsUpdatePutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamsUpdatePutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamsUpdatePutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamsUpdatePutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamsUpdatePutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamsUpdatePutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TeamsUpdatePutRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TeamsUpdatePutRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *TeamsUpdatePutRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamsUpdatePutRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamsUpdatePutRequest) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *TeamsUpdatePutRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamsUpdatePutRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamsUpdatePutRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



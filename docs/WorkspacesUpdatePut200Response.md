# WorkspacesUpdatePut200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Result** | [**WorkspacesUpdatePut200ResponseResult**](WorkspacesUpdatePut200ResponseResult.md) |  | 
**Errors** | **map[string]interface{}** |  | 
**Except** | **map[string]interface{}** |  | 
**Meta** | [**WorkspacesCreatePost201ResponseMeta**](WorkspacesCreatePost201ResponseMeta.md) |  | 

## Methods

### NewWorkspacesUpdatePut200Response

`func NewWorkspacesUpdatePut200Response(success bool, message string, result WorkspacesUpdatePut200ResponseResult, errors map[string]interface{}, except map[string]interface{}, meta WorkspacesCreatePost201ResponseMeta, ) *WorkspacesUpdatePut200Response`

NewWorkspacesUpdatePut200Response instantiates a new WorkspacesUpdatePut200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspacesUpdatePut200ResponseWithDefaults

`func NewWorkspacesUpdatePut200ResponseWithDefaults() *WorkspacesUpdatePut200Response`

NewWorkspacesUpdatePut200ResponseWithDefaults instantiates a new WorkspacesUpdatePut200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *WorkspacesUpdatePut200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WorkspacesUpdatePut200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WorkspacesUpdatePut200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *WorkspacesUpdatePut200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkspacesUpdatePut200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkspacesUpdatePut200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *WorkspacesUpdatePut200Response) GetResult() WorkspacesUpdatePut200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkspacesUpdatePut200Response) GetResultOk() (*WorkspacesUpdatePut200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkspacesUpdatePut200Response) SetResult(v WorkspacesUpdatePut200ResponseResult)`

SetResult sets Result field to given value.


### GetErrors

`func (o *WorkspacesUpdatePut200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WorkspacesUpdatePut200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WorkspacesUpdatePut200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.


### GetExcept

`func (o *WorkspacesUpdatePut200Response) GetExcept() map[string]interface{}`

GetExcept returns the Except field if non-nil, zero value otherwise.

### GetExceptOk

`func (o *WorkspacesUpdatePut200Response) GetExceptOk() (*map[string]interface{}, bool)`

GetExceptOk returns a tuple with the Except field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcept

`func (o *WorkspacesUpdatePut200Response) SetExcept(v map[string]interface{})`

SetExcept sets Except field to given value.


### GetMeta

`func (o *WorkspacesUpdatePut200Response) GetMeta() WorkspacesCreatePost201ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WorkspacesUpdatePut200Response) GetMetaOk() (*WorkspacesCreatePost201ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WorkspacesUpdatePut200Response) SetMeta(v WorkspacesCreatePost201ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



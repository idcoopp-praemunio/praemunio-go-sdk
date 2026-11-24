# FoldersDetailsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Result** | [**FoldersDetailsGet200ResponseResult**](FoldersDetailsGet200ResponseResult.md) |  | 
**Errors** | **map[string]interface{}** |  | 
**Except** | **map[string]interface{}** |  | 
**Meta** | [**WorkspacesCreatePost201ResponseMeta**](WorkspacesCreatePost201ResponseMeta.md) |  | 

## Methods

### NewFoldersDetailsGet200Response

`func NewFoldersDetailsGet200Response(success bool, message string, result FoldersDetailsGet200ResponseResult, errors map[string]interface{}, except map[string]interface{}, meta WorkspacesCreatePost201ResponseMeta, ) *FoldersDetailsGet200Response`

NewFoldersDetailsGet200Response instantiates a new FoldersDetailsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersDetailsGet200ResponseWithDefaults

`func NewFoldersDetailsGet200ResponseWithDefaults() *FoldersDetailsGet200Response`

NewFoldersDetailsGet200ResponseWithDefaults instantiates a new FoldersDetailsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *FoldersDetailsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FoldersDetailsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FoldersDetailsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *FoldersDetailsGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FoldersDetailsGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FoldersDetailsGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *FoldersDetailsGet200Response) GetResult() FoldersDetailsGet200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FoldersDetailsGet200Response) GetResultOk() (*FoldersDetailsGet200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FoldersDetailsGet200Response) SetResult(v FoldersDetailsGet200ResponseResult)`

SetResult sets Result field to given value.


### GetErrors

`func (o *FoldersDetailsGet200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FoldersDetailsGet200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FoldersDetailsGet200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.


### GetExcept

`func (o *FoldersDetailsGet200Response) GetExcept() map[string]interface{}`

GetExcept returns the Except field if non-nil, zero value otherwise.

### GetExceptOk

`func (o *FoldersDetailsGet200Response) GetExceptOk() (*map[string]interface{}, bool)`

GetExceptOk returns a tuple with the Except field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcept

`func (o *FoldersDetailsGet200Response) SetExcept(v map[string]interface{})`

SetExcept sets Except field to given value.


### GetMeta

`func (o *FoldersDetailsGet200Response) GetMeta() WorkspacesCreatePost201ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FoldersDetailsGet200Response) GetMetaOk() (*WorkspacesCreatePost201ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FoldersDetailsGet200Response) SetMeta(v WorkspacesCreatePost201ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



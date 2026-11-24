# FoldersContentGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Result** | [**FoldersContentGet200ResponseResult**](FoldersContentGet200ResponseResult.md) |  | 
**Errors** | **map[string]interface{}** |  | 
**Except** | **map[string]interface{}** |  | 
**Meta** | [**WorkspacesCreatePost201ResponseMeta**](WorkspacesCreatePost201ResponseMeta.md) |  | 

## Methods

### NewFoldersContentGet200Response

`func NewFoldersContentGet200Response(success bool, message string, result FoldersContentGet200ResponseResult, errors map[string]interface{}, except map[string]interface{}, meta WorkspacesCreatePost201ResponseMeta, ) *FoldersContentGet200Response`

NewFoldersContentGet200Response instantiates a new FoldersContentGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersContentGet200ResponseWithDefaults

`func NewFoldersContentGet200ResponseWithDefaults() *FoldersContentGet200Response`

NewFoldersContentGet200ResponseWithDefaults instantiates a new FoldersContentGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *FoldersContentGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FoldersContentGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FoldersContentGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *FoldersContentGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FoldersContentGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FoldersContentGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *FoldersContentGet200Response) GetResult() FoldersContentGet200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FoldersContentGet200Response) GetResultOk() (*FoldersContentGet200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FoldersContentGet200Response) SetResult(v FoldersContentGet200ResponseResult)`

SetResult sets Result field to given value.


### GetErrors

`func (o *FoldersContentGet200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FoldersContentGet200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FoldersContentGet200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.


### GetExcept

`func (o *FoldersContentGet200Response) GetExcept() map[string]interface{}`

GetExcept returns the Except field if non-nil, zero value otherwise.

### GetExceptOk

`func (o *FoldersContentGet200Response) GetExceptOk() (*map[string]interface{}, bool)`

GetExceptOk returns a tuple with the Except field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcept

`func (o *FoldersContentGet200Response) SetExcept(v map[string]interface{})`

SetExcept sets Except field to given value.


### GetMeta

`func (o *FoldersContentGet200Response) GetMeta() WorkspacesCreatePost201ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FoldersContentGet200Response) GetMetaOk() (*WorkspacesCreatePost201ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FoldersContentGet200Response) SetMeta(v WorkspacesCreatePost201ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



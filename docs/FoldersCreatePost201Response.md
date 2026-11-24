# FoldersCreatePost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Result** | [**FoldersCreatePost201ResponseResult**](FoldersCreatePost201ResponseResult.md) |  | 
**Errors** | **map[string]interface{}** |  | 
**Except** | **map[string]interface{}** |  | 
**Meta** | [**WorkspacesCreatePost201ResponseMeta**](WorkspacesCreatePost201ResponseMeta.md) |  | 

## Methods

### NewFoldersCreatePost201Response

`func NewFoldersCreatePost201Response(success bool, message string, result FoldersCreatePost201ResponseResult, errors map[string]interface{}, except map[string]interface{}, meta WorkspacesCreatePost201ResponseMeta, ) *FoldersCreatePost201Response`

NewFoldersCreatePost201Response instantiates a new FoldersCreatePost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersCreatePost201ResponseWithDefaults

`func NewFoldersCreatePost201ResponseWithDefaults() *FoldersCreatePost201Response`

NewFoldersCreatePost201ResponseWithDefaults instantiates a new FoldersCreatePost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *FoldersCreatePost201Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FoldersCreatePost201Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FoldersCreatePost201Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *FoldersCreatePost201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FoldersCreatePost201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FoldersCreatePost201Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *FoldersCreatePost201Response) GetResult() FoldersCreatePost201ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FoldersCreatePost201Response) GetResultOk() (*FoldersCreatePost201ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FoldersCreatePost201Response) SetResult(v FoldersCreatePost201ResponseResult)`

SetResult sets Result field to given value.


### GetErrors

`func (o *FoldersCreatePost201Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FoldersCreatePost201Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FoldersCreatePost201Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.


### GetExcept

`func (o *FoldersCreatePost201Response) GetExcept() map[string]interface{}`

GetExcept returns the Except field if non-nil, zero value otherwise.

### GetExceptOk

`func (o *FoldersCreatePost201Response) GetExceptOk() (*map[string]interface{}, bool)`

GetExceptOk returns a tuple with the Except field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcept

`func (o *FoldersCreatePost201Response) SetExcept(v map[string]interface{})`

SetExcept sets Except field to given value.


### GetMeta

`func (o *FoldersCreatePost201Response) GetMeta() WorkspacesCreatePost201ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FoldersCreatePost201Response) GetMetaOk() (*WorkspacesCreatePost201ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FoldersCreatePost201Response) SetMeta(v WorkspacesCreatePost201ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



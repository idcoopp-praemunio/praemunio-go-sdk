# DocumentsArchivePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** | ID | 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDocumentsArchivePostRequest

`func NewDocumentsArchivePostRequest(documentId string, ) *DocumentsArchivePostRequest`

NewDocumentsArchivePostRequest instantiates a new DocumentsArchivePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsArchivePostRequestWithDefaults

`func NewDocumentsArchivePostRequestWithDefaults() *DocumentsArchivePostRequest`

NewDocumentsArchivePostRequestWithDefaults instantiates a new DocumentsArchivePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *DocumentsArchivePostRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DocumentsArchivePostRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DocumentsArchivePostRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetReason

`func (o *DocumentsArchivePostRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DocumentsArchivePostRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DocumentsArchivePostRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DocumentsArchivePostRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *DocumentsArchivePostRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *DocumentsArchivePostRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



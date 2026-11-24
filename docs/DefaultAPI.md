# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthAuthTokenPost**](DefaultAPI.md#AuthAuthTokenPost) | **Post** /auth/auth-token | API KEY Authentification
[**DocumentsArchivePost**](DefaultAPI.md#DocumentsArchivePost) | **Post** /documents/archive | Archive Document
[**DocumentsDeleteDelete**](DefaultAPI.md#DocumentsDeleteDelete) | **Delete** /documents/delete | Delete Document
[**DocumentsDetailsGet**](DefaultAPI.md#DocumentsDetailsGet) | **Get** /documents/details | Get Document Info
[**DocumentsGet**](DefaultAPI.md#DocumentsGet) | **Get** /documents | Get Document By Workspace
[**DocumentsMovePost**](DefaultAPI.md#DocumentsMovePost) | **Post** /documents/move | Move Document
[**DocumentsRestorePost**](DefaultAPI.md#DocumentsRestorePost) | **Post** /documents/restore | Restore Document 
[**DocumentsStreamGet**](DefaultAPI.md#DocumentsStreamGet) | **Get** /documents/stream | Stream Document Copy
[**DocumentsUpdatePut**](DefaultAPI.md#DocumentsUpdatePut) | **Put** /documents/update | Update Document
[**DocumentsUploadPost**](DefaultAPI.md#DocumentsUploadPost) | **Post** /documents/upload | Upload Document 
[**FoldersContentGet**](DefaultAPI.md#FoldersContentGet) | **Get** /folders/content | Get Folder Content
[**FoldersCreatePost**](DefaultAPI.md#FoldersCreatePost) | **Post** /folders/create | Create Folder
[**FoldersDeleteDelete**](DefaultAPI.md#FoldersDeleteDelete) | **Delete** /folders/delete | Delete Folder
[**FoldersDetailsGet**](DefaultAPI.md#FoldersDetailsGet) | **Get** /folders/details | Get Folder Details
[**FoldersGet**](DefaultAPI.md#FoldersGet) | **Get** /folders | Get Folders 
[**FoldersMovePost**](DefaultAPI.md#FoldersMovePost) | **Post** /folders/move | Move Folder
[**FoldersUpdatePut**](DefaultAPI.md#FoldersUpdatePut) | **Put** /folders/update | Update Folder
[**TeamsCreatePost**](DefaultAPI.md#TeamsCreatePost) | **Post** /teams/create | Create Team
[**TeamsDashboardGet**](DefaultAPI.md#TeamsDashboardGet) | **Get** /teams/dashboard | Get Team Data
[**TeamsDeleteDelete**](DefaultAPI.md#TeamsDeleteDelete) | **Delete** /teams/delete | Delete Team
[**TeamsGet**](DefaultAPI.md#TeamsGet) | **Get** /teams | Get Account Teams
[**TeamsInfoGet**](DefaultAPI.md#TeamsInfoGet) | **Get** /teams/info | Get Team Details
[**TeamsLogsGet**](DefaultAPI.md#TeamsLogsGet) | **Get** /teams/logs | Get Team logs
[**TeamsMembersAddPost**](DefaultAPI.md#TeamsMembersAddPost) | **Post** /teams/members/add | Add Team Member Copy
[**TeamsMembersGet**](DefaultAPI.md#TeamsMembersGet) | **Get** /teams/members | Get Team Members
[**TeamsUpdatePut**](DefaultAPI.md#TeamsUpdatePut) | **Put** /teams/update | Update Team Data
[**TeamsWorkspacesGet**](DefaultAPI.md#TeamsWorkspacesGet) | **Get** /teams/workspaces | Get Team workspaces
[**WorkspacesCreatePost**](DefaultAPI.md#WorkspacesCreatePost) | **Post** /workspaces/create | Create New Workspace
[**WorkspacesDeleteDelete**](DefaultAPI.md#WorkspacesDeleteDelete) | **Delete** /workspaces/delete | Delete workspace
[**WorkspacesDetailsGet**](DefaultAPI.md#WorkspacesDetailsGet) | **Get** /workspaces/details | Get Workspace Details
[**WorkspacesDriveGet**](DefaultAPI.md#WorkspacesDriveGet) | **Get** /workspaces/drive | Get Account Drive
[**WorkspacesGet**](DefaultAPI.md#WorkspacesGet) | **Get** /workspaces | Get Workspaces By Team
[**WorkspacesTeamsGet**](DefaultAPI.md#WorkspacesTeamsGet) | **Get** /workspaces/teams | Get Team Workspaces
[**WorkspacesUpdatePut**](DefaultAPI.md#WorkspacesUpdatePut) | **Put** /workspaces/update | Update workspace



## AuthAuthTokenPost

> map[string]interface{} AuthAuthTokenPost(ctx).AuthAuthTokenPostRequest(authAuthTokenPostRequest).Execute()

API KEY Authentification



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authAuthTokenPostRequest := *openapiclient.NewAuthAuthTokenPostRequest("PublicKey_example", "SecretKey_example") // AuthAuthTokenPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AuthAuthTokenPost(context.Background()).AuthAuthTokenPostRequest(authAuthTokenPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AuthAuthTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAuthTokenPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AuthAuthTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authAuthTokenPostRequest** | [**AuthAuthTokenPostRequest**](AuthAuthTokenPostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsArchivePost

> map[string]interface{} DocumentsArchivePost(ctx).DocumentsArchivePostRequest(documentsArchivePostRequest).Execute()

Archive Document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	documentsArchivePostRequest := *openapiclient.NewDocumentsArchivePostRequest("DocumentId_example") // DocumentsArchivePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsArchivePost(context.Background()).DocumentsArchivePostRequest(documentsArchivePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsArchivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsArchivePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsArchivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsArchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentsArchivePostRequest** | [**DocumentsArchivePostRequest**](DocumentsArchivePostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsDeleteDelete

> map[string]interface{} DocumentsDeleteDelete(ctx).DocumentsDeleteDeleteRequest(documentsDeleteDeleteRequest).Execute()

Delete Document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	documentsDeleteDeleteRequest := *openapiclient.NewDocumentsDeleteDeleteRequest("Id_example") // DocumentsDeleteDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsDeleteDelete(context.Background()).DocumentsDeleteDeleteRequest(documentsDeleteDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsDeleteDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsDeleteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentsDeleteDeleteRequest** | [**DocumentsDeleteDeleteRequest**](DocumentsDeleteDeleteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsDetailsGet

> map[string]interface{} DocumentsDetailsGet(ctx).DocumentId(documentId).Execute()

Get Document Info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	documentId := "a3f8e9d2-1a4b-4c7e-9f3a-2b8c6d4e1f0a" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsDetailsGet(context.Background()).DocumentId(documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsDetailsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsGet

> map[string]interface{} DocumentsGet(ctx).WorkspaceId(workspaceId).Execute()

Get Document By Workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "94506d68-6e3b-479f-83d2-5401b779b565" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsGet(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsMovePost

> map[string]interface{} DocumentsMovePost(ctx).DocumentsMovePostRequest(documentsMovePostRequest).Execute()

Move Document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	documentsMovePostRequest := *openapiclient.NewDocumentsMovePostRequest("DocumentId_example", "WorkspaceId_example") // DocumentsMovePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsMovePost(context.Background()).DocumentsMovePostRequest(documentsMovePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsMovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsMovePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsMovePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsMovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentsMovePostRequest** | [**DocumentsMovePostRequest**](DocumentsMovePostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsRestorePost

> map[string]interface{} DocumentsRestorePost(ctx).DocumentsRestorePostRequest(documentsRestorePostRequest).Execute()

Restore Document 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	documentsRestorePostRequest := *openapiclient.NewDocumentsRestorePostRequest("DocumentId_example") // DocumentsRestorePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsRestorePost(context.Background()).DocumentsRestorePostRequest(documentsRestorePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsRestorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsRestorePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsRestorePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsRestorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentsRestorePostRequest** | [**DocumentsRestorePostRequest**](DocumentsRestorePostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsStreamGet

> map[string]interface{} DocumentsStreamGet(ctx).Body(body).DocumentId(documentId).Execute()

Stream Document Copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	documentId := "731a34c8-67b4-4fa4-ae44-594c9124f322" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsStreamGet(context.Background()).Body(body).DocumentId(documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsStreamGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsStreamGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsStreamGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsStreamGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **documentId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsUpdatePut

> map[string]interface{} DocumentsUpdatePut(ctx).DocumentsUpdatePutRequest(documentsUpdatePutRequest).Execute()

Update Document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	documentsUpdatePutRequest := *openapiclient.NewDocumentsUpdatePutRequest("Id_example") // DocumentsUpdatePutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsUpdatePut(context.Background()).DocumentsUpdatePutRequest(documentsUpdatePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsUpdatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsUpdatePut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsUpdatePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsUpdatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentsUpdatePutRequest** | [**DocumentsUpdatePutRequest**](DocumentsUpdatePutRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsUploadPost

> map[string]interface{} DocumentsUploadPost(ctx).WorkspaceId(workspaceId).FolderId(folderId).File1(file1).File2(file2).Execute()

Upload Document 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "731a34c8-67b4-4fa4-ae44-594c9124f322" // string |  (optional)
	folderId := "731a34c8-67b4-4fa4-ae44-594c9124f322" // string |  (optional)
	file1 := os.NewFile(1234, "some_file") // *os.File |  (optional)
	file2 := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DocumentsUploadPost(context.Background()).WorkspaceId(workspaceId).FolderId(folderId).File1(file1).File2(file2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DocumentsUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentsUploadPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DocumentsUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 
 **folderId** | **string** |  | 
 **file1** | ***os.File** |  | 
 **file2** | ***os.File** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersContentGet

> FoldersContentGet200Response FoldersContentGet(ctx).FolderId(folderId).Execute()

Get Folder Content



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	folderId := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FoldersContentGet(context.Background()).FolderId(folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FoldersContentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersContentGet`: FoldersContentGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FoldersContentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFoldersContentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **[]string** |  | 

### Return type

[**FoldersContentGet200Response**](FoldersContentGet200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersCreatePost

> FoldersCreatePost201Response FoldersCreatePost(ctx).FoldersCreatePostRequest(foldersCreatePostRequest).Execute()

Create Folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	foldersCreatePostRequest := *openapiclient.NewFoldersCreatePostRequest("WorkspaceId_example", "Name_example") // FoldersCreatePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FoldersCreatePost(context.Background()).FoldersCreatePostRequest(foldersCreatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FoldersCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersCreatePost`: FoldersCreatePost201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FoldersCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFoldersCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foldersCreatePostRequest** | [**FoldersCreatePostRequest**](FoldersCreatePostRequest.md) |  | 

### Return type

[**FoldersCreatePost201Response**](FoldersCreatePost201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersDeleteDelete

> map[string]interface{} FoldersDeleteDelete(ctx).FoldersDeleteDeleteRequest(foldersDeleteDeleteRequest).Execute()

Delete Folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	foldersDeleteDeleteRequest := *openapiclient.NewFoldersDeleteDeleteRequest("Id_example") // FoldersDeleteDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FoldersDeleteDelete(context.Background()).FoldersDeleteDeleteRequest(foldersDeleteDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FoldersDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersDeleteDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FoldersDeleteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFoldersDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foldersDeleteDeleteRequest** | [**FoldersDeleteDeleteRequest**](FoldersDeleteDeleteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersDetailsGet

> FoldersDetailsGet200Response FoldersDetailsGet(ctx).FolderId(folderId).Execute()

Get Folder Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	folderId := "7935763f-dac9-4bfe-be30-ff7dc0124f16" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FoldersDetailsGet(context.Background()).FolderId(folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FoldersDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersDetailsGet`: FoldersDetailsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FoldersDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFoldersDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string** |  | 

### Return type

[**FoldersDetailsGet200Response**](FoldersDetailsGet200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersGet

> FoldersGet200Response FoldersGet(ctx).WorkspaceId(workspaceId).Execute()

Get Folders 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "94506d68-6e3b-479f-83d2-5401b779b565" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FoldersGet(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FoldersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersGet`: FoldersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FoldersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFoldersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 

### Return type

[**FoldersGet200Response**](FoldersGet200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersMovePost

> map[string]interface{} FoldersMovePost(ctx).FoldersDeleteDeleteRequest(foldersDeleteDeleteRequest).Execute()

Move Folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	foldersDeleteDeleteRequest := *openapiclient.NewFoldersDeleteDeleteRequest("Id_example") // FoldersDeleteDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FoldersMovePost(context.Background()).FoldersDeleteDeleteRequest(foldersDeleteDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FoldersMovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersMovePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FoldersMovePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFoldersMovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foldersDeleteDeleteRequest** | [**FoldersDeleteDeleteRequest**](FoldersDeleteDeleteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersUpdatePut

> map[string]interface{} FoldersUpdatePut(ctx).FoldersUpdatePutRequest(foldersUpdatePutRequest).Execute()

Update Folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	foldersUpdatePutRequest := *openapiclient.NewFoldersUpdatePutRequest("Id_example", "Name_example", "Description_example") // FoldersUpdatePutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FoldersUpdatePut(context.Background()).FoldersUpdatePutRequest(foldersUpdatePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FoldersUpdatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersUpdatePut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FoldersUpdatePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFoldersUpdatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foldersUpdatePutRequest** | [**FoldersUpdatePutRequest**](FoldersUpdatePutRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreatePost

> map[string]interface{} TeamsCreatePost(ctx).TeamsCreatePostRequest(teamsCreatePostRequest).Execute()

Create Team



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamsCreatePostRequest := *openapiclient.NewTeamsCreatePostRequest("Name_example", "Type_example") // TeamsCreatePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsCreatePost(context.Background()).TeamsCreatePostRequest(teamsCreatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsCreatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamsCreatePostRequest** | [**TeamsCreatePostRequest**](TeamsCreatePostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsDashboardGet

> map[string]interface{} TeamsDashboardGet(ctx).Execute()

Get Team Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsDashboardGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsDashboardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsDashboardGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsDashboardGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDashboardGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsDeleteDelete

> map[string]interface{} TeamsDeleteDelete(ctx).TeamsDeleteDeleteRequest(teamsDeleteDeleteRequest).Execute()

Delete Team



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamsDeleteDeleteRequest := *openapiclient.NewTeamsDeleteDeleteRequest("TeamId_example") // TeamsDeleteDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsDeleteDelete(context.Background()).TeamsDeleteDeleteRequest(teamsDeleteDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsDeleteDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsDeleteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamsDeleteDeleteRequest** | [**TeamsDeleteDeleteRequest**](TeamsDeleteDeleteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGet

> map[string]interface{} TeamsGet(ctx).Execute()

Get Account Teams



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsInfoGet

> map[string]interface{} TeamsInfoGet(ctx).Body(body).TeamId(teamId).Execute()

Get Team Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	teamId := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsInfoGet(context.Background()).Body(body).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsInfoGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsInfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **teamId** | **[]string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsLogsGet

> map[string]interface{} TeamsLogsGet(ctx).TeamId(teamId).Page(page).Limit(limit).Execute()

Get Team logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamId := []string{"Inner_example"} // []string | team Id (optional)
	page := "1" // string |  (optional)
	limit := "50" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsLogsGet(context.Background()).TeamId(teamId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsLogsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsLogsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **[]string** | team Id | 
 **page** | **string** |  | 
 **limit** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsMembersAddPost

> map[string]interface{} TeamsMembersAddPost(ctx).TeamsMembersAddPostRequest(teamsMembersAddPostRequest).Execute()

Add Team Member Copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamsMembersAddPostRequest := *openapiclient.NewTeamsMembersAddPostRequest("TeamId_example") // TeamsMembersAddPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsMembersAddPost(context.Background()).TeamsMembersAddPostRequest(teamsMembersAddPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsMembersAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsMembersAddPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsMembersAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsMembersAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamsMembersAddPostRequest** | [**TeamsMembersAddPostRequest**](TeamsMembersAddPostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsMembersGet

> map[string]interface{} TeamsMembersGet(ctx).TeamId(teamId).Page(page).Limit(limit).Execute()

Get Team Members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamId := "1" // string |  (optional)
	page := float32(1) // float32 |  (optional)
	limit := float32(10) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsMembersGet(context.Background()).TeamId(teamId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsMembersGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsMembersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string** |  | 
 **page** | **float32** |  | 
 **limit** | **float32** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdatePut

> map[string]interface{} TeamsUpdatePut(ctx).TeamsUpdatePutRequest(teamsUpdatePutRequest).Execute()

Update Team Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamsUpdatePutRequest := *openapiclient.NewTeamsUpdatePutRequest("Name_example", "Type_example", "Id_example") // TeamsUpdatePutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsUpdatePut(context.Background()).TeamsUpdatePutRequest(teamsUpdatePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsUpdatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsUpdatePut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsUpdatePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamsUpdatePutRequest** | [**TeamsUpdatePutRequest**](TeamsUpdatePutRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsWorkspacesGet

> map[string]interface{} TeamsWorkspacesGet(ctx).TeamId(teamId).Execute()

Get Team workspaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamId := "64eefb76-8b3a-4f7d-92c0-7f76370d0ec0" // string | team Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TeamsWorkspacesGet(context.Background()).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TeamsWorkspacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsWorkspacesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TeamsWorkspacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsWorkspacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string** | team Id | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesCreatePost

> WorkspacesCreatePost201Response WorkspacesCreatePost(ctx).WorkspacesCreatePostRequest(workspacesCreatePostRequest).Execute()

Create New Workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspacesCreatePostRequest := *openapiclient.NewWorkspacesCreatePostRequest("Name_example", "Description_example", "TeamId_example") // WorkspacesCreatePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WorkspacesCreatePost(context.Background()).WorkspacesCreatePostRequest(workspacesCreatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkspacesCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesCreatePost`: WorkspacesCreatePost201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkspacesCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspacesCreatePostRequest** | [**WorkspacesCreatePostRequest**](WorkspacesCreatePostRequest.md) |  | 

### Return type

[**WorkspacesCreatePost201Response**](WorkspacesCreatePost201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesDeleteDelete

> map[string]interface{} WorkspacesDeleteDelete(ctx).WorkspacesDeleteDeleteRequest(workspacesDeleteDeleteRequest).Execute()

Delete workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspacesDeleteDeleteRequest := *openapiclient.NewWorkspacesDeleteDeleteRequest("WorkspaceId_example") // WorkspacesDeleteDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WorkspacesDeleteDelete(context.Background()).WorkspacesDeleteDeleteRequest(workspacesDeleteDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkspacesDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesDeleteDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkspacesDeleteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspacesDeleteDeleteRequest** | [**WorkspacesDeleteDeleteRequest**](WorkspacesDeleteDeleteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesDetailsGet

> WorkspacesDetailsGet200Response WorkspacesDetailsGet(ctx).WorkspaceId(workspaceId).Execute()

Get Workspace Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "94506d68-6e3b-479f-83d2-5401b779b565" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WorkspacesDetailsGet(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkspacesDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesDetailsGet`: WorkspacesDetailsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkspacesDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 

### Return type

[**WorkspacesDetailsGet200Response**](WorkspacesDetailsGet200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesDriveGet

> map[string]interface{} WorkspacesDriveGet(ctx).Execute()

Get Account Drive



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WorkspacesDriveGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkspacesDriveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesDriveGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkspacesDriveGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesDriveGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesGet

> map[string]interface{} WorkspacesGet(ctx).TeamId(teamId).Execute()

Get Workspaces By Team



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamId := "64eefb76-8b3a-4f7d-92c0-7f76370d0ec0" // string | team Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WorkspacesGet(context.Background()).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkspacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkspacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string** | team Id | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesTeamsGet

> WorkspacesTeamsGet200Response WorkspacesTeamsGet(ctx).TeamId(teamId).Execute()

Get Team Workspaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamId := "0f9dfdec-7cce-4ed3-b7d5-672d33b18e8e" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WorkspacesTeamsGet(context.Background()).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkspacesTeamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesTeamsGet`: WorkspacesTeamsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkspacesTeamsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string** |  | 

### Return type

[**WorkspacesTeamsGet200Response**](WorkspacesTeamsGet200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesUpdatePut

> WorkspacesUpdatePut200Response WorkspacesUpdatePut(ctx).WorkspacesUpdatePutRequest(workspacesUpdatePutRequest).Execute()

Update workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspacesUpdatePutRequest := *openapiclient.NewWorkspacesUpdatePutRequest("Id_example", "Name_example", "Description_example", "Status_example") // WorkspacesUpdatePutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WorkspacesUpdatePut(context.Background()).WorkspacesUpdatePutRequest(workspacesUpdatePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkspacesUpdatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesUpdatePut`: WorkspacesUpdatePut200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkspacesUpdatePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesUpdatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspacesUpdatePutRequest** | [**WorkspacesUpdatePutRequest**](WorkspacesUpdatePutRequest.md) |  | 

### Return type

[**WorkspacesUpdatePut200Response**](WorkspacesUpdatePut200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


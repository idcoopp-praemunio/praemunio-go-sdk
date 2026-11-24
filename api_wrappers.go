package praemunio

import (
	"context"
	"net/http"
	"os"
)

// PraemunioAPIClient wraps the generated API client.
type PraemunioAPIClient struct {
	*APIClient
}

// WorkspacesAPI provides access to workspace-related operations.
type WorkspacesAPI struct {
	client *PraemunioAPIClient
}

// NewWorkspacesAPI creates a new WorkspacesAPI instance.
func NewWorkspacesAPI(client *PraemunioAPIClient) *WorkspacesAPI {
	return &WorkspacesAPI{client: client}
}

// List lists all workspaces, optionally filtered by team.
func (a *WorkspacesAPI) List(ctx context.Context, teamID *string) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.WorkspacesGet(ctx)
	if teamID != nil {
		req = req.TeamId(*teamID)
	}
	return req.Execute()
}

// Drive gets the drive (root folders) for the authenticated user.
func (a *WorkspacesAPI) Drive(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.WorkspacesDriveGet(ctx).Execute()
}

// Teams gets teams associated with workspaces.
func (a *WorkspacesAPI) Teams(ctx context.Context, teamID *string) (*WorkspacesTeamsGet200Response, *http.Response, error) {
	req := a.client.DefaultAPI.WorkspacesTeamsGet(ctx)
	if teamID != nil {
		req = req.TeamId(*teamID)
	}
	return req.Execute()
}

// Details gets workspace details.
func (a *WorkspacesAPI) Details(ctx context.Context, workspaceID *string) (*WorkspacesDetailsGet200Response, *http.Response, error) {
	req := a.client.DefaultAPI.WorkspacesDetailsGet(ctx)
	if workspaceID != nil {
		req = req.WorkspaceId(*workspaceID)
	}
	return req.Execute()
}

// Create creates a new workspace.
func (a *WorkspacesAPI) Create(ctx context.Context, request WorkspacesCreatePostRequest) (*WorkspacesCreatePost201Response, *http.Response, error) {
	return a.client.DefaultAPI.WorkspacesCreatePost(ctx).WorkspacesCreatePostRequest(request).Execute()
}

// Update updates a workspace.
func (a *WorkspacesAPI) Update(ctx context.Context, request WorkspacesUpdatePutRequest) (*WorkspacesUpdatePut200Response, *http.Response, error) {
	return a.client.DefaultAPI.WorkspacesUpdatePut(ctx).WorkspacesUpdatePutRequest(request).Execute()
}

// Delete deletes a workspace.
func (a *WorkspacesAPI) Delete(ctx context.Context, request WorkspacesDeleteDeleteRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.WorkspacesDeleteDelete(ctx).WorkspacesDeleteDeleteRequest(request).Execute()
}

// TeamsAPI provides access to team-related operations.
type TeamsAPI struct {
	client *PraemunioAPIClient
}

// NewTeamsAPI creates a new TeamsAPI instance.
func NewTeamsAPI(client *PraemunioAPIClient) *TeamsAPI {
	return &TeamsAPI{client: client}
}

// List lists all teams.
func (a *TeamsAPI) List(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.TeamsGet(ctx).Execute()
}

// Dashboard gets teams dashboard data.
func (a *TeamsAPI) Dashboard(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.TeamsDashboardGet(ctx).Execute()
}

// Info gets team information.
func (a *TeamsAPI) Info(ctx context.Context, body interface{}, teamID []string) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.TeamsInfoGet(ctx)
	if body != nil {
		req = req.Body(body)
	}
	if teamID != nil {
		req = req.TeamId(teamID)
	}
	return req.Execute()
}

// Workspaces gets workspaces associated with teams.
func (a *TeamsAPI) Workspaces(ctx context.Context, teamID *string) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.TeamsWorkspacesGet(ctx)
	if teamID != nil {
		req = req.TeamId(*teamID)
	}
	return req.Execute()
}

// Logs gets team logs.
func (a *TeamsAPI) Logs(ctx context.Context, teamID []string, page, limit *string) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.TeamsLogsGet(ctx)
	if teamID != nil {
		req = req.TeamId(teamID)
	}
	if page != nil {
		req = req.Page(*page)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	return req.Execute()
}

// Create creates a new team.
func (a *TeamsAPI) Create(ctx context.Context, request TeamsCreatePostRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.TeamsCreatePost(ctx).TeamsCreatePostRequest(request).Execute()
}

// Update updates a team.
func (a *TeamsAPI) Update(ctx context.Context, request TeamsUpdatePutRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.TeamsUpdatePut(ctx).TeamsUpdatePutRequest(request).Execute()
}

// Delete deletes a team.
func (a *TeamsAPI) Delete(ctx context.Context, request TeamsDeleteDeleteRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.TeamsDeleteDelete(ctx).TeamsDeleteDeleteRequest(request).Execute()
}

// TeamsMembersAPI provides access to team member-related operations.
type TeamsMembersAPI struct {
	client *PraemunioAPIClient
}

// NewTeamsMembersAPI creates a new TeamsMembersAPI instance.
func NewTeamsMembersAPI(client *PraemunioAPIClient) *TeamsMembersAPI {
	return &TeamsMembersAPI{client: client}
}

// List lists team members with pagination.
func (a *TeamsMembersAPI) List(ctx context.Context, teamID *string, page, limit *float64) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.TeamsMembersGet(ctx)
	if teamID != nil {
		req = req.TeamId(*teamID)
	}
	if page != nil {
		req = req.Page(*page)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	return req.Execute()
}

// Add adds a member to a team.
func (a *TeamsMembersAPI) Add(ctx context.Context, request TeamsMembersAddPostRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.TeamsMembersAddPost(ctx).TeamsMembersAddPostRequest(request).Execute()
}

// FoldersAPI provides access to folder-related operations.
type FoldersAPI struct {
	client *PraemunioAPIClient
}

// NewFoldersAPI creates a new FoldersAPI instance.
func NewFoldersAPI(client *PraemunioAPIClient) *FoldersAPI {
	return &FoldersAPI{client: client}
}

// List lists folders in a workspace.
func (a *FoldersAPI) List(ctx context.Context, workspaceID string) (*FoldersGet200Response, *http.Response, error) {
	return a.client.DefaultAPI.FoldersGet(ctx).WorkspaceId(workspaceID).Execute()
}

// Content gets folder content.
func (a *FoldersAPI) Content(ctx context.Context, folderID []string) (*FoldersContentGet200Response, *http.Response, error) {
	req := a.client.DefaultAPI.FoldersContentGet(ctx)
	if folderID != nil {
		req = req.FolderId(folderID)
	}
	return req.Execute()
}

// Details gets folder details.
func (a *FoldersAPI) Details(ctx context.Context, folderID *string) (*FoldersDetailsGet200Response, *http.Response, error) {
	req := a.client.DefaultAPI.FoldersDetailsGet(ctx)
	if folderID != nil {
		req = req.FolderId(*folderID)
	}
	return req.Execute()
}

// Create creates a new folder.
func (a *FoldersAPI) Create(ctx context.Context, request FoldersCreatePostRequest) (*FoldersCreatePost201Response, *http.Response, error) {
	return a.client.DefaultAPI.FoldersCreatePost(ctx).FoldersCreatePostRequest(request).Execute()
}

// Update updates a folder.
func (a *FoldersAPI) Update(ctx context.Context, request FoldersUpdatePutRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.FoldersUpdatePut(ctx).FoldersUpdatePutRequest(request).Execute()
}

// Move moves a folder.
func (a *FoldersAPI) Move(ctx context.Context, request FoldersDeleteDeleteRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.FoldersMovePost(ctx).FoldersDeleteDeleteRequest(request).Execute()
}

// Delete deletes a folder.
func (a *FoldersAPI) Delete(ctx context.Context, request FoldersDeleteDeleteRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.FoldersDeleteDelete(ctx).FoldersDeleteDeleteRequest(request).Execute()
}

// DocumentsAPI provides access to document-related operations.
type DocumentsAPI struct {
	client *PraemunioAPIClient
}

// NewDocumentsAPI creates a new DocumentsAPI instance.
func NewDocumentsAPI(client *PraemunioAPIClient) *DocumentsAPI {
	return &DocumentsAPI{client: client}
}

// List lists documents in a workspace.
func (a *DocumentsAPI) List(ctx context.Context, workspaceID *string) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.DocumentsGet(ctx)
	if workspaceID != nil {
		req = req.WorkspaceId(*workspaceID)
	}
	return req.Execute()
}

// Details gets document details.
func (a *DocumentsAPI) Details(ctx context.Context, documentID *string) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.DocumentsDetailsGet(ctx)
	if documentID != nil {
		req = req.DocumentId(*documentID)
	}
	return req.Execute()
}

// Upload uploads a document.
func (a *DocumentsAPI) Upload(ctx context.Context, workspaceID, folderID *string, file1, file2 *os.File) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.DocumentsUploadPost(ctx)
	if workspaceID != nil {
		req = req.WorkspaceId(*workspaceID)
	}
	if folderID != nil {
		req = req.FolderId(*folderID)
	}
	if file1 != nil {
		req = req.File1(*file1)
	}
	if file2 != nil {
		req = req.File2(*file2)
	}
	return req.Execute()
}

// Stream streams a document.
func (a *DocumentsAPI) Stream(ctx context.Context, body interface{}, documentID *string) (map[string]interface{}, *http.Response, error) {
	req := a.client.DefaultAPI.DocumentsStreamGet(ctx)
	if body != nil {
		req = req.Body(body)
	}
	if documentID != nil {
		req = req.DocumentId(*documentID)
	}
	return req.Execute()
}

// Archive archives a document.
func (a *DocumentsAPI) Archive(ctx context.Context, request DocumentsArchivePostRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.DocumentsArchivePost(ctx).DocumentsArchivePostRequest(request).Execute()
}

// Restore restores an archived document.
func (a *DocumentsAPI) Restore(ctx context.Context, request DocumentsRestorePostRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.DocumentsRestorePost(ctx).DocumentsRestorePostRequest(request).Execute()
}

// Move moves a document.
func (a *DocumentsAPI) Move(ctx context.Context, request DocumentsMovePostRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.DocumentsMovePost(ctx).DocumentsMovePostRequest(request).Execute()
}

// Update updates a document.
func (a *DocumentsAPI) Update(ctx context.Context, request DocumentsUpdatePutRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.DocumentsUpdatePut(ctx).DocumentsUpdatePutRequest(request).Execute()
}

// Delete deletes a document.
func (a *DocumentsAPI) Delete(ctx context.Context, request DocumentsDeleteDeleteRequest) (map[string]interface{}, *http.Response, error) {
	return a.client.DefaultAPI.DocumentsDeleteDelete(ctx).DocumentsDeleteDeleteRequest(request).Execute()
}
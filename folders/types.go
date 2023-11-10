package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type PFetchFolderRequest struct {
	OrgID uuid.UUID;
	Token int
}

type PFetchFolderResponse struct {
	Folders []*Folder;
	Token int
}

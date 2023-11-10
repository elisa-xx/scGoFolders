package folders_test

import (
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

func TestGetAllFolders(t *testing.T) { //beginning the function with "Test" tells the testing framework this is a test function
	type args struct {
		req *folders.FetchFolderRequest
	}

	tests := []struct { // test case definition
		name string
		args args
		want *folders.FetchFolderResponse 
	}{ //test cases
		{ //test case
			name: "If org doesn't exist, expect empty folders",
			args: args{ //add values for creating the args struct
				req: &folders.FetchFolderRequest{
					OrgID: uuid.FromStringOrNil("doesn't exist"),
				},
			},
			want: &folders.FetchFolderResponse{
				Folders: []*folders.Folder{},
			},
		},

		{ //test case
			name: "If org has 1 folder, return 1 folder",
			args: args{ //add values for creating the args struct
				req: &folders.FetchFolderRequest{
					OrgID: uuid.FromStringOrNil("52214b35-f4da-461a-9f93-fbd3590e700f"),
				},
			},
			want: &folders.FetchFolderResponse{
				Folders: []*folders.Folder{
					{
						Id:  uuid.FromStringOrNil("9fc98418-7039-4443-a82b-84049ed25d38"),
						Name: "fitting-talisman",
						OrgId: uuid.FromStringOrNil("52214b35-f4da-461a-9f93-fbd3590e700f"),
						Deleted: false,
					},
				},
			},
		},

		{ //test case
			name: "If org has 2 folders, return 2 folders",
			args: args{ //add values for creating the args struct
				req: &folders.FetchFolderRequest{
					OrgID: uuid.FromStringOrNil("3b9a868b-8cd9-4b6b-ba23-fd1e08f3e2fa"),
				},
			},
			want: &folders.FetchFolderResponse{
				Folders: []*folders.Folder{
					{
						Id:  uuid.FromStringOrNil("71702b42-aee8-4c03-a05c-1a0cc5102a83"), //created a second (dummy)folder under this ID for testing.
						Name: "novel-human-robot",
						OrgId: uuid.FromStringOrNil("3b9a868b-8cd9-4b6b-ba23-fd1e08f3e2fa"),
						Deleted: true,
					},{
						Id:  uuid.FromStringOrNil("72813c53-aee8-4c03-a05c-1a0cc5102b92"),
						Name: "only-human-testing",
						OrgId: uuid.FromStringOrNil("3b9a868b-8cd9-4b6b-ba23-fd1e08f3e2fa"),
						Deleted: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := folders.GetAllFolders(tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

}

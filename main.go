package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	/* This original code block runs the improved GetAllFolders function without pagination
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	} */


	//This code block runs the extended GetSomeFolders with pagination implemented
	req := &folders.PFetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		Token: 1,
	}

	res, err := folders.GetSomeFolders(req)
	if err != nil {
		fmt.Printf("%v",err)
		return
	}

	folders.PrettyPrint(res)
}

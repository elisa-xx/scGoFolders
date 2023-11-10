package folders

import (
	"fmt"
)

func GetSomeFolders(req *PFetchFolderRequest) (*PFetchFolderResponse, error) { 
	itemsPerPage := 10
	startIndex := req.Token
	endIndex := req.Token + itemsPerPage -1
	response := []*Folder{}
	r, err := FetchAllFoldersByOrgID(req.OrgID) 
	if err != nil {
		fmt.Println("Error found")
		return nil, err          
	} else if startIndex <0 || startIndex>= len(r){
		return nil, fmt.Errorf("invalid start index")
	} else if endIndex <0  || endIndex>len(r){
		return nil, fmt.Errorf("invalid end index")
	} else if startIndex> endIndex{
		return nil, fmt.Errorf("one or more indices are invalid")
	}
	
	for i:=startIndex; i<=endIndex && i<len(r); i++{
		response = append(response, r[i])
	}
	
	ffr := &PFetchFolderResponse{Folders: response, Token: endIndex+1}  
	
	return ffr, nil
}



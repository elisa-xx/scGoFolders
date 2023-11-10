package folders

import (
	"fmt"

	"github.com/gofrs/uuid" //a package that enables the creating and parsing of unique IDs
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) { //take in a FetchFolderRequest struct (which contains a single OrgID field), return the value located at the *FetchFolderRequest (pointer) address
	r, err := FetchAllFoldersByOrgID(req.OrgID) //var r contains collection of Folder Pointers for Folders that have that OrgID, returned by the FetchAllFoldersByOrgID method. 
	if err != nil {
		fmt.Println("Error found")
		return nil, err          // return error and message if there's an error in fetching by OrgID
	}

	ffr := &FetchFolderResponse{Folders: r}  //Create FetchFolderResponse struct which has a Folders field (where val is the collection of Folder Pointers). Store struct in ffr
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) { //takes in an orgID of type uuid.UUID as argument. Returns a slice of pointers to Folder objects and an error.
	folders := GetSampleData() //from sample.json

	resFolder := []*Folder{}         // create empty slice for Folder pointers
	for _, folder := range folders { //iterate through whatever collection GetSampleData returns.
		if folder.OrgId == orgID { //if the OrgId of folder matches the input orgID,
			resFolder = append(resFolder, folder) //append pointer to resFolder which is a slice for FolderPointers.
		}
	}
	return resFolder, nil //return a slice of FolderPointers, and nil error
}




//********************************* This section below contains my comments alongside the original code****
/* var ( //declaring variables in bulk
	err error //var named err, of type error
	f1  Folder // var named f1, of type Folder
	fs  []*Folder // a collection of pointers to Folder objects
	)
	f := []Folder{}                           //declare and initialise a slice named f, containing Folder structs (but is empty now)
	r, _ := FetchAllFoldersByOrgID(req.OrgID) //var r contains results returned by the FetchAllFoldersByOrgID method. _ is the var name for the second return value which we want to ignore because the 2nd return value is an error and we don't need to use it.
	for _, v := range r {                     //k isn't used so can be replaced with _. range loop. k is the key, v is the value. They are the iteration variables. r is the collection we iterate through. There's no explicit initialisation or condition statement because the loop iterates over the entire collection r until there are no more elements.
		f = append(f, *v) //f returns the full collection of items in collection r, ie folder. (whatever FetchAllFoldersByOrdID returns)
	}
	var fp []*Folder       // another collection of pointers to Folder objects
	for _, v1 := range f { //f is a collection of folders, k1 changed to _
		fp = append(fp, &v1) //return the full collection of memory addresses
	}
	var ffr *FetchFolderResponse            //a new variable named ffr of type pointer-for-FetchFolderResponse-object
	ffr = &FetchFolderResponse{Folders: fp} //The memory address of FetchFolderREsponse object is assigned to the variable ffr (which is a pointer???). FetchFolderResponse object has a field named Folders with the value of fp.
	return ffr, nil */

	//Since the original code was getting a range of addresses, then getting the values at the addresses, then getting the addresses again from those values, I've simplified the original code by cutting out the middle steps.

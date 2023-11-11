Pagination Choice
//There were a lot of resources on implementing pagination with a database, and using other 3rd party packages. 
However, I boiled it down to the simplest level for the purposes of getting an MVP that embodied the concept and functionality of pagination.  
Instead of using a pre-built package, I wrote a function that would use the token as an indicator of which index position to use as a starting point when populating the slice.
GetSomeFolders would still take in a single request struct and return a response struct (and error if any). 
Instead of taking in a struct AND a token, and returning a struct AND a token (and error if any), I added a modified FetchFolderRequest struct and a modified FetchFolderResponse struct, both of which have additional fields for Token.
That way, if in future, you needed to add anymore inputs to the "GetFolders" function, you would be able to simply add a new field in the FetchFolderRequest/ FetchFolderResponse structs, and keep the actual function looking cleaner.
There are more sophisticated ways of implementing pagination, and I didn't choose them because we aren't dealing with a database in this instance, I wanted to keep things simple, and I was working with a pretty tight time constraint personally.

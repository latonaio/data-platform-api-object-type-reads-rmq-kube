package requests

type Text struct {
	ObjectType     		string  `json:"ObjectType"`
	Language          	string  `json:"Language"`
	ObjectTypeName		string  `json:"ObjectTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

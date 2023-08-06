package models

type CreateBranch struct {
	Name   string
	Adress string
}

type Branch struct {
	Id     string
	Name   string
	Adress string
}

type IdRequest struct {
	Id string
}

type GetAllBranchRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllBranch struct {
	Branches []Branch
	Count    int
}

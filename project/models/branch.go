package models

type CreateBranch struct {
	Name      string
	Address   string
	FoundedAt string
}
type Branch struct {
	Id        string
	Name      string
	Address   string
	FoundedAt string
	CreatedAt string
	Year      int
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

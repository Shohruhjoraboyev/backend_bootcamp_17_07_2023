package memory

import "backend_bootcamp_17_07_2023/lesson_8/project/storage"

type store struct {
	branches *branchRepo
}

func NewStorage() storage.StorageI {
	return &store{branches: NewBranchRepo()}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}

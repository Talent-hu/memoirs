package repository

import (
	"memoirs/repository/auth"
	"memoirs/repository/bank"
)

type RepositoryGroup struct {
	auth.UserRepository
	auth.RoleRepository
	auth.MenuRepository
	auth.DictRepository
	auth.DictItemRepository

	bank.SubjectRepository
}

var RepositoryGroupApp = new(RepositoryGroup)

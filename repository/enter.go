package repository

import (
	"memoirs/repository/auth"
	"memoirs/repository/bank"
)

type RepositoryGroup struct {
	auth.UserRepository
	auth.RoleRepository
	auth.MenuRepository
	bank.SubjectRepository
}

var RepositoryGroupApp = new(RepositoryGroup)

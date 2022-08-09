package auth

import "memoirs/repository"

type any = interface{}

var (
	userMapper     = repository.RepositoryGroupApp.UserRepository
	roleMapper     = repository.RepositoryGroupApp.RoleRepository
	menuMapper     = repository.RepositoryGroupApp.MenuRepository
	dictMapper     = repository.RepositoryGroupApp.DictRepository
	dictItemMapper = repository.RepositoryGroupApp.DictItemRepository
	areaMapper     = repository.RepositoryGroupApp.AreaRepository
)

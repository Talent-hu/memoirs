package repository

type RepositoryGroup struct {
	UserRepository
	RoleRepository
	MenuRepository
}

var RepositoryGroupApp = new(RepositoryGroup)

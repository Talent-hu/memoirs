package bank

import "memoirs/repository"

type any = interface{}

var (
	subjectMapper = repository.RepositoryGroupApp.SubjectRepository
)

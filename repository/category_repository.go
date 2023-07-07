package repository

import (
	"github.com/dickidarmawansaputra/belajar-go-unit-test/entity"
)

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}

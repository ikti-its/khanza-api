package usecase

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/ikti-its/khanza-api/internal/modules/datadokter/internal/repository"
	"github.com/ikti-its/khanza-api/internal/modules/datadokter/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/datadokter/internal/entity"
)

type UseCase struct {
	Repository repository.Repository
}

func NewUseCase(repo repository.Repository) *UseCase {
	return &UseCase{Repository: repo}
}

func (u *UseCase) Create(request *model.Request) (model.Response, error) {
	var Entity entity.Dokter
	copier.Copy(&Entity, &request)

	err := u.Repository.Insert(&Entity)
	if err != nil {
		return model.Response{}, fmt.Errorf("failed to create: %v", err)
	}

	var response model.Response
	copier.Copy(&response, &Entity)
	return response, nil
}

func (u *UseCase) GetAll() ([]model.Response, error) {
	List, err := u.Repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve: %v", err)
	}

	var response []model.Response
	for _, Entity := range List {
		var r model.Response
		copier.Copy(&r, &Entity)
		response = append(response, r)
	}
	return response, nil
}

func (u *UseCase) GetById(id string) (model.Response, error) {
	Entity, err := u.Repository.FindById(id)
	if err != nil {
		return model.Response{}, fmt.Errorf("data not found")
	}

	var response model.Response
	copier.Copy(&response, &Entity)
	return response, nil
}

func (u *UseCase) Update(id string, request *model.Request) (model.Response, error) {
	Entity, err := u.Repository.FindById(id)
	if err != nil {
		return model.Response{}, fmt.Errorf("not found")
	}

	copier.Copy(&Entity, &request)
	err = u.Repository.Update(&Entity)
	if err != nil {
		return model.Response{}, fmt.Errorf("failed to update: %v", err)
	}

	var response model.Response
	copier.Copy(&response, &Entity)
	return response, nil
}

func (u *UseCase) Delete(id string) error {
	err := u.Repository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete: %v", err)
	}
	return nil
}

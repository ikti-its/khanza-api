package usecase

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/ikti-its/khanza-api/internal/modules/pasienmeninggal/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/pasienmeninggal/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/pasienmeninggal/internal/repository"
)

type UseCase struct {
	Repository repository.Repository
}

func NewUseCase(repo repository.Repository) *UseCase {
	return &UseCase{Repository: repo}
}

func (u *UseCase) Create(request *model.Request) (model.Response, error) {
	var data entity.PasienMeninggal
	_ = copier.Copy(&data, &request)

	err := u.Repository.Insert(&data)
	if err != nil {
		return model.Response{}, fmt.Errorf("failed to create: %v", err)
	}

	var response model.Response
	_ = copier.Copy(&response, &data)
	return response, nil
}

func (u *UseCase) GetAll() ([]model.Response, error) {
	list, err := u.Repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve: %v", err)
	}

	var response []model.Response
	for _, data := range list {
		var r model.Response
		_ = copier.Copy(&r, &data)
		response = append(response, r)
	}
	return response, nil
}

func (u *UseCase) GetById(noRM string) (model.Response, error) {
	data, err := u.Repository.FindById(noRM)
	if err != nil {
		return model.Response{}, fmt.Errorf("data not found")
	}

	var response model.Response
	_ = copier.Copy(&response, &data)
	return response, nil
}

func (u *UseCase) Update(noRM string, request *model.Request) (model.Response, error) {
	data, err := u.Repository.FindById(noRM)
	if err != nil {
		return model.Response{}, fmt.Errorf("not found")
	}

	_ = copier.Copy(&data, &request)
	data.NoRkmMedis = noRM // jaga-jaga kalau request hilangkan kolom ini

	err = u.Repository.Update(&data)
	if err != nil {
		return model.Response{}, fmt.Errorf("failed to update: %v", err)
	}

	var response model.Response
	_ = copier.Copy(&response, &data)
	return response, nil
}

func (u *UseCase) Delete(noRM string) error {
	err := u.Repository.Delete(noRM)
	if err != nil {
		return fmt.Errorf("failed to delete: %v", err)
	}
	return nil
}

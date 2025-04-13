package usecase

import (
	"fmt"

	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/repository"
)

type AmbulansUseCase struct {
	Repository repository.AmbulansRepository
}

func NewAmbulansUseCase(repo repository.AmbulansRepository) *AmbulansUseCase {
	return &AmbulansUseCase{Repository: repo}
}

// Create a new ambulans entry
func (u *AmbulansUseCase) Create(request *model.AmbulansRequest) (model.AmbulansResponse, error) {
	ambulansEntity := entity.Ambulans{
		NoAmbulans: request.NoAmbulans,
		Status:     request.Status,
		Supir:      request.Supir,
	}

	err := u.Repository.Insert(&ambulansEntity)
	if err != nil {
		return model.AmbulansResponse{}, fmt.Errorf("failed to create ambulans: %v", err)
	}

	return model.AmbulansResponse{
		NoAmbulans: ambulansEntity.NoAmbulans,
		Status:     ambulansEntity.Status,
		Supir:      ambulansEntity.Supir,
	}, nil
}

// Retrieve all ambulans records from PostgreSQL
func (u *AmbulansUseCase) GetAll() ([]model.AmbulansResponse, error) {
	ambulansList, err := u.Repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ambulans: %v", err)
	}

	var response []model.AmbulansResponse
	for _, a := range ambulansList {
		response = append(response, model.AmbulansResponse{
			NoAmbulans: a.NoAmbulans,
			Status:     a.Status,
			Supir:      a.Supir,
		})
	}

	return response, nil
}

// Retrieve a specific ambulans record by NoAmbulans
func (u *AmbulansUseCase) GetByNoAmbulans(noAmbulans string) (model.AmbulansResponse, error) {
	ambulans, err := u.Repository.FindByNoAmbulans(noAmbulans)
	if err != nil {
		return model.AmbulansResponse{}, fmt.Errorf("ambulans not found")
	}

	return model.AmbulansResponse{
		NoAmbulans: ambulans.NoAmbulans,
		Status:     ambulans.Status,
		Supir:      ambulans.Supir,
	}, nil
}

// Update an existing ambulans record
func (u *AmbulansUseCase) Update(noAmbulans string, request *model.AmbulansRequest) (model.AmbulansResponse, error) {
	ambulans, err := u.Repository.FindByNoAmbulans(noAmbulans)
	if err != nil {
		return model.AmbulansResponse{}, fmt.Errorf("ambulans not found")
	}

	ambulans.Status = request.Status
	ambulans.Supir = request.Supir

	err = u.Repository.Update(&ambulans)
	if err != nil {
		return model.AmbulansResponse{}, fmt.Errorf("failed to update ambulans: %v", err)
	}

	return model.AmbulansResponse{
		NoAmbulans: ambulans.NoAmbulans,
		Status:     ambulans.Status,
		Supir:      ambulans.Supir,
	}, nil
}

// Delete an ambulans record by NoAmbulans
func (u *AmbulansUseCase) Delete(noAmbulans string) error {
	err := u.Repository.Delete(noAmbulans)
	if err != nil {
		return fmt.Errorf("failed to delete ambulans: %v", err)
	}
	return nil
}

func (u *AmbulansUseCase) Notify(req *entity.Ambulans) error {
	err := u.Repository.SetPending(req.NoAmbulans)
	if err != nil {
		return fmt.Errorf("failed to update ambulance status: %v", err)
	}
	return nil
}

func (u *AmbulansUseCase) GetPendingRequests() ([]entity.Ambulans, error) {
	return u.Repository.FindPendingRequests()
}

func (u *AmbulansUseCase) MarkRequestAccepted(noAmbulans string) error {
	return u.Repository.UpdateAmbulansStatus(noAmbulans, "accepted")
}

package usecase

import (
	"fmt"

	"github.com/ikti-its/khanza-api/internal/modules/kamar/internal/entity"     // Change the path to match Kamar module
	"github.com/ikti-its/khanza-api/internal/modules/kamar/internal/model"      // Change the path to match Kamar module
	"github.com/ikti-its/khanza-api/internal/modules/kamar/internal/repository" // Change the path to match Kamar module
)

type KamarUseCase struct {
	Repository repository.KamarRepository
}

func NewKamarUseCase(repo repository.KamarRepository) *KamarUseCase {
	return &KamarUseCase{Repository: repo}
}

// Create a new kamar entry
func (u *KamarUseCase) Create(request *model.KamarRequest) (model.KamarResponse, error) {

	// Convert request model to entity model
	kamarEntity := entity.Kamar{
		NomorBed:    request.NomorBed,
		KodeKamar:   request.KodeKamar,
		NamaKamar:   request.NamaKamar,
		Kelas:       request.Kelas,
		TarifKamar:  request.TarifKamar,
		StatusKamar: request.StatusKamar,
	}

	// Insert into database
	err := u.Repository.Insert(&kamarEntity)
	if err != nil {
		return model.KamarResponse{}, fmt.Errorf("failed to create kamar: %v", err)
	}

	// Return response
	return model.KamarResponse{
		NomorBed:    kamarEntity.NomorBed,
		KodeKamar:   kamarEntity.KodeKamar,
		NamaKamar:   kamarEntity.NamaKamar,
		Kelas:       kamarEntity.Kelas,
		TarifKamar:  kamarEntity.TarifKamar,
		StatusKamar: kamarEntity.StatusKamar,
	}, nil
}

// Retrieve all kamar records from PostgreSQL
func (u *KamarUseCase) GetAll() ([]model.KamarResponse, error) {
	kamarList, err := u.Repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve kamar: %v", err)
	}

	var response []model.KamarResponse
	for _, kamar := range kamarList {
		response = append(response, model.KamarResponse{
			NomorBed:    kamar.NomorBed,
			KodeKamar:   kamar.KodeKamar,
			NamaKamar:   kamar.NamaKamar,
			Kelas:       kamar.Kelas,
			TarifKamar:  kamar.TarifKamar,
			StatusKamar: kamar.StatusKamar,
		})
	}

	return response, nil
}

// Retrieve a specific kamar record by NomorBed
func (u *KamarUseCase) GetByNomorBed(nomorBed string) (model.KamarResponse, error) {
	kamar, err := u.Repository.FindByNomorBed(nomorBed)
	if err != nil {
		return model.KamarResponse{}, fmt.Errorf("kamar not found")
	}

	return model.KamarResponse{
		NomorBed:    kamar.NomorBed,
		KodeKamar:   kamar.KodeKamar,
		NamaKamar:   kamar.NamaKamar,
		Kelas:       kamar.Kelas,
		TarifKamar:  kamar.TarifKamar,
		StatusKamar: kamar.StatusKamar,
	}, nil
}

// Update an existing kamar record
func (u *KamarUseCase) Update(nomorBed string, request *model.KamarRequest) (model.KamarResponse, error) {
	kamar, err := u.Repository.FindByNomorBed(nomorBed)
	if err != nil {
		return model.KamarResponse{}, fmt.Errorf("kamar not found")
	}

	kamar.KodeKamar = request.KodeKamar
	kamar.NamaKamar = request.NamaKamar
	kamar.Kelas = request.Kelas
	kamar.TarifKamar = request.TarifKamar
	kamar.StatusKamar = request.StatusKamar

	err = u.Repository.Update(&kamar)
	if err != nil {
		return model.KamarResponse{}, fmt.Errorf("failed to update kamar: %v", err)
	}

	return model.KamarResponse{
		NomorBed:    kamar.NomorBed,
		KodeKamar:   kamar.KodeKamar,
		NamaKamar:   kamar.NamaKamar,
		Kelas:       kamar.Kelas,
		TarifKamar:  kamar.TarifKamar,
		StatusKamar: kamar.StatusKamar,
	}, nil
}

// Delete a kamar record by NomorBed
func (u *KamarUseCase) Delete(nomorBed string) error {
	err := u.Repository.Delete(nomorBed)
	if err != nil {
		return fmt.Errorf("failed to delete kamar: %v", err)
	}
	return nil
}

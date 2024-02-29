package service

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/kehadiran/entity"
	"github.com/fathoor/simkes-api/internal/kehadiran/model"
	"github.com/fathoor/simkes-api/internal/kehadiran/repository"
	"github.com/fathoor/simkes-api/internal/kehadiran/validation"
	shiftRepository "github.com/fathoor/simkes-api/internal/shift/repository"
	"github.com/google/uuid"
	"time"
)

type kehadiranServiceImpl struct {
	repository.KehadiranRepository
	shiftRepository.ShiftRepository
}

func (service *kehadiranServiceImpl) CheckIn(request *model.KehadiranRequest) model.KehadiranResponse {
	if valid := validation.ValidateKehadiranRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	shift, err := service.ShiftRepository.FindByNama(request.ShiftNama)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	tanggal, err := time.Parse("2006-01-02", request.Tanggal)
	exception.PanicIfError(err)

	jamMasuk := time.Now()
	date := time.Now().Format("2006-01-02")

	shiftMasuk := shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05")
	masuk, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", date, shiftMasuk), time.FixedZone("WIB", 7*60*60))
	exception.PanicIfError(err)

	var keterangan string
	if jamMasuk.After(masuk) {
		keterangan = "Terlambat"
	} else {
		keterangan = "Hadir"
	}

	kehadiran := entity.Kehadiran{
		ID:         uuid.New(),
		NIP:        request.NIP,
		Tanggal:    tanggal,
		ShiftNama:  request.ShiftNama,
		JamMasuk:   jamMasuk,
		Keterangan: keterangan,
	}

	if err := service.KehadiranRepository.Insert(&kehadiran); err != nil {
		exception.PanicIfError(err)
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.ShiftNama,
			JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (service *kehadiranServiceImpl) CheckOut(request *model.KehadiranRequest) model.KehadiranResponse {
	if valid := validation.ValidateKehadiranRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	shift, err := service.ShiftRepository.FindByNama(request.ShiftNama)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	kehadiran, err := service.KehadiranRepository.FindLatestByNIP(request.NIP)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	jamKeluar := time.Now()

	kehadiran.JamKeluar = jamKeluar

	if err := service.KehadiranRepository.Update(&kehadiran); err != nil {
		exception.PanicIfError(err)
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.ShiftNama,
			JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (service *kehadiranServiceImpl) GetAll() []model.KehadiranResponse {
	kehadiran, err := service.KehadiranRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.KehadiranResponse, len(kehadiran))
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			ID:      kehadiran.ID.String(),
			NIP:     kehadiran.NIP,
			Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
			Shift: model.KehadiranShiftResponse{
				Nama:      kehadiran.ShiftNama,
				JamMasuk:  kehadiran.Shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
				JamKeluar: kehadiran.Shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			},
			JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (service *kehadiranServiceImpl) GetByNIP(nip string) []model.KehadiranResponse {
	kehadiran, err := service.KehadiranRepository.FindByNIP(nip)
	exception.PanicIfError(err)

	response := make([]model.KehadiranResponse, len(kehadiran))
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			ID:      kehadiran.ID.String(),
			NIP:     kehadiran.NIP,
			Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
			Shift: model.KehadiranShiftResponse{
				Nama:      kehadiran.ShiftNama,
				JamMasuk:  kehadiran.Shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
				JamKeluar: kehadiran.Shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			},
			JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (service *kehadiranServiceImpl) GetByID(id string) model.KehadiranResponse {
	kehadiranID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	kehadiran, err := service.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.ShiftNama,
			JamMasuk:  kehadiran.Shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: kehadiran.Shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (service *kehadiranServiceImpl) Update(id string, request *model.KehadiranUpdateRequest) model.KehadiranResponse {
	if valid := validation.ValidateKehadiranUpdateRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	kehadiranID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	kehadiran, err := service.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	tanggal, err := time.Parse("2006-01-02", request.Tanggal)
	exception.PanicIfError(err)

	jamMasuk, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamMasuk), time.FixedZone("WIB", 7*60*60))
	exception.PanicIfError(err)

	jamKeluar, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamKeluar), time.FixedZone("WIB", 7*60*60))
	exception.PanicIfError(err)

	kehadiran.Tanggal = tanggal
	kehadiran.ShiftNama = request.ShiftNama
	kehadiran.JamMasuk = jamMasuk
	kehadiran.JamKeluar = jamKeluar
	kehadiran.Keterangan = request.Keterangan

	if err := service.KehadiranRepository.Update(&kehadiran); err != nil {
		exception.PanicIfError(err)
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.ShiftNama,
			JamMasuk:  kehadiran.Shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: kehadiran.Shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response

}

func (service *kehadiranServiceImpl) Delete(id string) {
	kehadiranID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	kehadiran, err := service.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	if err := service.KehadiranRepository.Delete(&kehadiran); err != nil {
		exception.PanicIfError(err)
	}
}

func NewKehadiranServiceProvider(repository *repository.KehadiranRepository, shiftRepository *shiftRepository.ShiftRepository) KehadiranService {
	return &kehadiranServiceImpl{*repository, *shiftRepository}
}

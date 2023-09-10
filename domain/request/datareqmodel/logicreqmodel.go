package datareqmodel

import (
	"backend/domain/model"
	"backend/domain/request"
	"time"
)

func RequestToModel(data request.MahasiswaRequest, tahun time.Time) model.Mahasiswa {
	return model.Mahasiswa{
		NamaMahasiswa: data.NamaMahasiswa,
		Nim:           data.Nim,
		KotaAsal:      data.KotaAsal,
		TahunKuliah:   tahun,
	}
}
func DosenRequestToModel(data request.DosenRequest) model.Dosen {
	return model.Dosen{

		NamaDosen: data.NamaDosen,
		KotaAsal:  data.KotaAsal,
		Nip:       data.Nip,
	}
}
func DosenModelToReq(data model.Dosen) request.DosenRequest {
	return request.DosenRequest{
		NamaDosen: data.NamaDosen,
		KotaAsal:  data.KotaAsal,
		Nip:       data.Nip,
	}
}
func ListModelToReq(data []model.Dosen) (datareq []request.DosenRequest) {
	for _, val := range data {
		datareq = append(datareq, DosenModelToReq(val))
	}
	return datareq
}
func MatkulRequestToModel(data request.MatkulRequest) model.MataKuliah {
	return model.MataKuliah{
		KodeMatakuliah: data.KodeMatakuliah,
		NamaMatakuliah: data.NamaMatakuliah,
		DosenNip:       data.Nip,
		JumlahSks:      data.JumlahSks,
	}
}
func MatkulmodelToreq(data model.MataKuliah) request.MatkulRequest {
	return request.MatkulRequest{
		KodeMatakuliah: data.KodeMatakuliah,
		NamaMatakuliah: data.NamaMatakuliah,
		Nip:            data.DosenNip,
		JumlahSks:      data.JumlahSks,
	}
}
func KuliahRequestToModel(data request.KuliahRequest) model.Kuliah {
	return model.Kuliah{
		KodeKuliah: data.KodeKuliah,
		MatkulKode: data.MatkulKode,
		IndukDosen: data.IndukDosen,
		IndukMhs:   data.IndukMhs,
		NilaiUts:   data.NilaiUts,
		NilaiUas:   data.NilaiUas,
	}
}

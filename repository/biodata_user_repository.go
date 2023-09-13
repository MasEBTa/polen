package repository

import (
	"database/sql"
	"polen/model"
)

type BiodataUser interface {
	Save(payload model.BiodataUser) error
	FindByNIK(nik string) (model.BiodataUser, error)
	FindAll() ([]model.BiodataUser, error)
	Update(payload model.BiodataUser) error
	DeleteById(id string) error
	FindById(id string) (model.BiodataUser, error)
}

type biodataUserRepository struct {
	db *sql.DB
}

// FindById implements BiodataUser.
func (b *biodataUserRepository) FindById(id string) (model.BiodataUser, error) {
	row := b.db.QueryRow(`SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id WHERE b.id =$1`, id)
	biodata := model.BiodataUser{}
	err := row.Scan(
		&biodata.Id,
		&biodata.UserCredential.Id,
		&biodata.UserCredential.Username,
		&biodata.UserCredential.Password,
		&biodata.UserCredential.Role,
		&biodata.UserCredential.IsActive,
		&biodata.NamaLengkap,
		&biodata.Nik,
		&biodata.NomorTelepon,
		&biodata.Pekerjaan,
		&biodata.TempatLahir,
		&biodata.TanggalLahir,
		&biodata.KodePos,
	)
	if err != nil {
		return model.BiodataUser{}, err
	}
	return biodata, nil
}

// DeleteById implements BiodataUser.
func (b *biodataUserRepository) DeleteById(id string) error {
	_, err := b.db.Exec("DELETE FROM biodata_user WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements BiodataUser.
func (b *biodataUserRepository) FindAll() ([]model.BiodataUser, error) {
	rows, err := b.db.Query(`SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id`)
	if err != nil {
		return nil, err
	}
	var biodatas []model.BiodataUser
	for rows.Next() {
		biodata := model.BiodataUser{}
		err := rows.Scan(
			&biodata.Id,
			&biodata.UserCredential.Id,
			&biodata.UserCredential.Username,
			&biodata.UserCredential.Password,
			&biodata.UserCredential.Role,
			&biodata.UserCredential.IsActive,
			&biodata.NamaLengkap,
			&biodata.Nik,
			&biodata.NomorTelepon,
			&biodata.Pekerjaan,
			&biodata.TempatLahir,
			&biodata.TanggalLahir,
			&biodata.KodePos,
		)
		if err != nil {
			return nil, err
		}
		biodatas = append(biodatas, biodata)
	}
	return biodatas, nil
}

// FindByNIK implements BiodataUser.
func (b *biodataUserRepository) FindByNIK(nik string) (model.BiodataUser, error) {
	row := b.db.QueryRow(`SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id WHERE b.nik =$1`, nik)
	biodata := model.BiodataUser{}
	err := row.Scan(
		&biodata.Id,
		&biodata.UserCredential.Id,
		&biodata.UserCredential.Username,
		&biodata.UserCredential.Password,
		&biodata.UserCredential.Role,
		&biodata.UserCredential.IsActive,
		&biodata.NamaLengkap,
		&biodata.Nik,
		&biodata.NomorTelepon,
		&biodata.Pekerjaan,
		&biodata.TempatLahir,
		&biodata.TanggalLahir,
		&biodata.KodePos,
	)
	if err != nil {
		return model.BiodataUser{}, err
	}
	return biodata, nil
}

// Save implements BiodataUser.
func (b *biodataUserRepository) Save(payload model.BiodataUser) error {
	_, err := b.db.Exec("INSERT INTO biodata_user VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		payload.Id,
		payload.UserCredential.Id,
		payload.NamaLengkap,
		payload.Nik,
		payload.NomorTelepon,
		payload.Pekerjaan,
		payload.TempatLahir,
		payload.TanggalLahir,
		payload.KodePos,
	)
	if err != nil {
		return err
	}
	return nil
}

// Update implements BiodataUser.
func (b *biodataUserRepository) Update(payload model.BiodataUser) error {
	_, err := b.db.Exec(`UPDATE biodata_user SET user_credential_id = $2, nama_lengkap = $3, nik = $4, nomor_telepon = $5 
	, pekerjaan = $6, tempat_lahir = $7, tanggal_lahir = $8, kode_pos = $9 WHERE id = $1`,
		payload.Id,
		payload.UserCredential.Id,
		payload.NamaLengkap,
		payload.Nik,
		payload.NomorTelepon,
		payload.Pekerjaan,
		payload.TempatLahir,
		payload.TanggalLahir,
		payload.KodePos,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewBiodataUserRepository(db *sql.DB) BiodataUser {
	return &biodataUserRepository{db: db}
}

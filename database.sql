/*Tabel User_credential*/
CREATE TABLE
  public.user_credential
(
  id VARCHAR(225) PRIMARY KEY NOT NULL,
  username VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(225) NOT NULL,
  role VARCHAR(50) NOT NULL UNIQUE,
  is_active BOOLEAN
);

/*Tabel Biodata User*/
CREATE TABLE public.biodata_user
(
  id VARCHAR(225) PRIMARY KEY NOT NULL,
  user_credential_id VARCHAR(225) NOT NULL REFERENCES public.user_credential(id),
  nama_lengkap VARCHAR(225) NOT NULL,
  nik VARCHAR(20) NOT NULL UNIQUE,
  nomor_telepon VARCHAR(20) NOT NULL UNIQUE,
  pekerjaan VARCHAR(225) NOT NULL,
  tempat_lahir VARCHAR(225) NOT NULL,
  tanggal_lahir DATE NOT NULL,
  kode_pos VARCHAR(10) NOT NULL
);

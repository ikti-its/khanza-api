-- Set Timezone
SET TIMEZONE='Asia/Jakarta';

-- Create Role Table
CREATE TABLE role (
    nama VARCHAR(20) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Jabatan Table
CREATE TABLE jabatan (
    nama VARCHAR(25) PRIMARY KEY,
    jenjang VARCHAR(25) NOT NULL,
    gaji_pokok NUMERIC NOT NULL DEFAULT 0,
    tunjangan NUMERIC NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Departemen Table
CREATE TABLE departemen (
    nama VARCHAR(25) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Pegawai Table
CREATE TYPE jenis_kelamin AS ENUM ('L', 'P');
CREATE TYPE status_kerja AS ENUM ('Tetap', 'Kontrak');
CREATE TYPE pendidikan AS ENUM ('SD', 'SMP', 'SMA', 'D3', 'S1', 'S2', 'S3');

CREATE TABLE pegawai (
    nip VARCHAR(5) PRIMARY KEY,
    nik VARCHAR(16) UNIQUE NOT NULL,
    nama VARCHAR(50) NOT NULL,
    jenis_kelamin jenis_kelamin NOT NULL,
    jabatan_nama VARCHAR(25) NOT NULL,
    departemen_nama VARCHAR(25) NOT NULL,
    status_kerja status_kerja NOT NULL,
    pendidikan pendidikan NOT NULL,
    tempat_lahir VARCHAR(20) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    alamat_lat NUMERIC NOT NULL DEFAULT 7.2575,
    alamat_lon NUMERIC NOT NULL DEFAULT 112.7521,
    telepon VARCHAR(15) NOT NULL,
    tanggal_masuk DATE NOT NULL,
    foto VARCHAR(255) NOT NULL DEFAULT '/assets/image/default.png',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (jabatan_nama) REFERENCES jabatan(nama),
    FOREIGN KEY (departemen_nama) REFERENCES departemen(nama)
);

-- Create Akun Table
CREATE TABLE akun (
    nip VARCHAR(5) PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_nama VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (nip) REFERENCES pegawai(nip),
    FOREIGN KEY (role_nama) REFERENCES role(nama)
);

-- Create Shift Table
CREATE TABLE shift (
    nama VARCHAR(10) PRIMARY KEY,
    jam_masuk TIMESTAMP WITH TIME ZONE NOT NULL,
    jam_keluar TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Jadwal Pegawai Table
CREATE TABLE jadwal_pegawai (
    nip VARCHAR(5) NOT NULL,
    tahun SMALLINT NOT NULL,
    bulan SMALLINT NOT NULL,
    hari SMALLINT NOT NULL,
    shift_nama VARCHAR(10) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (nip, tahun, bulan, hari),
    FOREIGN KEY (nip) REFERENCES pegawai(nip),
    FOREIGN KEY (shift_nama) REFERENCES shift(nama)
);

-- Create Kehadiran Table
CREATE TABLE kehadiran (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nip VARCHAR(5) NOT NULL,
    tanggal DATE NOT NULL DEFAULT CURRENT_DATE,
    shift_nama VARCHAR(10) NOT NULL,
    jam_masuk TIMESTAMP WITH TIME ZONE,
    jam_keluar TIMESTAMP WITH TIME ZONE,
    keterangan VARCHAR(25) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (nip) REFERENCES pegawai(nip),
    FOREIGN KEY (shift_nama) REFERENCES shift(nama)
);

-- Create Cuti Table
CREATE TABLE cuti (
    nip VARCHAR(5) NOT NULL,
    tanggal_mulai DATE NOT NULL,
    tanggal_selesai DATE NOT NULL,
    keterangan VARCHAR(255) NOT NULL,
    status BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (nip, tanggal_mulai, tanggal_selesai),
    FOREIGN KEY (nip) REFERENCES pegawai(nip)
);

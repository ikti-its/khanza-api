-- Set Timezone
SET TIMEZONE='Asia/Jakarta';

-- Create Role Table
CREATE TABLE role (
    id SERIAL PRIMARY KEY,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Akun Table
CREATE TABLE akun (
    nip VARCHAR(5) PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES role(id)
);

-- Create Jabatan Table
CREATE TABLE jabatan (
    id VARCHAR(10) PRIMARY KEY,
    jabatan VARCHAR(25) NOT NULL,
    jenjang VARCHAR(25) NOT NULL,
    gaji_pokok NUMERIC NOT NULL DEFAULT 0,
    tunjangan NUMERIC NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Departemen Table
CREATE TABLE departemen (
    id VARCHAR(10) PRIMARY KEY,
    departemen VARCHAR(25) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Pegawai Table
CREATE TABLE pegawai (
    nip VARCHAR(5) PRIMARY KEY,
    nik VARCHAR(20) UNIQUE NOT NULL,
    nama VARCHAR(50) NOT NULL,
    jenis_kelamin VARCHAR(1) NOT NULL,
    jabatan_id VARCHAR(10) NOT NULL,
    departemen_id VARCHAR(10) NOT NULL,
    status_kerja VARCHAR(20) NOT NULL,
    pendidikan VARCHAR(25) NOT NULL,
    tempat_lahir VARCHAR(20) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    alamat_lat FLOAT NOT NULL DEFAULT 7.2575,
    alamat_lon FLOAT NOT NULL DEFAULT 112.7521,
    no_telepon VARCHAR(15) NOT NULL,
    tanggal_masuk DATE NOT NULL,
    foto VARCHAR(255) NOT NULL DEFAULT '/resource/default.png',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (jabatan_id) REFERENCES jabatan(id),
    FOREIGN KEY (departemen_id) REFERENCES departemen(id)
);

-- Create Shift Table
CREATE TABLE shift (
    id VARCHAR(10) PRIMARY KEY,
    jam_masuk TIME WITH TIME ZONE NOT NULL,
    jam_keluar TIME WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Jadwal Pegawai Table
CREATE TABLE jadwal_pegawai (
    nip VARCHAR(5) NOT NULL,
    tahun SMALLINT NOT NULL,
    bulan SMALLINT NOT NULL,
    hari SMALLINT NOT NULL,
    shift_id VARCHAR(10) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (nip, tahun, bulan),
    FOREIGN KEY (nip) REFERENCES pegawai(nip),
    FOREIGN KEY (shift_id) REFERENCES shift(id)
);

-- Create Kehadiran Table
CREATE TABLE kehadiran (
    nip VARCHAR(5) NOT NULL,
    tanggal DATE NOT NULL DEFAULT CURRENT_DATE,
    tahun SMALLINT NOT NULL,
    bulan SMALLINT NOT NULL,
    shift_id VARCHAR(10) NOT NULL,
    jam_masuk TIME WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIME,
    jam_keluar TIME WITH TIME ZONE NOT NULL,
    keterangan VARCHAR(25) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (nip, tanggal),
    FOREIGN KEY (nip) REFERENCES pegawai(nip),
    FOREIGN KEY (shift_id) REFERENCES shift(id)
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

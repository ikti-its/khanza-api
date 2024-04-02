-- Role
CREATE TABLE IF NOT EXISTS role (
    id INT PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO role (id, nama) VALUES (1337, 'Developer'), (1, 'Admin'), (2, 'Pegawai');

-- MODUL C (Kepegawaian)
-- Jabatan
CREATE TABLE IF NOT EXISTS jabatan (
    id INT PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO jabatan (id, nama) VALUES (1000, 'Testing');
-- INSERT INTO jabatan (id, nama) VALUES (1, 'Direktur'), (2, 'Manager'), (3, 'Supervisor'), (4, 'Staff');

-- Departemen
CREATE TABLE IF NOT EXISTS departemen (
    id INT PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO departemen (id, nama) VALUES (1000, 'Testing');
-- INSERT INTO departemen (id, nama) VALUES (1, 'HRD'), (2, 'Marketing'), (3, 'Keuangan'), (4, 'Operasional');

-- Status Aktif Pegawai
CREATE TABLE IF NOT EXISTS status_aktif_pegawai (
    id VARCHAR(2) PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO status_aktif_pegawai (id, nama) VALUES ('A', 'Aktif'), ('BH', 'Berhenti dengan Hormat'), ('C', 'Cuti'), ('R', 'Resign'), ('BT', 'Berhenti dengan Tidak Hormat'), ('P', 'Pensiun'), ('W', 'Wafat');

-- Shift
CREATE TABLE IF NOT EXISTS shift (
    id VARCHAR(2) PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    jam_masuk TIME WITH TIME ZONE NOT NULL,
    jam_pulang TIME WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO shift (id, nama, jam_masuk, jam_pulang) VALUES ('NA', 'Belum Ditentukan', '07:00:00 +07:00', '07:00:00 +07:00'), ('P', 'Pagi', '07:00:00 +07:00', '15:00:00 +07:00'), ('S', 'Sore', '15:00:00 +07:00', '23:00:00 +07:00'), ('M', 'Malam', '23:00:00 +07:00', '07:00:00 +07:00');

-- Hari
CREATE TABLE IF NOT EXISTS hari (
    id INT PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO hari (id, nama) VALUES (1, 'Senin'), (2, 'Selasa'), (3, 'Rabu'), (4, 'Kamis'), (5, 'Jumat'), (6, 'Sabtu'), (7, 'Minggu');

-- Alasan Cuti
CREATE TABLE IF NOT EXISTS alasan_cuti (
    id VARCHAR(2) PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO alasan_cuti (id, nama) VALUES ('S', 'Sakit'), ('I', 'Izin'), ('CT', 'Cuti Tahunan'), ('CB', 'Cuti Besar'), ('CM', 'Cuti Melahirkan'), ('CU', 'Cuti Karena Alasan Penting');

-- Akun
CREATE TABLE IF NOT EXISTS akun (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    foto VARCHAR(255) NOT NULL,
    role INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updater UUID,
    FOREIGN KEY (role) REFERENCES ref.role (id)
);

INSERT INTO akun (email, password, foto, role) VALUES
('admin@fathoor.dev', '$2a$10$8jI.qKrVbXQjNzYX6KOIvukkYkNcmfYyPWiv9tuHE8vdg5EhjQBzy', 'http://localhost:8080/v1/file/img/default.png', 1),
('fathoor@fathoor.dev', '$2a$10$Yg78XjsfHtvpiZHzxjhtBeauNRpK928c1zKSXjPdV.jOE0q21qXgq', 'http://localhost:8080/v1/file/img/default.png', 1337);

CREATE INDEX IF NOT EXISTS idx_akun_foto ON akun (foto);
CREATE INDEX IF NOT EXISTS idx_akun_role ON akun (role);
CREATE INDEX IF NOT EXISTS idx_akun_updated_at ON akun (updated_at);
CREATE INDEX IF NOT EXISTS idx_akun_deleted_at ON akun (deleted_at);

-- Modul C (Kepegawaian)
-- Pegawai
CREATE TYPE jenis_kelamin AS ENUM ('L', 'P');
CREATE TYPE jenis_pegawai AS ENUM ('Tetap', 'Kontrak', 'Magang', 'Istimewa');

CREATE TABLE IF NOT EXISTS pegawai (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_akun UUID NOT NULL UNIQUE,
    nip VARCHAR(10) NOT NULL UNIQUE,
    nama VARCHAR(255) NOT NULL,
    jenis_kelamin jenis_kelamin NOT NULL,
    id_jabatan INT NOT NULL,
    id_departemen INT NOT NULL,
    id_status_aktif VARCHAR(2) NOT NULL,
    jenis_pegawai jenis_pegawai NOT NULL,
    telepon VARCHAR(255) NOT NULL,
    tanggal_masuk DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updater UUID,
    FOREIGN KEY (id_akun) REFERENCES akun (id),
    FOREIGN KEY (id_jabatan) REFERENCES ref.jabatan (id),
    FOREIGN KEY (id_departemen) REFERENCES ref.departemen (id),
    FOREIGN KEY (id_status_aktif) REFERENCES ref.status_aktif_pegawai (id),
    FOREIGN KEY (updater) REFERENCES akun (id)
);

CREATE INDEX IF NOT EXISTS idx_pegawai_jenis_kelamin ON pegawai (jenis_kelamin);
CREATE INDEX IF NOT EXISTS idx_pegawai_id_jabatan ON pegawai (id_jabatan);
CREATE INDEX IF NOT EXISTS idx_pegawai_id_departemen ON pegawai (id_departemen);
CREATE INDEX IF NOT EXISTS idx_pegawai_id_status_aktif ON pegawai (id_status_aktif);
CREATE INDEX IF NOT EXISTS idx_pegawai_jenis_pegawai ON pegawai (jenis_pegawai);
CREATE INDEX IF NOT EXISTS idx_pegawai_tanggal_masuk ON pegawai (tanggal_masuk);
CREATE INDEX IF NOT EXISTS idx_pegawai_updated_at ON pegawai (updated_at);
CREATE INDEX IF NOT EXISTS idx_pegawai_deleted_at ON pegawai (deleted_at);

CREATE TRIGGER init_jadwal_pegawai
AFTER INSERT ON pegawai
FOR EACH ROW
EXECUTE PROCEDURE trigger_init_jadwal_pegawai();

-- Alamat
CREATE TABLE IF NOT EXISTS alamat (
    id_akun UUID PRIMARY KEY,
    alamat VARCHAR(255) NOT NULL,
    alamat_lat NUMERIC NOT NULL DEFAULT 7.2575,
    alamat_lon NUMERIC NOT NULL DEFAULT 112.7521,
    kota VARCHAR(255) NOT NULL,
    kode_pos VARCHAR(5) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updater UUID,
    FOREIGN KEY (id_akun) REFERENCES akun (id),
    FOREIGN KEY (updater) REFERENCES akun (id)
);

CREATE INDEX IF NOT EXISTS idx_alamat_alamat ON alamat (alamat);
CREATE INDEX IF NOT EXISTS idx_alamat_alamat_lat ON alamat (alamat_lat);
CREATE INDEX IF NOT EXISTS idx_alamat_alamat_lon ON alamat (alamat_lon);
CREATE INDEX IF NOT EXISTS idx_alamat_kode_pos ON alamat (kode_pos);
CREATE INDEX IF NOT EXISTS idx_alamat_updated_at ON alamat (updated_at);
CREATE INDEX IF NOT EXISTS idx_alamat_deleted_at ON alamat (deleted_at);

-- Berkas Pegawai (Dokumen)
CREATE TYPE agama AS ENUM ('Islam', 'Kristen', 'Katolik', 'Hindu', 'Buddha', 'Konghucu', 'Lainnya');
CREATE TYPE pendidikan AS ENUM ('Tidak Sekolah', 'SD', 'SMP', 'SMA', 'D3', 'D4', 'S1', 'S2', 'S3');

CREATE TABLE IF NOT EXISTS berkas_pegawai (
    id_pegawai UUID PRIMARY KEY,
    nik VARCHAR(16) NOT NULL UNIQUE,
    tempat_lahir VARCHAR(255) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    agama agama NOT NULL,
    pendidikan pendidikan NOT NULL,
    ktp VARCHAR(255) NOT NULL,
    kk VARCHAR(255) NOT NULL,
    npwp VARCHAR(255) NOT NULL,
    bpjs VARCHAR(255),
    ijazah VARCHAR(255),
    skck VARCHAR(255),
    str VARCHAR(255),
    serkom VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updater UUID,
    FOREIGN KEY (id_pegawai) REFERENCES pegawai (id),
    FOREIGN KEY (updater) REFERENCES akun (id)
);

CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_agama ON berkas_pegawai (agama);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_pendidikan ON berkas_pegawai (pendidikan);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_ktp ON berkas_pegawai (ktp);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_kk ON berkas_pegawai (kk);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_npwp ON berkas_pegawai (npwp);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_bpjs ON berkas_pegawai (bpjs);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_ijazah ON berkas_pegawai (ijazah);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_skck ON berkas_pegawai (skck);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_str ON berkas_pegawai (str);
CREATE INDEX IF NOT EXISTS idx_berkas_pegawai_serkom ON berkas_pegawai (serkom);

-- Foto (Untuk keperluan presensi)
CREATE TABLE IF NOT EXISTS foto_pegawai (
    id_pegawai UUID PRIMARY KEY,
    foto VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updater UUID,
    FOREIGN KEY (id_pegawai) REFERENCES pegawai (id),
    FOREIGN KEY (updater) REFERENCES akun (id)
);

CREATE INDEX IF NOT EXISTS idx_foto_pegawai_foto ON foto_pegawai (foto);
CREATE INDEX IF NOT EXISTS idx_foto_pegawai_updated_at ON foto_pegawai (updated_at);
CREATE INDEX IF NOT EXISTS idx_foto_pegawai_deleted_at ON foto_pegawai (deleted_at);

-- Jadwal Pegawai
CREATE TABLE IF NOT EXISTS jadwal_pegawai (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_pegawai UUID NOT NULL,
    id_hari INT NOT NULL,
    id_shift VARCHAR(2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updater UUID,
    FOREIGN KEY (id_pegawai) REFERENCES pegawai (id),
    FOREIGN KEY (id_hari) REFERENCES ref.hari (id),
    FOREIGN KEY (id_shift) REFERENCES ref.shift (id),
    FOREIGN KEY (updater) REFERENCES akun (id)
);

CREATE INDEX IF NOT EXISTS idx_jadwal_pegawai_id_pegawai ON jadwal_pegawai (id_pegawai);
CREATE INDEX IF NOT EXISTS idx_jadwal_pegawai_id_shift ON jadwal_pegawai (id_shift);
CREATE INDEX IF NOT EXISTS idx_jadwal_pegawai_id_hari ON jadwal_pegawai (id_hari);
CREATE INDEX IF NOT EXISTS idx_jadwal_pegawai_updated_at ON jadwal_pegawai (updated_at);
CREATE INDEX IF NOT EXISTS idx_jadwal_pegawai_deleted_at ON jadwal_pegawai (deleted_at);

CREATE OR REPLACE FUNCTION default_jadwal_pegawai(IN new_id UUID)
RETURNS VOID AS $$
BEGIN
    INSERT INTO jadwal_pegawai (id_pegawai, id_hari, id_shift)
    SELECT new_id, hari.id, 'NA'
    FROM ref.hari;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trigger_init_jadwal_pegawai()
RETURNS TRIGGER AS $$
BEGIN
    PERFORM default_jadwal_pegawai(NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Presensi
CREATE TABLE IF NOT EXISTS presensi (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_pegawai UUID NOT NULL,
    id_jadwal_pegawai UUID NOT NULL,
    tanggal DATE DEFAULT CURRENT_DATE,
    jam_masuk TIME WITH TIME ZONE,
    jam_pulang TIME WITH TIME ZONE,
    keterangan VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updater UUID,
    FOREIGN KEY (id_pegawai) REFERENCES pegawai (id),
    FOREIGN KEY (id_jadwal_pegawai) REFERENCES jadwal_pegawai (id),
    FOREIGN KEY (updater) REFERENCES akun (id)
);

CREATE INDEX IF NOT EXISTS idx_presensi_id_pegawai ON presensi (id_pegawai);
CREATE INDEX IF NOT EXISTS idx_presensi_id_jadwal_pegawai ON presensi (id_jadwal_pegawai);
CREATE INDEX IF NOT EXISTS idx_presensi_tanggal ON presensi (tanggal);
CREATE INDEX IF NOT EXISTS idx_presensi_jam_masuk ON presensi (jam_masuk);
CREATE INDEX IF NOT EXISTS idx_presensi_jam_pulang ON presensi (jam_pulang);
CREATE INDEX IF NOT EXISTS idx_presensi_keterangan ON presensi (keterangan);
CREATE INDEX IF NOT EXISTS idx_presensi_updated_at ON presensi (updated_at);
CREATE INDEX IF NOT EXISTS idx_presensi_deleted_at ON presensi (deleted_at);

CREATE OR REPLACE FUNCTION check_keterangan_presensi()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.jam_masuk IS NOT NULL AND NEW.jam_pulang IS NOT NULL THEN
        SELECT s.nama INTO NEW.keterangan
        FROM ref.shift s
        JOIN jadwal_pegawai jp ON jp.id_shift = s.id
        WHERE jp.id = NEW.id_jadwal_pegawai;

        IF NEW.jam_masuk > ref.shift.jam_masuk THEN
            NEW.keterangan := NEW.keterangan || ' (Terlambat)';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_keterangan_presensi
BEFORE UPDATE ON presensi
FOR EACH ROW
EXECUTE PROCEDURE check_keterangan_presensi();

-- Cuti
CREATE TABLE IF NOT EXISTS cuti (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_pegawai UUID NOT NULL,
    tanggal_mulai DATE NOT NULL,
    tanggal_selesai DATE NOT NULL,
    id_alasan_cuti VARCHAR(2) NOT NULL,
    status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updater UUID,
    FOREIGN KEY (id_pegawai) REFERENCES pegawai (id),
    FOREIGN KEY (id_alasan_cuti) REFERENCES ref.alasan_cuti (id),
    FOREIGN KEY (updater) REFERENCES akun (id)
);

CREATE INDEX IF NOT EXISTS idx_cuti_id_pegawai ON cuti (id_pegawai);
CREATE INDEX IF NOT EXISTS idx_cuti_tanggal_mulai ON cuti (tanggal_mulai);
CREATE INDEX IF NOT EXISTS idx_cuti_tanggal_selesai ON cuti (tanggal_selesai);
CREATE INDEX IF NOT EXISTS idx_cuti_id_alasan_cuti ON cuti (id_alasan_cuti);
CREATE INDEX IF NOT EXISTS idx_cuti_status ON cuti (status);
CREATE INDEX IF NOT EXISTS idx_cuti_updated_at ON cuti (updated_at);
CREATE INDEX IF NOT EXISTS idx_cuti_deleted_at ON cuti (deleted_at);

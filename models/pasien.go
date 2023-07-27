package models

import "time"

type Pasien struct {
	ID                uint               `gorm:"primaryKey"`
	NamaLengkap       string             `gorm:"column:nama_lengkap"`
	NamaPanggilan     string             `gorm:"column:nama_panggilan"`
	NoTelp            string             `gorm:"column:no_telp"`
	JenisKelamin      string             `gorm:"column:jenis_kelamin"`
	Pekerjaan         string             `gorm:"column:pekerjaan"`
	TanggalLahir      time.Time          `gorm:"column:tanggal_lahir"`
	NoBPJS            string             `gorm:"column:no_bpjs"`
	NIK               string             `gorm:"column:nik"`
	TotalKunjungan    int                `gorm:"column:total_kunjungan"`
	KunjunganTerakhir time.Time          `gorm:"column:kunjungan_terakhir"`
	RiwayatAlergi     string             `gorm:"column:riwayat_alergi"`
	IsQueue           bool               `gorm:"column:is_queue"`
	DibuatPada        time.Time          `gorm:"column:dibuat_pada"`
	DibuatOleh        string             `gorm:"column:dibuat_oleh"`
	DiubahPada        time.Time          `gorm:"column:diubah_pada"`
	DiubahOleh        string             `gorm:"column:diubah_oleh"`
	RiwayatKunjungan  []RiwayatKunjungan `gorm:"foreignKey:PasienID"`
	HistoryPerubahan  []HistoryPerubahan `gorm:"foreignKey:PasienID"`
}
type RiwayatKunjungan struct {
	ID                 uint `gorm:"primaryKey"`
	PasienID           uint `gorm:"column:pasien_id"`
	PemeriksaanPerawat PemeriksaanPerawat
	PemeriksaanDokter  PemeriksaanDokter
}

type PemeriksaanPerawat struct {
	ID               uint      `gorm:"primaryKey"`
	TanggalKunjungan time.Time `gorm:"column:tanggal_kunjungan"`
	TinggiBadan      int       `gorm:"column:tinggi_badan"`
	BeratBadan       int       `gorm:"column:berat_badan"`
	LingkarPerut     int       `gorm:"column:lingkar_perut"`
	IMT              float64   `gorm:"column:IMT"`
	Sistole          int       `gorm:"column:sistole"`
	Diastole         int       `gorm:"column:diastole"`
	RespiratoryRate  int       `gorm:"column:respiratory_rate"`
	HeartRate        int       `gorm:"column:heart_rate"`
}

type PemeriksaanDokter struct {
	ID               uint      `gorm:"primaryKey"`
	TanggalKunjungan time.Time `gorm:"column:tanggal_kunjungan"`
	Diagnosa         string    `gorm:"column:diagnosa"`
	Saran            string    `gorm:"column:saran"`
	Obat             string    `gorm:"column:obat"`
}

type HistoryPerubahan struct {
	ID          uint      `gorm:"primaryKey"`
	PasienID    uint      `gorm:"column:pasien_id"`
	DirubahOleh string    `gorm:"column:dirubah_oleh"`
	DirubahPada time.Time `gorm:"column:dirubah_pada"`
	Perubahan   string    `gorm:"column:perubahan"`
}
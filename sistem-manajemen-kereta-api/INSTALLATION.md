# 📋 Panduan Instalasi Lengkap

## 🔧 Persiapan Environment

### 1. Install Go
- Download Go dari [golang.org](https://golang.org/dl/)
- Pilih versi terbaru untuk OS Anda
- Ikuti panduan instalasi sesuai OS

### 2. Verifikasi Instalasi Go
\`\`\`bash
go version
\`\`\`
Pastikan output menampilkan versi Go 1.21 atau lebih baru.

### 3. Setup GOPATH (Opsional)
\`\`\`bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
\`\`\`

## 🚀 Instalasi Project

### 1. Clone Repository
\`\`\`bash
git clone https://github.com/username/kereta-api.git
cd kereta-api
\`\`\`

### 2. Initialize Go Module
\`\`\`bash
go mod init kereta-api
go mod tidy
\`\`\`

### 3. Struktur Folder
Pastikan struktur folder seperti ini:
\`\`\`
kereta-api/
├── main.go
├── model/
│   └── models.go
├── controller/
│   ├── admin_controller.go
│   └── user_controller.go
├── templates/
│   ├── index.html
│   ├── admin.html
│   └── user.html
├── go.mod
├── .gitignore
└── README.md
\`\`\`

### 4. Jalankan Aplikasi
\`\`\`bash
go run main.go
\`\`\`

### 5. Akses Aplikasi
- Buka browser
- Kunjungi: `http://localhost:8080`

## 🔍 Troubleshooting

### Error: "cannot find module"
\`\`\`bash
go mod init kereta-api
go mod tidy
\`\`\`

### Error: "port already in use"
- Ubah port di `main.go`:
\`\`\`go
log.Fatal(http.ListenAndServe(":8081", nil))
\`\`\`

### Error: "template not found"
- Pastikan folder `templates/` ada di root directory
- Periksa nama file template sesuai dengan yang dipanggil di controller

## 🎯 Testing

### Test Manual
1. Akses halaman utama: `http://localhost:8080`
2. Test Admin Panel: `http://localhost:8080/admin`
3. Test User Panel: `http://localhost:8080/user`
4. Test CRUD operations di admin panel
5. Test booking di user panel

### Test Fungsionalitas
- [ ] Tambah kota baru
- [ ] Tambah stasiun baru
- [ ] Tambah rute baru
- [ ] Tambah jadwal baru
- [ ] Lihat daftar stasiun (user)
- [ ] Lihat jadwal (user)
- [ ] Pesan tiket (user)

## 📱 Deployment

### Local Development
\`\`\`bash
go run main.go
\`\`\`

### Build untuk Production
\`\`\`bash
go build -o kereta-api main.go
./kereta-api
\`\`\`

### Docker (Opsional)
\`\`\`dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
\`\`\`

## 🔒 Security Notes

- Aplikasi ini untuk pembelajaran, tidak untuk production
- Tidak ada autentikasi/autorisasi
- Data disimpan di memory (hilang saat restart)
- Untuk production, tambahkan database dan security layer

## 📞 Support

Jika mengalami masalah:
1. Periksa log error di terminal
2. Pastikan semua file ada di tempat yang benar
3. Verifikasi Go version
4. Buat issue di GitHub repository

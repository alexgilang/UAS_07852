# 🚄 Sistem Manajemen Kereta API

Sistem manajemen kereta api modern yang dibangun dengan **Go** dan **HTML/CSS**. Aplikasi ini menggunakan **Linked List** sebagai struktur data utama untuk mengelola data kota, stasiun, rute, jadwal, dan pemesanan tiket.

## ✨ Fitur Utama

### 👨‍💼 Admin Panel
- **Kelola Kota**: Tambah, edit, dan hapus data kota
- **Kelola Stasiun**: Manajemen data stasiun dengan informasi lengkap
- **Kelola Rute**: Atur rute perjalanan antar kota
- **Kelola Jadwal**: Tentukan jadwal keberangkatan dan kedatangan
- **Kelola Harga**: Atur harga tiket untuk setiap rute

### 👤 User Panel
- **Lihat Stasiun**: Informasi lengkap semua stasiun kereta api
- **Rute & Jadwal**: Cek jadwal dan rute perjalanan terkini
- **Pesan Tiket**: Sistem pemesanan tiket online yang mudah

## 🛠️ Teknologi yang Digunakan

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3, JavaScript
- **Struktur Data**: Linked List
- **Arsitektur**: MVC (Model-View-Controller)
- **Styling**: CSS dengan Glass Morphism & Gradient Design

## 📁 Struktur Project

\`\`\`
kereta-api/
├── main.go                 # Entry point aplikasi
├── model/
│   └── models.go          # Model data dan linked list operations
├── controller/
│   ├── admin_controller.go # Controller untuk admin
│   └── user_controller.go  # Controller untuk user
├── templates/
│   ├── index.html         # Halaman utama
│   ├── admin.html         # Dashboard admin
│   └── user.html          # Dashboard user
├── go.mod                 # Go module file
├── .gitignore            # Git ignore file
└── README.md             # Dokumentasi project
\`\`\`

## 🚀 Cara Menjalankan

### Prerequisites
- Go 1.21 atau lebih baru
- Web browser modern

### Langkah Instalasi

1. **Clone repository**
   \`\`\`bash
   git clone https://github.com/username/kereta-api.git
   cd kereta-api
   \`\`\`

2. **Initialize Go module**
   \`\`\`bash
   go mod init kereta-api
   go mod tidy
   \`\`\`

3. **Jalankan aplikasi**
   \`\`\`bash
   go run main.go
   \`\`\`

4. **Akses aplikasi**
   - Buka browser dan kunjungi: `http://localhost:8080`
   - Admin Panel: `http://localhost:8080/admin`
   - User Panel: `http://localhost:8080/user`

## 🎯 Endpoint API

### Main Routes
- `GET /` - Halaman utama
- `GET /admin` - Dashboard admin
- `GET /user` - Dashboard user

### Admin Routes
- `GET/POST /admin/cities` - Kelola kota
- `GET/POST /admin/stations` - Kelola stasiun
- `GET/POST /admin/routes` - Kelola rute
- `GET/POST /admin/schedules` - Kelola jadwal
- `GET /admin/prices` - Kelola harga

### User Routes
- `GET /user/stations` - Lihat daftar stasiun
- `GET /user/routes` - Lihat rute dan jadwal
- `GET/POST /user/booking` - Pesan tiket

## 💾 Struktur Data

Aplikasi ini menggunakan **Linked List** untuk menyimpan data:

### City (Kota)
\`\`\`go
type City struct {
    ID   int
    Name string
    Next *City
}
\`\`\`

### Station (Stasiun)
\`\`\`go
type Station struct {
    ID      int
    Name    string
    Address string
    Phone   string
    CityID  int
    Next    *Station
}
\`\`\`

### Route (Rute)
\`\`\`go
type Route struct {
    ID         int
    FromCityID int
    ToCityID   int
    Distance   float64
    Next       *Route
}
\`\`\`

### Schedule (Jadwal)
\`\`\`go
type Schedule struct {
    ID            int
    RouteID       int
    TrainName     string
    DepartureTime time.Time
    ArrivalTime   time.Time
    Price         float64
    Next          *Schedule
}
\`\`\`

### Booking (Pemesanan)
\`\`\`go
type Booking struct {
    ID          int
    ScheduleID  int
    UserName    string
    UserEmail   string
    Seats       int
    TotalPrice  float64
    BookingDate time.Time
    Next        *Booking
}
\`\`\`

## 🎨 Desain UI/UX

- **Theme**: Blue gradient dengan glass morphism effect
- **Responsive**: Mendukung desktop dan mobile
- **Modern**: Menggunakan CSS3 animations dan transitions
- **User-friendly**: Interface yang intuitif dan mudah digunakan

## 📊 Sample Data

Aplikasi sudah dilengkapi dengan sample data:
- 5 Kota (Jakarta, Bandung, Surabaya, Yogyakarta, Semarang)
- 6 Stasiun dengan informasi lengkap
- 5 Rute perjalanan
- 4 Jadwal kereta dengan harga

## 🔧 Pengembangan

### Menambah Fitur Baru
1. Tambahkan model di `model/models.go`
2. Buat controller di `controller/`
3. Tambahkan route di `main.go`
4. Buat template HTML di `templates/`

### Kustomisasi Tampilan
- Edit file CSS di dalam tag `<style>` pada template HTML
- Ubah warna tema dengan mengganti nilai gradient dan color variables
- Tambahkan animasi atau efek visual sesuai kebutuhan

## 🤝 Kontribusi

1. Fork repository ini
2. Buat branch fitur baru (`git checkout -b feature/AmazingFeature`)
3. Commit perubahan (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📝 License

Project ini dibuat untuk keperluan edukasi dan pembelajaran.

## 👨‍💻 Author

**[Nama Anda]**
- GitHub: [@username](https://github.com/username)
- Email: your.email@example.com

## 🙏 Acknowledgments

- Terima kasih kepada dosen dan teman-teman yang telah memberikan dukungan
- Inspirasi desain dari berbagai sumber UI/UX modern
- Komunitas Go Indonesia untuk pembelajaran dan sharing

---

⭐ **Jangan lupa berikan star jika project ini membantu!** ⭐

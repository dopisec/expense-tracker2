# Expense Tracker Fullstack

Aplikasi Expense Tracker fullstack dengan backend Go (Gin + Gorm + SQLite) dan frontend Next.js + Tailwind CSS.

## Struktur Project

```
.
├── backend/          # Backend Go API
│   ├── main.go
│   ├── config/
│   ├── models/
│   ├── handlers/
│   ├── middleware/
│   ├── utils/
│   └── data/        # SQLite database (auto-generated)
└── frontend/         # Frontend Next.js
    ├── app/
    ├── components/
    ├── lib/
    └── types/
```

## Backend Setup

### Prerequisites
- Go 1.21 atau lebih baru

### Installation & Run

**Penting:** Pastikan Go sudah terinstall dan ada di PATH. Jika menggunakan Go yang dikompilasi tanpa CGO, pastikan menggunakan driver pure Go.

```bash
cd backend
go clean -modcache  # Bersihkan cache modul (opsional, jika ada masalah)
go mod tidy         # Download dependencies dan update go.sum
go run main.go
```

Server akan berjalan di `http://localhost:8080`

**Troubleshooting CGO Error:**
Jika mendapatkan error "Binary was compiled with 'CGO_ENABLED=0'", pastikan:
1. File `go.mod` sudah memiliki `modernc.org/sqlite` di dependencies
2. File `go.mod` sudah mengecualikan semua versi `github.com/mattn/go-sqlite3`
3. Jalankan `go clean -modcache` lalu `go mod tidy` untuk membersihkan cache
4. Pastikan `backend/config/database.go` mengimpor `_ "modernc.org/sqlite"`

### API Endpoints

#### Authentication
- `POST /api/auth/register` - Register user baru
- `POST /api/auth/login` - Login dan dapatkan JWT token

#### Expenses (Protected - requires JWT token)
- `GET /api/expenses` - List semua expenses user
- `POST /api/expenses` - Create expense baru
- `PUT /api/expenses/:id` - Update expense
- `DELETE /api/expenses/:id` - Delete expense

### Database
SQLite database akan otomatis dibuat di `backend/data/expense.db` saat pertama kali menjalankan aplikasi.

## Frontend Setup

### Prerequisites
- Node.js 18+ dan npm/yarn

### Installation & Run

```bash
cd frontend
npm install
npm run dev
```

Aplikasi akan berjalan di `http://localhost:3000`

### Environment Variables (Optional)

Buat file `.env.local` di folder `frontend/` jika ingin mengubah API URL:

```
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## Fitur

- ✅ Autentikasi pengguna (Register & Login)
- ✅ JWT-based authentication
- ✅ CRUD Expenses
- ✅ Filter expenses berdasarkan user
- ✅ Modern UI dengan Tailwind CSS
- ✅ Responsive design

## Usage

1. Jalankan backend server terlebih dahulu
2. Jalankan frontend development server
3. Buka browser ke `http://localhost:3000`
4. Register akun baru atau login dengan akun yang sudah ada
5. Mulai menambahkan expenses!


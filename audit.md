## Grouping

```bash
Core
├── Identity
├── User
├── Post
├── Comment

Supporting
├── Like
├── Storage
├── Tag
├── Activity
├── Analytics
├── Notification

Infrastructure
├── Mail
├── JWT
├── Scheduler
├── Worker
├── Pipeline
├── Logger
├── Config
├── Database
├── Cache
├── Queue
```

## Core 
inti dari logic bisnis yang tersiolasi
- Identity: Entity/Identitas siapa penggunanya (User, Admin, dll) role dan permissions
    - Authentication: Logika Bisnis Untuk Memeastika Siapa Pengguna Yang Mengakses Resource 
    - Authorization : Logika Bisnis Siapa Yang Boleh Mengakses Resource
    - Credential : Bukti Kepemilkian
    - Session : Penyimpanan Sementara Untuk Menyimpan Informasi Yang Diperlukan
    - Verification : Logika Bisnis Untuk memastikan Kepemilikan

- User : Entity/identital yang dikenali 
    - Profile : Bertanggung Jawab Mengelolah data Public User sesuai Aturan bisnis
    - Activity: Bertanggung Jawab Menyimpan Aktivitas Pengguna Dalam Aplikasi Sesuai Kesepakan Policy
    - Configuration: Berangung Jawab Mengelolah Ketentuan User Terhadap Aplikasi Sesuai Policy

- Post : Entity/Konten Yang Dihasilkan Oleh User
    - Post: Beratanggung Jawab Membuat Konten
    - Update: Bertanggung Jawab Mengelolah Konten yang ada
    - Delete: Bertanggung Jawab Menghapus Kontent yang ada
    - Visibilty: Logika Bisnis Untuk Mengelola Visibilty Konten

- Commnet: Entity/Komenter Dari User Yang Dikenali
    - Post : Bertanggung Jawab Menmbuat Komentar Terhadap Parent
    - Update : Bertanggung Jawab Mengenelola Komentar yang sudah ada
    - Update : Bertanggung Jawab Melakukan Transaksi Pengahpus Komentar dan Anakan
    - Replay : Bertanggung Jawab Membuat Komentar Terhadap Parent Komentar yang Sudah Ada

# Supporting
Logika Tambahan untuk Membantu Logika Bisnis
- Like : Entity/Suka
    - Reacting : Bertanggung Jawab Membuat Like Dengan Referensi Parent
    - Toggle : Bertanggung Jawab Mengupdate Status Like Yang Sudah ada
    - Delete : Bertanggung Jawab Menghapus Entity Like Yang Sudah ada

- Storage: Tempat Penyimpanan Resource
    - Upload: Bertanggung Jawab Melakukan Transaksi Untuk Menyimpan Resource
    - Update: Bertanggung Jawab Melakukan Transaksi Untuk Membuat Temp Dan Rewrite Resurce Lama
    - Delete: Beranggung Jawab Melakukan transaksi Untuk menghapus Resurce Secara Aman
    - LifeCycle: Aturan Bisnis Bagaimana File Hidup Dan Konsisten
    - Cronjob : Bertanggung Jawab untuk Melakukan Tugas Pembersihan File Yang Mati(Orphan)

- Tag : Entity/Label Untuk Mengelompokkan Parent
    - Post: Bertanggung Jawab Membuat Label Baru dan Menjaga Duplikasi
    - Link: Menautkan Parent Dengan Daftar ID Label

- Activity : Entitiy/ Penimpanan Suatu Aktivitas Parent terhadapat Children
    - Create: Bertanggung Jawab Membuat Aktivitas Parent Sesuai Policy
    - Read: Membaca Aktivitas Sesuai Visibilty
    - Delete: Bertanggung Jawab Menghapus Aktivitas Sesuai Policy

- Analytics : Entitiy/ Penampung Data Statistik Parent
    - CronJob : Bertanggung Jawab Melakukan Pemasakan Data Untuk Siap Disajikan
    - View : Menyajikan Data Yang Sudah Dimasak Sesuai Parent ID
    
- Nontification : Logika Tambahan Untuk informasi
    - Push : Bertanggung Jawab Memberikan Informasi Sesuai Aktfitas Parent Ke Penerima
    - Seen : Bertanggung Jawab Untuk menandai Informasi Sudah Dibaca Oleh Penerima 

# Infrastructure
Logika Teknis Untuk Membantu Logika Bisnis
- Mail : Sistem Pengiriman Email
    - Send : Mengirim Email Ke Tujuan
    - Temp : Menyimpan Template Email
    - CronJob : Melakukan Tugas Berkala Untuk Mengirim Email
    
- JWT : JSON Web Token
    - Create : Membuat Token
    - Verify : Memverifikasi Token
    - Decode : Mendekode Token
    
- Scheduler : Penjadwalan Tugas
    - CronJob : Menjadwalkan Tugas
    - Job : Tugas Yang Akan Dijadwalkan
    - Queue : Antrian Tugas
    
- Worker : Pelaksana Tugas
    - Job : Tugas Yang Akan Dilaksanakan
    - Queue : Antrian Tugas
    
- Pipeline : Alur Kerja
    - Job : Tugas Yang Akan Dilaksanakan
    - Queue : Antrian Tugas
    
- Logger : Pencatat Aktivitas
    - Log : Catatan Aktivitas
    - CronJob : Melakukan Tugas Berkala Untuk Membersihkan Log
    
- Config : Konfigurasi
    - Read : Membaca Konfigurasi
    - Update : Mengupdate Konfigurasi
    - Reload : Memuat Ulang Konfigurasi
    
- Database : Basis Data
    - Read : Membaca Data
    - Update : Mengupdate Data
    - Delete : Menghapus Data
    
- Cache : Penyimpanan Sementara
    - Read : Membaca Data
    - Update : Mengupdate Data
    - Delete : Menghapus Data
    
- Queue : Antrian
    - Push : Menambahkan Data Ke Antrian
    - Pop : Mengambil Data Dari Antrian
    - Size : Mengambil Jumlah Data Dalam Antrian

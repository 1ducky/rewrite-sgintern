# Feture:

## auth
Login
Register
Forget Password
Login With Google
Verify Email
Refresh Token

## User
Profile(Public Visible)
Configure Account
History

## Post (User)
Draf Post
Upload Post
Upload Image
Update / Delete
Visibility
Statistic

## Post (Public)
Radom Post
Interaction Like
Share
Search Ny Tags
Search By Author


## comment
Add
Replay
Update
Delete
Like

## Storage
Upload
Get Key
Delete
Update
MetaData
CronJob

## Like Management
cronjob
add/update

# Technical

## Auth:
Domain Hanya Bertanggung jawab siapa pengunnanya 
- Login : Mengirim Credential Dan Melakukan validasi di server, server tidak mencocokan di tingkat db tapi mencocokan melalui primary(email) lalu melakukan validasi hash di server, setelah cocok server akan membuatkan field auth lalu membuat jwt dan dikirimkan menggunakan cookie httponly, frontend tidak memagang tanggung jawab auth, membiarkan browser yang menanganinya dan saat butuh auth client cukup mengirim cookie tersebut, dan server menyimpannya ke context agar bisa dibawah selama proses request
- Register : Mengirim credetial dan melakukan validasi di server, melakukan pengecekan dupikasi, tingkat password dan lainnya, jika tahap validasi berhasil server bisa membuat akun dengan langsung, dan jika ingin langsung login, usecase dialihkan ke Login dan melakukan validasi lagi dan mengirim token ke pengirim
- Forget Password (Opsional Sementara): mengirim email lalu server membuat token untuk refresh password lalu dikirimkan ke service email, service ini memiliki batasan 1 menit untuk 1 akun untuk menghindari spam, lalu setelah mengirim token yang valid selama belum kadaluarsa, user akan dialihkan untuk mengubah password, setelah berhasil diubah user akan dialihkan ke service login, dan mengirim penringatan email jika mencurigakan valid selama 1 hari, jika dicurigakan akan diberikan tindakan revoke semua token dan menganti password
- Login With Google: user akan login menggunakan google lalu server menyimpan credential yang didapatkan oleh google dengan email dan token
- Verivy Email : mengirim sebuah code valid jika code yang dikirim berhasil makan mengirim sinya verify email berhasil
- Revoke Token : mengubah versi Token dan saat request kembali dan mengalami unauthorized forntend bisa mengahpus token dan mengalihkan ke login

## User:
- Profile(Public Visibilty) : Mengirm Form untuk informasi string dan melakukan pengecekan auth dan authorized, menyimpan perubahan yang dikirim, avatar mengirim terpisah dan melakuan stream untuk menyimpannya ke image service dan mgirim path terbaru ke frontend untuk disimpan cache dan memngirim informasi profile terbaru saat frontend mengirim sinyal
- Configuration Account : Melihat aktivitas Token yang aktif, Revoke token yang dipilih dan memperbarui password
- History : Menampilkan daftar aktivitas user sesuai kategori, misal comment, like, upload Post

## Post (User):
- Draft Post : Menampilkan Daftar Post
- Upload Post : Mengirim form post dan disimpan ke db
- Upload Image : mengirim list image ke service storage dan memberikan id untuk mereferensikan
- Update : update Post berdasarakan id post dan id author, untuk update image akan menerima id img yang dihapus untuk menghapus referensi dan mengubah status menjadi orphan
- Delete : hapus post berdasarkan id post dan id author, dan image yang direferensikan diubah status ke orphan
- Visibilty : memberikan pilihan visibilitas (Public, draf)
- Statistic : Menampilkan statistik yang sudah di masak (cache) post berdasarkan id post

## Post (Public):
- Random Post : Menampilkan Post secara acak dengan format yang sudah di optimasi, dan visibilitas public
- Interaction Like : membuat Like atau update status like jika sudah ada berdasarkan id parent
- Share : membuat share post berdasarkan id parent
- Search By Tag : Menampilkan Post berdasarkan Tag 
- Search By Author : Menampilkan Post berdasarkan Author

## Comment:
- Add : Kirim Tanggapan berdasarkan Id Pengomentar dan Parent(Post/Comment) Id
- Replay : Mengirim balasan berdasarkan id parent(Comment) dan id pengomentar
- Update : Mengupdate comment Berdasarkan Comment id
- Delete : Menghapus Commentar berdasarkan Comment id beserta Child yang menggunakan parent id commentar
- Like : Memberikan Like dengan comment id(parent like)

## Storage:
- Upload Image : Mengirim file gambar dan menyimpan metadata ke storage service, mengubah status file menjadi inUse Dan wajib transaction, jika gagal melakukan rollback
- Get Key : Mendapatkan File berdasarkan id yang tersimpan
- Delete : Menghapus File sesuai Id, dilakukan oleh cronjob
- Update : buat temporary untuk menampung file baru, lalu rewrite metadata file lama dengan metadata file baru
- MetaData : Menyimpan Data File untuk keperluan analisis dan cronjob
- CronJob : Melakukan Daftar Tugaa secara berkala Untuk Keperluan Storage Service dan Monitoring

## Like Management:
- CronJob : Melakukan Tugas Berkala Untuk Mengevaluasi Data Like(Memasak/Cache) Untuk Keperluan statistic
- add/update : Membuat data Like Dan Update Status Like Sesuai Query bersarkan parent id dan uthor id

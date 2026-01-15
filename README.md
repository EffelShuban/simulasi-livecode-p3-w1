[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-2e0aaae1b6195c2367325f4f02e2d04e9abb55f0b24a779b69b11b9e10269abc.svg)](https://classroom.github.com/online_ide?assignment_repo_id=22106413&assignment_repo_type=AssignmentRepo)
# FTGO-AOH-P3-V1-SLC
## RULES
1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Livecode
   berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor livecode ini 0.
3. Clone repo ini kemudian buatlah **branch dengan nama kalian**.
4. Kerjakan pada file Golang (\*.go) yang telah disediakan.
5. Waktu pengerjaan: **120 menit** untuk **2 soal**.
6. **Pada text editor hanya ada file yang terdapat pada repository ini**.
7. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
8. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
   lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
9. Lakukan `git push origin <branch-name>` dan create pull request **hanya jika waktu Livecode telah usai (bukan ketika kalian sudah selesai
   mengerjakan)**. Tuliskan nama lengkap kalian saat membuat pull request dan assign buddy.
10. **Penilaian berbasis logika dan hasil akhir**. Pastikan keduanya sudah benar.




## Notes
Livecode ini memiliki bobot nilai sebagai berikut:

|Criteria| Meet Expectations                         | Points |
|---|-------------------------------------------|--------|
|Problem Solving - Functionality| Create, read, update, and delete workouts | 20 pts |
|   | Filter Data                               | 10 pts |
|   | Scheduler with CronJob                    | 10 pts |
|Data Management| Use of MongoDB                            | 20 pts |
|| Appropriate database schema design        | 10 pts |
|| Data consistency and integrity            | 5 pts  |
|Code Clarity| Consistent code style                     | 5 pts  |
|| Proper use of Go and Echo framework       | 10 pts |
|| Proper error handling and logging         | 10 pts |




#### KETENTUAN
Here are some dos and don'ts to consider when working on this livecode:

Do's:

- Do use proper HTTP status codes to indicate the outcome of an API request
- Do validate all user inputs to ensure data consistency and integrity
- Do use secure authentication and authorization mechanisms
- Do follow best practices for error handling and logging
- Do design the database schema based on the specific requirements of the app
- Do use proper indexing to optimize database queries
- Do use consistent code style and follow best practices for Go and Echo framework
- Do document the API endpoints and their expected inputs and outputs
- Do test the API thoroughly before deployment

Don'ts:

- Don't expose sensitive user data in the API responses
- Don't store plain text passwords in the database
- Don't hard-code secrets (e.g. API keys, database credentials) in the code
- Don't perform database operations in the API request handler functions
- Don't use deprecated or insecure versions of libraries and frameworks
- Don't store unnecessary data in the database
- Don't allow unauthorized users to access sensitive API endpoints or data
- Don't deploy the API without proper testing and security measures




# SIMULASI LIVECODE 1
## **Game Center**

## Objectives
- Mempelajari penggunaan MongoDB dengan Golang dan Echo framework.
- Mengimplementasikan cron job untuk menjalankan tugas tertentu pada waktu yang ditentukan.
- Membuat aplikasi "Game Center" yang memungkinkan user melakukan create, read, update dan delete data game serta melakukan Penjadwalan untuk update versi game.

#### Sebagai tambahan dari requirement yang sudah diberikan sebelumnya, Student juga diharapkan untuk memahami dan menerapkan konsep-konsep berikut:
- Cloud Deployment using GCP
Student diharapkan untuk mengimplementasikan Cloud Deployment menggunakan Google Cloud Platform (GCP).
Pastikan aplikasi Anda dapat diakses secara publik setelah deployment.
Sediakan dokumentasi sederhana mengenai langkah-langkah deployment yang Anda lakukan.
- Job Scheduling
Implementasikan konsep job scheduling untuk beberapa proses yang memerlukannya, seperti proses pembaharuan data atau pembersihan data yang tidak diperlukan.
- Unit Test
Buat unit test untuk memastikan bahwa setiap fungsi atau method dalam aplikasi Anda bekerja dengan semestinya.
- Docker
Kontainerisasi aplikasi Anda menggunakan Docker.
Pastikan Anda menyediakan Dockerfile dan dokumentasi singkat tentang bagaimana menjalankan aplikasi Anda menggunakan Docker.

## Directions

Buatlah sebuah RESTful API dengan menggunakan Golang Framework (Echo), cronjob dan database MongoDB untuk menyimpan data Game dengan schema berikut :

| Game |
|-------------------------|
| GameID INT          | 
| Title VARCHAR(255)      | 
| Description TEXT        | 
| ReleaseDate DATE        | 
| Version VARCHAR(20)    | 
| Platform VARCHAR(50)    | 
| UpdatedDate date    | 
| GoToUpdate int (1/0)    | 



### Step 1: 
Buat Database dengan nama Collection dan field seperti DB Schema diatas

### Step 2:
Buat endpoint untuk fungsi CRUD pada data Game

### Step 3:
1. Lakukan Penjadwalan dengan Cron Job untuk melakukan update versi (+1) dari versi sebelumnya
2. Perubahan versi hanya dilakukan ketika GoToUpdate bernilai 1 dan dilakukan pada jam 12 malam
3. Saat fungsi cronjob jalan maka akan merubah versi dan merubah updatedDate sesuai tanggal dilakukannya update versi
4. Contoh : versi sebelumnya V.1.0.0, maka versi selanjutnya adalah V.1.0.1

Notes:
- Don't rush through the problem or try to solve it all at once.
- Don't copy and paste code from external sources without fully understanding how it works.
- Don't hardcode values or rely on assumptions that may not hold true in all
  cases.
- Don't forget to handle error cases or edge cases, such as invalid input or unexpected behavior.
- Don't hesitate to refactor your code or make improvements based on feedback or new insights.


### Goodluck :)

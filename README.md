# Golang Web
Membuat golang web dapat menggunakan package `net/http`

## Server
Untuk membuat server di golang bisa menggunakan struct `Server` yang ada di package `net/http`
Dan untuk mejalankan server bisa dengan function `(SERVER)ListenAndServe()`

## Handler
Server hanya bertugas sebagai web server, sedangkan untuk menerima HTTP Request yang masuk, kita butuh yang namanya handler.<br>
Handler di golang di representasikan dalam interface, dimana dalam kontraknya terdapat sebuah function bernama ServeHTTP() yang digunakan sebagai function yang akan dieksekusi ketika menerima HTTP Request
## HandlerFunc
Salah satu implementasi dari interface Handler adalah HandlerFunc
Kita bisa menggunakan HandleFunc untuk membuat function HTTP

## ServeMux
ServeMux adalah implementasi Handler yang bisa mendukung multiple endpoint.
### URL Pattern
Jika dalam ServeMux kita menambahkan diakhirannya dengan garis artinya semua url tersebut akan menerima path dengan awalan tersebut<br>
Namun jika terdapat URL Pattern yang lebih panjang URL tersebut yang akan di prioritaskan

## Request
Request adalah struct yang merepresesntasikan HTTP Request yang dikirim oleh web browser <br>
Semua informasi request yang dikirim bisa kita dapatkan di Request, seperti URL, HTTP Method, HTTP Header, HTTP Body, dan lain-lain

## HTTP Test
HTTP Test adalah package khusus untuk membuat unit test terhadap fitur web yang kita buat <br>
Package ini berada di `net/http/httptest`<br>
Dengan package ini kita bisa menjalankan testing handler web tanpa harus menjalankan aplikasi webnya<br>
### httptest.NewRequest()
`NewRequest(method, url, body)`merupakan function yang digunakan untuk membuat http.Request<br>
kita bisa menentukan method, url, body yang akan kita kirim sebagai simulasi unit test <br>
Selain itu kita juga bisa menambahkan informasi tambahan seperti cookie, header, dan lain-lain
### httptest.NewRecorder()
`httptest.NewRecorder()` merupakan function yang digunakan untuk membuat ResponseRecorder<br>
ResponseRecorder merupakan sebuah struct bantuan untuk merekan HTTP Response dari hasil testing yang kita lakukan <br>

## Query Parameter
Untuk membuat query paramter kita bisa menggunakan `?name=value`pada URL-nya
### url.URL
Dalam parameter Request, terdapat atribut URL yang berisi data url.URL<br>
Kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan method `Query()` yang akan mengembalikan map
### Multiple Query Parameter
Dalam spesifikasi URL, Kita bisa menambahkan lebih dari satu query parameter<br>
Ini cocok sekali jika kita memang ingin mengirim banyak parameter data ke server, kita cukup menambahkan query parameter lainnya <br>
Untuk menambahkan query parameter kita bisa menggunakan tanda `&` lalu diikuti dengan query paramter berikutnya
### Multiple Value Query Parameter
Sebenarnya URL melakukan parsing query parameter dengan menyimpannya dalam `map[string][]string`<br>
Artinya dalam satu query parameter kita bisa memasukkan beberapa value <br>
Caranya dengan menambahkan query parameter dengan nama yang sama, namun berbeda value, misalnya: 
`name=Meizaluna&name=Wulandari`

## Header
Header adalah informasi tambahan yang biasa dikirim dari client ke server ataupun sebaliknya<br>
Jadi header tidak hanya ada pada HTTP Request tetapi juga ada pada HTTP Response <br>
Saat menggunakan browser, biasanya seacara otomatis header akan ditambahkan oleh browser, seperti informasi browser, jenis tipe konten yang diterima atau  dikirim oleh browser, dan lain-lain
### Request Header 
Untuk mengambil request header kita bisa menggunakan `Request.Header`<br>
Header berisi `map[string][]string`<br>
Berbeda dengan Query Parameter yang case sensitif, header key tidaklah case sensitif
### Response Header
Jika ingin menambahkan header pada response, kita bisa menggunakan `ResponseWriter.Header()`

## Form Post
Jika menggunakan method POST, maka semua data form akan dikirim via body HTTP Request <br>
### Request.PostForm
Semua data yang dikirim dari client, secara otomatis akan disimpan dalam attribute `Request.PostForm`<br>
Namun sebelum mengambil datanya kita harus melakukan parsing terlebih dahulu dengan function `Request.ParseForm`

## Response Code
Response Code merupakan representasi dari kode response(200, 404, 302, dll)<br>
Dari response code ini kita bisa melihat apakah sebuah request yang kita kirim itu sukses diproses oleh server atau gagal<br>
### Mengubah Response Code
Jika kita tidak menyebutkan response code, maka response codenya adalah 200 OK<br>
Jika kita ingin mengubahnya, kita bisa menggunakan function `ResponseWriter.WriteHeader(int)`<br>
Semua data status code juga sudah disediakan di golang, jadi kita bisa menggunakan variabel yang sadah di sediakan

## Cookie
### Stateless
HTTP merupakan stateless antara client dan server, artinya server tidak menyimpan data apapun dari untuk mengingat setiap request request dari client<br>
Hal ini bertujuan agar mudah melakukan scalability disisi server<br>
Namun jika ingin mengingat sebuah data request dari client seperti login bisa dengan memanfaatkan cookie
### Cookie
Cookie adalah fitur di HTTP dimana server bisa memberi response cookie (key-value) dan client akan menyimpan cookie tersebut di web browser<br>
Request selanjutnya, client akan selalu membawa cookie tersebut secara otomatis<br>
Dan Server secara otomatis akan selalu menerima data cookie yang dibawa oleh client setiap kali client mengirim request
### Membuat Cookie
Cookie merupakan data yang dibuat di server dan sengaja agar disimpan di web browser <br>
Untuk membuat cookie di server, kita bisa menggunakan function `http.SetCookie()`

## File Server 
Golang memiliki sebuah function yang bernama `FileServer` <br>
Dengan ini, kita bisa membuat handler di golang web yang digunakan sebagai static file server <br>
Dengan menggunakan `FileServer` kita tidak perlu manual me-load file lagi <br>
FileServer adalah Handler, jadi bisa kita tambahkan ke dalam `http.Server` atau `http.ServeMux`
### 404 Not Found
Jika coba jalankan `TestGetCookie()`, saat kita membuka misalnya `/static/index.html`, maka akan dapat error `404 Not Found`<br>
Hal ini dikarenakan FileServer akan membaca url, lalu mencari file berdasarkan urlnya, jadi jika kita membuat `/static/index.html`, maka FileServer akan mencari file `/resources/static/index.html`<br>
Hal ini menyebabkan 404 not found karena memang filenya tidak ditemukan <br>
Oleh karena itu, kita menggunakan function `httpStripPrefix()` untuk menghapus prefix url 
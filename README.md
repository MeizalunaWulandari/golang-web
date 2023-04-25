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
### 404 Not Found (FileServer)
Jika coba jalankan `TestGetCookie()`, saat kita membuka misalnya `/static/index.html`, maka akan dapat error `404 Not Found`<br>
Hal ini dikarenakan FileServer akan membaca url, lalu mencari file berdasarkan urlnya, jadi jika kita membuat `/static/index.html`, maka FileServer akan mencari file `/resources/static/index.html`<br>
Hal ini menyebabkan 404 not found karena memang filenya tidak ditemukan <br>
Oleh karena itu, kita menggunakan function `httpStripPrefix()` untuk menghapus prefix url 
### Golang Embed
Dalam golang embed kita bisa embed file ke dalam binary distribution file, hal ini mempermudah sehingga kita tidak perlu meng-copy static file lagi<br>
Golang Embed juga memiliki fitur yang bernama `embed.FS`, fitur ini bisa diintergrasikan dengan `FileServer`
### 404 Not Found (Embed)
Jika kita coba jalankan, dan coba buka `/static/index.html`, maka kita akan mendapatkan error `404 Not Found` <br>
Hal ini terjadi karena pada `golang embeed` nama folder ikut menjadi nama resourcenya, misalnya `resources/index.html`, jadi untuk mengaksesnya kita perlu menggunakan URL `/static/resources/index.html`<br>
Jika kita ingin langsung mengakses file `index.html` tanpa menggunakan `resources`, kita bisa mengunakan function `fs.Sub()` untuk mendapatkan sub directory

## ServeFile
Kadang ada kasus misalnya kita hanya ingin menggunakan static file sesuai dengan yang kita inginkan<br>
Hal bisa dilakukan menggunakan function `http.ServeFile()`
Dengan menggunakan function ini, kita bisa menemukan file mana yang ingin kita tulis ke http response
### Golang Embed
Parameter function `http.ServeFile` hanya berisi string file name, sehingga tidak bisa menggunakan golang embed<br>
Namun bukan berarti kita tidak bisa menggunakan golang embed, karena jika untuk melakukan load file, kita hanya butuh menggunakan package `fmt` dan `ResponseWriter` saja

## Templates
Di golang terdapat fitur HTML Template, yaitu fitur yang bisa kita gunakan untuk membuat HTML yang dinamis<br>
### HTML Template
Fitur HTML Template terdapat pada package `html/template`<br>
Sebelum menggunakan HTML Template, kita perlu membuat templatenya dulu<br>
Template bisa berupa string atau file<br>
Bagian dinamis pada HTML Template, adalah bagian yang menggunakan tanda `{{ }}`
### Membuat Template
Pada saat membuat template dengan string, kita bisa perlu memberi nama template nya<br>
Dan saat membuat text template, cukup buat text html, dan untuk konten yang dinamis kita bisa menggunakan `{{.}}` contonya : <br>
`<html><body>{{.}}</body></html>`
### Template Dari File
Selain membuat template dari string, kita juga bisa membuat template langsung dari file<br>
Hal ini mempermudah kita, karena bisa langsung membuat file HTML<br>
Saat membuat template menggunakan file, secara otomatis nama file akan menjadi nama templaten, misalnya jika kita punya file `simple.html`, maka nama templatenya adalah `simple.html`
### Template Directory
Jarang sekali kita menyebutkan file template satu persatu<br>
Alangkah baiknya  untuk template kita simpan disatu directory <br>
Golang Template mendukung proses load template dari directory<br>
Hal ini memudahkan kita, sehingga tidak perlu menyebutkan nama filenya satu persatu
### Template Dari Golang Embed
Karena sudah ada golang embed, jadi direkomendasikan menggunakan golang embed untuk menyimpan data file template<br>
Menggunakan golang embed menjadikan kita tidak perlu ikut mengcopy file template lagi, karena sudah otomatis di embed dalam distribution file golang.

## Template Data
Saat kita membuat template, kadang kita ingin menambahkan banyak data dinamis <br>
Hal ini bisa kita lakukan dengan cara menggunakan `struct` atau `map`<br>
Namun perlu dilakukan perubahan didalam text templatenya, kita  perlu memberi tahu Field atau key mana yang akan digunakan untuk mengisi data data dinamis di template tersebut<br>
Kita bisa menyebukan dengan cara seperti ini `{{.NamaField}}`

## Template Action
Golang template mendukung perintah action, seperti percabangan, perulangan dan lain-lain
### if else
`{{if.Value}}T1{{end}}`, jika `Value` tidak kosong maka `T1` akan dieksekusi, jika kosong, tidak ada yang dieksekusi<br>
`{{if.Value}}T1{{else}}T2{{end}}`,  jika `Value` tidak kosong, maka `T1` akan dieksekusi, Jika kosong maka `T2` yang akan dieksekusi<br>
`{{if.Value1}}T1{{else if.Value2}}T2{{else}}T3{{end}}` jika `Value` tidak kosong maka `T1` akan dieksekusi, Jika `Value2` tidak kosong, maka `T2` akan dieksekusi, jika semuanya kosong maka `T3` akan dieksekusi
### Operator Perbandingan 
Golang template juga mendukung operator perbandingan, ini cocok ketika kita butuh melakukan perbandingan number di if statement Operator Perbandingan di golang sedikit berbeda<br>
`eq` => `==`
`ne` => `!=`
`lt` => `<`
`le` => `<=`
`gt` => `>`
`ge` => `>=`
### Range
`Range` Digunakan untuk melakukan iterasi data template<br>
Tidak ada perulangan seperti biasa menggunakan for di golang template<br>
Yang bisa kita lakukan adalah menggunakan range untuk mengiterasi data tiap data array, slice, map atau channel<br>
`{{range $index, $element := .Value}} T1 {{end}}`, jika value memiliki data, maka T1 akan dieksekusi sebanyak element value, dan kita bisa menggunakan `$index` untuk mengakses `index` dan `$element` untuk mengakses `element`<br>
`{{range $index, $element := Value}}T1{{else}}T2{{end}}`, Sama seperti sebelumnya, namun jika `Value` tidak memiliki element apapun, maka `T2` yang akan dieksekusi
### With
Kadang kita sering membuat nested struct <br>
Jika menggunakan template, kita bisa mengaksesnya menggunakan `.Value.NestedValue`<br>
Di template terdapat action `With` yang bisa digunakan untuk mengubah scope dot menjadi object yang kita mau<br>
`{{with.Value}}T1{{end}}` jika value tidak kosong, T1 semua dot akan merefer ke value<br>
`{{with.Value}}T1{{else}}T2{{end}}` sama seperti sebelumjya, namun jika value kosong maka `T2` akan dieksekusi
### Komentar
Template juga mendukung komentar<br>
Komentar secara otomatis akan hilang ketika template text di parsing (tidak akan menjadi komentar HTML )<br>
Untuk membuat komentar sangat sederhana, kita bisa gunakan `{{/*Contoh Komentar*/}}`

## Template Layout
Saat kita membuat halaman website, kadang ada beberapa bagian yang selalu sama, misalnya header dan footer<br>
Best Practicenya jika terdapat bagian yang selalu sama, disarankan untuk disimpan pada template yang terpisah, agar bisa digunakan ditemplate lain<br>
Golang template mendukung import dari template lain
### Import Template
Untuk melakukan Import, kita bisa menggunakan perintah berikut:
`{{template "nama"}}` artinya kita akan mengimport template "nama" tanpa memberi data apapun<br>
`{{template "nama".Value}}` artinya kita akan mengimport template "nama" dengan memberikan data `Value`
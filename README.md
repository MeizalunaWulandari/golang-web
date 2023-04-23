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
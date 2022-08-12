package handler

import (
	"fmt"
	"golangniomic/echo/server"
	_ "mysql-master"
	"net/http"

	"github.com/labstack/echo"
)

type menu struct {
	Id_menu    string
	Nama_menu  string
	Deskripsi  string
	Jenis      string
	Harga      string
	Url_gambar string
}

var data []menu

func BacaData(c echo.Context) error {
	ambil_semua_menu()
	return c.JSON(http.StatusOK, data)
}

func TambahData(c echo.Context) error {
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
	}

	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("insert into tbl_menu values (?,?,?,?,?,?)", nil, nama, deskripsi, url_gambar, jenis, harga)

	if err != nil {
		fmt.Println("Menu Gagal Ditambahkan :(")
		return c.JSON(http.StatusOK, "Gagal menambahkan menu")
	} else {
		fmt.Println("Menu Berhasil Ditambahkan :D")
		return c.JSON(http.StatusOK, "Berhasil menambahkan menu")
	}
}

func ambil_semua_menu() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("Select * from tbl_menu")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

}

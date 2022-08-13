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
	defer db.Close()
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

func TambahOrder(c echo.Context) error {
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	var id = c.FormValue("id")
	var nama_pemesan = c.FormValue("nama_pemesan")
	var nomor_telepon = c.FormValue("nomor_telepon")
	var jumlah = c.FormValue("jumlah")
	var alamat = c.FormValue("alamat")

	_, err = db.Exec("insert into tbl_order values (?,?,?,?,?,?)", nil, id, nama_pemesan, nomor_telepon, alamat, jumlah)

	if err != nil {
		fmt.Println("Pesanan Gagal Dibuat :(")
		return c.HTML(http.StatusOK, "<strong>Gagal menambahkan Pemesanan</strong>")
	} else {
		fmt.Println("Pesanan Berhasil Dibuat :D")
		return c.HTML(http.StatusOK, "<script>alert('Berhasil menambahkan pesanan, silahkan tunggu telepon dari kami, terimakasih :D'); window.location = 'http://localhost:1323';</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func UpdateData(c echo.Context) error {
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	var id_menu = c.FormValue("Id_menu")
	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("update tbl_menu set nama_menu = ? , deskripsi = ? , harga = ? , jenis = ? , url_gambar = ? where id_menu = ?", nama, deskripsi, harga, jenis, url_gambar, id_menu)

	if err != nil {
		fmt.Println("Menu Gagal diubah :(")
		return c.JSON(http.StatusOK, "Gagal mengubah data")
	} else {
		fmt.Println("Menu Berhasil diubah :D")
		return c.JSON(http.StatusOK, "Berhasil mengubah data")
	}
}

//Delete Data 
func HapusData(c echo.Context) error {
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	var id_menu = c.FormValue("Id_menu")

	_, err = db.Exec("delete from tbl_menu where id_menu = ?", id_menu)

	if err != nil {
		fmt.Println("Menu Gagal dihapus :(")
		return c.JSON(http.StatusOK, "Gagal menghapus data")
	} else {
		fmt.Println("Menu Berhasil dihapus :D")
		return c.JSON(http.StatusOK, "Berhasil menghapus data")
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

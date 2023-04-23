package controllers

import (
	"github.com/YungBenn/go-mysql-test/internal/database"
	"github.com/YungBenn/go-mysql-test/internal/model"
	"github.com/gofiber/fiber/v2"
)

// INSERT MAHASISWA
func InsertMahasiswa(c *fiber.Ctx) error {
	mahasiswa := new(model.Mahasiswa)

	if err := c.BodyParser(mahasiswa); err != nil {
		panic(err)
	}

	nama := mahasiswa.Nama
	usia := mahasiswa.Usia
	gender := mahasiswa.Gender
	jurusan := mahasiswa.Jurusan
	hobi := mahasiswa.Hobi

	db := database.ConnectDB()
	defer db.Close()

	result, err := db.Exec("INSERT INTO mahasiswa (nama, usia, gender, tanggal_registrasi) VALUES (?, ?, ?, NOW())", nama, usia, gender)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal memasukkan data mahasiswa": err.Error(),
		})
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal mendapatkan ID mahasiswa": err.Error(),
		})
	}

	_, err = db.Exec("INSERT INTO jurusan (nama) VALUES (?)", jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal memasukkan data jurusan": err.Error(),
		})
	}

	_, err = db.Exec("INSERT INTO hobi (nama) VALUES (?)", hobi)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal memasukkan data hobi": err.Error(),
		})
	}

	idHobi, err := result.LastInsertId()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal mendapatkan ID hobi": err.Error(),
		})
	}

	_, err = db.Exec("INSERT INTO mahasiswa_hobi (id_mahasiswa, id_hobi) VALUES (?, ?)", id, idHobi)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal memasukkan data mahasiswa_hobi": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Berhasil memasukkan data mahasiswa",
	})

}

// UPDATE MAHASISWA BY ID
func UpdateMahasiswa(c *fiber.Ctx) error {
	mahasiswa := new(model.Mahasiswa)
	if err := c.BodyParser(mahasiswa); err != nil {
		panic(err)
	}

	db := database.ConnectDB()
	defer db.Close()

	id := c.Params("id")
	nama := mahasiswa.Nama
	usia := mahasiswa.Usia
	gender := mahasiswa.Gender
	jurusan := mahasiswa.Jurusan
	hobi := mahasiswa.Hobi

	_, err := db.Exec("UPDATE mahasiswa SET nama=?, usia=?, gender=? WHERE id=?", nama, usia, gender, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal mengupdate data mahasiswa": err.Error(),
		})
	}

	_, err = db.Exec("UPDATE jurusan SET nama=? WHERE id=?", jurusan, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal mengupdate data jurusan": err.Error(),
		})
	}

	_, err = db.Exec("UPDATE hobi SET nama=? WHERE id=?", hobi, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal mengupdate data hobi": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Berhasil mengupdate data mahasiswa",
	})
}

// GET ALL MAHASISWA
func GetAllMahasiswa(c *fiber.Ctx) error {
	db := database.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM mahasiswa")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal mendapatkan data mahasiswa": err.Error(),
		})
	}
	defer rows.Close()

	var mahasiswas []model.GetMahasiswa
	for rows.Next() {
		var mahasiswa model.GetMahasiswa
		err := rows.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TanggalReg)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Gagal memindai data mahasiswa": err.Error(),
			})
		}
		mahasiswas = append(mahasiswas, mahasiswa)
	}

	for i, m := range mahasiswas {
		rows, err = db.Query("SELECT j.nama, h.nama FROM jurusan j INNER JOIN mahasiswa_hobi mh ON j.id=mh.id_mahasiswa INNER JOIN hobi h ON h.id=mh.id_hobi WHERE mh.id_mahasiswa=?", m.ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Gagal mendapatkan data jurusan dan hobi": err.Error(),
			})
		}
		defer rows.Close()

		for rows.Next() {
			var jurusan, hobi string
			err := rows.Scan(&jurusan, &hobi)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"Gagal memindai data jurusan dan hobi": err.Error(),
				})
			}
			mahasiswas[i].Jurusan = append(mahasiswas[i].Jurusan, jurusan)
			mahasiswas[i].Hobi = append(mahasiswas[i].Hobi, hobi)
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"data": mahasiswas,
	})
}

// GET DETAIL MAHASISWA BY ID
func GetDetailMahasiswa(c *fiber.Ctx) error {
	db := database.ConnectDB()
	defer db.Close()

	id := c.Params("id")

	var mahasiswa model.GetMahasiswa
	err := db.QueryRow("SELECT * FROM mahasiswa WHERE id=?", id).Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TanggalReg)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal mendapatkan data mahasiswa": err.Error(),
		})
	}

	rows, err := db.Query("SELECT j.nama, h.nama FROM jurusan j INNER JOIN mahasiswa_hobi mh ON j.id=mh.id_mahasiswa INNER JOIN hobi h ON h.id=mh.id_hobi WHERE mh.id_mahasiswa=?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal mendapatkan data jurusan dan hobi": err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var jurusan, hobi string
		err := rows.Scan(&jurusan, &hobi)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Gagal memindai data jurusan dan hobi": err.Error(),
			})
		}
		mahasiswa.Jurusan = append(mahasiswa.Jurusan, jurusan)
		mahasiswa.Hobi = append(mahasiswa.Hobi, hobi)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": mahasiswa,
	})

}

// DELETE MAHASISWA BY ID
func DeleteMahasiswa(c *fiber.Ctx) error {
	db := database.ConnectDB()
	defer db.Close()

	id := c.Params("id")

	_, err := db.Exec("DELETE FROM mahasiswa_hobi WHERE id_hobi=?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal menghapus data hobi mahasiswa": err.Error(),
		})
	}

	_, err = db.Exec("DELETE FROM hobi WHERE id=?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal menghapus hobi mahasiswa": err.Error(),
		})
	}

	_, err = db.Exec("DELETE FROM jurusan WHERE id=?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal menghapus jurusan mahasiswa": err.Error(),
		})
	}

	_, err = db.Exec("DELETE FROM mahasiswa WHERE id=?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Gagal menghapus data mahasiswa": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Data mahasiswa berhasil dihapus",
	})
}

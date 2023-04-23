package routes

import (
	"github.com/YungBenn/go-mysql-test/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	mahasiswa := app.Group("/mahasiswa")

	mahasiswa.Post("/insert", controllers.InsertMahasiswa)
	mahasiswa.Get("/", controllers.GetAllMahasiswa)
	mahasiswa.Get("/:id", controllers.GetDetailMahasiswa)
	mahasiswa.Put("/update/:id", controllers.UpdateMahasiswa)
	mahasiswa.Delete("/delete/:id", controllers.DeleteMahasiswa)
}

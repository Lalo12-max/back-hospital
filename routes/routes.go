package routes

import (
	"hospital-system/handlers"
	"hospital-system/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api/v1", middleware.GeneralAPIRateLimit())

	// Rutas de autenticación (sin protección)
	auth := api.Group("/auth")
	auth.Post("/register", middleware.RegisterRateLimit(), handlers.Register)
	auth.Post("/login", middleware.AuthRateLimit(), handlers.Login)
	auth.Post("/refresh", middleware.AuthRateLimit(), handlers.RefreshToken)
	auth.Get("/mfa/setup/:user_id", handlers.InitialMFASetup)
	auth.Post("/mfa/setup/:user_id/verify", handlers.VerifyInitialMFASetup)

	// Rutas protegidas con JWT y Role Guard
	protected := api.Group("/", middleware.JWTMiddleware(), middleware.RoleGuard(), middleware.MedicalRateLimit())

	// CRUD Usuarios - Solo Admin tiene acceso completo
	usuarios := protected.Group("/usuarios")
	usuarios.Get("/", handlers.GetUsuarios)
	usuarios.Get("/:id", handlers.GetUsuario)
	usuarios.Post("/", middleware.AdminRateLimit(), handlers.CreateUsuario)
	usuarios.Put("/:id", handlers.UpdateUsuario)
	usuarios.Delete("/:id", middleware.AdminRateLimit(), handlers.DeleteUsuario)

	// CRUD Consultorios - Admin: completo, Médico/Enfermera: solo lectura
	consultorios := protected.Group("/consultorios")
	consultorios.Get("/", handlers.GetConsultorios)
	consultorios.Get("/:id", handlers.GetConsultorio)
	consultorios.Post("/", middleware.AdminRateLimit(), handlers.CreateConsultorio)
	consultorios.Put("/:id", middleware.AdminRateLimit(), handlers.UpdateConsultorio)
	consultorios.Delete("/:id", middleware.AdminRateLimit(), handlers.DeleteConsultorio)

	// CRUD Consultas - Admin/Médico: completo, Enfermera: sin delete, Paciente: solo lectura
	consultas := protected.Group("/consultas")
	consultas.Get("/", handlers.GetConsultas)
	consultas.Get("/:id", handlers.GetConsulta)
	consultas.Post("/", handlers.CreateConsulta)
	consultas.Put("/:id", handlers.UpdateConsulta)
	consultas.Delete("/:id", handlers.DeleteConsulta)

	// CRUD Expedientes - Admin/Médico: completo, Enfermera: sin delete, Paciente: solo lectura
	expedientes := protected.Group("/expedientes")
	expedientes.Get("/", handlers.GetExpedientes)
	expedientes.Get("/:id", handlers.GetExpediente)
	expedientes.Post("/", handlers.CreateExpediente)
	expedientes.Put("/:id", handlers.UpdateExpediente)
	expedientes.Delete("/:id", middleware.AdminRateLimit(), handlers.DeleteExpediente)

	// CRUD Horarios - Admin/Médico: completo, Enfermera: solo lectura
	horarios := protected.Group("/horarios")
	horarios.Get("/", handlers.GetHorarios)
	horarios.Get("/:id", handlers.GetHorario)
	horarios.Post("/", handlers.CreateHorario)
	horarios.Put("/:id", handlers.UpdateHorario)
	horarios.Delete("/:id", middleware.AdminRateLimit(), handlers.DeleteHorario)

	// CRUD Recetas - Admin/Médico/Enfermera: completo, Paciente: solo lectura
	recetas := protected.Group("/recetas")
	recetas.Get("/", handlers.GetRecetas)
	recetas.Get("/:id", handlers.GetReceta)
	recetas.Post("/", handlers.CreateReceta)
	recetas.Put("/:id", handlers.UpdateReceta)
	recetas.Delete("/:id", middleware.AdminRateLimit(), handlers.DeleteReceta)

	// Ruta de salud (sin protección)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Hospital System API is running",
		})
	})

	// Ruta para verificar estado de rate limit (solo admin)
	api.Get("/rate-limit/status", middleware.JWTMiddleware(), middleware.RequireRole("admin"), func(c *fiber.Ctx) error {
		clientIP := c.Query("ip")
		if clientIP == "" {
			clientIP = c.IP()
		}

		status := middleware.GetRateLimitStatus(clientIP)
		return c.JSON(fiber.Map{
			"ip":     clientIP,
			"status": status,
		})
	})

	// Logs - Solo admin
	logs := api.Group("/logs", middleware.JWTMiddleware(), middleware.RequireRole("admin"))
	logs.Get("/", handlers.GetLogs)
	logs.Delete("/cleanup", handlers.DeleteOldLogs)

	// Endpoint para obtener permisos del usuario actual
	api.Get("/user/permissions", middleware.JWTMiddleware(), func(c *fiber.Ctx) error {
		userTipo := c.Locals("user_tipo").(string)
		permissions := middleware.GetUserPermissions(userTipo)
		
		return c.JSON(fiber.Map{
			"user_type": userTipo,
			"permissions": permissions,
		})
	})
}

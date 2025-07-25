package handlers

import (
	"hospital-system/config"
	"hospital-system/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetRecetas(c *fiber.Ctx) error {
	db := config.GetDB()

	
	userID := c.Locals("user_id").(int)
	userTipo := c.Locals("user_tipo").(string)

	var query string
	var args []interface{}

	
	if userTipo == "paciente" {
		query = `SELECT id_receta, fecha, medico_id, medicamento, dosis, consultorio_id, paciente_id 
				 FROM Recetas WHERE paciente_id = $1`
		args = []interface{}{userID}
	} else {
		
		query = `SELECT id_receta, fecha, medico_id, medicamento, dosis, consultorio_id, paciente_id 
				 FROM Recetas`
		args = []interface{}{}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener recetas"})
	}
	defer rows.Close()

	var recetas []models.Receta
	for rows.Next() {
		var receta models.Receta
		err := rows.Scan(&receta.IDReceta, &receta.Fecha,
			&receta.MedicoID, &receta.Medicamento,
			&receta.Dosis, &receta.ConsultorioID,
			&receta.PacienteID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error al escanear receta"})
		}
		recetas = append(recetas, receta)
	}

	return c.JSON(recetas)
}

func GetReceta(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	// Obtener información del usuario autenticado
	userID := c.Locals("user_id").(int)
	userTipo := c.Locals("user_tipo").(string)

	db := config.GetDB()
	var receta models.Receta

	var query string
	var args []interface{}

	// Si es paciente, solo puede ver sus propias recetas
	if userTipo == "paciente" {
		query = `SELECT id_receta, fecha, medico_id, medicamento, dosis, consultorio_id, paciente_id 
				 FROM Recetas WHERE id_receta = $1 AND paciente_id = $2`
		args = []interface{}{id, userID}
	} else {
		// Otros roles pueden ver cualquier receta
		query = `SELECT id_receta, fecha, medico_id, medicamento, dosis, consultorio_id, paciente_id 
				 FROM Recetas WHERE id_receta = $1`
		args = []interface{}{id}
	}

	err = db.QueryRow(query, args...).Scan(
		&receta.IDReceta, &receta.Fecha,
		&receta.MedicoID, &receta.Medicamento,
		&receta.Dosis, &receta.ConsultorioID,
		&receta.PacienteID)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Receta no encontrada"})
	}

	return c.JSON(receta)
}

func CreateReceta(c *fiber.Ctx) error {
	var receta models.Receta
	if err := c.BodyParser(&receta); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error al parsear datos: " + err.Error()})
	}

	db := config.GetDB()

	err := db.QueryRow(`
        INSERT INTO Recetas (fecha, medico_id, medicamento, dosis, consultorio_id, paciente_id) 
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_receta`,
		receta.Fecha, receta.MedicoID,
		receta.Medicamento, receta.Dosis,
		receta.ConsultorioID, receta.PacienteID).Scan(&receta.IDReceta)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al crear receta: " + err.Error()})
	}

	return c.Status(201).JSON(receta)
}

func UpdateReceta(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	var receta models.Receta
	if err := c.BodyParser(&receta); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error al parsear datos"})
	}

	db := config.GetDB()

	_, err = db.Exec(`
        UPDATE Recetas 
        SET fecha = $1, medico_id = $2, medicamento = $3, dosis = $4, consultorio_id = $5, paciente_id = $6 
        WHERE id_receta = $7`,
		receta.Fecha, receta.MedicoID,
		receta.Medicamento, receta.Dosis,
		receta.ConsultorioID, receta.PacienteID, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar receta"})
	}

	receta.IDReceta = id
	return c.JSON(receta)
}

func DeleteReceta(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	db := config.GetDB()

	_, err = db.Exec("DELETE FROM Recetas WHERE id_receta = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar receta"})
	}

	return c.JSON(fiber.Map{"message": "Receta eliminada exitosamente"})
}

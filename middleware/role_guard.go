package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// RoleBasedAccessControl define las reglas de acceso por tipo de usuario
type AccessRule struct {
	Resource string
	Actions  []string
}

// Definir las reglas de acceso por tipo de usuario
var rolePermissions = map[string][]AccessRule{
	"admin": {
		{Resource: "usuarios", Actions: []string{"read", "create", "update", "delete"}},
		{Resource: "consultorios", Actions: []string{"read", "create", "update", "delete"}},
		{Resource: "consultas", Actions: []string{"read", "create", "update", "delete"}},
		{Resource: "expedientes", Actions: []string{"read", "create", "update", "delete"}},
		{Resource: "horarios", Actions: []string{"read", "create", "update", "delete"}},
		{Resource: "recetas", Actions: []string{"read", "create", "update", "delete"}},
		{Resource: "logs", Actions: []string{"read", "delete"}},
	},
	"medico": {
		{Resource: "consultorios", Actions: []string{"read"}},
		{Resource: "consultas", Actions: []string{"read", "create", "update", "delete"}},
		{Resource: "expedientes", Actions: []string{"read", "create", "update"}},
		{Resource: "horarios", Actions: []string{"read", "create", "update", "delete"}},
		{Resource: "recetas", Actions: []string{"read", "create", "update", "delete"}},
	},
	"enfermera": {
		{Resource: "consultorios", Actions: []string{"read"}},
		{Resource: "consultas", Actions: []string{"read", "create", "update"}},
		{Resource: "expedientes", Actions: []string{"read", "create", "update"}},
		{Resource: "horarios", Actions: []string{"read"}},
		{Resource: "recetas", Actions: []string{"read", "create", "update"}},
	},
	"paciente": {
		{Resource: "consultas", Actions: []string{"read"}},
		{Resource: "expedientes", Actions: []string{"read"}},
		{Resource: "recetas", Actions: []string{"read"}},
		{Resource: "usuarios", Actions: []string{"read"}},      // Corregido: Actions en lugar de Action
		{Resource: "consultorios", Actions: []string{"read"}},  // Corregido: Actions en lugar de Action
	},
}

// RoleGuard middleware que verifica permisos basados en el tipo de usuario
func RoleGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Printf("[GUARD] 🛡️ Verificando acceso para: %s %s\n", c.Method(), c.Path())

		// Obtener el tipo de usuario del contexto
		userTipo, ok := c.Locals("user_tipo").(string)
		if !ok {
			fmt.Printf("[GUARD] ❌ No se pudo obtener el tipo de usuario\n")
			return c.Status(403).JSON(fiber.Map{
				"error": "Tipo de usuario no válido",
			})
		}

		// Extraer recurso y acción de la ruta
		resource, action := extractResourceAndAction(c)
		if resource == "" {
			// Si no es una ruta de recurso conocida, permitir acceso
			return c.Next()
		}

		fmt.Printf("[GUARD] 🔍 Usuario: %s, Recurso: %s, Acción: %s\n", userTipo, resource, action)

		// Verificar si el usuario tiene permisos para este recurso y acción
		if hasPermission(userTipo, resource, action) {
			fmt.Printf("[GUARD] ✅ Acceso permitido\n")
			return c.Next()
		}

		fmt.Printf("[GUARD] ❌ Acceso denegado\n")
		return c.Status(403).JSON(fiber.Map{
			"error": "No tienes permisos para acceder a este recurso",
			"details": fiber.Map{
				"user_type": userTipo,
				"resource":  resource,
				"action":    action,
			},
		})
	}
}

// extractResourceAndAction extrae el recurso y la acción de la ruta HTTP
func extractResourceAndAction(c *fiber.Ctx) (string, string) {
	path := c.Path()
	method := c.Method()

	// Extraer el recurso de la ruta (ej: /api/v1/usuarios -> usuarios)
	pathParts := strings.Split(strings.Trim(path, "/"), "/")
	var resource string

	// Buscar el recurso en la ruta
	for _, part := range pathParts {
		switch part {
		case "usuarios", "consultorios", "consultas", "expedientes", "horarios", "recetas", "logs":
			resource = part
			break
		}
	}

	if resource == "" {
		return "", ""
	}

	// Mapear método HTTP a acción
	var action string
	switch method {
	case "GET":
		action = "read"
	case "POST":
		action = "create"
	case "PUT", "PATCH":
		action = "update"
	case "DELETE":
		action = "delete"
	default:
		action = "read"
	}

	return resource, action
}

// hasPermission verifica si un tipo de usuario tiene permisos para un recurso y acción específicos
func hasPermission(userType, resource, action string) bool {
	rules, exists := rolePermissions[userType]
	if !exists {
		return false
	}

	for _, rule := range rules {
		if rule.Resource == resource {
			for _, allowedAction := range rule.Actions {
				if allowedAction == action {
					return true
				}
			}
			return false
		}
	}

	return false
}

// GetUserPermissions devuelve los permisos de un tipo de usuario (útil para el frontend)
func GetUserPermissions(userType string) []AccessRule {
	if rules, exists := rolePermissions[userType]; exists {
		return rules
	}
	return []AccessRule{}
}

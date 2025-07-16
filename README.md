# 🏥 Sistema de Gestión Hospitalaria

Sistema de gestión de citas médicas desarrollado con Go y Fiber (backend) y Angular (frontend), conectado a Supabase PostgreSQL.

## 🚀 Características

- **Gestión de Usuarios**: Pacientes, médicos, enfermeras y administradores
- **Autenticación MFA**: Sistema de autenticación de dos factores con TOTP
- **Gestión de Consultorios**: CRUD completo de consultorios médicos
- **Sistema de Consultas**: Programación y gestión de citas médicas
- **Expedientes Médicos**: Historial clínico de pacientes
- **Recetas Médicas**: Gestión de prescripciones
- **Horarios**: Control de disponibilidad médica
- **API REST**: Endpoints bien estructurados con validación
- **CORS**: Habilitado para desarrollo frontend
- **Logging**: Sistema de logs integrado
- **Rate Limiting**: Control de límites de peticiones
- **Middleware de Seguridad**: Validación y autenticación robusta

## 📁 Estructura del Proyecto
Backend-Base-de-datos-main/
├── .env                    # Variables de entorno
├── .gitignore             # Archivos ignorados por Git
├── .vscode/
│   └── settings.json      # Configuración de VS Code
├── README.md              # Documentación del proyecto
├── CHANGELOG.md           # Registro de cambios
├── go.mod                 # Dependencias de Go
├── go.sum                 # Checksums de dependencias
├── main.go                # Punto de entrada de la aplicación
├── package.json           # Configuración de Node.js (si aplica)
├── config/
│   └── database.go        # Configuración de base de datos
├── models/                # Modelos de datos
│   ├── usuario.go         # Modelo de usuario
│   ├── consultorio.go     # Modelo de consultorio
│   ├── consulta.go        # Modelo de consulta
│   ├── expediente.go      # Modelo de expediente
│   ├── receta.go          # Modelo de receta
│   ├── horario.go         # Modelo de horario
│   └── log.go             # Modelo de logs
├── handlers/              # Controladores de la API
│   ├── auth.go            # Handlers de autenticación
│   ├── mfa.go             # Handlers específicos de MFA
│   ├── usuarios.go        # Handlers de usuarios
│   ├── consultorios.go    # Handlers de consultorios
│   ├── consultas.go       # Handlers de consultas
│   ├── expedientes.go     # Handlers de expedientes
│   ├── recetas.go         # Handlers de recetas
│   ├── horarios.go        # Handlers de horarios
│   └── logs.go            # Handlers de logs
├── middleware/            # Middlewares
│   ├── auth.go            # Middleware de autenticación
│   ├── logger.go          # Middleware de logging
│   ├── ratelimit.go       # Middleware de rate limiting
│   ├── response_validator.go # Validador de respuestas
│   └── role_guard.go      # Middleware de roles
├── utils/                 # Utilidades
│   ├── jwt.go             # Utilidades JWT
│   ├── mfa.go             # Utilidades MFA
│   ├── password.go        # Utilidades de contraseñas
│   └── response_codes.go  # Códigos de respuesta
├── routes/
│   └── routes.go          # Configuración de rutas
└── schemas/
└── response_schemas.go # Esquemas de respuesta


## 🛠️ Tecnologías

### Backend
- **Go** 1.21+
- **Fiber** v2 - Framework web rápido y minimalista
- **Supabase PostgreSQL** - Base de datos en la nube
- **JWT** para autenticación
- **TOTP** para autenticación de dos factores
- **bcrypt** para hash de contraseñas
- **GORM** para ORM (si aplica)

## 📋 Requisitos

- **Go** 1.21 o superior
- **Git** para control de versiones
- Cuenta de **Supabase** configurada
- Editor de código (recomendado: VS Code)

## 🛠️ Instalación

### 1. Clonar el repositorio
```bash
git clone https://github.com/Lalo12-max/back-hospital.git
cd back-hospital
```

### 2. Configurar variables de entorno
Crea un archivo `.env` en la raíz del proyecto:

```env
# Configuración de la base de datos
DATABASE_URL=tu_url_de_supabase
DB_HOST=tu_host
DB_PORT=5432
DB_USER=tu_usuario
DB_PASSWORD=tu_contraseña
DB_NAME=tu_base_de_datos

# Configuración JWT
JWT_SECRET=tu_clave_secreta_jwt
JWT_EXPIRES_IN=24h

# Configuración del servidor
PORT=3000
ENVIRONMENT=development

# Configuración CORS
CORS_ORIGINS=http://localhost:4200,http://localhost:3000
```

### 3. Instalar dependencias
```bash
go mod download
```

### 4. Ejecutar la aplicación
```bash
# Modo desarrollo
go run main.go

# O compilar y ejecutar
go build -o hospital-system.exe
.\hospital-system.exe
```

El servidor estará disponible en: `http://localhost:3000`

## 🚀 Uso

### Endpoints principales

#### Autenticación
- `POST /api/v1/auth/register` - Registro de usuario
- `POST /api/v1/auth/login` - Inicio de sesión
- `POST /api/v1/auth/enable-mfa` - Habilitar MFA
- `POST /api/v1/auth/verify-mfa` - Verificar código MFA

#### Gestión de Usuarios
- `GET /api/v1/usuarios` - Listar usuarios
- `GET /api/v1/usuarios/:id` - Obtener usuario específico
- `POST /api/v1/usuarios` - Crear usuario
- `PUT /api/v1/usuarios/:id` - Actualizar usuario
- `DELETE /api/v1/usuarios/:id` - Eliminar usuario

#### Consultorios
- `GET /api/v1/consultorios` - Listar consultorios
- `POST /api/v1/consultorios` - Crear consultorio
- `PUT /api/v1/consultorios/:id` - Actualizar consultorio
- `DELETE /api/v1/consultorios/:id` - Eliminar consultorio

#### Consultas Médicas
- `GET /api/v1/consultas` - Listar consultas
- `POST /api/v1/consultas` - Programar consulta
- `PUT /api/v1/consultas/:id` - Actualizar consulta
- `DELETE /api/v1/consultas/:id` - Cancelar consulta

#### Expedientes
- `GET /api/v1/expedientes` - Listar expedientes
- `GET /api/v1/expedientes/:id` - Ver expediente específico
- `POST /api/v1/expedientes` - Crear expediente
- `PUT /api/v1/expedientes/:id` - Actualizar expediente

#### Recetas
- `GET /api/v1/recetas` - Listar recetas
- `POST /api/v1/recetas` - Crear receta
- `PUT /api/v1/recetas/:id` - Actualizar receta

#### Horarios
- `GET /api/v1/horarios` - Ver horarios
- `POST /api/v1/horarios` - Crear horario
- `PUT /api/v1/horarios/:id` - Actualizar horario

## 🔐 Autenticación y Seguridad

### Registro con MFA
1. Registra un nuevo usuario en `/api/v1/auth/register`
2. El sistema generará un código QR y clave secreta para MFA
3. Configura tu aplicación de autenticación (Google Authenticator, Authy, etc.)
4. Verifica el código MFA para completar el registro

### Inicio de Sesión
1. Envía credenciales a `/api/v1/auth/login`
2. Si MFA está habilitado, proporciona el código de 6 dígitos
3. Recibe el token JWT para autenticación en futuras peticiones

### Middleware de Seguridad
- **Rate Limiting**: Previene ataques de fuerza bruta
- **CORS**: Configurado para desarrollo seguro
- **JWT Validation**: Verificación de tokens en rutas protegidas
- **Role-based Access**: Control de acceso basado en roles

## 📊 Logging y Monitoreo

El sistema incluye logging completo de:
- Peticiones HTTP
- Errores de autenticación
- Operaciones de base de datos
- Acciones de usuarios

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 📞 Contacto

- **Desarrollador**: Lalo12-max
- **Repositorio**: [https://github.com/Lalo12-max/back-hospital](https://github.com/Lalo12-max/back-hospital)

## 🔄 Changelog

Ver `CHANGELOG.md` para un historial detallado de cambios y versiones.
# 🛍️ Bazar API + Frontend

Aplicación completa para gestión de un bazar, con backend en **Go (Gin + SQLite)** y frontend sencillo en **HTML/JavaScript**.

## 🌟 Características

- **Backend**:
  - API REST con CRUD para `productos`, `clientes` y `ventas`.
  - Swagger integrado (`/swagger/index.html`).
  - SQLite como base de datos.
- **Frontend**:
  - Interfaz sencilla para listar, agregar productos y crear ventas.
  - Conexión directa a la API.

## 🚀 Instalación  

1. Clona el repositorio:  

   ```bash  
   git clone https://github.com/tu-usuario/bazar-go.git  
   ```

2. Inicia el backend en local:

   ```bash
   cd backend
   go mod download
   make run
   ```

3. Abre el frontend en local:

   Navega a [http://localhost:8080](http://localhost:8080) en tu navegador.

---

## 🌐 Producción

El servicio está desplegado en **Render** y accesible en la siguiente URL:

- **Frontend**: [https://bazar-go.onrender.com](https://bazar-go.onrender.com)
- **Swagger**: [https://bazar-go.onrender.com/swagger/index.html](https://bazar-go.onrender.com/swagger/index.html)

---

## 📂 Estructura del Proyecto

```
bazar-go/
├── backend/          # Código Go (API)
├── frontend/         # Interfaz web (HTML/JS)
├── Makefile          # Comandos útiles
└── README.md
```

---

## 🔌 Endpoints clave (API)

| Método | Ruta               | Descripción              |
|--------|--------------------|--------------------------|
| GET    | /api/v1/products   | Listar todos los productos |
| POST   | /api/v1/products   | Crear un producto        |
| PUT    | /api/v1/products/:id   | Modificar un producto |
| DELETE | /api/v1/products/:id | Eliminar un producto    |

---

## 🛠️ Contribuir

Si deseas contribuir al proyecto, sigue estos pasos:

1. **Pruebas locales**:
   - Realiza tus pruebas en un entorno local utilizando `localhost`. Esto asegura que los cambios no afecten el servicio en producción.
   - Usa la URL `http://localhost:8080` para acceder al backend y frontend en local.

2. **Pruebas en producción**:
   - **No realices pruebas directamente en el entorno de producción** (`https://bazar-go.onrender.com`), ya que esto podría afectar a los usuarios finales.

3. **Enviar cambios**:
   - Crea un fork del repositorio.
   - Realiza tus cambios en una nueva rama.
   - Envía un pull request con una descripción detallada de los cambios.

---

## 🚀 Despliegue

### Local:
Ejecuta el siguiente comando en la carpeta `backend` para iniciar el servidor:

```bash
make run
```

### Producción:
El despliegue en producción utiliza Docker y Render. Consulta la [wiki](https://github.com/tu-usuario/bazar-go/wiki) para más detalles.

---

## 📝 Notas adicionales

- **Swagger**:
  - La documentación Swagger está disponible en:
    - Local: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
    - Producción: [https://bazar-go.onrender.com/swagger/index.html](https://bazar-go.onrender.com/swagger/index.html)

- **Base de datos**:
  - En local, se utiliza SQLite como base de datos. Asegúrate de que el archivo de base de datos esté correctamente configurado en el entorno local.


-- Crear la tabla de Usuarios
CREATE TABLE usuarios
(
    id                   SERIAL PRIMARY KEY,
    nombre               VARCHAR(100)                                                    NOT NULL,
    email                VARCHAR(100) UNIQUE                                             NOT NULL,
    password_hash        VARCHAR(255)                                                    NOT NULL,       -- Para almacenar el hash de la contraseña
    tipo_usuario         VARCHAR(50) CHECK (tipo_usuario IN ('lector', 'bibliotecario')) NOT NULL,
    estado               VARCHAR(50) CHECK (estado IN ('activo', 'inactivo')) DEFAULT 'activo',          -- Estado de la cuenta
    fecha_registro       TIMESTAMP                                            DEFAULT CURRENT_TIMESTAMP, -- Fecha de registro
    ultimo_inicio_sesion TIMESTAMP                                            DEFAULT CURRENT_TIMESTAMP  -- Última fecha de inicio de sesión
);

-- Crear la tabla de Autores
CREATE TABLE autores
(
    id     SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

-- Crear la tabla de Categorías
CREATE TABLE categorias
(
    id     SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

-- Crear la tabla de Libros
CREATE TABLE libros
(
    id                SERIAL PRIMARY KEY,
    titulo            VARCHAR(150) NOT NULL,
    autor_id          INT REFERENCES autores (id) ON DELETE CASCADE,
    categoria_id      INT REFERENCES categorias (id),
    fecha_publicacion DATE,
    disponible        BOOLEAN DEFAULT TRUE
);

-- Crear la tabla de Préstamos
CREATE TABLE prestamos
(
    id               SERIAL PRIMARY KEY,
    libro_id         INT REFERENCES libros (id) ON DELETE CASCADE,
    lector_id        INT REFERENCES usuarios (id) ON DELETE CASCADE,
    fecha_prestamo   DATE                                                    DEFAULT CURRENT_DATE,
    fecha_devolucion DATE,
    fecha_devuelto   DATE,
    estado           VARCHAR(50) CHECK (estado IN ('pendiente', 'devuelto')) DEFAULT 'pendiente'
);

CREATE TABLE usuarios
(
    id                   SERIAL PRIMARY KEY,
    nombre               VARCHAR(100) NOT NULL,
    rut                  VARCHAR(12) UNIQUE NOT NULL,       -- RUT completo, incluyendo puntos y guion
    email                VARCHAR(100) UNIQUE NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    password_hash        VARCHAR(255) NOT NULL,             -- Para almacenar el hash de la contraseña
    tipo_usuario         VARCHAR(50) CHECK (tipo_usuario IN ('lector', 'bibliotecario')) NOT NULL,
    estado               VARCHAR(50) CHECK (estado IN ('activo', 'inactivo')) DEFAULT 'activo',
    fecha_registro       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ultimo_inicio_sesion TIMESTAMP
);

CREATE TABLE autores
(
    id     SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    UNIQUE(nombre) -- Para asegurar que no se dupliquen autores
);

CREATE TABLE categorias
(
    id     SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    UNIQUE(nombre) -- Para asegurar que no se dupliquen categorías
);

CREATE TABLE libros
(
    id                SERIAL PRIMARY KEY,
    titulo            VARCHAR(150) NOT NULL,
    autor_id          INT REFERENCES autores (id) ON DELETE CASCADE,
    categoria_id      INT REFERENCES categorias (id),
    fecha_publicacion DATE CHECK (fecha_publicacion <= CURRENT_DATE), -- No permitir fechas futuras
    disponible        BOOLEAN DEFAULT TRUE,
    UNIQUE(titulo, autor_id) -- Evitar duplicación de títulos por autor
);

CREATE TABLE prestamos
(
    id               SERIAL PRIMARY KEY,
    libro_id         INT REFERENCES libros (id) ON DELETE CASCADE,
    usuario_id       INT REFERENCES usuarios (id) ON DELETE CASCADE,
    fecha_prestamo   DATE DEFAULT CURRENT_DATE,
    fecha_devolucion DATE CHECK (fecha_devolucion >= fecha_prestamo), -- Asegurar que la fecha de devolución no sea antes de la de préstamo
    fecha_devuelto   DATE,
    estado           VARCHAR(50) CHECK (estado IN ('pendiente', 'devuelto')) DEFAULT 'pendiente',
    CONSTRAINT fk_libro_disponible CHECK ((estado = 'pendiente' AND fecha_devuelto IS NULL) OR estado = 'devuelto') -- Controlar que el libro solo esté en préstamo cuando esté pendiente
);

CREATE TABLE parqueaderos (
    id TEXT PRIMARY KEY,
    nombre TEXT NOT NULL,
    ubicacion TEXT NOT NULL,
    capacidad INTEGER NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE espacios (
    id TEXT PRIMARY KEY,
    parqueadero_id TEXT NOT NULL,
    numero TEXT NOT NULL,
    disponible BOOLEAN NOT NULL DEFAULT TRUE,
    reservado_hasta DATETIME,

    FOREIGN KEY (parqueadero_id)
        REFERENCES parqueaderos(id)
        ON DELETE CASCADE
);

CREATE TABLE ocupaciones (
    id TEXT PRIMARY KEY,
    espacio_id TEXT NOT NULL,
    placa TEXT NOT NULL,
    entrada DATETIME NOT NULL,
    salida DATETIME,
    activa BOOLEAN NOT NULL DEFAULT TRUE,

    FOREIGN KEY (espacio_id)
        REFERENCES espacios(id)
        ON DELETE CASCADE
);

-- ==========================================
-- MÓDULO: TRANSPORTE INTERNO - CRISTINA
-- ==========================================

CREATE TABLE rutas (
    id     TEXT PRIMARY KEY,
    nombre      TEXT NOT NULL,
    descripcion TEXT NOT NULL
);

CREATE TABLE paradas (
    id_parada    TEXT PRIMARY KEY,
    nombre       TEXT NOT NULL,
    latitud      REAL NOT NULL,
    longitu      REAL NOT NULL,
    ruta_id     TEXT NOT NULL
);

CREATE TABLE carritos (
    id     TEXT PRIMARY KEY,
    nombre_carrito TEXT NOT NULL,
    capacidad      INTEGER NOT NULL,
    estado         TEXT NOT NULL,
    ruta_id        TEXT NOT NULL
);

CREATE TABLE locaciones (
    id TEXT PRIMARY KEY,
    latitud     REAL NOT NULL,
    longitud    REAL NOT NULL,
    time_stamp  TEXT NOT NULL,
    carrito_id  TEXT NOT NULL
);

CREATE TABLE solicitudes (
    id   TEXT PRIMARY KEY,
    cedula_usuario TEXT NOT NULL,
    cant_personas  INTEGER NOT NULL,
    punto_destino  TEXT NOT NULL,
    estado         TEXT NOT NULL,
    carrito_id     TEXT
);
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
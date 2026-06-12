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

-- Tabla base necesaria para la relación con Solicitud (Módulo de Acceso)
CREATE TABLE IF NOT EXISTS Usuario (
    Cedula int(10) NOT NULL PRIMARY KEY,
    Nombre_usuario char(50) NOT NULL,
    Contrasena char(50) NOT NULL,
    Email char(50) NOT NULL,
    Rol char(50) NOT NULL
);

-- 1. Tabla: Ruta
CREATE TABLE IF NOT EXISTS Ruta (
    ID_ruta INTEGER PRIMARY KEY AUTOINCREMENT,
    Nombre string NOT NULL,
    Descripcion string NOT NULL
);

-- 2. Tabla: Paradas
CREATE TABLE IF NOT EXISTS Paradas (
    ID_parada INTEGER PRIMARY KEY AUTOINCREMENT,
    FK2_ID_ruta INTEGER NOT NULL,
    Nombre string NOT NULL,
    latitud float NOT NULL,
    longitu float NOT NULL, -- Mantenemos el nombre 'longitu' tal como está en tu diagrama
    FOREIGN KEY (FK2_ID_ruta) REFERENCES Ruta(ID_ruta) ON DELETE CASCADE
);

-- 3. Tabla: Carrito
CREATE TABLE IF NOT EXISTS Carrito (
    ID_carrito INTEGER PRIMARY KEY AUTOINCREMENT,
    FK1_ID_ruta INTEGER NOT NULL,
    Nombre_carrito string NOT NULL,
    Capacidad int NOT NULL,
    Estado string NOT NULL,
    FOREIGN KEY (FK1_ID_ruta) REFERENCES Ruta(ID_ruta)
);

-- 4. Tabla: Locacion (Historial de ubicación del carrito)
CREATE TABLE IF NOT EXISTS Locacion (
    ID_locacion INTEGER PRIMARY KEY AUTOINCREMENT,
    FK1_ID_carrito INTEGER NOT NULL,
    latitud float NOT NULL,
    longitud float NOT NULL,
    time_stamp datetime NOT NULL,
    FOREIGN KEY (FK1_ID_carrito) REFERENCES Carrito(ID_carrito) ON DELETE CASCADE
);

-- 5. Tabla: Solicitud
CREATE TABLE IF NOT EXISTS Solicitud (
    ID_solicitud INTEGER PRIMARY KEY AUTOINCREMENT,
    FK1_Cedula_usuario int(10) NOT NULL,
    FK2_ID_carrito int NOT NULL,
    Cant_personas int NOT NULL,
    Punto_destino string NOT NULL,
    Estado string NOT NULL,
    FOREIGN KEY (FK1_Cedula_usuario) REFERENCES Usuario(Cedula),
    FOREIGN KEY (FK2_ID_carrito) REFERENCES Carrito(ID_carrito)
);
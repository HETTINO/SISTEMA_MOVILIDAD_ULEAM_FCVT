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

========================================== 
-- MÓDULO: ACCESO ENTRADA Y SALIDA - SHIRLEY -- 
==========================================

-- Crear tabla de Usuarios
CREATE TABLE IF NOT EXISTS usuario (
    cedula VARCHAR(10) PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    contrasena VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL,
    rol VARCHAR(50) NOT NULL
);

-- Crear tabla de Vehículos
CREATE TABLE IF NOT EXISTS vehiculo (
    placa VARCHAR(8) PRIMARY KEY,
    id_usuario VARCHAR(10) NOT NULL,
    tipo_vehiculo VARCHAR(50) NOT NULL,
    marca VARCHAR(50) NOT NULL,
    modelo VARCHAR(50) NOT NULL,
    color VARCHAR(30) NOT NULL,
    año INT NOT NULL,
    FOREIGN KEY (id_usuario) REFERENCES usuario(cedula) ON DELETE CASCADE
);

-- Crear tabla de Puntos de Acceso
CREATE TABLE IF NOT EXISTS punto_de_acceso (
    id_punto_acceso INTEGER PRIMARY KEY AUTOINCREMENT,
    frecuencia VARCHAR(50) NOT NULL,
    ubicacion VARCHAR(100) NOT NULL
);

-- Crear tabla de Accesos (Entradas y Salidas)
CREATE TABLE IF NOT EXISTS acceso (
    id_acceso INTEGER PRIMARY KEY AUTOINCREMENT,
    placa_vehiculo VARCHAR(8) NOT NULL,
    id_punto_acceso INT NOT NULL,
    tiempo_entrada DATETIME NOT NULL,
    tiempo_salida DATETIME, -- Puede ser NULL si el vehículo aún no ha salido
    estado VARCHAR(50) NOT NULL,
    observaciones TEXT,     -- Puede ser opcional (NULL) si no hay novedades
    FOREIGN KEY (placa_vehiculo) REFERENCES vehiculo(placa) ON DELETE CASCADE,
    FOREIGN KEY (id_punto_acceso) REFERENCES punto_de_acceso(id_punto_acceso) ON DELETE RESTRICT
);
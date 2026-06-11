-- name: ListarParqueaderos :many
SELECT * FROM parqueaderos;

-- name: ObtenerParqueadero :one
SELECT * FROM parqueaderos
WHERE id = ?;

-- name: CrearParqueadero :exec
INSERT INTO parqueaderos (
    id,
    nombre,
    ubicacion,
    capacidad,
    activo
)
VALUES (?, ?, ?, ?, ?);

-- name: EliminarParqueadero :exec
DELETE FROM parqueaderos
WHERE id = ?;

-- Actualizar parqueadero
UPDATE parqueaderos
SET nombre = ?, ubicacion = ?, capacidad = ?, activo = ?
WHERE id = ?
RETURNING id, nombre, ubicacion, capacidad, activo;

-- name: ListarEspacios :many
SELECT * FROM espacios;

-- name: ObtenerEspacio :one
SELECT *
FROM espacios
WHERE id = ?;

-- name: CrearEspacio :exec
INSERT INTO espacios (
    id,
    parqueadero_id,
    numero,
    disponible,
    reservado_hasta
)
VALUES (?, ?, ?, ?, ?);

-- name: ActualizarEspacio :one
UPDATE espacios
SET
    numero = ?,
    disponible = ?
WHERE id = ?
RETURNING *;

-- name: EliminarEspacio :exec
DELETE FROM espacios
WHERE id = ?;

-- name: ListarOcupaciones :many
SELECT *
FROM ocupaciones;

-- name: ObtenerOcupacion :one
SELECT *
FROM ocupaciones
WHERE id = ?;

-- name: CrearOcupacion :exec
INSERT INTO ocupaciones (
    id,
    espacio_id,
    placa,
    entrada,
    salida,
    activa
)
VALUES (?, ?, ?, ?, ?, ?);

-- name: ActualizarOcupacion :one
UPDATE ocupaciones
SET
    salida = ?,
    activa = ?
WHERE id = ?
RETURNING *;

-- name: EliminarOcupacion :exec
DELETE FROM ocupaciones
WHERE id = ?;



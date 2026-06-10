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

-- name: ListarEspacios :many
SELECT * FROM espacios;

UPDATE parqueaderos
SET nombre = ?, ubicacion = ?, capacidad = ?, activo = ?
WHERE id = ?
RETURNING id, nombre, ubicacion, capacidad, activo;
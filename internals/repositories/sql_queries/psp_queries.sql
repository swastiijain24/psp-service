-- name: GetPspRegistration :one
SELECT *
FROM psp_registrations 
WHERE psp_id = $1;

-- name: IsActive :one
SELECT EXISTS (
    SELECT 1 FROM psp_registrations
    WHERE psp_id = $1
      AND is_active = TRUE
);
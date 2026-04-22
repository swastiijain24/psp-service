-- name: GetVpaMapping :one
SELECT vpa_id, account_id, bank_code, is_active
FROM vpa_map
WHERE vpa_id = $1 AND is_active = TRUE
LIMIT 1;

-- name: CreateVpaMapping :one
INSERT INTO vpa_map (
    vpa_id, account_id, bank_code
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: DeactivateVpa :exec
UPDATE vpa_map
SET is_active = false, updated_at = CURRENT_TIMESTAMP
WHERE vpa_id = $1;

-- name: CheckVpaExists :one
SELECT EXISTS (
    SELECT 1 
    FROM vpa_map 
    WHERE vpa_id = $1 
      AND is_active = true
);
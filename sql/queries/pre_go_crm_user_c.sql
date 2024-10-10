-- name: GetUserByEmailSQLC :one
SELECT user_id, user_email FROM pre_go_crm_user_c WHERE user_email = ? LIMIT 1;

-- name: UpdateUserStatusByUserId :exec
UPDATE `pre_go_crm_user_c`
SET user_status = $2,
    user_updated_at = $3
WHERE user_id = $1;

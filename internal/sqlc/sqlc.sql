-- name: CreateUser :one
INSERT INTO "user" (id, email, first_name, last_name, password_hash)
VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING RETURNING *;

-- name: GetUser :one
SELECT *
FROM "user"
WHERE email = $1;

-- name: CreateWorkspace :one
INSERT INTO "workspace" (id, name, region)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING *;

-- name: GetWorkspaces :many
SELECT workspace.*
FROM "workspace"
         JOIN "workspace_member" ON workspace.id = workspace_member.workspace_id
WHERE workspace_member.user_id = $1
ORDER BY workspace.id DESC LIMIT $2;

-- name: CreateWorkspaceMember :exec
INSERT INTO "workspace_member" (workspace_id, user_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: GetWorkspaceMembers :many
SELECT "user".*
FROM "user"
         JOIN "workspace_member" ON "user".id = workspace_member.user_id
WHERE workspace_member.workspace_id = $1
  AND (sqlc.arg(cursor)::int8 = 0 OR "user".id < sqlc.arg(cursor)::int8)
ORDER BY "user".id DESC LIMIT $2;

-- name: CreateUser :one
INSERT INTO "user" (id, slack_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING *;

-- name: GetUser :one
SELECT *
FROM "user"
WHERE slack_id = $1;

-- name: CreateWorkspace :one
INSERT INTO "workspace" (id, slack_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING *;

-- name: GetWorkspace :one
SELECT *
FROM "workspace"
WHERE slack_id = $1;

-- name: CreateWorkspaceMember :exec
INSERT INTO "workspace_member" (workspace_id, user_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: GetWorkspaceMembers :many
SELECT "user".*
FROM "user"
         JOIN "workspace_member" ON "user".id = workspace_member.user_id
WHERE workspace_member.workspace_id = $1
ORDER BY "user".id DESC LIMIT $2;

-- name: GetWorkspaceMembersC :many
SELECT "user".*
FROM "user"
         JOIN "workspace_member" ON "user".id = workspace_member.user_id
WHERE workspace_member.workspace_id = sqlc.arg(workspace_id)
  AND "user".id < sqlc.arg(cursor)
ORDER BY "user".id DESC LIMIT $1;

-- name: CreateCheckIn :one
INSERT INTO "check_in" (id, workspace_id, slack_channel, question, schedule)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: CreateCheckInAnswer :one
INSERT INTO "check_in_answer" (id, workspace_id, check_in_id, user_id, answer)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

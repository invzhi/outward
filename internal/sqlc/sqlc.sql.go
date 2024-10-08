// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sqlc.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (id, email, first_name, last_name, password_hash)
VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING RETURNING id, created_at, email, first_name, last_name, password_hash
`

type CreateUserParams struct {
	ID           int64
	Email        string
	FirstName    string
	LastName     string
	PasswordHash pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.PasswordHash,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.PasswordHash,
	)
	return i, err
}

const createWorkspace = `-- name: CreateWorkspace :one
INSERT INTO "workspace" (id, name, region)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id, created_at, name, region
`

type CreateWorkspaceParams struct {
	ID     int64
	Name   string
	Region int32
}

func (q *Queries) CreateWorkspace(ctx context.Context, arg CreateWorkspaceParams) (Workspace, error) {
	row := q.db.QueryRow(ctx, createWorkspace, arg.ID, arg.Name, arg.Region)
	var i Workspace
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Region,
	)
	return i, err
}

const createWorkspaceMember = `-- name: CreateWorkspaceMember :exec
INSERT INTO "workspace_member" (workspace_id, user_id, role)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING
`

type CreateWorkspaceMemberParams struct {
	WorkspaceID int64
	UserID      int64
	Role        int32
}

func (q *Queries) CreateWorkspaceMember(ctx context.Context, arg CreateWorkspaceMemberParams) error {
	_, err := q.db.Exec(ctx, createWorkspaceMember, arg.WorkspaceID, arg.UserID, arg.Role)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, email, first_name, last_name, password_hash
FROM "user"
WHERE email = $1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.PasswordHash,
	)
	return i, err
}

const getWorkspaceMembers = `-- name: GetWorkspaceMembers :many
SELECT "user".id, "user".created_at, "user".email, "user".first_name, "user".last_name, "user".password_hash
FROM "user"
         JOIN "workspace_member" ON "user".id = workspace_member.user_id
WHERE workspace_member.workspace_id = $1
  AND ($3::int8 = 0 OR "user".id < $3::int8)
ORDER BY "user".id DESC LIMIT $2
`

type GetWorkspaceMembersParams struct {
	WorkspaceID int64
	Limit       int32
	Cursor      int64
}

func (q *Queries) GetWorkspaceMembers(ctx context.Context, arg GetWorkspaceMembersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getWorkspaceMembers, arg.WorkspaceID, arg.Limit, arg.Cursor)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.PasswordHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWorkspaces = `-- name: GetWorkspaces :many
SELECT workspace.id, workspace.created_at, workspace.name, workspace.region
FROM "workspace"
         JOIN "workspace_member" ON workspace.id = workspace_member.workspace_id
WHERE workspace_member.user_id = $1
ORDER BY workspace.id DESC
`

func (q *Queries) GetWorkspaces(ctx context.Context, userID int64) ([]Workspace, error) {
	rows, err := q.db.Query(ctx, getWorkspaces, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Workspace
	for rows.Next() {
		var i Workspace
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Region,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

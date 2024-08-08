// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sqlc.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (id, email, first_name, last_name)
VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING RETURNING id, created_at, email, first_name, last_name, password
`

type CreateUserParams struct {
	ID        ID
	Email     string
	FirstName string
	LastName  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
	)
	return i, err
}

const createWorkspace = `-- name: CreateWorkspace :one
INSERT INTO "workspace" (id, name, region)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id, created_at, name, region
`

type CreateWorkspaceParams struct {
	ID     ID
	Name   string
	Region string
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
INSERT INTO "workspace_member" (workspace_id, user_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING
`

type CreateWorkspaceMemberParams struct {
	WorkspaceID ID
	UserID      ID
}

func (q *Queries) CreateWorkspaceMember(ctx context.Context, arg CreateWorkspaceMemberParams) error {
	_, err := q.db.Exec(ctx, createWorkspaceMember, arg.WorkspaceID, arg.UserID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, email, first_name, last_name, password
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
		&i.Password,
	)
	return i, err
}

const getWorkspaceMembers = `-- name: GetWorkspaceMembers :many
SELECT "user".id, "user".created_at, "user".email, "user".first_name, "user".last_name, "user".password
FROM "user"
         JOIN "workspace_member" ON "user".id = workspace_member.user_id
WHERE workspace_member.workspace_id = $1
ORDER BY "user".id DESC LIMIT $2
`

type GetWorkspaceMembersParams struct {
	WorkspaceID ID
	Limit       int32
}

func (q *Queries) GetWorkspaceMembers(ctx context.Context, arg GetWorkspaceMembersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getWorkspaceMembers, arg.WorkspaceID, arg.Limit)
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
			&i.Password,
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

const getWorkspaceMembersC = `-- name: GetWorkspaceMembersC :many
SELECT "user".id, "user".created_at, "user".email, "user".first_name, "user".last_name, "user".password
FROM "user"
         JOIN "workspace_member" ON "user".id = workspace_member.user_id
WHERE workspace_member.workspace_id = $2
  AND "user".id < $3
ORDER BY "user".id DESC LIMIT $1
`

type GetWorkspaceMembersCParams struct {
	Limit       int32
	WorkspaceID ID
	Cursor      ID
}

func (q *Queries) GetWorkspaceMembersC(ctx context.Context, arg GetWorkspaceMembersCParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getWorkspaceMembersC, arg.Limit, arg.WorkspaceID, arg.Cursor)
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
			&i.Password,
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
ORDER BY workspace.id DESC LIMIT $2
`

type GetWorkspacesParams struct {
	UserID ID
	Limit  int32
}

func (q *Queries) GetWorkspaces(ctx context.Context, arg GetWorkspacesParams) ([]Workspace, error) {
	rows, err := q.db.Query(ctx, getWorkspaces, arg.UserID, arg.Limit)
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

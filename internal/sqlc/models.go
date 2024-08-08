// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type EmailContent struct {
	ID        ID
	ProjectID ID
	CreatedAt time.Time
	Payload   string
}

type EmailLayout struct {
	ID        ID
	ProjectID ID
	CreatedAt time.Time
	Payload   string
}

type EmailTemplate struct {
	ID        ID
	ProjectID ID
	CreatedAt time.Time
	Payload   string
}

type EmailTemplateHistory struct {
	ID              ID
	EmailTemplateID ID
	ProjectID       ID
	CreatedAt       time.Time
	Payload         string
}

type Environment struct {
	ID        ID
	ProjectID ID
	CreatedAt time.Time
	Name      string
}

type Project struct {
	ID          ID
	WorkspaceID ID
	CreatedAt   time.Time
	Name        string
}

type ProjectMember struct {
	ProjectID ID
	UserID    ID
	Role      string
}

type User struct {
	ID        ID
	CreatedAt time.Time
	Email     string
	FirstName string
	LastName  string
	Password  pgtype.Text
}

type Workspace struct {
	ID        ID
	CreatedAt time.Time
	Name      string
	Region    string
}

type WorkspaceMember struct {
	WorkspaceID ID
	UserID      ID
	Role        string
}

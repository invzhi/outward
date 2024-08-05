CREATE TABLE "user"
(
    id            BIGINT    NOT NULL PRIMARY KEY,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email         TEXT      NOT NULL,
    first_name    TEXT      NOT NULL,
    last_name     TEXT      NOT NULL,
    password_hash TEXT      NULL
);

CREATE UNIQUE INDEX "user_idx_email" ON "user" (email);

CREATE TABLE "workspace"
(
    id         BIGINT    NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       TEXT      NOT NULL,
    region     INTEGER   NOT NULL
);

CREATE TABLE "workspace_member"
(
    workspace_id BIGINT  NOT NULL,
    user_id      BIGINT  NOT NULL,
    role         INTEGER NOT NULL,
    PRIMARY KEY (workspace_id, user_id)
);

CREATE TABLE "project"
(
    id           BIGINT    NOT NULL PRIMARY KEY,
    workspace_id BIGINT    NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name         TEXT      NOT NULL
);

CREATE TABLE "project_member"
(
    project_id BIGINT NOT NULL,
    user_id    BIGINT NOT NULL,
    role       TEXT   NOT NULL,
    PRIMARY KEY (project_id, user_id)
);

CREATE TABLE "environment"
(
    id         BIGINT    NOT NULL PRIMARY KEY,
    project_id BIGINT    NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       TEXT      NOT NULL
);

CREATE TABLE "email_layout"
(
    id         BIGINT    NOT NULL PRIMARY KEY,
    project_id BIGINT    NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    payload    TEXT      NOT NULL
);

CREATE TABLE "email_content"
(
    id         BIGINT    NOT NULL PRIMARY KEY,
    project_id BIGINT    NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    payload    TEXT      NOT NULL
);

CREATE TABLE "email_template"
(
    id         BIGINT    NOT NULL PRIMARY KEY,
    project_id BIGINT    NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    payload    TEXT      NOT NULL
);

CREATE TABLE "email_template_history"
(
    id                BIGINT    NOT NULL PRIMARY KEY,
    email_template_id BIGINT    NOT NULL,
    project_id        BIGINT    NOT NULL,
    created_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    payload           TEXT      NOT NULL
);

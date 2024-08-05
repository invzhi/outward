CREATE TABLE "user"
(
    id         BIGINT    NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email      TEXT      NOT NULL,
    password   TEXT      NOT NULL
);

CREATE UNIQUE INDEX "user_idx_slack_id" ON "user" (slack_id);

CREATE TABLE "workspace"
(
    id         BIGINT    NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    slack_id   TEXT      NOT NULL
);

CREATE UNIQUE INDEX "workspace_idx_slack_id" ON "workspace" (slack_id);

CREATE TABLE "workspace_member"
(
    workspace_id BIGINT NOT NULL,
    user_id      BIGINT NOT NULL,
    PRIMARY KEY (workspace_id, user_id)
);

CREATE TABLE "email_layout"
(
    id            BIGINT    NOT NULL PRIMARY KEY,
    workspace_id  BIGINT    NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    slack_channel TEXT      NOT NULL,
    question      TEXT      NOT NULL,
    schedule      JSONB     NOT NULL
);

CREATE TABLE "email_template"
(
    id           BIGINT    NOT NULL PRIMARY KEY,
    workspace_id BIGINT    NOT NULL,
    check_in_id  BIGINT    NOT NULL,
    user_id      BIGINT    NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    answer       TEXT      NOT NULL
);

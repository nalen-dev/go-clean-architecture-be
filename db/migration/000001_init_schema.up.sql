CREATE TABLE "users"(
    "id"            BIGSERIAL   PRIMARY KEY,
    "name"          VARCHAR     NOT NULL,
    "email"         VARCHAR     NOT NULL,
    "password"      VARCHAR     NOT NULL,
    "created_at"    timestamp with time zone  NOT NULL DEFAULT (now())
) 
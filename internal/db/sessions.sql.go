// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sessions.sql

package db

import (
	"context"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
    id,
    title,
    message_count,
    prompt_tokens,
    completion_tokens,
    cost,
    updated_at,
    created_at
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    strftime('%s', 'now'),
    strftime('%s', 'now')
) RETURNING id, title, message_count, prompt_tokens, completion_tokens, cost, updated_at, created_at
`

type CreateSessionParams struct {
	ID               string  `json:"id"`
	Title            string  `json:"title"`
	MessageCount     int64   `json:"message_count"`
	PromptTokens     int64   `json:"prompt_tokens"`
	CompletionTokens int64   `json:"completion_tokens"`
	Cost             float64 `json:"cost"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.queryRow(ctx, q.createSessionStmt, createSession,
		arg.ID,
		arg.Title,
		arg.MessageCount,
		arg.PromptTokens,
		arg.CompletionTokens,
		arg.Cost,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.MessageCount,
		&i.PromptTokens,
		&i.CompletionTokens,
		&i.Cost,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = ?
`

func (q *Queries) DeleteSession(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteSessionStmt, deleteSession, id)
	return err
}

const getSessionByID = `-- name: GetSessionByID :one
SELECT id, title, message_count, prompt_tokens, completion_tokens, cost, updated_at, created_at
FROM sessions
WHERE id = ? LIMIT 1
`

func (q *Queries) GetSessionByID(ctx context.Context, id string) (Session, error) {
	row := q.queryRow(ctx, q.getSessionByIDStmt, getSessionByID, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.MessageCount,
		&i.PromptTokens,
		&i.CompletionTokens,
		&i.Cost,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listSessions = `-- name: ListSessions :many
SELECT id, title, message_count, prompt_tokens, completion_tokens, cost, updated_at, created_at
FROM sessions
ORDER BY created_at DESC
`

func (q *Queries) ListSessions(ctx context.Context) ([]Session, error) {
	rows, err := q.query(ctx, q.listSessionsStmt, listSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.MessageCount,
			&i.PromptTokens,
			&i.CompletionTokens,
			&i.Cost,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSession = `-- name: UpdateSession :one
UPDATE sessions
SET
    title = ?,
    prompt_tokens = ?,
    completion_tokens = ?,
    cost = ?
WHERE id = ?
RETURNING id, title, message_count, prompt_tokens, completion_tokens, cost, updated_at, created_at
`

type UpdateSessionParams struct {
	Title            string  `json:"title"`
	PromptTokens     int64   `json:"prompt_tokens"`
	CompletionTokens int64   `json:"completion_tokens"`
	Cost             float64 `json:"cost"`
	ID               string  `json:"id"`
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error) {
	row := q.queryRow(ctx, q.updateSessionStmt, updateSession,
		arg.Title,
		arg.PromptTokens,
		arg.CompletionTokens,
		arg.Cost,
		arg.ID,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.MessageCount,
		&i.PromptTokens,
		&i.CompletionTokens,
		&i.Cost,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

-- +goose Up
ALTER TABLE psp_registrations
ALTER COLUMN psp_id TYPE UUID USING psp_id::uuid;
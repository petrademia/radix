# RADIX Stage 1 — In-memory store

Status: **frozen**.

## Commands (minimum)

`PING`, `GET`, `SET`, `DEL`, `EXISTS`, `EXPIRE`, `TTL`, `KEYS` (optional slow).

## Required

- String keys/values
- Passive + passive TTL expiry
- Memory bound + LRU (or documented eviction) when over maxmemory

## Forbidden

- Embedding real Redis as the implementation

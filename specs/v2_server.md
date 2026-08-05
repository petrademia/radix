# RADIX Stage 2 — TCP server

Status: **frozen**.

## Purpose

Accept Redis clients over TCP; pipeline multiple commands per connection.

## Required

- `radixd -addr :6379` (port configurable)
- Concurrent connections (goroutines OK)
- Compare a subset of commands against `redis-benchmark` or `redis-cli` smoke scripts

## Guardrail

`make smoke` runs built-in client tests against a started server.

# RADIX

**RESP-Addressable Data In-memory eXchange**

In-memory Redis-compatible cache: RESP protocol over TCP, TTL expiration, and LRU eviction.

## Languages (locked)

**Go**

Canonical matrix: [`../SYSTEMS_LANGUAGE_MATRIX.md`](../SYSTEMS_LANGUAGE_MATRIX.md)

## Goals

- RESP2/RESP3 parser
- Concurrent key-value store with memory bounds
- Active and passive TTL expiration
- Compare against `redis-benchmark` for a core command set

## Status

Scaffold only (harness not built yet).

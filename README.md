# RADIX

**RESP-Addressable Data In-memory eXchange**

Go Redis-compatible cache learning harness. **You implement** RESP, store, and TCP server.

## Languages (locked)

**Go**

## Specs

- [specs/v0_resp.md](specs/v0_resp.md)
- [specs/v1_store.md](specs/v1_store.md)
- [specs/v2_server.md](specs/v2_server.md)

Progress: [docs/PROGRESS.md](docs/PROGRESS.md)

## Quick start

```bash
make expect-fail
# edit internal/resp/resp.go
make test
```

## Status

Harness scaffold. Not a working Redis yet.

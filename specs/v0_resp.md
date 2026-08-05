# RADIX Stage 0 — RESP parser

Status: **frozen**. Language: **Go** only.

## Purpose

Parse Redis Serialization Protocol (RESP2) from a byte stream into values.

## API

```text
Decode(reader) -> (Value, error)
Encode(writer, Value) error
```

Value kinds: SimpleString, Error, Integer, BulkString, Array (null bulk/array).

## Tests

Round-trip fixtures under `testdata/*.resp`.

# go-sequence

A Go library that builds a slice of positive integers out of its string representation.
Sequence is expected to be a comma-separated list of intervals. An interval is either a number, or a pair of numbers separated by a dash.

The library functionality can be explored with the provided executable.

Example:

```shell
go install ./cmd/sequence
sequence '1,3,6-8'     # 1 3 6 7 8
sequence '2003-1999'   # 2003 2002 2001 2000 1999
sequence '1-3,7,5-2'   # 1 2 3 7 5 4 3 2
```

# LSGO

A minimal implementation of `ls` equivalent in Go.

## Usage

```bash
go build .
./lsgo
```

Commands are as follows:

```bash
./lsgo -h
Usage: lsgo [OPTION]... [FILE]...
-ls	List files in current directory
-l	List files in current directory with long format
-la	List files in current directory with long format and hidden files
-a	List files in current directory with hidden files
-r	List files in current directory in reverse order
-s	List files in current directory with size
-S	List files in current directory with size in reverse order

```

Exemption: Working with `*` and `?` is not supported yet.

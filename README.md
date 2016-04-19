# Roman - Project specific command alias

> When in Rome, do as the Romans do.

## Install

```sh
go get github.com/typester/roman
```

## Usage

Create `.roman.yml` on your project root, edit like following:

```yaml
- name: go
  exec: docker run -it --rm -v $ROMAN_ROOT:/app -w /app/$ROMAN_REL -e GOPATH=/app golang go
```

Then instead of typing above huge command lines, just type

```sh
roman go
```

Arguments are passed too, so you can:

```sh
roman go build
```

or

```sh
roman go get -u -v ...
```

But can't:

```sh
roman gofmt -s -w .
```

You need `gofmt` definition in `.roman.yml` if you want to do this.

## Environment Variables

Roman set custom environment variables listed below.
These are usable in `.roman.yml`

- `ROMAN_CONFIG`: path to `.roman.yml`
- `ROMAN_ROOT`: directory that located `.roman.yml` (generally it's project root)
- `ROMAN_REL`: relative path from ROMAN_ROOT to working directory

## TODO

- [ ] more options, dry-run etc
- [ ] exec tests

## Author

Daisuke Murase (typester)

## License

MIT

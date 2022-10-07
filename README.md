# imgi

A CLI tool that prints images' information.

## Usage

```bash
imgi [DIR] # DIR: directory containing image files (default: "./")
```

`imgi` will print images' information like this:

```txt
images:
  - name: 1.jpg
    width: 2160
    height: 2880
  - name: 2.png
    width: 2160
    height: 2880
  - name: 3.webp
    width: 2160
    height: 2880
```

## Flags

- --format value, -f value

  set output format, available value: "yaml", "toml", "json" (default: "yaml")

- --copy, -c

  copy output to clipboard (default: false)

## Supported image formats

- JPEG (.jpg, .jpeg)

- PNG (.png)

- WEBP (.webp)

- GIF (.gif)

## Install from source code

```bash
go install github.com/mys1024/imgi@latest
```

## License

MIT

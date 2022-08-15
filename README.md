# imgi

A CLI tool that prints images' information.

## Usage

```bash
imgi [DIR]
```

`DIR`: Directory containing image files (default: "./")

**imgi** prints images' information in **YAML** format:

```txt
dir: /root
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

## Supported image formats

- JPEG (.jpg, .jpeg)

- PNG (.png)

- WEBP (.webp)

- GIF (.gif)

## Install from source code

```bash
go install github.com/mys1024/imgi@latest
```

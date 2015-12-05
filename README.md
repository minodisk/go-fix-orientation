# go-fix-orientation [![Circle CI](https://circleci.com/gh/minodisk/go-fix-orientation/tree/master.svg?style=svg)](https://circleci.com/gh/minodisk/go-fix-orientation/tree/master)

Apply Exif orientation tag to pixels.

## Installation

```bash
go get github.com/minodisk/go-fix-orientation/processor
```

## Usage

```go
r, err := os.Open("path/to/image.jpg")
if err != nil {
  panic(err)
}
fixed, err := processor.Process(r)
if err != nil {
  panic(err)
}
```

# mkrepeat

`mkrepeat` is a cli tool to repeat a (audio) fail `n` times.

The main purpose is to have a convenience wrapper around `ffmpeg` if you want to have one (audio) file that loops.

## Synopsis

```sh
mkrepeat [-n <nr>] <filepath>
```

## Dependencies

* `go` compiler
* `ffmpeg`

## Quick start

* To have it system wide automatically:

```sh
go install
mkrepeat -n 5 $HOME/Music/song.mp3
```

* To make it system wide manually (just a single example):

```sh
go build
doas ln -s $(pwd)/mkrepeat /usr/local/bin/mkrepeat
mkrepeat -n 5 $HOME/Music/song.mp3
```

* To run it in place:
```sh
go run main.go -n 5 $HOME/Music/song.mp3
```

## Author

Meelis Utt

## Simple File Server built on golang

No css, no js, no nothing. Just listing files.

#### run on local device
```
$ go install
$ sfs [-path PATH, default to .] [-port PORT, default to 8000]
```

#### run with docker
```
$ make docker
$ docker run -it --rm -p 80:8000 -v /path/to/server/root:/data gosfs
```

#### run pre-built docker image from docker hub
```
$ docker run -it --rm -p 80:8000 -v /path/to/server/root:/data 21314/gosfs
```
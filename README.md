## Simple File Server built on golang

No css, no js, no nothing. Just listing files.

#### run on local device
```
$ go build -o /usr/local/bin/sfs
$ sfs [-path PATH] [-port PORT]
```

#### run with docker
```
$ build -t sfs .
$ docker run -it --rm -p 80:8000 -v /path/to/server/root:/data sfs
```

#### run pre-built docker image from docker hub
```
$ docker run -it --rm -p 80:8000 -v /path/to/server/root:/data 21314/gosfs
```
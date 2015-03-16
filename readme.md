## Bridge

Exports environment configuration variables from etcd to applications.

```
$ etcdctl ls /example.com/app/
/example.com/app/S3_BUCKET

$ etcdctl get /example.com/app/S3_BUCKET
app.example.com

$ go run bridge.go -debug -path=/example.com/app /bin/bash -c 'echo starting && sleep 10 && echo $S3_BUCKET'
2015/03/15 22:21:24 Application path: /example.com/app, etcd endpoint http://127.0.0.1:4001
2015/03/15 22:21:24 Exporting S3_BUCKET as app.example.com
2015/03/15 22:21:24 Executing /bin/bash -c echo starting && sleep 10 && echo $S3_BUCKET
starting
app.example.com
```

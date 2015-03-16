## Bridge

Configures and executes applications that expect environment variables from etcd data. 

### Usage

```
$ ./bridge -h
Usage of ./bridge:
  -debug=false: log environment variable values
  -etcd_host="http://127.0.0.1:4001": etcd cluster endpoint
  -path="/example.com/app_name": Path to application variables

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

Store your configuration under /example.com/application_name/VAR_NAME and it will be exported to the child process.

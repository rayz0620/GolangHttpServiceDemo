# Task Description Overview 

Write a micro-service to gather stats of a file system. Estimated to take 30 mins to complete. 

## Specs 

1. Run on any computer with Docker installed
2. Service accepts HTTP POST request at ‘/’ and port: 8080
3. Tester posts data: ‘path’ and ‘timeout’ to ‘/’
4. Service responses with a JSON of keys: ‘time’, ‘path’, ‘timeout’ and ‘duration’. 

## Query 

path: file path inside container
 timeout: timeout of reading from fs in seconds 

**Sample query**: 

```shell
curl -d 'path=/root/main' -d 'timeout=3'  http://127.0.0.1:8080
```

**Response**

* time (string): server time (after gathering stats), in format: ‘2018-03-15 16:55:54’
* timeout (int64): timeout from input
* path (string): file path from input
* duration (int64): time of reading the file at ‘path’ in micro second 

**Sample response**: 

```json
 {"time":"2018-03-15 09:03:22","path":"/root/main","timeout":3,"duration":5953}
```

## Submission 

Should contain a Dockerfile, plus the source code
Commit your solution to Github/Gitlab/Bitbucket and send us the link of the repo 
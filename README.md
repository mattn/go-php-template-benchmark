# go-php-template-benchmark

template engine benchmark for Go(html/template) and PHP

## Result

PHP 7.3 smarty

```
$ ab -n 10000 http://localhost:8081/
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        Apache/2.4.52
Server Hostname:        localhost
Server Port:            8081

Document Path:          /
Document Length:        20041 bytes

Concurrency Level:      1
Time taken for tests:   37.157 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      202370000 bytes
HTML transferred:       200410000 bytes
Requests per second:    269.13 [#/sec] (mean)
Time per request:       3.716 [ms] (mean)
Time per request:       3.716 [ms] (mean, across all concurrent requests)
Transfer rate:          5318.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     3    4   0.4      4      18
Waiting:        3    4   0.4      3      18
Total:          3    4   0.4      4      18
ERROR: The median and mean for the waiting time are more than twice the standard
       deviation apart. These results are NOT reliable.

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      4
  80%      4
  90%      4
  95%      4
  98%      5
  99%      5
 100%     18 (longest request)
```

Go 1.21 html/tmeplate

```
$ ab -n 10000 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
apr_socket_recv: Connection refused (111)
mattn@dolphin:~/dev/cloudflare-nostr-nullpoga$ ab -n 10000 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        17039 bytes

Concurrency Level:      1
Time taken for tests:   17.750 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      171350000 bytes
HTML transferred:       170390000 bytes
Requests per second:    563.40 [#/sec] (mean)
Time per request:       1.775 [ms] (mean)
Time per request:       1.775 [ms] (mean, across all concurrent requests)
Transfer rate:          9427.51 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    2   0.3      2      21
Waiting:        1    1   0.1      1       3
Total:          2    2   0.3      2      21

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      2
  75%      2
  80%      2
  90%      2
  95%      2
  98%      2
  99%      3
 100%     21 (longest request)
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)

# go-php-template-benchmark

template engine benchmark for Go(html/template) and PHP

## Result

PHP 8.2 fpm smarty

```
$ ab -n 10000 -c 10 http://localhost:8081/
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        nginx/1.27.0
Server Hostname:        localhost
Server Port:            8081

Document Path:          /
Document Length:        20041 bytes

Concurrency Level:      10
Time taken for tests:   8.913 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      202040000 bytes
HTML transferred:       200410000 bytes
Requests per second:    1121.95 [#/sec] (mean)
Time per request:       8.913 [ms] (mean)
Time per request:       0.891 [ms] (mean, across all concurrent requests)
Transfer rate:          22136.55 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   2.5      0     176
Processing:     3    8   9.8      8     184
Waiting:        2    8   9.8      7     184
Total:          3    9  10.1      8     186

Percentage of the requests served within a certain time (ms)
  50%      8
  66%      8
  75%      9
  80%      9
  90%     10
  95%     12
  98%     14
  99%     16
 100%    186 (longest request)
```

Go 1.21 html/template

```
$ ab -n 10000 -c 10 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        17039 bytes

Concurrency Level:      10
Time taken for tests:   5.966 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      171350000 bytes
HTML transferred:       170390000 bytes
Requests per second:    1676.18 [#/sec] (mean)
Time per request:       5.966 [ms] (mean)
Time per request:       0.597 [ms] (mean, across all concurrent requests)
Transfer rate:          28048.16 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       5
Processing:     2    6   2.2      5      28
Waiting:        1    4   1.9      4      27
Total:          2    6   2.2      6      29

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      6
  75%      7
  80%      7
  90%      8
  95%     10
  98%     11
  99%     13
 100%     29 (longest request)
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)

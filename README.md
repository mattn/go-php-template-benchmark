# go-php-template-benchmark

template engine benchmark for Go(html/template) and PHP

## Result

PHP 8.2 fpm smarty

```
$ ab -n 10000 -c 10 http://localhost:8081/
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        nginx/1.27.0
Server Hostname:        localhost
Server Port:            8081

Document Path:          /
Document Length:        20041 bytes

Concurrency Level:      10
Time taken for tests:   3.899 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      202040000 bytes
HTML transferred:       200410000 bytes
Requests per second:    2564.86 [#/sec] (mean)
Time per request:       3.899 [ms] (mean)
Time per request:       0.390 [ms] (mean, across all concurrent requests)
Transfer rate:          50605.87 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     1    4   1.0      4      16
Waiting:        1    3   1.0      3      15
Total:          2    4   1.0      4      16

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      4
  80%      4
  90%      5
  95%      6
  98%      7
  99%      7
 100%     16 (longest request)
```

Go 1.21 html/tmeplate

```
$ ab -n 10000 -c 10 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        17039 bytes

Concurrency Level:      10
Time taken for tests:   6.945 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      171350000 bytes
HTML transferred:       170390000 bytes
Requests per second:    1439.82 [#/sec] (mean)
Time per request:       6.945 [ms] (mean)
Time per request:       0.695 [ms] (mean, across all concurrent requests)
Transfer rate:          24093.05 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:     2    7   3.6      6      75
Waiting:        1    4   2.1      3      43
Total:          2    7   3.6      6      75

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      7
  75%      8
  80%      8
  90%     10
  95%     12
  98%     15
  99%     19
 100%     75 (longest request)
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)

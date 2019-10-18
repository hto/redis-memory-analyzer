# Redis / ElastiCache Memory Analyzer
- Redis or Aws ElastiCache Redis Memory Analyzer.
- Groups all keys with 'SCAN' by name. It then lists the dimensions it occupies in memory with the group prefixes that it determines. 
- Export to csv.

## Warning
Using Example Redis Commands
```
SCAN 0 MATCH '*' COUNT 999999
MEMORY USAGE users
```
- [Redis SCAN command](https://redis.io/commands/scan)
- [Redis MEMORY USAGE command](https://redis.io/commands/memory-usage)

## RunTime
```
25-30 minutes(avg) for Aws ElastiCache 24 millions Keys report
```

## Usage

### Start using it

Download and install it:

```sh
$ go get github.com/hto/redis-memory-analysis
```

```sh
$ go build
```

### Runnnig alternative
```sh
$ ./redis-memory-analyzer 
$ ./redis-memory-analyzer -host=rediscluster-0001111.cache.amazonaws.com
$ ./redis-memory-analyzer -port=6390
$ ./redis-memory-analyzer -port=6390 -password=1234
$ ./redis-memory-analyzer -port=6390 -password=1234 -reportPath=analysis
```

## Example Report

|Key|Count                        |Size  |
|---|-----------------------------|------|
|users:data:*|109270                       |832.476 MB|
|users:products:pid:*|1410800                      |784.237 MB|
|users:inbox:*|593842                       |761.199 MB|
|users:energy:*|1176558                      |532.499 MB|
|customer:bundle:*|1417628                      |511.679 MB|
|users:packages:*|314613                       |426.036 MB|
|toma:*|1015310                      |323.098 MB|
|hto:power:*|869738                       |309.491 MB|
|users:powers:*|1404458                      |297.675 MB|
|users:reward:*|1544428                      |229.860 MB|
|cards|1                            |200.563 MB|
|users:video:*|575571                       |192.188 MB|

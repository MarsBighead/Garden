# Go package Aiqiyi

Aiqiyi provide a struct for aiqiyi request and response.

The package file is automatic generated by official protobuf tool - protoc.

## Ubuntu Platform

##########################

Run gen.sh with command line:

```bash
user@ubuntu:~$ bash gen.sh
```

## Other Platform

Can try short command sh for bash or [x]sh

```bash
appadmins-Mac-mini-2:aiqiyi user$ sh gen.sh
```

## Refernce

##########################

1. Similiar document to understand golang protobuf struct automatic generation mechanism.
  appcoachs.net/x/xiaodu/README.md

2. aiqiyi.proto

- We get aiqiyi.proto according  merge bid_request.proto and bid_response.proto.

- The union package name in .proto file is changed to aiqiyi.
  
- Remove unneccessary lines content "extensions 100 to max;" in .proto file to avoid introduce bugs in aiqiyi.pb.go.

2017.02.17
Author: Paul Duan

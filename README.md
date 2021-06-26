# go-queue

這邊用了兩種方式, 實作 queue
一個是 傳統資料結構
另一個是 使用 golang channel 實作

golang channel 最後會印出錯誤, 本程式碼 先專注於 queue 的實作, 
而有關 channel 的 error handling, 會在 go-channel 另行介紹處理

## build

```shell
$ > make build PLATFORM=<PLATFORM> STAGE=<STAGE>
```

Note:
You can build for a specific platform by setting PLATFORM.
PLATFORM:
* local
* linux
* windows/amd64.
You can build with config file by setting STAGE
STAGE:
* stg, for stg.go
* prod, for prod.go

## Launch Develop Environment

It would launch /bin/bash in golang:1-buster (ubuntu)
```shell
$ > make devenv
```

### Update package

```shell
$ > make update
```

## Clean

```shell
$ > make clean
```


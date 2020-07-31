# zeiss-machinestream
zeiss machinestream test project

## How to run

Download pre-compile excutable file from release

- upzip the target release
- then run follow script

```
chmod +x machinestream_mac_amd64

./machinestream_mac_amd64

or use specific http port if is conflict with exist http service, don't forget add colon(:) before port number
./machinestream_mac_amd64 --port :8080

```

## Build it by yourself or for other platform

 You need install [golang](https://golang.org/doc/install) on your laptop first
```
#build from source code
> make

#build for windows x64 platform
> env GOOS=windows GORACH=amd64 go build -ldflags="-s -w" -o bin/machinestream_win_amd64 *.go
```

## View machine stream data
> Open your favorite web browser, open link http://localhost:8090 to check the http server is properly running.
> open link http://localhost:8090/machines to fetch machine status data stream

Download Wechat unified login ppt from [here](https://zeisssharefolder.s3-us-west-1.amazonaws.com/WeChat+login.pptx)

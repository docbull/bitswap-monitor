## IPFS Bitswap Monitor
Bitswap monitor for [IPFS](https://ipfs.io) nodes.

> This software is based on [IPFS monitor](https://github.com/asabya/ipfs-monitor) software that uses HTTP API.
> 
> You need to install IPFS node that uses command line (i.e., IPFS daemon, ipfs-go, or ipfs-js) to run this software.

![image](https://user-images.githubusercontent.com/59289320/179771011-e95ca2f6-259c-42db-bcac-0eb80a32099b.png)

### How to Download
```
$ git clone https://github.com/docbull/bitswap-monitor.git
$ cd bitswap-monitor
```

### Hwo to Run

If you didn't configured or modified IPFS daemon settings, type following commands; it connects to IPFS node `http://localhost:5001` using HTTP API.
```
$ go run main.go
```

Otherwise, specify customized HTTP address by adding as a flag; if you changed http port number from `5001` to `5051`:
```
$ go run main.go http://localhost:5051
```
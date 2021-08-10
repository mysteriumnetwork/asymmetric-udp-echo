asymmetric-udp-echo
===================

Receives datagram on receiver port, extracts Port and UUID and sends UUID back to originating address on specified Port. Useful for checking whether UDP port is open to external connections or not.

## Installation

#### Binaries

Pre-built binaries are available [here](https://github.com/mysteriumnetwork/asymmetric-udp-echo/releases/latest).

#### Build from source

Alternatively, you may install application from source. Run the following within the source directory:

```
make install
```

## Usage

Ideally, this application is intended to use two distinct IP addresses: one to receive datagrams, another to send. Adjust it with `-bind-receiver` and `-bind-sender` options.

#### Autostart

Place this text into `/etc/systemd/system/asymmetric-echo-udp.service`:

```
[Unit]
Description=asymmetric UDP echo
After=syslog.target network.target

[Service]
Type=simple
User=nobody
Group=nogroup
ExecStart=/usr/local/bin/asymmetric-udp-echo -bind-receiver FIRST_IP:4589 -bind-sender SECOND_IP:0
Restart=always
KillMode=process
TimeoutStartSec=5
TimeoutStopSec=5

[Install]
WantedBy=multi-user.target
```

replacing `FIRST_IP` and `SECOND_IP` with actual values and path to binary to actual location. Then issue following commands:

```
systemctl daemon-reload
systemctl enable asymmetric-echo-udp.service
systemctl start  asymmetric-echo-udp.service
```

## Synopsis

```
# asymmetric-udp-echo -h
Usage of asymmetric-udp-echo:
  -bind-receiver string
    	socket address for request datagrams (default "0.0.0.0:4589")
  -bind-sender string
    	socket address for response datagrams (default "0.0.0.0:0")
  -reuse-socket
    	reuse response socket
  -version
    	show program version and exit
```

# rpp

** This is a quick and dumb script (aka prototype). Use it at your own risk. **

RPP: *R*aspberry *P*re*P*rovision. This helps you to initiate your Raspberry Pi enablig SSH and WiFi access by default. This is to install a headless Raspberry pi.

This is helpfull if you want to start using your Raspberry Pi right away without the need to connect a screen, keyboard and network cable to start using it.
Just turn it on, and access the device using ssh.

## How to use

You need to pass the volumen disk where the boot is located, ssid and password of the WiFi network _in that order_.

**Example**
`rpp "/mnt/f" "HOME-WIFI" "p455w0rd"`

You can then discovery what's the IP address of the device using nmap like `nmap -sn 192.168.1.0/24`

**Note:** This runs on Linux and WSL on Windows. This will only take effect on a freshly Raspbian flashed card.

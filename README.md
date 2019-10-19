# rpp
RPP: *R*aspberry *P*re*P*rovision. This helps you to initiate your Raspberry Pi enablig SSH and WiFi access by default.

This is helpfull if you want to start using your Raspberry Pi right away without the need to connect a screen, keyboard and network cable to start using it.
Just turn it on, and access to it using ssh.

## How to use

You need to pass the volumen disk where the boot is located, ssid and password of the WiFi network _in that order_.

**Example on Linux**
`rpp "/mnt/f" "HOME-WIFI" "p455w0rd"`

You can then discovery what it's the Raspberry Pi IP address using nmap like `nmap -sn 192.168.1.0/24`

**Note:** This will only work on a freshly Raspbian flashed card.

# rpp

**This is a quick and dumb script (aka prototype). Use it at your own risk.**

RPP: **R**aspberry **P**re**P**rovision. This helps you to initiate your Raspberry Pi enablig SSH and WiFi access by default.
This is helpfull if you want to start using your Raspberry Pi right away without the need to connect a screen, keyboard and network cable. This is to install a [headless Raspberry pi](https://www.raspberrypi.org/documentation/configuration/wireless/headless.md).

After running this program, all you have to do is turn on your raspberry and access the device using ssh.

## How to use

First you need to flash your microSD card with a fresh Raspbian.

Then you  run this program passing the volumen disk where the boot is located, ssid and password of the WiFi network _in that order_.

**Example on Linux**
`rpp "/mnt/f/" "HOME-WIFI" "p455w0rd"`

**Example on Windows**
`rpp.exe "F:\" "HOME-WIFI" "p455w0rd"`

That's it!

Now just turn on the Raspberry. You can then discovery what's the IP address of the device using nmap like `nmap -sn 192.168.1.0/24` and access to it using the default username (pi) and password (raspberry)

**Note:** This will only take effect on a freshly Raspbian flashed card.

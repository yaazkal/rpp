package main

import (
	"log"
	"os"
)

func enable_ssh(base string) {
	var ssh_file, err = os.Create(base + "ssh")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer ssh_file.Close()
	log.Println("SSH enabled")
}

func enable_wifi(base string, ssid string, password string) {
	var wifi_file, err = os.Create(base + "wpa_supplicant.conf")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	data := []byte(`update_config=1
ctrl_interface=/var/run/wpa_supplicant

network={
 ssid="` + ssid + `"
 psk="` + password + `"
}`)

	_, err = wifi_file.Write(data)

	defer wifi_file.Close()

	log.Printf("Wifi enabled. The device will connect to %v\n", ssid)
}

func main() {
	// Check if arguments are correctly passed
	if len(os.Args) != 4 {
		log.Fatalln("Please use three arguments in the correct order, see README.md")
	}

	baseDirectory := os.Args[1]
	ssid := os.Args[2]
	ssid_password := os.Args[3]
	log.Printf("The boot directory to use is: %v\n", baseDirectory)

	enable_ssh(baseDirectory)
	enable_wifi(baseDirectory, ssid, ssid_password)
}

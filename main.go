/*
MIT License

Copyright (c) 2021 Martin Stuckenbröker

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
package main

import (
	"flag"
	"fmt"

	"github.com/barnybug/go-tradfri"
)

func Connect(ip string, key string) (*tradfri.Client, error) {
	client := tradfri.NewClient(ip)
	err := client.LoadPSK()
	if err != nil {
		client.Key = key
	}
	err = client.Connect()
	if err == nil {
		client.SavePSK()
	} else {
		return nil, err
	}
	return client, nil
}

func GetDeviceInfo(c *tradfri.Client, id int) (string, int, error) {
	var status int
	var device string
	if id&(1<<17) == 0 {
		info, err := c.GetDeviceDescription(id)
		if err != nil {
			return "", -1, err
		}
		status = *info.LightControl[0].Power
		device = info.DeviceName

	} else {
		info, err := c.GetGroupDescription(id)
		if err != nil {
			return "", -1, err
		}
		status = info.Power
		device = info.GroupName
	}
	return device, status, nil

}

func GetStatus(c *tradfri.Client, id int) {
	_, status, err := GetDeviceInfo(c, id)
	if err != nil {
		fmt.Println("-1")
		return
	}
	fmt.Println(status)
	return
}

func Switch(c *tradfri.Client, id int, mode int) {
	device, status, err := GetDeviceInfo(c, id)
	if err != nil {
		return
	}

	change := tradfri.LightControl{}
	power := 0

	if mode < 0 || mode > 1 {
		if status == 0 {
			power = 1
		}
	} else {
		power = mode
	}

	change.Power = &power
	fmt.Printf("Switching %s to %d...", device, power)
	if id&(1<<17) == 0 {
		c.SetDevice(id, change)
	} else {
		c.SetGroup(id, change)
	}

}

func List(c *tradfri.Client) {
	listd, err := c.ListDevices()
	if err != nil {
		fmt.Printf("Cannot list devices of client %s", c.Ident)
		return
	}
	for _, d := range listd {
		dtype := "other"
		if d.LightControl != nil {
			dtype = "Light"
		}
		fmt.Printf("Device ID: %d / Device Name: %s / Device Type: %s\n", d.DeviceID, d.DeviceName, dtype)
	}
	listg, err := c.ListGroups()
	if err != nil {
		fmt.Printf("Cannot list devices of client %s", c.Ident)
		return
	}
	for _, g := range listg {
		fmt.Printf("Group ID: %d / Group Name: %s\n", g.GroupID, g.GroupName)
	}

}

func main() {
	var id int
	var ip string
	var key string
	var command string
	var mode string

	flag.IntVar(&id, "id", 0, "ID of group or device to switch")
	flag.StringVar(&ip, "ip", "", "IP of gateway")
	flag.StringVar(&key, "key", "", "Security code of gateway (only for initial connect)")
	flag.StringVar(&command, "command", "list", "command, values [list | switch | get]")
	flag.StringVar(&mode, "mode", "toggle", "mode, values: [on | off | toggle]")
	flag.Parse()
	if ip == "" {
		fmt.Println("-1")
		fmt.Println("Missing IP")
	}

	client, err := Connect(ip, key)
	if err != nil {
		fmt.Println("Unknown error!")
	}
	if command == "switch" {
		imode := -1
		switch mode {
		case "toggle":
			imode = -1
		case "on":
			imode = 1
		case "off":
			imode = 0
		}
		Switch(client, id, imode)
	} else if command == "get" {
		GetStatus(client, id)
	} else if command == "list" {
		List(client)
	} else {
		fmt.Println("Unkown command!")
	}
}

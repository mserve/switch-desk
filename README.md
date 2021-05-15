# switch-desk

*switch-desk* is a simple command line utility to switch or toggle the power of IKEA Homesmart (f.k.a. Tradfri) devices with a simple command on the command line.

## Usage

### Enumerate devices and groups

    switch-desk -ip=192.168.123.234 -key=SECRETCODEFROMBACKOFGATEWAY 

Result:

    Device ID: 65543 / Device Name: TRADFRI on/off switch / Device Type: other
    Device ID: 65540 / Device Name: Desk / Device Type: Light
    Device ID: 65547 / Device Name: Ceiling / Device Type: Light
    Device ID: 65545 / Device Name: Left / Device Type: Light
    Device ID: 65548 / Device Name: Right / Device Type: Light
    Device ID: 65549 / Device Name: TRADFRI remote control / Device Type: other
    Group ID: 132574 / Group Name: Desk
    Group ID: 132576 / Group Name: SuperGroup
    Group ID: 132577 / Group Name: Living Room

### Toggle a device or group

    switch-desk -ip=192.168.123.234 -command=switch -id=65540 -mode=toggle
    switch-desk -ip=192.168.123.234 -command=switch -id=132574 -mode=toggle


### Switch on or off a device or group


    switch-desk -ip=192.168.123.234 -command=switch -id=65547 -mode=on
    switch-desk -ip=192.168.123.234 -command=switch -id=132577 -mode=off

### Get a devices status

* Remark: This command is useful for your personal scripting efforts

## Problems

If *switch-desk* does not work as expected, check the IP address and the security code of your gateway. If both are correct, try to ping your gateway to ensure network connectivity.


## FAQ

### How do I control *any other feature* of my *some device*?

You can't: *switch-desk* has only a single purpose: switch on or off the power. If you need further control, check e.g.
* *go-tradri* at https://github.com/barnybug/go-tradfri
* *tradri-go* at https://github.com/eriklupander/tradfri-go/   

###
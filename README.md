# switch-desk 🔀🔌💡

``switch-desk`` is a simple command line utility to 🔀 switch or  toggle the 🔌 power of IKEA Homesmart (f.k.a. Tradfri) devices 💡 with a simple command on the command line.

## Usage

The key has only to be set at the first call of ``switch-desk``. At first connect, an individual *pre-shared key* (PSK) is generated and stored in the users home directory in the file ``~/.tradfri-psk``. This behaviour is controlled by the library github.com/barnybug/go-tradfri and cannot be changed within ``switch-desk``. 

For the sake of idempotency, it is recommended to always add the key when calling ``switch-desk``. A new PSK is only generated if required.

### Enumerate devices and groups

*Remark: Do no forget to add the key at least the first time you connect to your gateway*

    switch-desk -ip=192.168.123.234 -key=SECRETCODEFROMBACKOFGATEWAY 

**Result:**

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

    switch-desk -ip=192.168.123.234 -command=switch -id=65540  -mode=toggle
    switch-desk -ip=192.168.123.234 -command=switch -id=132574 -mode=toggle


### Switch on or off a device or group


    switch-desk -ip=192.168.123.234 -command=switch -id=65547  -mode=on
    switch-desk -ip=192.168.123.234 -command=switch -id=132577 -mode=off

### Get a devices status

*Remark: This command is useful for your personal scripting efforts*

    switch-desk -ip=192.168.123.234 -command=get -id=65547


**Result:**

    1

*Possible results are*

| Code  |            |
|:-----:|-------------|
| -1    | Device not found | 
|  0    | Group/Device is powered OFF | 
|  1    | Group/Device is powered ON   | 

## Building the tool

    go get .
    go build switch-desk.go


## FAQ

### How do I control *any other feature* of my *some device*?

You can't: *switch-desk* has only a single purpose: switch on or off the power. If you need further control, check e.g.
* *go-tradri* at https://github.com/barnybug/go-tradfri
* *tradri-go* at https://github.com/eriklupander/tradfri-go/   

### I get an error message

If ``switch-desk`` does not work as expected, check the IP address and the security code (key)of your gateway. If both are correct, try to ping your gateway to ensure network connectivity.

### Can you add *any other feature* to this tool?
In short: **no**. Please use another tool if you want a specific feature. I developed this little program only to learn a litte bit of [golang](https://golang.org) and to quickly switch on and off the light of my desk by double-clinking on a link on my desktop.

### Why did you call your tool ``switch-desk``?
Well, because I developed it to ``switch`` the light above my ``desk`` by using a link on my desktop - ``switch-desk`` was born.
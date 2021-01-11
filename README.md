# Raspberry Pi 3d Christmas Tree 

Go code to support the [Raspberry Pi 3d Chistmas Tree](https://thepihut.com/products/3d-xmas-tree-for-raspberry-pi) as a learning alternative to the provided python code. 
Relies on [Gobot.io](https://gobot.io/)

## Example Usage

### Download, Build and Deploy with all Lights

    git clone
    cd rpi3christmastree

Update Makefile to your Raspberry Pi User and IP.
If using a Raspbery Pi 2 or older, update Makefile to GOARM=6

    make deploy

## Run on Raspberry Pi

    ssh [rpi-user]@[rpi-ip]
    ./christmastree -star=true -style=tree

## Usage of ./christmastree:

    ./christmastree -h
    
    -blink (default false)
    Enable blink effect for full tree. Use speed to adjust blink speed.

    -speed int (default 500)
    Delay for roation steps in milliseconds.

    -star (default true)
    Tree star light on.

    -style string (default tree)
    Use tree, branch, side or off to select lighting style.
    
    Examples:
        Rotate through branches at 25ms per step with star blinking:
        ./christmastree -star=true -style=branch -speed=25
    
        Turn Off Tree:
        ./christmastree -style=off
    
        Blink whole tree with blinking star:
        ./christmastree -star=true -style=tree -blink=true -speed=250
         (default "tree")

### Packages Used

* Gobot ([gobot.io](https://github.com/hybridgroup/gobot/))
* Command-line flags ([flag](https://golang.org/pkg/flag/))
* Time ([time](https://golang.org/pkg/time/))
* OS ([os](https://golang.org/pkg/os/))

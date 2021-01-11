# Raspberry Pi 3d Christmas Tree 

This repository contains a collection of Go programs and libraries that
demonstrate the language, standard libraries, and tools.

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

    christmastree -h

  -blink
        Enable blink effect for full tree. Use speed to adjust blink speed.
        (default false)
  -speed int
        Delay for roation steps in milliseconds.
         (default 500)
  -star
        Tree star light on.
         (default true)
  -style string
        Use tree, branch, side or off to select lighting style.
    
    Examples:
        Rotate through branches at 25ms per step with star blinking:
        ./christmastree -star=true -style=branch -speed=25
    
        Turn Off Tree:
        ./christmastree -style=off
    
        Blink whole tree with blinking start:
        ./christmastree -star=true -style=tree -blink=true -speed=250
         (default "tree")

### Packages Used

* Gobot ([gobot.io](https://github.com/hybridgroup/gobot/))
* Command-line flags ([flag](https://golang.org/pkg/flag/))
* Time ([time](https://golang.org/pkg/time/))
* OS ([os](https://golang.org/pkg/os/))

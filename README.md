# ESPmon0

TinyGo firmware for an ESP-WROOM-32 development board. It counts received
802.11 frames on one configured Wi-Fi channel and prints one CSV-like status
line per second over the board's USB-to-UART serial connection. It does not
save or forward packet contents.

## Build and flash

Install TinyGo with the `esp32-mini32` target, then from this directory run:

```sh
tinygo flash -target=esp32-mini32 -monitor .
```

The monitor should show lines like:

```text
ESPmon0,ready,channel=6
ESPmon0,count,channel=6,window_ms=1000,frames=23
```

Change `channel` in `main.go` to monitor a different 2.4 GHz channel (1–13).
The radio listens on only that channel during each sample window.

## How it works

1. The short delay after reset gives the serial monitor time to connect.
2. `espradio.Enable` initializes the radio, and `espradio.Start` starts its
   Wi-Fi driver.
3. Each `SniffCountOnChannel` call listens in monitor mode on the chosen
   channel for one second, counts received 802.11 frames, and returns the
   count. The radio does not listen on other channels during that window.
4. `fmt.Printf` writes a line to TinyGo's standard output, which the board's
   USB-to-UART bridge sends to the computer. The line contains the channel,
   window length, and frame count. It is a count-only API: this program does
   not receive or print individual frames.

## Using another serial-console program

`-monitor` opens the serial port itself. To read the output with another
program, flash without `-monitor`, then open the board's serial port at
115200 baud. For example, with Python and `pyserial` installed:

```python
import serial

with serial.Serial("/dev/ttyUSB0", 115200, timeout=1) as port:
    for line in port:
        print(line.decode("utf-8", errors="replace").rstrip())
```

Use the port shown for your board (for example, `/dev/ttyUSB0` on Linux or a
`COM` port on Windows). This starter reports frame counts only; it does not
provide per-frame metadata or payloads. Test only with networks and devices
you own or are authorized to monitor.

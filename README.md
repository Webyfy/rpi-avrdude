# rpi-avrdude

Fixing the lack of a DTR pin on the Raspberry Pi for the purposes of uploading hex files to the arduino via RPi's UART.

This project is based on https://github.com/Siytek/avrdude-rpi, which ultimately is based on https://github.com/deanmao/avrdude-rpi (via multiple levels of forking)

## Documentation

Documenation is avaialable [here](git.reach-iot.com/docs/rpi-avrdude)


## References:
* [HOW TO USE RASPBERRY PI GPIO SERIAL PORT TO PROGRAM ARDUINO](https://siytek.com/raspberry-pi-gpio-arduino/) - _Siytek(Simon)_
* [Fixing the DTR pin](http://www.deanmao.com/2012/08/12/fixing-the-dtr-pin/) - _Dean Mao_
  > Since the page is no longer available on the web,a copy of its text is included [here](https://git.reach-iot.com/docs/rpi-avrdude/fixing_the_dtr_pin.md)
* [Gobot](https://github.com/hybridgroup/gobot)
* [GPIO Programming: Using the sysfs Interface](https://www.ics.com/blog/gpio-programming-using-sysfs-interface)
* [GPIO (/sys/class/gpio) is operated by file IO under Linux](https://programmer.help/blogs/gpio-sys-class-gpio-is-operated-by-file-io-under-linux.html)
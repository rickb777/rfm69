//go:build linux && arm64 && !pizero

package rfm69

// Configuration for Raspberry Pi with Adafruit RFM69HCW bonnet:
// https://www.adafruit.com/product/4072
// https://learn.adafruit.com/adafruit-radio-bonnets/pinouts

const (
	spiDevice    = "/dev/spidev0.1"
	spiSpeed     = 4_000_000 // Hz
	interruptPin = 29        // GPIO for receive interrupts (DIO0)
	resetPin     = 22        // GPIO for hardware reset
)

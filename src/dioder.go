//Most parts of this file are taken from Lakrizz' example on GitHub
// https://github.com/lakrizz/go-pidioder/blob/master/main.go

package main

import (
	"os"
	"bufio"
	"strconv"
	"errors"
)

const (
	REDPIN   = "18"
	BLUEPIN  = "17"
	GREENPIN = "4"
)

var (
	CURRENT_RED int
	CURRENT_GREEN int
	CURRENT_BLUE int
)

func FloatToString(floatValue float64) string {
	return strconv.FormatFloat(floatValue, 'f', 6, 64)
}

func setColor(channel string, value float64) error {
	piBlasterCommand := channel + "=" + FloatToString(value) + "\n"

	file, error := os.OpenFile("/tmp/devPiBlaster", os.O_RDWR, os.ModeNamedPipe)

	if error != nil {
		panic(error)
	}

	defer file.Close()

	stream := bufio.NewWriter(file)

	_, error = stream.WriteString(piBlasterCommand)

	if error != nil {
		panic(error)
	}

	stream.Flush()

	return nil
}


func setChannelInteger(value int, channel string) error {
	if value > 255 {
		return errors.New("can't go over 255. sorry mate.")
	}

	if value < 0 {
		return errors.New("can't go below 0. sorry mate.")
	}

	floatval := float64(value) / 255.0

	setColor(channel, float64(floatval))

	return nil
}

func SetRed(value int) error {
	return setChannelInteger(value, REDPIN)
}

func SetGreen(value int) error {
	return setChannelInteger(value, GREENPIN)
}

func SetBlue(value int) error {
	return setChannelInteger(value, BLUEPIN)
}

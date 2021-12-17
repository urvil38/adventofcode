package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

var inputFile = flag.String("inputFile", "inputs/day16.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
	p2(input)
}

func p1(s string) {
	buf := bytes.NewBuffer([]byte(s))
	vs, _, _ := parsePacket(buf)
	fmt.Println(vs)
}

func p2(s string) {
	buf := bytes.NewBuffer([]byte(s))
	_, value, _ := parsePacket(buf)
	fmt.Println(value)
}

var ErrPacketParse = errors.New("Unable to parse packet")

func parsePacket(buf *bytes.Buffer) (versionSum int, value int, err error) {
	if buf.Len() < 7 {
		return 0, 0, ErrPacketParse
	}

	ver := binaryToInt(string(buf.Next(3)))
	versionSum = ver
	typeId := binaryToInt(string(buf.Next(3)))

	if typeId == 4 {
		// parse literal packet
		value = parseLiteralPacket(buf)
	} else {
		// parse operator packet
		verSum, val := parseOperatorPacket(typeId, buf)
		versionSum += verSum
		value = val
	}

	return versionSum, value, nil
}

func parseLiteralPacket(buf *bytes.Buffer) (value int) {
	vs := ""
	for {
		cont := buf.Next(1)
		if string(cont) == "0" {
			vs += string(buf.Next(4))
			break
		} else {
			vs += string(buf.Next(4))
		}
	}
	return binaryToInt(vs)
}

func parseOperatorPacket(typeId int, buf *bytes.Buffer) (versionSum int, value int) {
	lengthTypeId := string(buf.Next(1))
	var values []int
	if lengthTypeId == "0" {
		subPacketLen := binaryToInt(string(buf.Next(15)))
		subPacket := bytes.NewBuffer(buf.Next(subPacketLen))
		for {
			ver, val, err := parsePacket(subPacket)
			if err != nil {
				break
			}
			versionSum += ver
			values = append(values, val)
		}
	} else {
		subPacketCount := binaryToInt(string(buf.Next(11)))
		for i := 0; i < subPacketCount; i++ {
			ver, val, _ := parsePacket(buf)
			versionSum += ver
			values = append(values, val)
		}
	}

	value = operateOnValues(typeId, values)
	return versionSum, value
}

func operateOnValues(typeId int, values []int) (value int) {
	switch typeId {
	case 0: // sum values
		for _, v := range values {
			value += v
		}
	case 1: // product
		value = 1
		for _, v := range values {
			value *= v
		}
	case 2: // min
		value, _ = minMax(values)
	case 3: // max
		_, value = minMax(values)
	case 5: // >
		if values[0] > values[1] {
			value = 1
		}
	case 6: // <
		if values[0] < values[1] {
			value = 1
		}
	case 7: // ==
		if values[0] == values[1] {
			value = 1
		}
	}
	return
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func binaryToInt(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}

func parseInput() string {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	s := string(b)
	var input string
	for _, c := range s {
		in, _ := strconv.ParseInt(string(c), 16, 64)
		input += fmt.Sprintf("%04b", in)
	}
	return input
}

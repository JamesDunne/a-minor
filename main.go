// main
package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type MidiPrograms struct {
	MidiPrograms []MidiProgram `yaml:"midi_programs"`
}

type MidiProgram struct {
	Midi  int    `yaml:"midi"`
	Amps  []Amp  `yaml:"amps"`
	Songs []Song `yaml:"songs"`
}

type Amp struct {
	Name string `yaml:"name"`
}

type Song struct {
	Name string `yaml:"name"`
}

func main() {
	allPrograms := MidiPrograms{}

	f, err := os.Open("midi.v6.yml")
	if err != nil {
		panic(err)
	}

	y := yaml.NewDecoder(f)
	err = y.Decode(&allPrograms)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", allPrograms)
}

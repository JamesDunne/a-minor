// main
package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

////////

type Configuration struct {
	Channel         int                       `yaml:"channel"`
	Controllers     map[string]int            `yaml:"controllers"`
	ControllerMetas map[string]ControllerMeta `yaml:"controller_metas"`
	MidiPrograms    []MidiProgram             `yaml:"midi_programs"`
}

type ControllerMeta struct {
	Display string `yaml:",omitempty"`
	Kind    string
}

type MidiProgram struct {
	Midi  int    `yaml:"midi"`
	Amps  []Amp  `yaml:"amps"`
	Songs []Song `yaml:"songs"`
}

type Amp struct {
	Name string `yaml:"name"`

	GainCC   int     `yaml:"gain_controller_cc"`
	VolumeCC int     `yaml:"volume_controller_cc"`
	Blocks   []Block `yaml:"blocks"`
}

type Block struct {
	Name      string `yaml:"name"`
	EnabledCC int    `yaml:"enabled_switch_cc"`
	XYCC      *int   `yaml:"xy_switch_cc"`
}

type Song struct {
	Name string `yaml:"name"`
}

func main() {
	f, err := os.Open("midi.v6.yml")
	if err != nil {
		panic(err)
	}

	allPrograms := Configuration{}

	y := yaml.NewDecoder(f)
	err = y.Decode(&allPrograms)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", allPrograms)
}

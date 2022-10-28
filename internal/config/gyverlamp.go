package config

import (
	"github.com/FirinKinuo/gyverlampUDP"
	"net"
)

type gyverLampYAML struct {
	IP    string `yaml:"IP"`
	Port  uint16 `yaml:"port"`
	Group uint16 `yaml:"group"`
	GLKey string `yaml:"GLKey"`
}

func newGyverLampYaml() gyverLampYAML {
	gyverLamp := gyverLampYAML{
		IP:    net.IP{255, 255, 255, 255}.String(),
		Group: gyverlampUDP.DefaultGroup,
		GLKey: string(gyverlampUDP.DefaultGLKey),
	}
	defaultPort := gyverlampUDP.ComputeUDPPort([]byte(gyverLamp.GLKey), gyverLamp.Port)
	gyverLamp.Port = defaultPort

	return gyverLamp
}

// GetGyverLampUDP returns a pointer to a structure gyverlampUDP.GyverLamp
func (g gyverLampYAML) GetGyverLampUDP() *gyverlampUDP.GyverLamp {
	return gyverlampUDP.NewGyverLamp(net.ParseIP(g.IP), g.Port, []byte(g.GLKey))
}

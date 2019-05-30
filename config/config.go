package config

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/juju/errors"
)

// Config represents common configuration of mtg.
type Config struct {
	Debug      bool
	Verbose    bool
	SecureMode bool
	SecureOnly bool

	ReadBufferSize  int
	WriteBufferSize int

	BindPort       uint16
	PublicIPv4Port uint16
	PublicIPv6Port uint16

	BindIP     net.IP
	PublicIPv4 net.IP
	PublicIPv6 net.IP

	AntiReplayMaxSize      int
	AntiReplayEvictionTime time.Duration

	AdTag []byte
}

// URLs contains links to the proxy (tg://, t.me) and their QR codes.
type URLs struct {
	TG        string `json:"tg_url"`
	TMe       string `json:"tme_url"`
	TGQRCode  string `json:"tg_qrcode"`
	TMeQRCode string `json:"tme_qrcode"`
}

// IPURLs contains links to both ipv4 and ipv6 of the proxy.
type IPURLs struct {
	IPv4      URLs   `json:"ipv4"`
	IPv6      URLs   `json:"ipv6"`
	BotSecret string `json:"secret_for_mtproxybot"`
}

// BindAddr returns connection for this server to bind to.
func (c *Config) BindAddr() string {
	return getAddr(c.BindIP, c.BindPort)
}

// UseMiddleProxy defines if this proxy has to connect middle proxies
// which supports promoted channels or directly access Telegram.
func (c *Config) UseMiddleProxy() bool {
	return len(c.AdTag) > 0
}

func getAddr(host fmt.Stringer, port uint16) string {
	return net.JoinHostPort(host.String(), strconv.Itoa(int(port)))
}

// NewConfig returns new configuration. If required, it manages and
// fetches data from external sources. Parameters passed to this
// function, should come from command line arguments.
func NewConfig(debug, verbose bool, // nolint: gocyclo
	writeBufferSize, readBufferSize uint32,
	bindIP, publicIPv4, publicIPv6 net.IP,
	bindPort, publicIPv4Port, publicIPv6Port uint16,
	secureOnly bool,
	antiReplayMaxSize int, antiReplayEvictionTime time.Duration,
	adtag []byte) (*Config, error) {
	secureMode := secureOnly

	var err error
	if publicIPv4 == nil {
		publicIPv4, err = getGlobalIPv4()
		if err != nil {
			publicIPv4 = nil
		} else if publicIPv4.To4() == nil {
			return nil, errors.Errorf("IP %s is not IPv4", publicIPv4.String())
		}
	}
	if publicIPv4Port == 0 {
		publicIPv4Port = bindPort
	}

	if publicIPv6 == nil {
		publicIPv6, err = getGlobalIPv6()
		if err != nil {
			publicIPv6 = nil
		} else if publicIPv6.To4() != nil {
			return nil, errors.Errorf("IP %s is not IPv6", publicIPv6.String())
		}
	}
	if publicIPv6Port == 0 {
		publicIPv6Port = bindPort
	}

	conf := &Config{
		Debug:                  debug,
		Verbose:                verbose,
		SecureOnly:             secureOnly,
		BindIP:                 bindIP,
		BindPort:               bindPort,
		PublicIPv4:             publicIPv4,
		PublicIPv4Port:         publicIPv4Port,
		PublicIPv6:             publicIPv6,
		PublicIPv6Port:         publicIPv6Port,
		AdTag:                  adtag,
		SecureMode:             secureMode,
		ReadBufferSize:         int(readBufferSize),
		WriteBufferSize:        int(writeBufferSize),
		AntiReplayMaxSize:      antiReplayMaxSize,
		AntiReplayEvictionTime: antiReplayEvictionTime,
	}

	return conf, nil
}

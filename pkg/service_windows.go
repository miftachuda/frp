//go:build windows
// +build windows

// Windows Service Handler

package ntlm_proxy

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"os"

	"github.com/kardianos/service"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

type program struct {
	exit chan struct{}
}

var svcFlag string

func init() {
	if flag.Lookup("service") == nil {
		flag.StringVar(&svcFlag, "service", "", "Control the Windows System service.")
	}
}

func RunWindows() {
	svcConfig := &service.Config{
		Name:        "gontlm-proxy",
		DisplayName: "GoNTLM Proxy",
		Description: "GoNTLM Forwarding Proxy",
		Arguments:   []string{"-proxy", ProxyServer},
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(svcFlag) != 0 {
		err := service.Control(s, svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}

	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Start(s service.Service) (err error) {
	if service.Interactive() {
		if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) {
			log.SetFormatter(&log.TextFormatter{ForceColors: true})
			log.SetOutput(colorable.NewColorableStdout())
		}
		log.Debug("Running in terminal.")
	} else {
		log.Debug("Running under service manager.")
	}
	p.exit = make(chan struct{})

	go p.run()

	return
}

func (p *program) run() (err error) {
	// Run the Proxy
	go Run()
	// Wait for Exit Signal
	<-p.exit
	return
}

func (p *program) Stop(s service.Service) (err error) {
	close(p.exit)
	return
}

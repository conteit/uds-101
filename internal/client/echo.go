package client

import (
	"bufio"
	"io"
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

func EchoTo(addr string) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		s, err := net.Dial("unix", addr)
		if err != nil {
			log.Fatal().
				Err(err).
				Str("address", addr).
				Msg("failed connecting to echo server")
			close(c)
		}
		defer s.Close()

		go func() {
			io.Copy(s, os.Stdout)
		}()

		r := bufio.NewReader(os.Stdin)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				close(c)
			} else if err != nil {
				log.Fatal().
					Err(err).
					Msg("failed reading stdin")
				close(c)
			}
			s.Write([]byte(line + "\n"))
		}

	}()
	return c
}

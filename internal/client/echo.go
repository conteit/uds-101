package client

import (
	"bufio"
	"io"
	"net"
	"os"
	"strings"

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
			r := bufio.NewReader(s)
			for {
				line, err := r.ReadString('\n')
				if err == io.EOF {
					close(c)
					return
				} else if err != nil {
					log.Fatal().
						Err(err).
						Msg("failed reading from socket")
					close(c)
					return
				}
				if strings.Index(line, "<uid:") == 0 {
					log.Info().Str("uid", line[5:len(line)-1]).Msg("connected")
				} else if line != "\n" {
					log.Info().Msg(strings.ReplaceAll(line, "\n", ""))
				}
			}
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
			s.Write([]byte(line))
		}

	}()
	return c
}

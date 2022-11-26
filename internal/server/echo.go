package server

import (
	"io"
	"net"

	"github.com/rs/zerolog/log"
)

func EchoOn(addr string) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		s, err := net.Listen("unix", addr)
		if err != nil {
			log.Fatal().
				Err(err).
				Str("address", addr).
				Msg("cannot spawn echo server")
			close(c)
			return
		}
		defer s.Close()

		for {
			conn, err := s.Accept()
			if err != nil {
				log.Error().Err(err).Msg("failed accepting incoming connection")
			} else {
				go echo(conn)
			}
		}
	}()
	return c
}

func echo(c net.Conn) {
	defer c.Close()
	log.Info().Str("remoteAddr", c.RemoteAddr().String()).Msg("client connected")
	io.Copy(c, c)
}

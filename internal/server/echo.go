package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/ksuid"
)

func EchoOn(addr string) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		if err := os.RemoveAll(addr); err != nil {
			log.Fatal().Err(err).Msg("failed initializing socket file")
			close(c)
			return
		}

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

func newUid() string {
	uid, _ := ksuid.NewRandom()
	return uid.String()
}
func echo(c net.Conn) {
	defer c.Close()
	uid := newUid()
	log.Info().Str("uid", uid).Msg("client connected")
	defer func() {
		log.Info().Str("uid", uid).Msg("client disconnected")
	}()
	c.Write([]byte(fmt.Sprintf("<uid:%s\n", uid)))
	r := bufio.NewReader(c)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			log.Error().Err(err).Msg("terminating connection")
			return
		}

		fmt.Print(line)
		c.Write([]byte(line + "\n"))
	}
}

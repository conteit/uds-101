/*
 * Copyright © 2022 Paolo Contessi
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

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

		log.Info().Str("uid", uid).Str("content", strings.ReplaceAll(line, "\n", "")).Msg("message received")
		c.Write([]byte(line + "\n"))
	}
}

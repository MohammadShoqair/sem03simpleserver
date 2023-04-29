package main

import (

	"fmt"
	"github.com/MohammadShoqair/funtemps/conv"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/MohammadShoqair/is105sem03/mycrypt"
)

func main() {

	var wg sync.WaitGroup

	server, err := net.Listen("tcp", "172.17.0.3:8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bundet til %s", server.Addr().String())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Println("før server.Accept() kallet")
			conn, err := server.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				for {
					buf := make([]byte, 1024)
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return // fra for løkke
					}
dekryptertMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
log.Println("Dekrypter melding: ", string(dekryptertMelding))

switch msg := string(dekryptertMelding); msg {
                                        case "ping":

							response1 := "pong"
							     encryptedResponse := mycrypt.Krypter([]rune(response1), mycrypt.ALF_SEM03, 4)
							_,err=c.Write([]byte(string(encryptedResponse)))
                                         case "Kjevik;SN39040;18.03.2022 01:50;6":
                                  		parts := strings.Split(msg, ";")

				            	  if len(parts) < 4 {
							log.Println("Invalid input message")
							return
						               }
                     				t, err := strconv.ParseFloat(parts[len(parts)-1], 64)
                     					if err != nil {
                 				        log.Println(err) }

				               f := conv.CelsiusToFarhenheit(t)
						response:= fmt.Sprintf("Kjevik;SN39040;18.03.2022 01:50;%0.2f", f)
						     encryptedResponse := mycrypt.Krypter([]rune(response), mycrypt.ALF_SEM03, 4)
							_,err = c.Write([]byte(string(encryptedResponse)))
					default:
						_, err = c.Write(buf[:n])
					}
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return // fra for løkke
					}
				}
			}(conn)
		}
	}()
	wg.Wait()
}

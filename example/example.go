package main

import (
	"flag"
	"log"
	"net"

	fgbgp "github.com/BasilFillan/fgbgp/server"
	server "github.com/BasilFillan/fgbgp/server"

	"github.com/BasilFillan/fgbgp/messages"
)

type Collector struct {
}

func (col *Collector) Notification(msg *messages.BGPMessageNotification, n *server.Neighbor) bool {
	log.Printf("%v", msg)
	return true
}

func (col *Collector) ProcessReceived(v interface{}, n *server.Neighbor) (bool, error) {
	log.Printf("%v", v)
	return true, nil
}

func (col *Collector) ProcessSend(v interface{}, n *server.Neighbor) (bool, error) {
	log.Printf("%v", v)
	return true, nil
}

func (col *Collector) ProcessUpdateEvent(e *messages.BGPMessageUpdate, n *server.Neighbor) (add bool) {
	log.Printf("%v", e)
	return true
}

func (col *Collector) DisconnectedNeighbor(n *server.Neighbor) {
	log.Printf("%v", n)

}

func (col *Collector) NewNeighbor(on *messages.BGPMessageOpen, n *fgbgp.Neighbor) bool {
	log.Printf("%v %v", on, n)
	return true
}

func (col *Collector) OpenSend(on *messages.BGPMessageOpen, n *fgbgp.Neighbor) bool {
	log.Printf("%v %v", on, n)
	return true
}

func (col *Collector) OpenConfirm() bool {
	log.Printf("OpenConfirm")
	return true
}

var BgpAddr = flag.String("addr", "[::]:179", "BGPAddr")

func main() {
	flag.Parse()
	m := server.NewManager(65001, net.ParseIP("10.0.0.1"), false, false)
	m.UseDefaultUpdateHandler(10)
	col := Collector{}
	m.SetEventHandler(&col)
	m.SetUpdateEventHandler(&col)
	err := m.NewServer(*BgpAddr)
	if err != nil {
		log.Fatal(err)
	}
	m.Start()
	// log.Println("Running")
	// exit := make(chan bool)
	// <-exit
}

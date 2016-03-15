package main

import (
    "net"
)

func handleConnection(conn *net.UDPConn) {
    var buffer [1200]byte
    defer conn.Close()
    for {
        i, a, e := conn.ReadFromUDP(buffer[:])
        
        if e != nil {
            return
        }
        
        i, e = conn.WriteToUDP(buffer[:i], a)
        if e != nil {
            return
        }
    }
}

func runUDPServer(addr *net.UDPAddr) {
    
    for {
        conn, err := net.ListenUDP("udp4", addr)
        
        if err != nil {
            continue
        }
        
        handleConnection(conn)
        
    }
    
}


func main() {
    a, e := net.ResolveUDPAddr("udp4", ":6666")
    
    if e != nil {
        panic("Failed to resolve port 6666")
    }
    
    runUDPServer(a)
}
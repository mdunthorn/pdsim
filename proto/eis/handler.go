/*
 * handler.go
 */
package handler

import (
	"io"
    "log"
	"net"
    "time"
)

func Handle(c net.Conn) {
    log.Print("start handling a connection from " + c.RemoteAddr().String())

    buf := make([]byte, 0, 4096)
    tmp := make([]byte, 256)
    bytes_read := 0
    for {
        for {
            n, err := c.Read(tmp)
            if err != nil {
                if err != io.EOF {
                    log.Print("read error: ", err)
                }
                break
            }
            log.Printf("read %d bytes", n)
            buf = append(buf, tmp[:n]...)
            bytes_read += n
            if (bytes_read >= 5) {
                break
            }
        }
        _, err := send_vol_long(c)
        if (err != nil) {
            break
        }
        send_volume(c)
        time.Sleep(10 * time.Millisecond)
        send_occupancy(c)
        time.Sleep(10 * time.Millisecond)
        send_vol_med_1(c)
        time.Sleep(10 * time.Millisecond)
        send_vol_xlong(c)
        time.Sleep(10 * time.Millisecond)
        send_speed(c)
        time.Sleep(10 * time.Millisecond)
    }
    log.Printf("stop handling a connection from %s", c.RemoteAddr().String())
}

func send_vol_long(c net.Conn) (int, error) {
    var buf[]byte
    buf = []byte{0xff, 0x1b, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
    written, err := c.Write(buf)
    if err != nil {
        log.Print("error writing")
        return -1, err
    } else {
        log.Printf("wrote %d bytes", written)
        return written, nil
    }
}

func send_volume(c net.Conn) {
    var buf[]byte
    buf = []byte{0xff, 0x10, 0x09, 0x0a, 0x0c, 0x0e, 0x10, 0x00, 0x00, 0x00, 0x00, 0x07, 0x3b}
    c.Write(buf)
    log.Printf("wrote %d bytes", len(buf))
}

func send_occupancy(c net.Conn) {
    var buf[]byte
    buf = []byte{0xff, 0x11, 0x09, 0x05, 0x06, 0x07, 0x08, 0x00, 0x00, 0x00, 0x00, 0x01, 0x1b}
    c.Write(buf)
    log.Printf("wrote %d bytes", len(buf))
}

func send_vol_med_1(c net.Conn) {
    var buf[]byte
    buf = []byte{0xff, 0x20, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
    c.Write(buf)
    log.Printf("wrote %d bytes", len(buf))
}

func send_vol_xlong(c net.Conn) {
    var buf[]byte
    buf = []byte{0xff, 0x36, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
    c.Write(buf)
    log.Printf("wrote %d bytes", len(buf))
}

func send_speed(c net.Conn) {
    var buf[]byte
    buf = []byte{0xff, 0x12, 0x0b, 0x5f, 0x64, 0x69, 0x6e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9a}
    c.Write(buf)
    log.Printf("wrote %d bytes", len(buf))
}


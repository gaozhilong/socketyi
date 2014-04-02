package socketyi

import (
	"net"
	"github.com/golang/glog"
	//"time"
    "errors"
    "fmt"
	//"bytes"
)

type Connection struct {
	conn net.Conn
	so *SocketYi
	id string
	queue chan interface{}
	receives chan interface{}
}

func newConnection(so * SocketYi, conn net.Conn) (con *Connection) {
	con = &Connection{
		conn : conn,
		so : so,
		id : CreatSessionID(),
		queue : make(chan interface{}, so.config.QueueLength),
		receives : make(chan interface{}, so.config.QueueLength),
	}
	return con
}

func (con *Connection) RemoteAddr() string {
	return con.conn.RemoteAddr().String()
}

func (con *Connection) Read() {
	buf := make([]byte, con.so.config.ReadBufferSize)
	 for {
        read_len, err := con.conn.Read(buf)
		if err != nil {
            glog.V(0).Infoln(con, "Receive")
            break
        }
		if read_len == 0 {
            break // connection already closed by client
        } else {
			con.Receive(buf[0:read_len])
		}
		buf = make([]byte, con.so.config.ReadBufferSize) // clear last read content
	 }
}

func (con *Connection) Send(data interface{}) error {
	select {
		case con.queue <- data:
		default:
			err := errors.New("send queue is full")
			return err
	}
	return nil
}

func (con *Connection) Receive(data interface{}) {
	select {
	case con.receives <- data:
	default:
		err := errors.New("receive queue is full")
		fmt.Println(err)
	}
}

func (con *Connection) Close() error {
	return nil
}

func (con *Connection) disconnect() {
}

func (con *Connection) flush() {
	var msg interface{}
	for msg = range con.queue {
		mss := msg.([]byte)
		con.conn.Write([]byte(mss))
	}
}


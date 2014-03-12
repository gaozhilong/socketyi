package socketyi

import (
	"net"
	"github.com/golang/glog"
	"time"
    //"strconv"
    "fmt"
)

type Connection struct {
	conn net.Conn
	so *SocketYi
	id string
}

func newConnection(so * SocketYi, conn net.Conn) (con *Connection) {
	con = &Connection{
		conn : conn,
		so : so,
		id : CreatSessionID(),	
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
        	fmt.Println("sdadsda")
            daytime := time.Now().String()
            con.conn.Write([]byte(daytime)) 
        }
		fmt.Println("ssssss")
        buf = make([]byte, con.so.config.ReadBufferSize) // clear last read content
     }
}

func (con *Connection) Send(data interface{}) error {
	return nil
}

func (con *Connection) Receive(data interface{}) {
}

func (con *Connection) Close() error {
	return nil
}

func (con *Connection) disconnect() {
}


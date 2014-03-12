package socketyi

import (
	"net"
	
)

type SocketYi struct {
	connections map[string]*Connection
	config Config
}

func NewSocketYi(config *Config) (server *SocketYi){
	if config == nil {
		config = &Default
	}
	server = &SocketYi{
		connections : make(map[string]*Connection),
		config : *config,
	}
	return server
}

func (so *SocketYi) Broadcast(data interface{}) {
	
}

func (so *SocketYi) Handle(con *Connection) {
	//conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
    //defer con.Close()  // close connection before exit
    con.Read()
}

func (so *SocketYi) ListenAndServe(addr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
    CheckError(so,err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    CheckError(so,err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        con := newConnection(so, conn)
        go so.Handle(con)
    }
}

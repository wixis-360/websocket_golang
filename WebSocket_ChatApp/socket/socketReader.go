package socket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// SocketReader struct
type socketReader struct {
	con  *websocket.Conn
	mode int
	name string
}

//upgrader struct
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// socketReader slice
var savedSocketReader []*socketReader

func SocketReaderCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("Socket request")
	if savedSocketReader == nil {
		savedSocketReader = make([]*socketReader, 0)
	}

	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		r.Body.Close()

	}()
	con, _ := upgrader.Upgrade(w, r, nil)

	//create new socketReader struct and assign in to a new variable
	ptrSocketReader := &socketReader{
		con: con,
	}
	//ptrSocketReader assign in to a savedSocketReader slice
	savedSocketReader = append(savedSocketReader, ptrSocketReader)

	ptrSocketReader.startThread()
}

func (i *socketReader) Broadcast(str string) {
	for _, g := range savedSocketReader {

		if g == i {
			// no send message to himself
			continue
		}

		if g.mode == 1 {
			// no send message to connected user before user write his name
			continue
		}
		g.WriteMsg(i.name, str)
	}
}

func (i *socketReader) Read() {
	_, b, er := i.con.ReadMessage()
	if er != nil {
		panic(er)
	}
	log.Println(i.name + " :" + string(b))
	log.Println(i.mode)
	// if user name exist
	if i.mode == 1 {
		i.name = string(b)
		i.WriteMsg("System", "Welcome "+i.name+ ", Please write a message and we will broadcast it to other users.")
		i.mode = 2 // real msg mode

		return
	}

	i.Broadcast(string(b))

	log.Println(i.name + " :" + string(b))
}

func (i *socketReader) WriteMsg(name string, str string) {
	i.con.WriteMessage(websocket.TextMessage, []byte("<b>"+name+" :</b> "+str))
}

func (i *socketReader) startThread() {
	i.WriteMsg("System", "Please write your name first, Otherwise you can't see messages from other users...")
	i.mode = 1 // get user name

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)
			}
			log.Println("Thread SocketReader Finish")
		}()

		for {
			i.Read()
		}

	}()
}

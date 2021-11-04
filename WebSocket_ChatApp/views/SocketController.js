class SocketController {
    constructor(){
        this.mysocket =  null;
        // get the value of div element in html page
        this.vMsgContainer = document.getElementById("msgContainer");
        // get the value of input field html page
        this.vMsgInput = document.getElementById("input");
    }

    showMessage(text, myself){
        // create new div element in html page
        let div = document.createElement("div");
        div.innerHTML = text;
        let self = (myself)? "self" : "";
        div.className="msg " + self;
        this.vMsgContainer.appendChild(div);
    }

    send(){
        let txt = this.vMsgInput.value;
        // set txt value to the showMessage method
        this.showMessage("<b>Me :</b> " + txt,true);
        this.mysocket.send(txt);
        this.vMsgInput.value = ""
    }

    keypress(e){
        if (e.keyCode == 13) {
            this.send();
        }
    }

    connectSocket(){
        console.log("Socket");
        //make sure the port matches with your golang code
        let socket = new WebSocket("ws://localhost:8080/socket");
        this.mysocket = socket;

        socket.onmessage = (e)=>{
            this.showMessage(e.data,false);

        }
        socket.onopen =  ()=> {
            console.log("Socket Opened...")
        };
        socket.onclose = ()=>{
            console.log("Socket Closed...")
        }
    }
}

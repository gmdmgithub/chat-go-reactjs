// api/index.js
const socket = new WebSocket("ws://localhost:8080/ws");

let userId = null;

let connect = cb => {
    console.log("connecting");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
        const data = JSON.parse(msg.data)
        console.log("Message is: ",data);
        if(data.usr_id){
            userId = data.usr_id;
            console.log(`Get client ID as ${userId}`);
            
        }
        cb(msg);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event, userId);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error, userId);
    };
};

let sendMsg = (msg, id) => {

    const message = {
        value: msg,
        ID: id
    }
    console.log("sending msg: ", JSON.stringify(message));
    socket.send(JSON.stringify(message));
};

export {
    connect,
    sendMsg,
    userId
};
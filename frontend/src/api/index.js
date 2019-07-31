// api/index.js
const socket = new WebSocket("ws://localhost:8080/ws");

const connect = cb => {
    console.log("connecting ...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
        const data = JSON.parse(msg.data);
        console.log("Message is: ", data);
        cb(msg);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

const sendMsg = (msg, id) => {

    const message = {
        value: msg,
        ID: id
    };
    console.log("sending msg: ", JSON.stringify(message));
    socket.send(JSON.stringify(message));
};

export {
    connect,
    sendMsg
};
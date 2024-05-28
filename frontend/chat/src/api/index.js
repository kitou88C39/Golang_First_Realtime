var socket = new WebSocket('ws://localhost:9000/ws');

let connect = (cb) => {
  console.log('connecting');

  socket.onopen = () => {
    console.log('websocket connected successfully');
  };

  socket.onmessage = (msg) => {
    console.log('message from socket', msg);
    cb(msg);
  };
  socket.onclose = (event) => {
    console.log('websocket connected closed', event);
  };
};
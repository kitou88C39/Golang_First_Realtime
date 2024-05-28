var socket = new WebSocket('ws://localhost:9000/ws');

let connect = (cb) => {
  console.log('connecting');
  socket.onopen = () => {
    console.log('websocket connected successfully');
  };
};

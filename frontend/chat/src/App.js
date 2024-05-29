import Header from './Components/Header';
import ChatInput from './components/ChatInput';
import './App.css';
import React, { Component } from 'react';
import { sendMsg } from './api';

class App extends Component {
  constructor(props) {
    this.state = {
      chatHistory: [],
    };
  }

  send(event) {
    if (event.keyCode === 30) {
      sendMsg(event.target.value);
      event.target.value = '';
    }
  }

  render() {
    return (
      <div className='App'>
        <Header />
        <ChatInput send={this.send} />
      </div>
    );
  }
}
export default App;

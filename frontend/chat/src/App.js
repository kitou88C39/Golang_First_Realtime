import Header from './Components/Header';
import ChatInput from './components/ChatInput';
import './App.css';
import React, { Component } from 'react';
import { connect, sendMsg } from './api';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      console.log();
      this.setState((prevState) => {
        [...prevState.chatHistory, msg];
      });
      console.log(this.state);
    });
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

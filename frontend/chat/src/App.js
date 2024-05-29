import Header from './Components/Header';
import ChatInput from './components/ChatInput';
import './App.css';
import React, { Component } from 'react';

class App extends Component {
  constructor(props){
    this.state = {
    chatHistory:[],
    }
  }
  return (
    <div className='App'>
      <Header />
      <ChatInput send={this.send} />
    </div>
  );
}

export default App;

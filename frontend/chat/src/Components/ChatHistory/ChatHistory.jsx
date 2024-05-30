import React, { Component } from 'react';
import Message from '../Message';

class ChatHistory extends Component {
  render() {
    const messages = this.props.ChatHistory.map((msg) => (
      <Message key={msg.timeStamp} message={msg.data} />
    ));
    return <div className='ChatHistory'></div>;
  }
}

export default ChatHistory;

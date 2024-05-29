import React, { Component } from 'react';

class ChatHistory extends Component {
    render() {
        return (
          <div className='ChatInput'>
            <input onKeyDown={this.props.send} placeholder='Enter a message...' />
          </div>
        );
};

export default ChatHistory;

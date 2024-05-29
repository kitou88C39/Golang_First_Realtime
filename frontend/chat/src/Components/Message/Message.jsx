import React, { Component } from 'react';

class Message extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    return <div className='Message'>{this.state.message.body}</div>;
  }
}

export default Message;

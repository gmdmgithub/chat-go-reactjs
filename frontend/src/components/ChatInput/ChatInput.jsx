import React from 'react';
import { Component } from 'react';
import "./ChatInput.css";

class ChatInput extends Component {
    render() {
      return (
        <div className="ChatInput">
          <input onKeyDown={this.props.send} placeholder="Type a message and hit enter"/>
        </div>
      );
    }
  }
  
  export default ChatInput;
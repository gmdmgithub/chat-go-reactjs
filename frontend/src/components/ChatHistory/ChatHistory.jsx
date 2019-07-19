import React, { Component } from "react";
import "./ChatHistory.css";
import Message from "../Message/Message";

class ChatHistory extends Component {
  render() {
    // const messages = this.props.chatHistory.map((msg, index) => (
    //   <p key={index}> <Message message={msg.data} /></p>
    // ));
    console.log(this.props.chatHistory);
    const messages = this.props.chatHistory.map((msg, index) => <Message key={index} message={msg.data} />);

    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {messages}
      </div>
    );
  }
}

export default ChatHistory;
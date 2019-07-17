import React, {Component } from 'react';
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';

class App extends Component {

  state = {
    chatHistory:[
      {
        data:"how are you"
      }
    ]
  };

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state);
    });
  }

  send() {
    console.log("Sending hello to the backend");
    sendMsg("hello - react frontend");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button className="appButton" onClick={this.send}>Send message!</button>
      </div>
    );
  }
}

export default App;
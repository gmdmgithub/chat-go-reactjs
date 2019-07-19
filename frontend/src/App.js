import React, {Component } from 'react';
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';

class App extends Component {

  state = {
    chatHistory:[
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

  send(event) {
    if(event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }
  sendOnClick(){
    //get a user data
    sendMsg("Hi there")
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
        <div className="AppBtn">
          <button className="appButton" onClick={this.sendOnClick}> Send some message!</button>
        </div>
      </div>
    );
  }
}

export default App;
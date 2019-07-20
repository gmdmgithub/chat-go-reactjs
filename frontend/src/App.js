import React, {Component } from 'react';
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory:[],
      userId: null
    };
    // just define send(event) as send = (event)=>.... bind not needed
    // this.send = this.send.bind(this);
  }

  // special method for react after component is rendered - opposite method is: componentWillUnmount(){}
  componentDidMount() {
    connect((msg) => {
      const data = JSON.parse(msg.data)
      console.log("componentDidMount - New Message", data.usr_id)
      this.userId = data.usr_id;
      
      this.setState(state => ({
        chatHistory: [...state.chatHistory, msg],
        userId:data.usr_id
      }))
      console.log("componentDidMount - current state is",(this.state));
    });
  }

  send = (event) => {
    if(event.keyCode === 13) {
      console.log("send - current state is",(this.state));
      console.log("Inspected value", event.target.value);
      
      sendMsg(event.target.value, this.state.userId);
      event.target.value = "";
    }
  }
  sendOnClick = ()=>{
    //get a user data
    sendMsg("Hi there", this.state.userId)
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
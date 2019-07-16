import React, {Component } from 'react';
import "./App.css";
import { connect, sendMsg } from "./api";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("Sending hello to the backend");
    sendMsg("hello - react frontend");
  }

  render() {
    return (
      <div className="App">
        <button className="appButton" onClick={this.send}>Send message!</button>
      </div>
    );
  }
}

export default App;
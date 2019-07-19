import React, { Component } from "react";

import "./Message.css";

class Message extends Component {
    constructor(props) {
      super(props);
      
      let temp = '';
      
      if(this.props.message){
        try{
        temp = JSON.parse(this.props.message);
        }catch(e){
            console.log(`Exception not a JSON message: ${e}`);
        }
      }
      this.state = {
        message: temp
      };
    }
  
    render() {
      return <div className="Message">{this.state.message.body}</div>;
    }
  }
  
  export default Message;
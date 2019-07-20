import React, { Component } from "react";

import "./Message.css";

class Message extends Component {
    constructor(props) {
        super(props);

        let temp = '';

        if (this.props.message) {
            console.log("My message to parse is", this.props.message);

            try {
                temp = JSON.parse(this.props.message);
                console.log("My message after parse is", this.props.message);
            } catch (e) {
                console.log(`Exception not a JSON message: ${e}`);
            }
        }
        this.state = {
            message: temp
        };
    }

    render() {
        return <div className="Message">{this.state.message.body} <span className="me"> from user ({this.state.message.usr_id})</span></div>;
    }
}

export default Message;
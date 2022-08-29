import React, { Component } from "react";

export default class AppHeader extends Component {
  render() {
    return (
      <div>
        <h1>{this.props.title}</h1>
        {/* <h1>{this.props.subject}</h1>
        <h1>{this.props.favourite_color}</h1> */}
      </div>
    );
  }
}

import React, { Component, Fragment } from "react";

export default class AppFooter extends Component {
  render() {
    // 現在の年を取得
    const currentYear = new Date().getFullYear();
    return (
      <Fragment>
        <hr />
        <p>Copyrigth &copy; 2020 - {currentYear} Acme Ltd.</p>
      </Fragment>
    );
  }
}

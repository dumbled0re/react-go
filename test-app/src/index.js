import React, { Component } from "react";
import ReactDOM from "react-dom";
import AppHeader from "./AppHeader";
import AppFooter from "./AppFooter";
import AppContent from "./AppContent";

// css
import "bootstrap/dist/css/bootstrap.min.css";
// js
import "bootstrap/dist/js/bootstrap.bundle.min.js";
// index.cssの読み込み
import "./index.css";

class App extends Component {
  render() {
    const myProps = {
      title: "My Cool App!",
      subject: "My subject",
      favourite_color: "red",
    };
    return (
      <div className="app">
        {/* myPropsを展開 */}
        <AppHeader {...myProps} />
        <AppContent />
        <AppFooter />
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById("root"));

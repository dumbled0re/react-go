import React, { Component } from "react";

export default class AppContent extends Component {
  anotherFunction = () => {
    console.log("another function");
  };

  leftParagraph = () => {
    console.log("left the paragraph");
  };

  fetchList = () => {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => response.json())
      .then((json) => {
        console.log(json);
        let posts = document.getElementById("post-list");

        json.forEach(function (obj) {
          let li = document.createElement("li");
          li.appendChild(document.createTextNode(obj.title));
          posts.appendChild(li);
        });
      });
  };

  render() {
    return (
      <p>
        This is the content.
        <hr />
        <br />
        <p
          // マウスが当たったら
          onMouseEnter={this.anotherFunction}
          // マウスが離れたら
          onMouseLeave={this.leftParagraph}
        >
          This is some text
        </p>
        {/* ボタンがクリックされたらfetchListを実行 */}
        <button onClick={this.fetchList} className="btn btn-primary">
          Fetch Data
        </button>
        <hr />
        <ul id="post-list"></ul>
      </p>
    );
  }
}

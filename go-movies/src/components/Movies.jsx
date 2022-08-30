import React, { Component, Fragment } from "react";
import { Link } from "react-router-dom";

export default class Movies extends Component {
  state = { movies: [] };

  // コンポーネントがマウント(配置)された直前に呼び出されるメソッド
  componentDidMount() {
    this.setState({
      movies: [
        { id: 1, title: "The Shawshank Redemption", runtime: 142 },
        { id: 2, title: "The Godfather", runtime: 175 },
        { id: 3, title: "The Dark Knight", runtime: 153 },
      ],
    });
  }

  render() {
    return (
      <Fragment>
        <h2>Choose a movie</h2>
        <ul>
          {this.state.movies.map((m) => (
            <li key={m.id}>
              <Link to={`/movies/${m.id}`}>{m.title}</Link>
            </li>
          ))}
        </ul>
      </Fragment>
    );
  }
}
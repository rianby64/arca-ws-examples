import React, { Component } from 'react';
import 'semantic-ui-css/semantic.min.css';
import TableArca from './Table';

class App extends Component {
  render() {
    return (
      <div className="App">
        <TableArca></TableArca>
      </div>
    );
  }
}

new WebSocket("ws://" + document.location.host + "/arca-node");

export default App;

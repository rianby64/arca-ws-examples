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

void function() {
  const arca = new WebSocket("ws://" + document.location.host + "/arca-node");
  arca.onopen = () => {
    arca.send(JSON.stringify({
      Method: 'read',
      ID: `id-for-request-ViewTable1`,
      Context: {
          Source: 'ViewTable1'
      }
    }));
    arca.onmessage = (e) => {
      console.log(e.data);
    }
  }
}();

export default App;

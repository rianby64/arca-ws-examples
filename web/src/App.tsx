import React, { Component } from 'react';
import './App.css';
import TableArca from './Table';

class App extends Component {
  state: {
    tables: any,
    ViewTable1: any[],
    ViewTable2: any[],
  };

  private ws: WebSocket;
  private requests: any[];
  private wsConnected: boolean;

  constructor(props: any) {
    super(props);
    this.state = {
      tables: [
        {
          source: 'ViewTable1',
          rows: [],
          fields: ['Num1', 'Num2', 'I'],
          headers: ['ID', 'Num1', 'Num2', 'I'],
        },
        {
          source: 'ViewTable2',
          rows: [],
          fields: ['Num3', 'Num4', 'I'],
          headers: ['ID', 'Num3', 'Num4', 'I'],
        },
      ],
      ViewTable1: [],
      ViewTable2: [],
    };
    this.requests = [];
    this.wsConnected = false;

    this.ws = new WebSocket("ws://" + document.location.host + "/arca-node");
    this.ws.onopen = () => {
      this.wsConnected = true;
      this.ws.onmessage = (e) => {
        const response = JSON.parse(e.data);
        const source = response.Context.Source;
        if (response.Method == 'read') {
          const rows = response.Result;
          this.setState((state: any) => {
            return {
              ...state,
              [source]: rows,
            }
          });
        } else if (response.Method == 'update') {
          const rowUpdated = response.Result;
          this.setState((state: any) => {
            const rows = state[source];
            if (!rows) return state;
            return {
              ...state,
              [source]: rows.map((row: any) => {
                if (row.ID == rowUpdated.ID) {
                  return rowUpdated;
                }
                return row;
              }),
            };
          });
        } else if (response.Method == 'insert') {
          const rowInserted = response.Result;
          this.setState((state: any) => {
            const rows = state[source];
            if (!rows) return state;
            return {
              ...state,
              [source]: [...rows, rowInserted]
            };
          });
        } else if (response.Method == 'delete') {
          const rowDeleted = response.Result;
          this.setState((state: any) => {
            const rows = state[source];
            if (!rows) return state;
            return {
              ...state,
              [source]: rows.filter((row: any) => {
                if (row.ID !== rowDeleted.ID) {
                  return row;
                }
              })
            };
          });
        }
      };
      this.requests.forEach(request => {
        this.sendRequest(request);
      });
      this.requests.length = 0;
    };
  };

  sendRequest = (request: any) => {
    if (this.wsConnected) {
      this.ws.send(JSON.stringify(request));
    } else {
      this.requests.push(request);
    }
  }

  render() {
    return (
      <div className="App">
        <TableArca
          rows={this.state.ViewTable1}
          headers={this.state.tables[0].headers}
          fields={this.state.tables[0].fields}
          send={this.sendRequest}
          source={this.state.tables[0].source}
        />
        <TableArca
          rows={this.state.ViewTable2}
          headers={this.state.tables[1].headers}
          fields={this.state.tables[1].fields}
          send={this.sendRequest}
          source={this.state.tables[1].source}
        />
      </div>
    );
  };
}

export default App;

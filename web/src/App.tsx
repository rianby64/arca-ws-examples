import React, { Component } from 'react';
import './App.css';
import TableArca from './Table';
import Buttom from './Button';
import Modal from './Modal';

class App extends Component {
  state: {
    modal: boolean,
    request: any,
    ViewTable1: any[],
  };

  private ws: WebSocket;
  private requests: any[];
  private wsConnected: boolean;

  constructor(props: any) {
    super(props);
    this.state = {
      modal: false,
      request: {},
      ViewTable1: [],
    };
    this.requests = [];
    this.wsConnected = false;

    this.modalSwitcher = this.modalSwitcher.bind(this);
    this.getRequestMethod = this.getRequestMethod.bind(this);
    this.sendRequest = this.sendRequest.bind(this);

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
          };
      };
      this.requests.forEach(request => {
        this.sendRequest(request);
      });
      this.requests.length = 0;
    };
  };

  sendRequest(request: any) {
    if (this.wsConnected) {
      this.ws.send(JSON.stringify(request));
    } else {
      this.requests.push(request);
    }
  }

  modalSwitcher() {
    this.setState((state: any) => {
      const newState = {
        modal: !state.modal,
      };

      return newState;
    });
  };

  getRequestMethod(method: object) {
    this.setState(() => {
      const newState = {
        request: method,
      };

      return newState;
    });
  };

  render() {
    return (
      <div className="App">
        <Buttom openModal={this.modalSwitcher} />
        <TableArca
          getRequestMethod={this.getRequestMethod}
          rows={this.state.ViewTable1}
          headers={['ID', 'Num1', 'Num2', 'I']}
          fields={['Num1', 'Num2', 'I']}
          send={this.sendRequest}
          source="ViewTable1"
        />
        { this.state.modal ?
          <Modal
            closeModal={this.modalSwitcher}
            request={this.state.request}
          /> :
          null
        }
      </div>
    );
  };
}

export default App;

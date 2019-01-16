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
        if (response.Context.Source == 'ViewTable1') {
          if (response.Method == 'read') {
            const ViewTable1 = response.Result;
            this.setState((state: any) => {
              return {
                ...state,
                ViewTable1,
              }
            });
          } else if (response.Method == 'update') {
            const rowUpdated = response.Result;
            this.setState((state: any) => {
              const ViewTable1 = state.ViewTable1.map((row: any) => {
                if (row.ID == rowUpdated.ID) {
                  return rowUpdated;
                }
                return row;
              });
              return {
                ...state,
                ViewTable1,
              };
            });
          } else if (response.Method == 'insert') {
            const rowInserted = response.Result;
            this.setState((state: any) => {
              const { ViewTable1 } = state;
              return {
                ...state,
                ViewTable1: [...ViewTable1, rowInserted]
              };
            });
          } else if (response.Method == 'delete') {
            const rowDeleted = response.Result;
            this.setState((state: any) => {
              const ViewTable1 = state.ViewTable1.filter((row: any) => {
                if (row.ID !== rowDeleted.ID) {
                  return row;
                }
              });
              return {
                ...state,
                ViewTable1
              };
            });
          }
        }
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
          send={this.sendRequest}
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

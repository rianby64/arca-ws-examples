import React, { Component } from 'react';
import './App.css';
import TableArca from './Table';
import Buttom from './Button';
import Modal from './Modal';

class App extends Component {
  state: {
    modal: boolean,
    request: object,
  }

  constructor(props: any) {
    super(props);
    this.state = {
      modal: false,
      request: {},
    };

    this.modalSwitcher = this.modalSwitcher.bind(this);
    this.getRequestMethod = this.getRequestMethod.bind(this);
  }

  modalSwitcher() {
    this.setState((state: any) => {
      const newState = {
        modal: !state.modal,
      };

      return newState;
    });
  }

  getRequestMethod(method: object) {
    this.setState(() => {
      const newState = {
        request: method,
      };

      return newState;
    });
  }

  render() {
    return (
      <div className="App">
        <Buttom openModal={this.modalSwitcher} />
        <TableArca getRequestMethod={this.getRequestMethod} ></TableArca>
        { this.state.modal ? <Modal closeModal={this.modalSwitcher} request={this.state.request} /> : null }
      </div>
    );
  }
}

export default App;

import React, { Component, ChangeEvent, FormEvent } from 'react';

class Modal extends Component<any> {
  state: {
    Num1: any,
    Num2: any,
    I: any,
  }

  constructor(props: any) {
    super(props);

    this.state = {
      Num1: '',
      Num2: '',
      I: '',
    }

    this.createRow = this.createRow.bind(this);
    this.cancelCreating = this.cancelCreating.bind(this);
    this.changeNum1 = this.changeNum1.bind(this);
    this.changeNum2 = this.changeNum2.bind(this);
    this.changeI = this.changeI.bind(this);
  }

  createRow(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const { Num1, Num2, I } = this.state;

    const data = {
      Num1: +Num1,
      Num2: +Num2,
      I: +I,
    }

    this.props.request(data);
    this.props.closeModal();
  }

  cancelCreating() {
    this.props.closeModal();
  }

  changeNum1(event: ChangeEvent<HTMLInputElement>) {
    const newState = {
      ...this.state,
      Num1: event.target.value
    };
    this.setState(newState);
  }

  changeNum2(event: ChangeEvent<HTMLInputElement>) {
    const newState = {
      ...this.state,
      Num2: event.target.value
    };
    this.setState(newState);
  }

  changeI(event: ChangeEvent<HTMLInputElement>) {
    const newState = {
      ...this.state,
      I: event.target.value
    };
    this.setState(newState);
  }

  render() {
    return (
      <div className="modal">
        <form onSubmit={this.createRow} autoComplete="off">
          <h2>Creating new row</h2>
          <label htmlFor="num1">Num1</label>
          <input id="num1" value={this.state.Num1} onChange={this.changeNum1} />
          <label htmlFor="num2">Num2</label>
          <input id="num2" value={this.state.Num2} onChange={this.changeNum2} />
          <label htmlFor="I">I</label>
          <input type="text" id="I" value={this.state.I} onChange={this.changeI} />
          <button type="submit">submit</button>
          <button type="button" onClick={this.cancelCreating}>close</button>
        </form>
      </div>
    );
  }
}

export default Modal;

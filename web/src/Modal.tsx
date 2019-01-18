import React, { Component, FormEvent } from 'react';

class Modal extends Component<any> {
  state: {
    first: any,
    second: any,
    third: any,
  }

  constructor(props: any) {
    super(props);

    this.state = {
      first: '',
      second: '',
      third: '',
    }

    this.createRow = this.createRow.bind(this);
    this.cancelCreating = this.cancelCreating.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  createRow(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const { first, second, third } = this.state;
    const currentTable = this.props.tablesData.filter((table: any) => {
      if (table.source === this.props.table) {
        return table;
      }
    })
    const fields = currentTable[0].fields;

    const data = {
      [fields[0]]: +first,
      [fields[1]]: +second,
      [fields[2]]: +third,
    }

    this.props.request(data, this.props.table);
    this.props.closeModal('');
  }

  cancelCreating() {
    this.props.closeModal('');
  }

  handleChange(event: any) {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  render() {
    const currentTable = this.props.tablesData.filter((table: any) => {
      if (table.source === this.props.table) {
        return table;
      }
    })
    const fields = currentTable[0].fields;

    return (
      <div className="modal">
      <form onSubmit={this.createRow} autoComplete="off">
        <h2>{`Creating row in ${this.props.table}`}</h2>
        <label htmlFor="first">{fields[0]}</label>
        <input name="first" value={this.state.first} onChange={this.handleChange} />
        <label htmlFor="second">{fields[1]}</label>
        <input name="second" value={this.state.second} onChange={this.handleChange} />
        <label htmlFor="third">{fields[2]}</label>
        <input name="third" value={this.state.third} onChange={this.handleChange} />
        <button type="submit">submit</button>
        <button type="button" onClick={this.cancelCreating}>close</button>
      </form>
    </div>
    );
  }
}

export default Modal;

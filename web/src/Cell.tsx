import React, { Component, ChangeEvent, FormEvent } from 'react'

class CellArca extends Component<any> {
  state: {
    edit: boolean,
    cell: {
      value: any,
      field: string,
      ID: string,
    }
  }

  constructor(props: any) {
    super(props);
    this.state = {
      edit: false,
      cell: {
        value: this.props.value,
        ID: this.props.ID,
        field: this.props.field,
      }
    };

    this.beginRedact = this.beginRedact.bind(this);
    this.redact = this.redact.bind(this);
    this.endRedact = this.endRedact.bind(this);

    this.submit = this.submit.bind(this);
  }
  beginRedact() {
    this.setState({
      edit: true,
    });
  }

  redact(e: ChangeEvent<HTMLInputElement>) {
    const newState = {
      cell: {
        ...this.state.cell,
        value: e.target.value
      }
    };
    this.setState(newState);
  }

  endRedact() {
    this.setState({
      edit: false,
    });
  }

  submit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    this.props.onRedact(this.state.cell);
    this.endRedact();
  }

  render() {
    const { cell: { field, value }, edit } = this.state;
    return edit ?
    (
      <form onSubmit={this.submit}>
        <input
          onChange={this.redact}
          name={field}
          value={value}
          placeholder={`${field}...`} />
      </form>
    ) :
    (
      <span onClick={this.beginRedact}>{this.props.value}</span>
    );
  }
}

export default CellArca;

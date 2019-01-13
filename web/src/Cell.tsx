import React, { Component, ChangeEvent, FormEvent } from 'react'
import { Form } from 'semantic-ui-react'

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
    this.submit = this.submit.bind(this);
    this.redact = this.redact.bind(this);
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

  submit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    this.setState({
      edit: false,
    });

    console.log(this.state.cell, 'want to update');
  }

  render() {
    const { cell: { field, value }, edit } = this.state;
    return edit ?
      (
      <Form onSubmit={this.submit}>
        <Form.Input
          onChange={this.redact}
          name={field}
          value={value}
          fluid placeholder={`${field}...`} />
      </Form>
      ) :
      (
      <span onClick={this.beginRedact}>{value}</span>
      );
  }
}

export default CellArca;

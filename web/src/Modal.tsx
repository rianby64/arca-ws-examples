import React, { Component, FormEvent } from 'react';

class Modal extends Component<any> {
  state: {
    modal: boolean,
    row: any,
  }

  constructor(props: any) {
    super(props);
    const { fields } = this.props;

    const row = fields.reduce((acc: any, field: string) => {
      const value = this.props.row[field];
      acc[field] = value ? value : 0;
      return acc;
    }, {});

    this.state = {
      modal: false,
      row,
    };
  }

  onSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const { row } = this.state;
    const { fields } = this.props;

    const newRow = fields.reduce((acc: any, field: string) => {
      const value = row[field];
      acc[field] = value ? +value : 0;
      return acc;
    }, {});

    this.props.onSubmit(newRow);
    this.close();
  }

  close = () => {
    this.setState((state) => {
      return {
        ...state,
        modal: false,
      }
    });
  }

  handleChange = (event: any) => {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    this.setState((state: any) => {
      return {
        ...state,
        row: {
          ...state.row,
          [name]: value,
        }
      };
    });
  }

  open = () => {
    this.setState((state: any) => {
      return {
        ...state,
        modal: true,
      }
    });
  }

  render() {
    const { fields } = this.props;
    return (
      <div>
        <button onClick={this.open}>Create row</button>
        { this.state.modal ?
          (<div className="modal">
            <form onSubmit={this.onSubmit} autoComplete="off">
              <h2>{`Creating row in ${this.props.title}`}</h2>
              { fields.map((field: string, key: number) => (
                <label key={key}>{field}
                  <input name={field} value={this.state.row[field]} onChange={this.handleChange} />
                </label>
              ))}
              <button type="submit">submit</button>
              <button type="button" onClick={this.close}>close</button>
            </form>
          </div>) :
          null
        }
      </div>
    );
  }
}

export default Modal;

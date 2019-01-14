import React, { Component } from 'react'
import CellArca from './Cell';

export interface IRowArca {
  ID: string,
  Num1: number,
  Num2: number,
  I: number
};

class RowArca extends Component<any> {
  constructor(props: any) {
    super(props);

    this.onDelete = this.onDelete.bind(this);
  }

  onDelete() {
    const { row } = this.props;
    this.props.onDelete(row);
  }

  render() {
    const { row, onRedact } = this.props;
    return (
      <tr>
        <td>{row.ID}</td>
        <td>
          <CellArca
            onRedact={onRedact}
            ID={row.ID}
            value={row.Num1}
            field='Num1'
          />
        </td>
        <td>
          <CellArca
            onRedact={onRedact}
            ID={row.ID}
            value={row.Num2}
            field='Num2'
          />
        </td>
        <td>
          <CellArca
            onRedact={onRedact}
            ID={row.ID}
            value={row.I}
            field='I'
          />
        </td>
        <td>
          {
            this.props.onDelete ?
            (
              <button onClick={this.onDelete} >X</button>
            ) :
            null
          }
        </td>
      </tr>
    );
  }
}

export default RowArca;

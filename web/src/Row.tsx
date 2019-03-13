import React, { Component } from 'react'
import CellArca from './Cell';

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
    const { fields, row, onRedact } = this.props;
    return (
      <tr>
        <td>{row.ID}</td>
        { fields.map((field: any, key: number) => (
          <td key={key}>
            <CellArca
              onRedact={onRedact}
              ID={row.ID}
              value={row[field]}
              field={field}
            />
          </td>))
        }
        <td>
          {
            this.props.onDelete ?
            (
              <button onClick={this.onDelete}>X</button>
            ) :
            null
          }
        </td>
      </tr>
    );
  }
}

export default RowArca;

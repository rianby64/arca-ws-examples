import React, { Component } from 'react'
import RowArca, { IRowArca } from './Row';

class TableArca extends Component<any> {
  constructor(props: any) {
    super(props);
  }

  componentDidMount() {
    this.props.getRequestMethod(this.createRow);
    this.readRows();
  }

  readRows = () => {
    const request = {
      Method: 'read',
      ID: `id-for-request-ViewTable1`,
      Context: {
        Source: 'ViewTable1',
      },
    };
    this.props.send(request);
  }

  createRow = (row: any) => {
    const request = {
      Method: 'insert',
      Context: {
        Source: 'ViewTable1',
      },
      Params: {...row, ID: ''},
    };
    this.props.send(request);
  }

  updateRow = (cell: any) => {
    const request = {
      Method: 'update',
      Context: {
        Source: 'ViewTable1',
      },
      Params: {
        ID: cell.ID,
        [cell.field]: Number(cell.value),
      },
    };
    this.props.send(request);
  }

  deleteRow = (cell: any) => {
    const request = {
      Method: 'delete',
      Context: {
        Source: 'ViewTable1',
      },
      Params: {
        ID: cell.ID,
      },
    };
    this.props.send(request);
  }

  render() {
    const { rows } = this.props;
    return (
    <table>
      <thead>
        <tr>
          <td>ID</td>
          <td>Num1</td>
          <td>Num2</td>
          <td>I</td>
          <td></td>
        </tr>
      </thead>

      <tbody>
        {
          rows.map((row: any) =>
            <RowArca
              key={row.ID}
              row={row}
              onRedact={this.updateRow}
              onDelete={this.deleteRow}
            />
          )
        }
      </tbody>
    </table>
    );
  }
}

export default TableArca;

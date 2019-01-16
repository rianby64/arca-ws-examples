import React, { Component } from 'react'
import RowArca from './Row';

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
      ID: `id-for-request`,
      Context: {
        Source: this.props.source,
      },
    };
    this.props.send(request);
  }

  createRow = (row: any) => {
    const request = {
      Method: 'insert',
      Context: {
        Source: this.props.source,
      },
      Params: {...row, ID: ''},
    };
    this.props.send(request);
  }

  updateRow = (cell: any) => {
    const request = {
      Method: 'update',
      Context: {
        Source: this.props.source,
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
        Source: this.props.source,
      },
      Params: {
        ID: cell.ID,
      },
    };
    this.props.send(request);
  }

  render() {
    const { rows, headers } = this.props;
    return (
    <table>
      <thead>
        <tr>
          { headers.map((header: any, index: number) => (
              <td key={index}>{header}</td>
            ))
          }
          <td></td>
        </tr>
      </thead>

      <tbody>
        {
          rows.map((row: any) =>
            <RowArca
              fields={this.props.fields}
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

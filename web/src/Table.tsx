import React, { Component } from 'react'
import RowArca from './Row';
import Modal from './Modal';

class TableArca extends Component<any> {
  state = {
    newRow: {},
  };

  constructor(props: any) {
    super(props);
    this.state = {
      newRow: {},
    };
  }

  componentDidMount() {
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
    <div>
      <Modal
        title={this.props.source}
        fields={this.props.fields}
        row={this.state.newRow}
        onSubmit={this.createRow}
      />
      <table>
        <thead>
          <tr>
            { headers.map((header: any, index: number) => (
                <th key={index}>{header}</th>
              ))
            }
            <th></th>
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
    </div>
    );
  }
}

export default TableArca;

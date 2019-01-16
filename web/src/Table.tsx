import React, { Component } from 'react'
import RowArca, { IRowArca } from './Row';

class TableArca extends Component<any> {
  private ws: WebSocket;
  state: {
    rows: IRowArca[]
  };

  constructor(props: any) {
    super(props);

    this.state = {
      rows: [],
    };

    this.createRow = this.createRow.bind(this);
    this.updateRow = this.updateRow.bind(this);
    this.deleteRow = this.deleteRow.bind(this);

    this.ws = new WebSocket("ws://" + document.location.host + "/arca-node");
    this.ws.onopen = () => {
      this.ws.send(JSON.stringify({
        Method: 'read',
        ID: `id-for-request-ViewTable1`,
        Context: {
          Source: 'ViewTable1',
        },
      }));
      this.ws.onmessage = (e) => {
        const response = JSON.parse(e.data);
        if (response.Context.Source == 'ViewTable1') {
          if (response.Method == 'read') {
            const rows = response.Result;
            this.setState({ rows });
          } else if (response.Method == 'update') {
            const rowUpdated = response.Result;
            this.setState((state: any) => {
              const rows = state.rows.map((row: any) => {
                if (row.ID == rowUpdated.ID) {
                  return rowUpdated;
                }
                return row;
              });
              return { rows };
            });
          } else if (response.Method == 'insert') {
            const rowInserted = response.Result;
            this.setState((state: any) => {
              const rows = state.rows;
              rows.push(rowInserted);
              return { rows };
            });
          } else if (response.Method == 'delete') {
            const rowDeleted = response.Result;
            this.setState((state: any) => {
              const rows = state.rows.filter((row: any) => {
                if (row.ID !== rowDeleted.ID) {
                  return row;
                }
              });
              return { rows };
            });
          }
        }
      }
    }
  }

  componentDidMount() {
    this.props.getRequestMethod(this.createRow);
  }

  createRow(row: any) {
    const request = {
      Method: 'insert',
      Context: {
        Source: 'ViewTable1',
      },
      Params: {...row, ID: ''},
    };
    this.ws.send(JSON.stringify(request));
  }

  updateRow(cell: any) {
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
    this.ws.send(JSON.stringify(request));
  }

  deleteRow(cell: any) {
    const request = {
      Method: 'delete',
      Context: {
        Source: 'ViewTable1',
      },
      Params: {
        ID: cell.ID,
      },
    };
    this.ws.send(JSON.stringify(request));
  }

  render() {
    const { rows } = this.state;
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
          rows.map(row =>
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

import React, { Component } from 'react'
import RowArca, { IRowArca } from './Row';

class TableArca extends Component {
  private ws: WebSocket;
  state: {
    newRow: IRowArca,
    rows: IRowArca[]
  };

  constructor(props: any) {
    super(props);

    this.state = {
      newRow: {
        ID: "",
        Num1: 0,
        Num2: 0,
        I: 0,
      },
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
            console.log(response, 'insert');
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

  createRow(cell: any) {
    this.setState((state: any) => {
      const newState = {
        ...state,
        newRow: {
          ...state.newRow,
          [cell.field]: Number(cell.value),
        },
      };
      const { newRow } = newState;

      if (newRow.Num1 && newRow.Num2) {
        const request = {
          Method: 'insert',
          Context: {
            Source: 'ViewTable1',
          },
          Params: {...newRow, ID: 0},
        };
        this.ws.send(JSON.stringify(request));
      }
      return newState;
    });
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
    const { rows, newRow } = this.state;
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
        <RowArca
          key={newRow.ID}
          row={newRow}
          onRedact={this.createRow}
        />
      </tbody>
    </table>
    );
  }
}

export default TableArca;

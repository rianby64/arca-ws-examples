import React, { Component } from 'react'
import { Table } from 'semantic-ui-react'
import RowArca, { IRowArca } from './Row';

class TableArca extends Component {
  private ws: WebSocket;
  state: {
    rows: IRowArca[]
  };

  constructor(props: any) {
    super(props);

    this.state = {
      rows: []
    };

    this.updateRow = this.updateRow.bind(this);

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
          }
        }
      }
    }
  }

  updateRow(row: any) {
    const request = {
      Method: 'update',
      Context: {
        Source: 'ViewTable1',
      },
      Params: {
        ID: row.ID,
        [row.field]: Number(row.value),
      },
    };
    this.ws.send(JSON.stringify(request));
  }

  render() {
    const { rows } = this.state;
    return (
    <Table celled>
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell>ID</Table.HeaderCell>
          <Table.HeaderCell>Num1</Table.HeaderCell>
          <Table.HeaderCell>Num2</Table.HeaderCell>
          <Table.HeaderCell>I</Table.HeaderCell>
        </Table.Row>
      </Table.Header>

      <Table.Body>
        {
          rows.map(row =>
            <RowArca
              key={row.ID}
              row={row}
              onRedact={this.updateRow}
            />
          )
        }
      </Table.Body>
    </Table>
    );
  }
}

export default TableArca;

import React, { Component } from 'react'
import { Table, Button, Tab } from 'semantic-ui-react'
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
    this.prepareAdd = this.prepareAdd.bind(this);

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

  createRow(row: any) {
    console.log('create row', row);
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

  deleteRow(row: any) {
    const request = {
      Method: 'delete',
      Context: {
        Source: 'ViewTable1',
      },
      Params: {
        ID: row.ID,
      },
    };
    console.log(request, "to delete");
  }

  prepareAdd() {
    console.log('add a new empty row');
  }

  render() {
    const { rows, newRow } = this.state;
    return (
    <Table celled>
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell colSpan="5">
            <Button
              icon='add'
              onClick={this.prepareAdd}
            />
          </Table.HeaderCell>
        </Table.Row>
        <Table.Row>
          <Table.HeaderCell>ID</Table.HeaderCell>
          <Table.HeaderCell>Num1</Table.HeaderCell>
          <Table.HeaderCell>Num2</Table.HeaderCell>
          <Table.HeaderCell>I</Table.HeaderCell>
          <Table.HeaderCell></Table.HeaderCell>
        </Table.Row>
      </Table.Header>

      <Table.Body>
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
      </Table.Body>
    </Table>
    );
  }
}

export default TableArca;

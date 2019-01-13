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
            Source: 'ViewTable1'
        }
      }));
      this.ws.onmessage = (e) => {
        const rows = JSON.parse(e.data).Result;
        this.setState({ rows });
      }
    }
  }

  updateRow(row: any) {
    console.log('update row', row);
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

import React, { Component } from 'react'
import { Table } from 'semantic-ui-react'
import CellArca from './Cell';

export interface IRowArca {
  ID: string,
  Num1: number,
  Num2: number,
  I: number
};

class RowArca extends Component<any> {
  render() {
    const { row } = this.props;
    return (
      <Table.Row>
        <Table.Cell>{row.ID}</Table.Cell>
        <Table.Cell>
          <CellArca
            ID={row.ID}
            value={row.Num1}
            field='Num1'
          />
        </Table.Cell>
        <Table.Cell>
          <CellArca
            ID={row.ID}
            value={row.Num2}
            field='Num2'
          />
        </Table.Cell>
        <Table.Cell>
          <CellArca
            ID={row.ID}
            value={row.I}
            field='I'
          />
        </Table.Cell>
      </Table.Row>
    );
  }
}

export default RowArca;

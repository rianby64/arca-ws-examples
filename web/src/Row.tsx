import React, { Component } from 'react'
import { Table } from 'semantic-ui-react'
import { any } from 'prop-types';

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
        <Table.Cell>{row.Num1}</Table.Cell>
        <Table.Cell>{row.Num2}</Table.Cell>
        <Table.Cell>{row.I}</Table.Cell>
      </Table.Row>
    );
  }
}

export default RowArca;

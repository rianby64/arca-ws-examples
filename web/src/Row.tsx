import React, { Component } from 'react'
import { Table, Button } from 'semantic-ui-react'
import CellArca from './Cell';

export interface IRowArca {
  ID: string,
  Num1: number,
  Num2: number,
  I: number
};

class RowArca extends Component<any> {
  constructor(props: any) {
    super(props);

    this.onDelete = this.onDelete.bind(this);
  }

  onDelete() {
    const { row } = this.props;
    this.props.onDelete(row);
  }

  render() {
    const { row, onRedact } = this.props;
    return (
      <Table.Row>
        <Table.Cell collapsing>{row.ID}</Table.Cell>
        <Table.Cell>
          <CellArca
            onRedact={onRedact}
            ID={row.ID}
            value={row.Num1}
            field='Num1'
          />
        </Table.Cell>
        <Table.Cell>
          <CellArca
            onRedact={onRedact}
            ID={row.ID}
            value={row.Num2}
            field='Num2'
          />
        </Table.Cell>
        <Table.Cell collapsing>
          <CellArca
            onRedact={onRedact}
            ID={row.ID}
            value={row.I}
            field='I'
          />
        </Table.Cell>
        <Table.Cell collapsing>
          <Button icon='delete' onClick={this.onDelete} />
        </Table.Cell>
      </Table.Row>
    );
  }
}

export default RowArca;

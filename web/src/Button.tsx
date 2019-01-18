import React, { Component } from 'react';

class Button extends Component<any> {
  constructor(props: any) {
    super(props);

    this.handleModal = this.handleModal.bind(this);
  }

  handleModal() {
    this.props.openModal(this.props.table);
  }
  
  render() {
    return (
      <button onClick={this.handleModal} className="button new-row">Create Row</button>
    );
  }
}

export default Button;

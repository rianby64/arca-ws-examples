import React, { Component } from 'react';

class Button extends Component<any> {
  constructor(props: any) {
    super(props);
  }
  
  render() {
    return (
      <button onClick={this.props.openModal} className="button new-row">Create Row</button>
    );
  }
}

export default Button;

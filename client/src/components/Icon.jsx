import React from 'react';
import './Input.css';
export default class Icon extends React.Component {
    constructor(props) {
        super(props);
        this.state = {img: this.props.img}
        this.handleClick = this.props.onClick
    }

    render() {
        return (
            <div className="ds-col-2 ds-ml-4 icon-bg">
                <img alt="add-member" className="icon" src={this.state.img} height="28" onClick={this.handleClick}/>
            </div>
        )
    }
}
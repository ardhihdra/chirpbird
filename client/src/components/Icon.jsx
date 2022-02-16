import React from 'react';
export default class Icon extends React.Component {
    constructor(props) {
        super(props);
        this.state = {img: this.props.img, className: this.props.className}
        this.handleClick = this.props.onClick
    }

    render() {
        return (
            <div className={`${this.state.className} ml-4`}>
                <img 
                    alt="add-member"
                    className="icon-bg"
                    src={this.state.img}
                    height="36"
                    width="36"
                    onClick={this.handleClick}/>
            </div>
        )
    }
}
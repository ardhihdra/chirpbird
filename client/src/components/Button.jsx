import React from 'react';
import './Input.css';

export default class Button extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: props.value}
        this.handleSubmitParent = this.props.onClick
    }
       
    handleChange = (e) => {
        this.setState({value: e.target.value})
        this.handleSubmitParent(e)
    }

    styling = () => {
        const button = {
            cursor: 'pointer',
            // minWidth: '6.5rem',
            width: '100%',
            height: '2.5rem',
            marginLeft: 'auto',
            marginRight: 'auto',
        }
        const largeButton = {
            width: '100%',
        }
        return {
            button: button,
            largeButton: largeButton,
            buttonClear: {
                ...button,
                backgroundColor: 'white',
                border: '1px solid var(--main-button)',
                color: 'var(--main-button)',
                cursor: 'pointer'
            },
            buttonSolid: {
                ...button,
                backgroundColor: 'var(--main-button)',
                borderColor: 'var(--main-button)',
                color: 'white',
            },
            cancelButtonSolid: {
                ...button,
                backgroundColor: 'white',
                borderColor: 'var(--red)',
                color: 'var(--red)'
            },
        }
    }

    render() {
        const label = this.props.label
        const type = this.props.type || 'button'
        const title = this.props.title
        const styling = this.styling()
        const size = this.props.size
        const style = this.props.style
        const color = this.props.color
        let background = styling.buttonSolid
        let divClass = `px-2 border rounded-xl text-base shadow-xl shadow-gray-400 ${this.props.className}`
        if(style === 'clear') background = styling.buttonClear
        if(size === 'lg') background = {...background, ...styling.largeButton}
        if(color === 'danger') background = {...background, ...styling.cancelButtonSolid}

        return (
            <button 
                className={divClass}
                style={background}
                title={title}
                type={type} 
                onClick={this.handleSubmitParent}
            >{label}
            </button>
        )
    }
}
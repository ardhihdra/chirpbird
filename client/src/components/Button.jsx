import React from 'react';
import './Input.css';

export default class Button extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: props.value}
        this.handleSubmitParent = this.props.onSubmit
    }
       
    handleChange = (e) => {
        this.setState({value: e.target.value})
        this.handleSubmitParent(e)
    }

    styling = () => {
        const button = {
            cursor: 'pointer',
            width: '6.5rem',
            // width: '100%',
            height: '2.5rem',
            marginLeft: 'auto',
            marginRight: 'auto',
            border: '2px solid',
            borderRadius: 'var(--border-radius-norm)',
        }
        const largeButton = {
            width: '100%',
            // width: '8rem',
            fontSize: 'larger'
        }
        return {
            button: button,
            buttonClear: {
                ...button,
                backgroundColor: 'white',
                borderColor: 'var(--blue-soft)',
                color: 'var(--blue-soft)',
                cursor: 'pointer'
            },
            largeButton: largeButton,
            buttonSolid: {
                ...button,
                ...largeButton,
                backgroundColor: 'var(--blue-soft)',
                borderColor: 'var(--blue-soft)',
                color: 'white',
            },
            cancelButtonSolid: {
                ...button,
                ...largeButton,
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
        let divClass = `${this.props.className} ds-m-2 ds-mb-5`
        if(size === 'lg') divClass += styling.largeButton + ' '
        if(style === 'clear') {
            divClass += styling.buttonClear
        } else {
            divClass += styling.buttonSolid
        }
        return (
            <button className={divClass} style={styling.buttonClear} title={title}
                type={type} onClick={this.state.handleSubmitParent}>{label}</button>
        )
    }
}
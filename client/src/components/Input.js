import React from 'react';
import './Input.css';
export default class Input extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: props.value}
        this.handleChangeParent = this.props.onChange
        this.handleKeyUp = this.props.onKeyUp
    }
       
    handleChange = (e) => {
        this.setState({value: e.target.value})
        this.handleChangeParent(e)
    }

    styling = () => {
        return {
            label: {
                fontSize: 'smaller',
                position: 'relative',
                top: this.props.type !== 'textarea' ?'-4rem': '-9.4rem',
                left: '-36%',
                width: 'fit-content',
                backgroundColor: 'transparent',
                borderRadius: 'var(--border-radius-sm)',
                color: 'var(--primary-font)',
            },
            smallinput: {
                width: '80px',
                marginLeft: 'auto',
                marginRight: 'auto'
            }
        }
    }

    render() {
        const value = this.state.value
        const name = this.props.name
        const label = this.props.label
        const required = this.props.required
        const type = this.props.type || 'text'
        const placeholder = this.props.placeholder
        let min = null
        let max = null 
        let styling = this.styling()
        let style = {}
        const divClass = `${this.props.className} px-1 rounded text-primary-blue`
        if(type === 'number') {
            min = this.props.min
            max = this.props.max 
            style = {...styling.smallinput}
        }
        return (
            <div>
                { type !== 'textarea' ? 
                    <input 
                        className={divClass}
                        style={style}
                        name={name}
                        type={type}
                        value={value}
                        min={min}
                        max={max}
                        placeholder={placeholder}
                        onChange={this.handleChange}
                        onKeyUp={this.handleKeyUp}
                        required={required}
                    />
                :
                    <textarea 
                        className={divClass}
                        style={style}
                        name={name}
                        type={type}
                        value={value}
                        placeholder={placeholder}
                        onChange={this.handleChange}
                        required={required}
                    />
                }
                <span style={styling.label}>{label}</span>
            </div>
        )
    }
}
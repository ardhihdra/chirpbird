import React from 'react';
import './Input.css';
export default class Input extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: props.value}
        this.handleChangeParent = this.props.onChange
    }
       
    handleChange = (e) => {
        this.setState({value: e.target.value})
        this.handleChangeParent(e)
    }

    styling() {
        return {
            label: {
                fontSize: 'smaller',
                position: 'relative',
                top: '1rem',
                left: '-30%',
                width: 'fit-content',
                backgroundColor: this.props.label ? 'white': 'transparent',
                borderRadius: 'var(--border-radius-sm)',
                padding: '1px 2px',
                color: 'var(--primary-font)'
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
        const min = type === 'number' ? this.props.min: null
        const max = type === 'number' ? this.props.max: null 
        const placeholder = this.props.placeholder
        const style = this.styling()
        return (
            <div className={this.props.className}>
                <span style={style.label}>{label}</span>
                { type !== 'textarea' ? 
                    <input className="ds-ml-3 ds-p-1" name={name} type={type} value={value} 
                        min={min} max={max} style={type === 'number' ? style.smallinput: {}}
                        placeholder={placeholder} onChange={this.handleChange} required={required}/>
                :
                    <textarea className="ds-p-2" name={name} type={type} value={value} 
                    min={min} max={max} style={type === 'number' ? style.smallinput: {}}
                    placeholder={placeholder} onChange={this.handleChange} required={required}/>
                }
            </div>
        )
    }
}
import React, {Component} from 'react';
import reactCSS from 'reactcss'
import {CirclePicker} from 'react-color';

import './NameInput.css';

class NameInput extends Component {
    constructor(props) {
        super(props);
        this.state = {
            value: '',
            onChange: props.onChange,
            onSubmit: props.onSubmit,
            displayColorPicker: false,
            color: {
                r: '241',
                g: '112',
                b: '19',
                a: '1',
            },
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleClick = () => {
        this.setState({displayColorPicker: !this.state.displayColorPicker})
    };

    handleClose = () => {
        this.setState({displayColorPicker: false})
    };

    handleChangeColor = (color) => {
        this.setState({color: color.rgb})
        this.handleClose()
    };

    handleChange(event) {
        this.setState({value: event.target.value});
        this.state.onChange(event.target.value);
    }

    rgbToUint = (color) => {
        return (color.r << 24 >>> 0) + (color.g << 16) + (color.b << 8);
    }

    handleSubmit(event) {
        this.state.onSubmit(this.state.value, this.rgbToUint(this.state.color));
        event.preventDefault();
    }

    render() {
        const styles = reactCSS({
            'default': {
                color: {
                    width: '36px',
                    height: '36px',
                    borderRadius: '2px',
                    background: `rgb(${this.state.color.r}, ${this.state.color.g}, ${this.state.color.b})`,
                },
                swatch: {
                    padding: '5px',
                    background: '#fff',
                    borderRadius: '1px',
                    boxShadow: '0 0 0 1px rgba(0,0,0,.1)',
                    display: 'inline-block',
                    cursor: 'pointer',
                },
                popover: {
                    position: 'absolute',
                    backgroundColor: '#fff',
                    borderRadius: '5px',
                    zIndex: '2',
                },
                cover: {
                    position: 'fixed',
                    top: '0px',
                    right: '0px',
                    bottom: '0px',
                    left: '0px',
                },
            },
        });
        return (
            <div style={{display: "flex", flexDirection: "column"}}>
                <form className="name-input-form" onSubmit={this.handleSubmit}>
                    <div style={{display: "flex", flexDirection: "row"}}>
                        Username:
                        <input className="name-input-input" type="text" value={this.state.value}
                               onChange={this.handleChange}/>
                    </div>
                    <div style={{display: "flex", flexDirection: "row"}}>
                        Color:
                        <div style={styles.swatch} onClick={this.handleClick}>
                            <div style={styles.color}/>
                        </div>
                        {this.state.displayColorPicker ? <div style={styles.popover}>
                            <div style={styles.cover} onClick={this.handleClose}/>
                            <CirclePicker color={this.state.color} onChange={this.handleChangeColor}/>
                        </div> : null}
                    </div>
                    <input className="name-input-button" disabled={this.props.isDisabled} type="submit" value="Submit"/>
                </form>
            </div>

        );
    }
}

export default NameInput;
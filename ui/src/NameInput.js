import React, {Component} from 'react';

import './NameInput.css';

class NameInput extends Component {
    constructor(props) {
        super(props);
        this.state = {value: '', onChange: props.onChange, onSubmit: props.onSubmit};

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        this.setState({value: event.target.value});
        this.state.onChange(event.target.value);
    }

    handleSubmit(event) {
        this.state.onSubmit(this.state.value);
        event.preventDefault();
    }

    render() {
        return (
            <div>
                Username:
                {this.props.nameResponse ?
                <div>{this.props.nameResponse}</div>
                : <div><form className="name-input-form" onSubmit={this.handleSubmit}>
                    <input className="name-input-input" type="text" value={this.state.value} onChange={this.handleChange} />
                    <input className="name-input-button" disabled={this.props.isDisabled} type="submit" value="Submit" />
                </form>
                </div>}
            </div>

        );
    }
}

export default NameInput;
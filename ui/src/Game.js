import React, {Component} from 'react';
import './Game.css';

export default class Game extends Component {

    constructor(props) {
        super(props);

        this.canvasRef = React.createRef();
    }

    componentDidMount() {
        const canvas = this.canvasRef.current;
        const context = canvas.getContext('2d');
        context.fillRect(0, 0, canvas.width, canvas.height);
    }

    render() {
        return (
            <div className="Game">
                <canvas ref={this.canvasRef}></canvas>
            </div>
        );
    }
}
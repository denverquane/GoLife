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

    componentDidUpdate(prevProps, prevState, snapshot) {
        let oldTick = prevProps.board && prevProps.board.tick ? prevProps.board.tick : -1
        if (this.props.board && this.props.board.tick !== oldTick) {
            const canvas = this.canvasRef.current;
            const context = canvas.getContext('2d');
            context.fillRect(0, 0, canvas.width, canvas.height);
            context.fillStyle = "#FF00FF";
            let cWidth = canvas.width / this.props.board.width;
            let cHeight = canvas.height / this.props.board.height;
            for (let y = 0; y < this.props.board.height; y++) {
                for (let x = 0; x < this.props.board.width; x++) {
                    let elemIndex = y * (this.props.board.width) + x
                    if (this.props.board.data[elemIndex] & 128) {
                        context.fillRect(x*cWidth, y*cHeight, cWidth-1, cHeight-1);
                    }
                }
            }
            context.fillStyle = "#000000";
        }
    }

    render() {
        return (
            <div className="Game" >
                <canvas width={800} height={800} ref={this.canvasRef}></canvas>
                Tick: {this.props.board ? this.props.board.tick : 0}
            </div>
        );
    }
}
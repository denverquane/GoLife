import React, {Component} from 'react';
import './Game.css';

const ALIVE = 0x000000FF;

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
        if (this.props.tick !== prevProps.tick) {
            const canvas = this.canvasRef.current;
            const context = canvas.getContext('2d');
            context.fillRect(0, 0, canvas.width, canvas.height);
            let cWidth = canvas.width / this.props.width;
            let cHeight = canvas.height / this.props.height;
            for (let y = 0; y < this.props.height; y++) {
                for (let x = 0; x < this.props.width; x++) {
                    let elemIndex = y * (this.props.width) + x;
                    let cell = this.props.boardData[elemIndex];
                    let aliveness = cell & ALIVE
                    if (aliveness > 0) {
                        //this is so dumb; fighting JS' unsigned integer stupidity
                        let r = ((cell >> 8) & (ALIVE << 16)) >> 16;
                        let g = ((cell >> 8) & (ALIVE << 8)) >> 8;
                        let b = (cell >> 8) & ALIVE;
                        context.fillStyle = 'rgba(' + r + ', ' + g + ',' + b + ',' + (aliveness+64)/255.0 + ')';
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
                <canvas width={1600} height={900} ref={this.canvasRef}></canvas>
                Tick: {this.props.tick}
            </div>
        );
    }
}
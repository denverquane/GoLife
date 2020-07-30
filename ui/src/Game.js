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

    equal(arr1, arr2) {
        if (arr1.length !== arr2.length) {
            return false
        }
        for (let i=0; i < arr1.length; i++) {
            if (arr1[i] !== arr2[i]) {
                return false
            }
        }
        return true
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (this.props.paused !== prevProps.paused || this.props.width !== prevProps.width || this.props.tick !== prevProps.tick || !this.equal(this.props.boardData, prevProps.boardData)) {
            console.log("Updating")
            const canvas = this.canvasRef.current;
            if (this.props.paused !== prevProps.paused) {
                if (!this.props.paused) {
                    canvas.style.cursor = 'not-allowed';
                } else {
                    canvas.style.cursor = 'crosshair';
                }
            }
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
                <canvas width={this.props.canvasWidth} height={this.props.canvasHeight} ref={this.canvasRef} onClick={this.props.onClick}></canvas>
                Tick: {this.props.tick}
            </div>
        );
    }
}
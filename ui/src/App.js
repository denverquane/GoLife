import React, {Component} from 'react';
import './App.css';
import Game from './Game';
import NameInput from './NameInput';

const Messages = require('./proto/message/message_pb');

const UNCONNECTED = 0;
const CONNECTED = 1;
const REGISTERED = 2;

export const CANVAS_BASE_WIDTH = 800;
export const CANVAS_BASE_HEIGHT = 800;

const DEBUG_DONT_REGISTER_FOR_DATA = false;

let BASE_URL = process.env.REACT_APP_SERVICE_URL;
if (!BASE_URL || BASE_URL === "") {
    console.log("REACT_APP_SERVICE_URL not provided; defaulting to localhost:5000")
    BASE_URL = "localhost:5000"
}

class App extends Component {
    constructor(props) {
        super(props);

        this.state = {
            ws: null,
            playersOnline: null,
            localUsername: null,
            remoteUsername: null,
            color: null,
            rles: null,
            currentRLE: null,

            wsMessage: null,

            gameState: UNCONNECTED,
            boardData: null,
            boardTick: 0,
            boardWidth: 0,
            boardHeight: 0,
            paused: false,
        };
        this.onChangeUsername = this.onChangeUsername.bind(this);
        this.onSubmitUsername = this.onSubmitUsername.bind(this);
        this.onTogglePause = this.onTogglePause.bind(this);
        this.onCanvasClick = this.onCanvasClick.bind(this);
        this.onEnterRLEMode = this.onEnterRLEMode.bind(this);
    }

    componentDidMount() {
        this.connect();
    }

    timeout = 250; // Initial timeout duration as a class variable

    /**
     * @function connect
     * This function establishes the connect with the websocket and also ensures constant reconnection if connection closes
     */
    connect = () => {
        var ws = new WebSocket("ws://" + BASE_URL + "/ws");
        let that = this; // cache the this
        var connectInterval;

        // websocket onopen event listener
        ws.onopen = () => {
            console.log("connected websocket main component");

            this.setState({ws: ws, gameState: CONNECTED});

            that.timeout = 250; // reset timer to 250 on open of websocket connection
            clearTimeout(connectInterval); // clear Interval on on open of websocket connection
        };

        ws.onmessage = (event) => {
            event.data.arrayBuffer().then(buffer => {
                let message = Messages.Message.deserializeBinary(new Uint8Array(buffer));
                switch (message.getType()) {
                    case Messages.MessageType.WORLD_DATA:
                        let WorldMessage = Messages.WorldData.deserializeBinary(message.getContent())
                        if (WorldMessage.getHeight() !== 0 && WorldMessage.getWidth() !== 0) {
                            this.setState({
                                boardWidth: WorldMessage.getWidth(),
                                boardHeight: WorldMessage.getHeight(),
                                boardData: WorldMessage.getDataList(),
                                boardTick: WorldMessage.getTick(),
                                paused: WorldMessage.getPaused()
                            })
                        } else {
                            this.setState({
                                boardData: WorldMessage.getDataList(),
                                boardTick: WorldMessage.getTick(),
                                paused: WorldMessage.getPaused()
                            })
                        }

                        break;
                    case Messages.MessageType.REGISTER:
                        let RegisterMessage = Messages.Player.deserializeBinary(message.getContent())
                        this.setState({gameState: REGISTERED, remoteUsername: RegisterMessage.getName()})
                        break;
                    case Messages.MessageType.PLAYERS:
                        let PlayersMessage = Messages.Players.deserializeBinary(message.getContent())
                        this.setState({playersOnline: PlayersMessage.getPlayersList()})
                        break;
                    case Messages.MessageType.RLE_OPTIONS:
                        let RlesMessage = Messages.RLEs.deserializeBinary(message.getContent())
                        this.setState({rles: RlesMessage.getRlesList()})
                        break;
                }
            });
        }

        // websocket onclose event listener
        ws.onclose = e => {
            console.log(
                `Socket is closed. Reconnect will be attempted in ${Math.min(
                    10000 / 1000,
                    (that.timeout + that.timeout) / 1000
                )} second.`,
                e.reason
            );
            this.setState({
                gameState: UNCONNECTED,
                rles: null,
                remoteUsername: null,
                localUsername: null,
                playersOnline: null
            })

            that.timeout = that.timeout + that.timeout; //increment retry interval
            connectInterval = setTimeout(this.check, Math.min(10000, that.timeout)); //call check function after timeout
        };

        // websocket onerror event listener
        ws.onerror = err => {
            console.error(
                "Socket encountered error: ",
                err.message,
                "Closing socket"
            );

            ws.close();
        };
    };

    /**
     * utilized by the @function connect to check if the connection is close, if so attempts to reconnect
     */
    check = () => {
        const {ws} = this.state;
        if (!ws || ws.readyState === WebSocket.CLOSED) this.connect(); //check if websocket instance is closed, if so call `connect` function.
    };

    onChangeUsername(username) {
        //console.log("Change: " + username);
        this.setState({localUsername: username});
    }

    onSubmitUsername(username, color) {
        let regMsg = new Messages.Player();
        regMsg.setName(username);
        regMsg.setColor(color);
        console.log("Submitting username " + username + " and color " + color)
        let innerBytes = regMsg.serializeBinary();

        let msg = new Messages.Message();
        msg.setType(Messages.MessageType.REGISTER);
        msg.setContent(innerBytes);
        let bytes = msg.serializeBinary();

        this.state.ws.send(bytes);
        this.setState({color: color})
    }

    onTogglePause() {
        let cmdMsg = new Messages.Command();
        cmdMsg.setType(Messages.CommandType.TOGGLE_PAUSE)
        let innerBytes = cmdMsg.serializeBinary()
        let msg = new Messages.Message();
        msg.setType(Messages.MessageType.COMMAND);
        msg.setContent(innerBytes);
        let bytes = msg.serializeBinary();

        this.state.ws.send(bytes);
    }

    onEnterRLEMode(name) {
        if (this.state.currentRLE && this.state.currentRLE.getName() === name) {
            this.setState({currentRLE: null})
        } else {
            this.state.rles.forEach((item, idx) => {
                if (item.getName() === name) {
                    this.setState({currentRLE: item})
                }
            })
        }
    }

    onCanvasClick(event) {
        let x = event.nativeEvent.offsetX;
        let y = event.nativeEvent.offsetY;
        let cellX = Math.floor(x / ((this.state.boardWidth + CANVAS_BASE_WIDTH) / this.state.boardWidth))
        let cellY = Math.floor(y / ((this.state.boardHeight + CANVAS_BASE_HEIGHT) / this.state.boardHeight))
        //console.log(cellX, cellY)

        let cmdMsg = new Messages.Command();

        if (this.state.currentRLE) {
            cmdMsg.setType(Messages.CommandType.PLACE_RLE);
            cmdMsg.setText(this.state.currentRLE.getName());
            cmdMsg.setX(cellX)
            cmdMsg.setY(cellY)
        } else {
            cmdMsg.setType(Messages.CommandType.MARK_CELL)
            cmdMsg.setX(cellX)
            cmdMsg.setY(cellY)
        }
        let innerBytes = cmdMsg.serializeBinary()
        let msg = new Messages.Message();
        msg.setType(Messages.MessageType.COMMAND);
        msg.setContent(innerBytes);
        let bytes = msg.serializeBinary();

        this.state.ws.send(bytes);
    }


    render() {
        return (
            <div className="App">
                <header className="App-header">
                    {
                        this.state.gameState === REGISTERED ?
                            <div className="App-header-left">
                                <div>Players Online:</div>
                                {this.state.playersOnline
                                    ? <div style={{alignSelf: "center"}}>
                                        {this.state.playersOnline.map(function (item, i) {
                                            const colorString = item.getColor().toString(16).toUpperCase().substr(0, 6);
                                            return <div key={i} style={{display: "flex", flexDirection: "row"}}>
                                                <div>
                                                    <div style={{
                                                        width: '45px',
                                                        height: '45px',
                                                        borderRadius: '2px',
                                                        background: `#${colorString}`,
                                                    }}/>
                                                </div>
                                                <div style={{paddingLeft: '10px'}}>
                                                    {item.getName()}
                                                </div>

                                            </div>
                                        })}
                                    </div>
                                    : <div/>}
                            </div> : <div className="App-header-left"/>
                    }

                    <div className="App-header-middle">
                        {/*<img src={logo} className="App-logo" alt="logo" />*/}
                        <div className="App-name">GoLife</div>
                        <div className="App-slogan">Interactive Multiplayer Cellular Automata!</div>
                        {this.state.gameState === UNCONNECTED ? <div>DISCONNECTED</div> : <div/>}
                    </div>
                    <div className="App-header-right">
                        {this.state.rles ? <div>
                            <div>
                                Patterns available:
                            </div>
                            {this.state.rles.map((item, i) => {
                                return <div key={i} style={{display: "flex", flexDirection: "row"}}>
                                    <div>
                                        {/*<div style={ {*/}
                                        {/*    width: '45px',*/}
                                        {/*    height: '45px',*/}
                                        {/*    borderRadius: '2px',*/}
                                        {/*    background: `#${colorString}`,*/}
                                        {/*}} />*/}
                                    </div>
                                    <button disabled={!this.state.paused}
                                            onClick={() => this.onEnterRLEMode(item.getName())}>
                                        {item.getName()}
                                    </button>

                                </div>
                            })}</div> : <div/>
                        }
                    </div>

                </header>
                <div className="App-content">
                    {
                        this.state.gameState === REGISTERED || (this.state.gameState === CONNECTED && DEBUG_DONT_REGISTER_FOR_DATA)
                            ? <div>
                                <button onClick={() => {
                                    this.onTogglePause()
                                }}>
                                    Toggle Pause
                                </button>
                                <Game boardData={this.state.boardData} tick={this.state.boardTick}
                                      width={this.state.boardWidth} height={this.state.boardHeight}
                                      onClick={this.onCanvasClick}
                                    //we do this addition to guarantee every cell has a 1 pixel border
                                      canvasWidth={CANVAS_BASE_WIDTH + this.state.boardWidth}
                                      canvasHeight={CANVAS_BASE_HEIGHT + this.state.boardHeight}
                                      paused={this.state.paused}
                                      currentRLE={this.state.currentRLE}
                                color={this.state.color}/>
                            </div>
                            : this.state.gameState !== UNCONNECTED ? <div>
                                <NameInput
                                    isDisabled={this.state.gameState === UNCONNECTED || this.state.localUsername === this.state.remoteUsername}
                                    onSubmit={this.onSubmitUsername} onChange={this.onChangeUsername}
                                />
                            </div> : <div/>
                    }
                </div>
            </div>
        )
    }
}

export default App;

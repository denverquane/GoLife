import React, {Component} from 'react';
import logo from './logo.svg';
import './App.css';
import Game from './Game';
import NameInput from './NameInput';
const Messages = require('./proto/message/message_pb');

const UNCONNECTED = 0;
const CONNECTED = 1;
const REGISTERED = 2;

const DEBUG_DONT_REGISTER_FOR_DATA = true;

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

            wsMessage: null,

            gameState: UNCONNECTED,
            boardData: null,
            boardWidth: 0,
            boardHeight: 0
        };
        this.onChangeUsername = this.onChangeUsername.bind(this);
        this.onSubmitUsername = this.onSubmitUsername.bind(this);
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

            this.setState({ ws: ws, gameState: CONNECTED });

            that.timeout = 250; // reset timer to 250 on open of websocket connection
            clearTimeout(connectInterval); // clear Interval on on open of websocket connection
        };

        ws.onmessage = (event) => {
            event.data.arrayBuffer().then(buffer => {
                let message = Messages.Message.deserializeBinary(new Uint8Array(buffer));
                if (message.getType() === Messages.MessageType.WORLD_INFO) {
                    let WorldInfo = Messages.WorldInfo.deserializeBinary(message.getContent())
                    console.log("Received dimensions of grid: %dw,%dh", WorldInfo.getWidth(), WorldInfo.getHeight())
                    this.setState({boardWidth: WorldInfo.getWidth(), boardHeight: WorldInfo.getHeight() })
                } else if (message.getType() === Messages.MessageType.WORLD_DATA) {
                    let WorldMessage = Messages.WorldData.deserializeBinary(message.getContent())
                    let board = {
                        width: this.state.boardWidth,
                        height: this.state.boardHeight,
                        data: WorldMessage.getData(),
                        tick: WorldMessage.getTick(),
                    }
                    this.setState({board: board})
                } else if (message.getType() === Messages.MessageType.REGISTER) {
                    let RegisterMessage = Messages.Player.deserializeBinary(message.getContent())
                    this.setState({gameState: REGISTERED, remoteUsername: RegisterMessage.getName()})
                } else if (message.getType() === Messages.MessageType.PLAYERS) {
                    let PlayersMessage = Messages.Players.deserializeBinary(message.getContent())
                    this.setState({playersOnline: PlayersMessage.getPlayersList()})
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
            this.setState({gameState: UNCONNECTED})

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
        const { ws } = this.state;
        if (!ws || ws.readyState === WebSocket.CLOSED) this.connect(); //check if websocket instance is closed, if so call `connect` function.
    };

    onChangeUsername(username) {
        console.log("Change: " + username);
        this.setState({localUsername: username});
    }

    onSubmitUsername(username) {
        let regMsg = new Messages.Player();
        regMsg.setName(username);
        let innerBytes = regMsg.serializeBinary();

        let msg = new Messages.Message();
        msg.setType(Messages.MessageType.REGISTER);
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
                      <div>Players Online: </div>
                      {this.state.playersOnline ? <div>{this.state.playersOnline.length}</div> : <div/>}
                  </div> : <div className="App-header-left"/>
          }

          <div className="App-header-middle">
              {/*<img src={logo} className="App-logo" alt="logo" />*/}
              <div className="App-name">GoLife</div>
              <div className="App-slogan">Interactive Multiplayer Cellular Automata!</div>
              {this.state.gameState === UNCONNECTED ? <div>DISCONNECTED</div> : <div/>}
          </div>
          <div className="App-header-right">
              <NameInput isDisabled={this.state.gameState === UNCONNECTED || this.state.localUsername === this.state.remoteUsername}
                         onSubmit={this.onSubmitUsername}  onChange={this.onChangeUsername}
              nameResponse={this.state.remoteUsername}/>
          </div>

      </header>
        <div className="App-content">
            {
                this.state.gameState === REGISTERED || (this.state.gameState === CONNECTED && DEBUG_DONT_REGISTER_FOR_DATA)
                    ? <Game board={this.state.board}/>
                    : this.state.gameState !== UNCONNECTED ? <div>Please enter a username!</div> : <div/>
            }
        </div>
    </div>
  )}
}

export default App;

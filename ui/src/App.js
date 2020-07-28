import React, {Component} from 'react';
import logo from './logo.svg';
import './App.css';
import Game from './Game';
import NameInput from './NameInput';
import { load } from "protobufjs";
import proto from "./proto/message/message.proto";

const UNCONNECTED = 0;
const CONNECTED = 1;
const REGISTERED = 2;

var ProtoMessage;
var RegisterNameMessage;

class App extends Component {
    constructor(props) {
        super(props);

        this.state = {
            ws: null,
            playersOnline: null,
            localUsername: null,
            remoteUsername: null,

            wsMessage: null,

            gameState: UNCONNECTED
        };
        this.onChangeUsername = this.onChangeUsername.bind(this);
        this.onSubmitUsername = this.onSubmitUsername.bind(this);
    }

    componentDidMount() {
        load(proto, function(err, root) {
            if (err)
                throw err;

            // Obtain a message type
            RegisterNameMessage = root.lookupType("message.RegisterName");
            ProtoMessage = root.lookupType("message.Message");

            // // Exemplary payload
            // var payload = {name: "AwesomeString"};
            //
            // // Verify the payload if necessary (i.e. when possibly incomplete or invalid)
            // var errMsg = RegisterNameMessage.verify(payload);
            // if (errMsg)
            //     throw Error(errMsg);
            //
            // // Create a new message
            // var message = AwesomeMessage.create(payload); // or use .fromObject if conversion is necessary
            //
            // // Encode a message to an Uint8Array (browser) or Buffer (node)
            // var buffer = AwesomeMessage.encode(message).finish();
            // // ... do something with buffer

            // // Decode an Uint8Array (browser) or Buffer (node) to a message
            // var message = AwesomeMessage.decode(buffer);
            // // ... do something with message
            //
            // // If the application uses length-delimited buffers, there is also encodeDelimited and decodeDelimited.
            //
            // // Maybe convert the message back to a plain object
            // var object = AwesomeMessage.toObject(message, {
            //     longs: String,
            //     enums: String,
            //     bytes: String,
            //     // see ConversionOptions
            // });
        });
        this.connect();
    }

    timeout = 250; // Initial timeout duration as a class variable

    /**
     * @function connect
     * This function establishes the connect with the websocket and also ensures constant reconnection if connection closes
     */
    connect = () => {
        var ws = new WebSocket("ws://localhost:5000/ws");
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
            console.log(event);
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
        var payload = {name: username};

        var errMsg = RegisterNameMessage.verify(payload);
            if (errMsg)
                throw Error(errMsg);
        var message = RegisterNameMessage.create(payload); // or use .fromObject if conversion is necessary
        var buffer = RegisterNameMessage.encode(message).finish();

        var payload2 = {type: 0, content: buffer}
        var errMsg2 = ProtoMessage.verify(payload2);
        if (errMsg2)
            throw Error(errMsg);
        var message2 = ProtoMessage.create(payload2); // or use .fromObject if conversion is necessary
        var buffer2 = ProtoMessage.encode(message2).finish();


        console.log("Sending: " + JSON.stringify(payload2));
        this.state.ws.send(buffer2);
    }

    render() {
  return (
    <div className="App">
      <header className="App-header">
          <div className="App-header-left">
              <div>Players Online: </div>
              <div>{this.state.playersOnline}</div>
          </div>
          <div className="App-header-middle">
              <img src={logo} className="App-logo" alt="logo" />
              <div className="App-name">GoLife</div>
              <div className="App-slogan">Interactive Multiplayer Cellular Automata!</div>
              {this.state.gameState === UNCONNECTED ? <div>DISCONNECTED</div> : <div/>}
          </div>
          <div className="App-header-right">
              <NameInput isDisabled={this.state.gameState === UNCONNECTED|| this.state.localUsername === this.state.remoteUsername}
                         onSubmit={this.onSubmitUsername}  onChange={this.onChangeUsername}
              nameResponse={this.state.nameResponse}/>
          </div>

      </header>
        <div className="App-content">
            {
                this.state.gameState === REGISTERED
                    ? <Game/>
                    : <div/>
            }
        </div>
    </div>
  )}
}

export default App;

import { browser } from "$app/environment";
import SocketIO from "@lib/SocketIO";
import * as globals from "@global";

let io: SocketIO | undefined;

if (browser) {
    const socket: SocketIO = new SocketIO(new URL("/ws", globals.serverHost));
    socket.onError((ev) => {
        console.error(ev.type);
    });
    io = socket;
}

export default io as SocketIO;

export interface SocketMessageRequest {
	eventKey: string;
	to?: string;
	message?: string;
	data?: Object | string;
	except: string[];
}
export type SocketIOListener = (message?: string, ev?: MessageEvent) => void;
export type ID_t = string;

export class SocketEmiter {
    #to?: string;
    #except: string[];
    constructor(public readonly socketClient: SocketIO) {
        this.#to = undefined;
        this.#except = [];
    }
    public to(ID: ID_t): SocketEmiter {
        this.#to = ID;
        return this;
    }
    public except(ID: ID_t): SocketEmiter {
        this.#except.push(ID);
        return this;
    }
    public emit(eventKey: string, message?: string) {
        this.socketClient.socket.send(
            JSON.stringify({
                eventKey,
                message,
                to: this.#to,
            } as SocketMessageRequest)
        );
    }
}

export default class SocketIO {
    socket: WebSocket;
    #listeners: Record<string, SocketIOListener> = {};
    #ID: ID_t = "";

    constructor(public readonly url: URL | string) {
        this.socket = new WebSocket(url);

        this.#listeners["connected"] = (msg, ev) => {
            this.#ID = msg as any;
        };

        this.socket.onmessage = (ev) => {
            const res: SocketMessageRequest = JSON.parse(ev.data);
            if (this.#listeners[res.eventKey]) {
                this.#listeners[res.eventKey](res.message, ev);
            }
        };
    }

    get ID() {
        return this.#ID;
    }

    public close() {
        this.socket.close();
    }

    public onConnect(call: SocketIOListener) {
        this.#listeners["connected"] = (msg, ev) => {
            this.#ID = msg as any;
            call(msg, ev);
        };
    }

    public onError(call: (ev: Event) => any) {
        this.socket.onerror = call;
    }

    public onClose(call: (ev: CloseEvent) => any) {
        this.socket.onclose = (ev) => {
            call(ev);
        };
    }

    public emit(eventKey: string, message?: string) {
        this.socket.send(
            JSON.stringify({
                eventKey,
                message,
                except: [],
            } as SocketMessageRequest)
        );
    }

    public to(ID: ID_t): SocketEmiter {
        const res = new SocketEmiter(this);
        res.to(ID);
        return res;
    }

    public broadcast(): SocketEmiter {
        const res = new SocketEmiter(this);
        res.except(this.ID);
        return res;
    }

    public except(ID: ID_t): SocketEmiter {
        const res = new SocketEmiter(this);
        res.except(ID);
        return res;
    }

    public on(eventId: string, call: SocketIOListener) {
        this.#listeners[eventId] = call;
    }
}

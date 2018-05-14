// noinspection JSUnresolvedFunction

class CMMonitor {

    constructor(url) {
        this.url = url
        this.mid = url
    }

    connect() {
        const socket = new WebSocket(this.url);

        socket.addEventListener('open', function (event) {
            socket.send('Hello Server!');
        });

        socket.addEventListener('message', function (event) {
            console.log('Message from server ', event.data);
        });

        this.socket = socket
    }

    send(cmd, params) {
        this.socket.send(JSON.stringify({
            command: cmd,
            params: Array.from(params)
        }))
    }

    start(cmd) {
        this.send('start', [cmd])
    }

    interrupt(pid) {
        this.send('interrupt', [pid.toString()])
    }

    kill(pid) {
        this.send('kill', [pid.toString()])
    }
}

class CMMonitors {
    constructor() {
        this.map = new Map()
    }

    addLocal() {
        this.add(`${window.location.host}`)
    }

    add(host) {
        const m = new CMMonitor(`ws://${host}/nodes`)
        m.connect()
        this.map.set(m.mid, m)
    }

    get(mid) {
        return this.map.get(mid)
    }

}

window.monitors = new CMMonitors()

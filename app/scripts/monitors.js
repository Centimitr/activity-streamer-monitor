// noinspection JSUnresolvedFunction

class CMMonitor {

    constructor(url) {
        this.url = url
        this.mid = url
    }

    connect() {
        const socket = new WebSocket(this.url);

        socket.addEventListener('open', event => this.send('test'));

        socket.addEventListener('message', function (event) {
            console.log('Message from server ', event.data);
        });

        this.socket = socket
    }

    send(cmd, params) {
        console.log(cmd, params)
        this.socket.send(JSON.stringify({
            command: cmd,
            params: params ? Array.from(params) : []
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

    get list() {
        return Array.from(this.map.values())
    }

    addLocal() {
        this.add(`${window.location.host}`)
    }

    add(host) {
        if (!host) {
            return
        }
        const m = new CMMonitor(`ws://${host}/nodes`)
        m.connect()
        this.map.set(m.mid, m)
    }

    get(mid) {
        return this.map.get(mid)
    }

}

window.monitors = new CMMonitors()

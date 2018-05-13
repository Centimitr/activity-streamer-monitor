// noinspection JSUnresolvedFunction

class Monitor {

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

    terminate(pid) {
        this.socket.send(`terminate ${pid}`)
    }

    kill(pid) {
        this.socket.send(`kill ${pid}`)
    }
}

class Monitors {
    constructor() {
        this.map = new Map()
    }

    addLocal() {
        this.add(`${window.location.host}`)
    }

    add(host) {
        const m = new Monitor(`ws://${host}/nodes`)
        m.connect()
        this.map.set(m.mid, m)
    }

    get(mid) {
        return this.map.get(mid)
    }

}

const monitors = new Monitors()
monitors.addLocal()

// noinspection JSUnresolvedFunction
angular.module('app', [])
    .controller('MonitorsController', function () {
        const mctrl = this
        mctrl.nodes = []
        mctrl.monitors = monitors
        // mctrl.terminate = (mid, pid) => monitors.get(mid).terminate(pid)
        // mctrl.kill = (mid, pid) => monitors.get(mid).kill(pid)
        // mctrl.refresh = function () {
        //     return $http.get('/nodes').then(data => mctrl.nodes = data)
        // }
        // unused
        // mctrl.keepRefresh = function (interval) {
        //     return setInterval(() => monitor.refresh(), interval)
        // }
        // mctrl.stopRefresh = function (timer) {
        //     clearInterval(timer)
        // }
        // monitor.keepRefresh(2000)

    });
monitors.addLocal()

// noinspection JSUnresolvedFunction
angular.module('app', [])
    .controller('MonitorsController', function () {
        const mctrl = this
        mctrl.tab = 1
        mctrl.tmp = {
            jarDir: '/Users/shixiao/Playground/COMP90015/out/artifacts/jar/',
            monitorHostToConnect: null,
            selectedHistoryRecord: null,
            history: ['java -jar '],
            clientToStart: {
                hostname: 'localhost',
                port: 7001,
                username: 'shixiao',
                secret: null
            },
            serverToStart: {
                hostname: 'localhost',
                port: 700,
                remoteHostname: 'localhost',
                remotePort: 700,
                secret: 'blablasecret'
            },
            params(...params) {
                const fvs = []
                for (let i = 0; i < params.length / 2; i++) {
                    fvs.push({
                        flag: params[i * 2],
                        value: params[i * 2 + 1]
                    })
                }
                const param = (flag, v) => v ? `-${flag} '${v}'` : ''
                return ' ' + (fvs.map(fv => param(fv.flag, fv.value)).join(' '))
            },
            get clientLaunchCmd() {
                const tmp = mctrl.tmp
                const c = tmp.clientToStart
                return `java -jar ${tmp.jarDir}Client.jar ${tmp.params('rh', c.hostname, 'rp', c.port, 'u', c.username, 's', c.secret)}`
            },
            get serverLaunchCmd() {
                const tmp = mctrl.tmp
                const s = tmp.serverToStart
                return `java -jar ${tmp.jarDir}Server.jar ${tmp.params('lh', s.hostname, 'lp', s.port, 'rh', s.remoteHostname, 'rp', s.remotePort, 'a', 5000, 's', s.secret)}`
            }
        }
        mctrl.selectedMonitor = null
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

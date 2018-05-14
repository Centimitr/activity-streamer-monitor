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

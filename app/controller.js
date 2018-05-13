// noinspection JSUnresolvedFunction
angular.module('app', [])
    .controller('MonitorController', function ($http) {
        const monitor = this
        monitor.nodes = []
        monitor.refresh = function () {
            return $http.get('/nodes').then(data => monitor.nodes = data)
        }
        // unused
        monitor.keepRefresh = function (interval) {
            return setInterval(() => monitor.refresh(), interval)
        }
        monitor.stopRefresh = function (timer) {
            clearInterval(timer)
        }
        // monitor.keepRefresh(2000)
    });

var responseTimesAvg = [];
var responseTime99 = [];
var requestsPerSec = [];

var optsResponseTimes = {
    chart: {
        type: 'area'
    },
    title: {
        text: 'Response time'
    },
    xAxis: {
        allowDecimals: false,
        labels: {
            formatter: function () {
                return this.value;
            }
        }
    },
    yAxis: {
        title: {
            text: 'millis'
        },
        labels: {
            formatter: function () {
                return this.value;
            }
        }
    },
    plotOptions: {
        area: {
            pointStart: 0,
            marker: {
                enabled: false,
                symbol: 'circle',
                radius: 2,
                states: {
                    hover: {
                        enabled: true
                    }
                }
            }
        }
    },
    series: [{
        name: '99% percentile',
        data: []
    }, {
        name: 'Average',
        data: []
    }]
};

var optsRequests = {
    chart: {
        type: 'area'
    },
    title: {
        text: 'Requests per second'
    },
    xAxis: {
        allowDecimals: false,
        labels: {
            formatter: function () {
                return this.value;
            }
        }
    },
    yAxis: {
        title: {
            text: 'Requests'
        },
        labels: {
            formatter: function () {
                return this.value;
            }
        }
    },
    plotOptions: {
        area: {
            pointStart: 0,
            marker: {
                enabled: false,
                symbol: 'circle',
                radius: 2,
                states: {
                    hover: {
                        enabled: true
                    }
                }
            }
        }
    },
    series: [{
        name: 'request per second',
        data: []
    }]
};

const PUSH_MAX_VALUE = 60;
var push = (arr, val) => {
    arr.push(val);
    if (arr.length > PUSH_MAX_VALUE) {
        arr.shift();
    }
};

var chart = Highcharts.chart('container', optsResponseTimes);
var chartRequests = Highcharts.chart('request-container', optsRequests);

var fetchAndUpdate = () => {
    superagent.get("http://localhost:8089").end((err, res) => {
        if (err != null) {
            console.log(err);return;
        }
        var body = JSON.parse(res.text);
        var mean = body.map(m => m.latencies.mean / (1000 * 1000)).reduce((acc, cur) => acc + cur, 0) / body.length;
        var lat99th = body.map(m => m.latencies["99th"] / (1000 * 1000)).reduce((acc, cur) => acc + cur, 0) / body.length;

        push(responseTimesAvg, mean);
        push(responseTime99, lat99th);

        chart.options.series[0].data = responseTime99;
        chart.options.series[1].data = responseTimesAvg;

        var rate = body.map(m => m.rate).reduce((acc, cur) => acc + cur, 0);
        push(requestsPerSec, rate);
        chartRequests.options.series[0].data = requestsPerSec;

        window.requestAnimationFrame(() => {
            chart.update(chart.options, true);
            chartRequests.update(chartRequests.options, true);
            console.log("Should update");
        });

        setTimeout(() => {
            fetchAndUpdate();
        }, 1000);
    });
};

setTimeout(() => {
    fetchAndUpdate();
}, 1000);

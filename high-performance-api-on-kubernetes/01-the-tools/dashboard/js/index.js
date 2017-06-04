
var responseTimesAvg = [];
var responseTime99 = [];

var opts = {
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
                return this.value; // clean, unformatted number for year
            }
        }
    },
    yAxis: {
        title: {
            text: 'Response times'
        },
        labels: {
            formatter: function () {
                return this.value / 1000 + 'k';
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

var chart = Highcharts.chart('container', opts);

var fetchAndUpdate = () => {
    superagent.get("http://localhost:8089").end((err, res) => {
        if (err != null) {
            console.log(err);return;
        }
        var body = JSON.parse(res.text);
        var mean = body.map(m => m.latencies.mean / (1000 * 1000)).reduce((acc, cur) => acc + cur, 0) / body.length;
        var lat99th = body.map(m => m.latencies["99th"] / (1000 * 1000)).reduce((acc, cur) => acc + cur, 0) / body.length;
        responseTimesAvg.push(mean);
        responseTime99.push(lat99th);

        chart.options.series[0].data = responseTime99;
        chart.options.series[1].data = responseTimesAvg;

        window.requestAnimationFrame(() => {
            chart.update(chart.options, true);
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

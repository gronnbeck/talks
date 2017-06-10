var responseTimesAvg = []
var responseTime99 = []
var requestsPerSec = []

const statusChart = (id) => Highcharts.chart(id, {
    chart: {
        plotBackgroundColor: null,
        plotBorderWidth: null,
        plotShadow: false,
        type: 'pie'
    },
    title: {
        text: 'Status codes'
    },
    plotOptions: {
        pie: {
            allowPointSelect: true,
            cursor: 'pointer',
            dataLabels: {
                enabled: true,
                format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                style: {
                    color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                }
            }
        }
    },
    series: [{
        name: 'Status Code',
        colorByPoint: true,
        data: []
    }]
});


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
    series: [
      {
         name: '99% percentile',
         data: []
     },
     {
        name: 'Average',
        data: []
    },
  ]
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
    series: [
      {
         name: 'request per second',
         data: []
     },
  ]
};

const PUSH_MAX_VALUE = 60
var push = (arr, val) => {
  arr.push(val)
  if (arr.length > PUSH_MAX_VALUE) {
    arr.shift()
  }
}

const statusCodesData = (payloads) => {
  const statusCodes = payloads.map((payload) => payload.status_codes)
  const combined = statusCodes.reduce((acc, curr) => {
    const keys = Object.keys(curr)
    for (i = 0; i < keys.length; i++) {
      const key = keys[i]
      if (acc[key] === undefined || acc[key] === null) {
        acc[key] = curr[key]
      } else {
        acc[key] += curr[key]
      }
    }
    return acc
  }, {})
  var data = []
  const keys = Object.keys(combined)
  for (i = 0; i < keys.length; i++) {
    const key = keys[i]
    data.push({name: key, y: combined[key]})
  }

  return data
}

var chart = Highcharts.chart('container', optsResponseTimes);
var chartRequests = Highcharts.chart('request-container', optsRequests);
var piechart = statusChart('status-chart')

var fetchAndUpdate = () => {
  superagent.get("http://localhost:8089").end((err,res) => {
  if (err != null) { console.log(err); return }
    var body = JSON.parse(res.text)
    var mean = body.map((m) => m.latencies.mean / (1000 * 1000))
                   .reduce((acc, cur) => acc+cur,0) / body.length;
   var lat99th = body.map((m) => m.latencies["99th"] / (1000 * 1000))
                  .reduce((acc, cur) => acc+cur,0) / body.length;

    push(responseTimesAvg, mean)
    push(responseTime99, lat99th)

    chart.options.series[0].data = responseTime99
    chart.options.series[1].data = responseTimesAvg


    var rate = body.map((m) => m.rate).reduce((acc, cur) => acc+cur, 0);
    push(requestsPerSec, rate)
    chartRequests.options.series[0].data = requestsPerSec

    piechart.options.series[0].data = statusCodesData(body)


    window.requestAnimationFrame(() => {
      document.getElementById("response-time-99-number").innerHTML = lat99th.toFixed(2)
      document.getElementById("response-time-mean-number").innerHTML = mean.toFixed(2)
      document.getElementById("request-per-sec-number").innerHTML = rate.toFixed(2)

      chart.update(chart.options, true)
      chartRequests.update(chartRequests.options, true)
      piechart.update(piechart.options, true)
      console.log("Should update")
    })

    setTimeout(() => {
      fetchAndUpdate()
    }, 1000)
  })
}

setTimeout(() => {
  fetchAndUpdate()
}, 1000)

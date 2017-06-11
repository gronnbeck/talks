
var chart = Highcharts.chart('container', {
    chart: {
        type: 'line'
    },
    title: {
        text: 'Currency'
    },
    xAxis: {
        categories: ["1", "2"],
    },
    yAxis: {
        title: {
            text: 'USD'
        }
    },
    series: [{
        name: 'weirdcurrency',
        data: []
    }]
});

const PUSH_MAX_VALUE = 60
var push = (arr, val) => {
  arr.push(val)
  if (arr.length > PUSH_MAX_VALUE) {
    arr.shift()
  }
}

const parseBody = (res) => {
  if (res.body != null) {
    return res.body
  }
  return JSON.parse(res.text)
}

var values = [];
var xAxis = [0];

const fetch = () => {
  superagent.get("http://localhost:8080").end((err,res) => {
    const body = parseBody(res)
    if (body.error != null) {
      console.log("Error from API: " + body.error)
      return
    }
    const current = body.current
    push(values, current)
    push(xAxis, xAxis[xAxis.length-1] + 1)
    chart.options.series[0].data = values
    chart.options.xAxis[0].categories = xAxis

    window.requestAnimationFrame(() => {
      chart.update(chart.options, true)
    })
  })
}

const scheduleFetching = () => {
  const fetchFunc = () => {
    fetch()
    setTimeout(fetchFunc, 1000)
  }

  fetchFunc()
}

const onload = () => {
  scheduleFetching()
}

window.onload = onload

(function()
{

    google.charts.load('current', {packages: ['corechart','line']});

    function drawChart()
    {
      var data = new google.visualization.DataTable();

      data.addColumn("string", "Year");
      data.addColumn("number", "Average");

      var rows = [];
      for(var i = 0; i < rawData.length; i++)
      {
        rows.push([""+rawData[i].Year, rawData[i].Amount]);
      }
      data.addRows(rows);

      var options = {
        hAxis: {
          title: "Year"
        },
        vAxis: {
          title: "Average Amount"
        },
        height: 300
      };

      var chart = new google.visualization.LineChart(document.getElementById("graph"));
      chart.draw(data, options);
    }

    google.charts.setOnLoadCallback(drawChart);

})();

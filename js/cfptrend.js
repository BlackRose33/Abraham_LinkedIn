(function(){

    google.charts.load('current', {packages: ['corechart','line']});

    function drawChart()
    {

      var data1 = new google.visualization.DataTable();
      var data2 = new google.visualization.DataTable();

      data1.addColumn("string", "Year");
      data1.addColumn("number", "Amount");

      data2.addColumn("string", "Year");
      data2.addColumn("number", "Participant Count");

      var rows1 = [];
      var rows2 = [];

      for(var i = 0; i < rawData.length; i++)
      {
        rows1.push([""+rawData[i].Year, rawData[i].Amount]);
        rows2.push([""+rawData[i].Year, rawData[i].Count]);
      }

      data1.addRows(rows1);
      data2.addRows(rows2);

      var options1 = {
        titlePosition: "none",
        hAxis: {
          title: "Year"
        },
        vAxis: {
          title: "Amount Matched"
        }
      };

      var options2 = {
        titlePosition: "none",
        hAxis: {
          title: "Year"
        },
        vAxis: {
          title: "Participant Count"
        }
      };

      var chart1 = new google.visualization.LineChart(document.getElementById("graph1"));
      chart1.draw(data1, options1);

      var chart2 = new google.visualization.LineChart(document.getElementById("graph2"));
      chart2.draw(data2, options2);
    }

    google.charts.setOnLoadCallback(drawChart);

})();

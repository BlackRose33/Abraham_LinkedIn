(function()
{

  if(!doChart)
  {
    return;
  }

  google.charts.load('current', {packages: ['corechart','line']});

  var indexMap = {
    "2001C" : 1,
    "2001E" : 2,
    "2003C" : 3,
    "2003E" : 4,
    "2005C" : 5,
    "2005E" : 6,
    "2009C" : 7,
    "2009E" : 8,
    "2013C" : 9,
    "2013E" : 10,
  };

  var rowMap = {
    1: ["Jan", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    2: ["Feb", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    3: ["Mar", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    4: ["Apr", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    5: ["May", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    6: ["Jun", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    7: ["Jul", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    8: ["Aug", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    9: ["Sep", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    10: ["Oct", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    11: ["Nov", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    12: ["Dec", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
  };

  var appearanceMap = {};
  var finalRows = [];

  for(var i = 0; i < rawData.length; i++)
  {
    var item = rawData[i];
    var year = item.Year;
    var month = item.Month;
    var amount = item.Amount;
    var key = year + ((item.IsContributions)? "C" : "E");
    if(indexMap[key])
    {
      appearanceMap[key] = true;
      if(month < 1 || month > 12)
      {
        // Dunno which month.. spread equally
        for(var j = 1; j <= 12; j++)
        {
          rowMap[j][indexMap[key]] += amount / 12;
        }
      }
      else
      {
        rowMap[month][indexMap[key]] += amount;
      }
    }
  }

  var years = [2001, 2003, 2005, 2009, 2013];

  function drawChart() {
    var dataContrib = new google.visualization.DataTable();
    var dataExpense = new google.visualization.DataTable();

    dataContrib.addColumn("string", "Month");
    dataExpense.addColumn("string", "Month");

    var rowsContrib = [["Jan"], ["Feb"], ["Mar"], ["Apr"], ["May"], ["Jun"],
      ["Jul"], ["Aug"], ["Sep"], ["Oct"], ["Nov"], ["Dec"]];
    var rowsExpense = [["Jan"], ["Feb"], ["Mar"], ["Apr"], ["May"], ["Jun"],
      ["Jul"], ["Aug"], ["Sep"], ["Oct"], ["Nov"], ["Dec"]];

    for(var i = 0; i < years.length; i++)
    {
      if(appearanceMap[years[i] + "C"])
      {
        for(var j = 0; j < 12; j++)
        {
          rowsContrib[j].push(rowMap[j+1][indexMap[years[i] + "C"]]);
        }
        dataContrib.addColumn("number", years[i]);
      }
      if(appearanceMap[years[i] + "E"])
      {
        for(var k = 0; k < 12; k++)
        {
          rowsExpense[k].push(rowMap[k+1][indexMap[years[i] + "E"]]);
        }
        dataExpense.addColumn("number", years[i]);
      }
    }

    dataContrib.addRows(rowsContrib);
    dataExpense.addRows(rowsExpense);

    var options = {
      titlePosition: "none",
      "hAxis": {
        title: "Month"
      },
      "vAxis": {
        title: "Amount"
      },
      "height" : 250
    };


    var options2 = {
      titlePosition: "none",
      "hAxis": {
        title: "Month"
      },
      "vAxis": {
        title: "Amount"
      },
      "height" : 250
    };

    var chart = new google.visualization.LineChart(document.getElementById('contrib_graph'));
    chart.draw(dataContrib, options);

    var chart2 = new google.visualization.LineChart(document.getElementById('expense_graph'));
    chart2.draw(dataExpense, options2);
  }

  google.charts.setOnLoadCallback(drawChart);

})();

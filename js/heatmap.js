$(document).ready(function() {

  var candidateList = [];
  var candidateMap = {};

  APIGetCandidateList(function(data)
  {
    for(var i = 0; i < data.length; i++)
    {
      var candidate = data[i];
      var li = document.createElement("li");
      li.className = "candidate-link";
      li.id = "candidate-" + candidate.ID;
      li.appendChild(document.createTextNode(candidate.FirstName + " " +
        candidate.LastName));
      li.onclick = function(candidateID)
      {
        window.location.href = candidateID;
      }.bind(window, candidate.ID);
      $("#candidateList").append(li);

      var key = candidate.FirstName + " " + candidate.LastName;
      candidateList.push(key);
      if(candidateMap[key])
      {
        candidateMap[key].push(candidate.ID);
      }
      else
      {
        candidateMap[key] = [candidate.ID];
      }
    }
  }, function()
  {
    alert("No candidate list found.");
  });

  window.getMatchingCandidates = function(query)
  {
    var results = [];
    var parts = query.toLowerCase().split(" ");

    mainLoop:
    for(var i = 0; i < candidateList.length; i++)
    {
      var checks = candidateList[i].toLowerCase().split(" ");
      for(var j = 0; j < checks.length; j++)
      {
        for(var k = 0; k < parts.length; k++)
        {
          if(checks[j].indexOf(parts[k]) > -1)
          {
            results.push(candidateList[i]);
            continue mainLoop;
          }
        }
      }
    }

    return results;
  };

  window.filter = function(elem)
  {
    var query = elem.value;
    if(query.length == 0)
    {
      $(".candidate-link").show();
      return;
    }
    var matches = getMatchingCandidates(query);
    $(".candidate-link").hide();
    for(var i = 0; i < matches.length; i++)
    {
      var ids = candidateMap[matches[i]];
      for(var j = 0; j < ids.length; j++)
      {
        $("#candidate-" + ids[j]).show();
      }
    }
  };

});

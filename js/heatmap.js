$(document).ready(function() {

  APIGetCandidateList(function(data)
  {
    for(var i = 0; i < data.length; i++)
    {
      var candidate = data[i];
      var li = document.createElement("li");
      li.appendChild(document.createTextNode(candidate.FirstName + " " +
        candidate.LastName));
      li.onclick = function(candidateID)
      {
        window.location.href = candidateID;
      }.bind(window, candidate.ID);
      $("#candidateList").append(li);
    }
  }, function()
  {
    alert("No candidate list found.");
  });

});

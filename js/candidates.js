(function($) {

  var getCandidateList = function(success, error)
  {
    $.ajax({
      url: "/api/candidates/",
      dataType: "json",
      success: success,
      error: error
    });
  };

  window.APIGetCandidateList = getCandidateList;

})( jQuery );

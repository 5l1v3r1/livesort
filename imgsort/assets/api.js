(function() {

  function addPair(lesser, greater, cb) {
    url = '/add?lesser=' + encodeURIComponent(lesser) + '&greater=' +
      encodeURIComponent(greater);
    apiCall(url, function(err, obj) {
      if (err) {
        cb(err);
      } else {
        cb(null, obj.pair[0], obj.pair[1]);
      }
    });
  }

  function getPair(cb) {
    apiCall('/pair', function(err, obj) {
      if (err) {
        cb(err);
      } else {
        cb(null, obj.pair[0], obj.pair[1]);
      }
    });
  }

  function apiCall(url, cb) {
    var req = new XMLHttpRequest();
    req.onload = function() {
      var obj = JSON.parse(req.responseText);
      if (obj.error) {
        cb(obj.error);
      } else {
        cb(null, obj);
      }
    };
    req.onerror = function() {
      cb('failed to connect');
    };
    req.open('GET', url);
    req.send(null);
  }

  window.sortAPI = {addPair: addPair, getPair: getPair};

})();

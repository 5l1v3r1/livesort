(function() {

  window.addEventListener('load', function() {
    var picker = new window.Picker();
    picker.onpick = function(lesser, greater) {
      window.sortAPI.addPair(lesser, greater, function(err, next1, next2) {
        if (err) {
          picker.showError(err);
        } else {
          picker.show(next1, next2);
        }
      });
    };

    window.sortAPI.getPair(function(err, next1, next2) {
      if (err) {
        picker.showError(err);
      } else {
        picker.show(next1, next2);
      }
    })
  });

})();

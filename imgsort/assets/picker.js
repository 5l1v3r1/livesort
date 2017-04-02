(function() {

  var LEFT_KEY = 37;
  var RIGHT_KEY = 39;

  function Picker() {
    this._element = document.getElementById('picker');
    this._error = document.getElementById('error');
    this._left = document.getElementById('image-left');
    this._right = document.getElementById('image-right');
    this._leftName = null;
    this._rightName = null;
    this._element.className = 'no-options';

    this.onpick = function() {};

    this._left.onclick = this._pick.bind(this, false);
    this._right.onclick = this._pick.bind(this, true);
    window.addEventListener('keydown', this._keyEvent.bind(this));
  }

  Picker.prototype.show = function(name1, name2) {
    if (!name1) {
      this.showError('list has been sorted');
      return;
    }
    this._element.className = 'options';
    this._left.src = '/image?name=' + encodeURIComponent(name1);
    this._right.src = '/image?name=' + encodeURIComponent(name2);
    this._leftName = name1;
    this._rightName = name2;
  };

  Picker.prototype.showError = function(err) {
    this._element.className = 'error';
    this._leftName = null;
    this._rightName = null;
    this._error.textContent = err;
  };

  Picker.prototype._pick = function(isRight) {
    var lesser, greater;
    if (isRight) {
      lesser = this._leftName;
      greater = this._rightName;
    } else {
      lesser = this._rightName;
      greater = this._leftName;
    }
    this._leftName = null;
    this._rightName = null;
    this._element.className = 'no-options';
    this.onpick(lesser, greater);
  };

  Picker.prototype._keyEvent = function(e) {
    if (!this._leftName) {
      return;
    }

    if (e.which === LEFT_KEY || e.which === RIGHT_KEY) {
      e.preventDefault();
      e.stopPropagation();
      this._pick(e.which === RIGHT_KEY);
    }
  };

  window.Picker = Picker;

})();

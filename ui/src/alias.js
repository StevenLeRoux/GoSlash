(function () {
  var alias = [];

  function render() {
    var list = document.getElementById('alias-list');

    alias.forEach(function (a) {
      var row = document.createElement('div');
      row.classList.add('alias');
      row.innerHTML = '<h3>' + a.alias + '</h3>' +
        '<p class="target">&nbsp;-&nbsp;' + a.target + '</p>' +
        '<p>' + a.description + '</p>' +
        '<span class="icon"><?xml version="1.0" encoding="UTF-8"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"><svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" width="32" height="32" viewBox="0 0 24 24"><path d="M4,11V13H16L10.5,18.5L11.92,19.92L19.84,12L11.92,4.08L10.5,5.5L16,11H4Z" /></svg></span>';

      row.addEventListener('click', function () {
        window.location.href = '/' + a.alias;
      });

      list.appendChild(row);
    });
  }

  var oReq = new XMLHttpRequest();
  oReq.onload = function (e) {
    alias = e.target.response;
    render();
  };
  oReq.open('GET', 'http://127.0.0.1:8080/dump?' + new Date().getTime(), true);
  oReq.responseType = 'json';
  oReq.send();
})();

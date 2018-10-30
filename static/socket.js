'use strict';
var conn = new WebSocket("ws://" + document.location.host + "/ws");
conn.onclose = function (evt) {
    console.log(evt);
};
conn.onmessage = function (evt) {
    console.log(evt);
};
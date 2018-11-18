'use strict';
const conn = new WebSocket("ws://" + document.location.host + "/ws");

conn.onmessage = function (evt) {
    const data = JSON.parse(evt.data);
    console.log(data);
};
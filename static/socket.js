'use strict';
const conn = new WebSocket("ws://" + document.location.host + "/ws");
conn.onclose = function (evt) {
    console.log(evt);
};
conn.onmessage = function (evt) {
    const data = JSON.parse(evt.data);
    const messageNode = document.createElement('div');
    messageNode.innerText = data.message;
    document.querySelector('#results').appendChild(messageNode);
};
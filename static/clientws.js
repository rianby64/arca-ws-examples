'use strict';
document.querySelector('form').addEventListener('submit', e => {
    e.preventDefault();
    const fd = new FormData(e.target);
    const message = fd.get('message');

    conn.send(message);
});
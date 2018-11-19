'use strict';
conn.onopen = () => {
    const message = {
        Jsonrpc: '2.0',
        Method: 'getUsers',
        ID: 'whatever'
    };
    conn.send(JSON.stringify(message));
}
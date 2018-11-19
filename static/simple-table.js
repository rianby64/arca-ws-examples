'use strict';
conn.onopen = () => {
    const message = {
        Jsonrpc: '2.0',
        Method: 'getUsers',
        id: 'id-1'
    };
    conn.send(JSON.stringify(message));
}
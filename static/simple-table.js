'use strict';

const message = {
    Jsonrpc: '2.0',
    Method: 'getTable',
    Params: {
        Message: 'this is my message',
        A: ['xx', 'yy', 'zz']
    },
    id: 'my id'
}

conn.onopen = () => {
    conn.send(JSON.stringify(message));
}
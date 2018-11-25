'use strict';
const conn = new WebSocket("ws://" + document.location.host + "/ws");
conn.messageHandlers = {};

conn.onmessage = (e) => {
    const data = JSON.parse(e.data);
    const handler = conn.messageHandlers[data.Context.source] || function() {};
    handler(data);
}

const blockEdit = (e) => {
    e.preventDefault();
    e.stopPropagation();
    e.stopImmediatePropagation();
};

// Need to implement some Redux here in this thing...
function setupTable(tableid, rowid, source, fields) {
    const table = document.querySelector(`#${tableid}`);
    const tbody = table.querySelector('tbody');
    const insertButton = table.querySelector('[action="insert"]');
    const tmplRow = document.querySelector(`[id="${rowid}"]`);
    const tmplCell = document.querySelector('[id="cell"]');

    const processCell = (row, key, data) => {
        const td = row.querySelector(`[key="${key}"]`);
        const cell = document.importNode(tmplCell.content, true);
        const span = cell.querySelector('span');
        const input = cell.querySelector('input');
        const form = cell.querySelector('form')

        span.textContent = data[key] ? data[key] : '-';
        input.value = data[key] ? data[key] : '';
        input.name = key;
        cell.querySelector('input[name="ID"]').value = data.ID;

        const toggleSpanToForm = () => {
            span.hidden = true;
            form.hidden = false;
        };

        span.addEventListener('click', toggleSpanToForm);

        form.addEventListener('submit', e => {
            e.preventDefault();
            span.hidden = false;
            form.hidden = true;

            const tr = td.closest('tr');
            tr.setAttribute('disabled', '');
            tr.addEventListener('click', blockEdit, true);
            span.removeEventListener('click', toggleSpanToForm);

            let fd;
            const data = new FormData(e.target).toJSON();
            if (Number(data.ID) > 0) {
                data.ID = Number(data.ID);
                span.textContent = data[key] ? data[key] : '-';

                fd = {
                    Jsonrpc: "2.0",
                    Method: 'update',
                    Params: data
                };
            } else {
                fd = {
                    Jsonrpc: "2.0",
                    Method: 'insert',
                    Params: data
                };
            }
            conn.send(JSON.stringify({...fd, 
                context: {
                    source
                }
            }));
        });

        td.appendChild(cell);
    };

    const processRow = (data, row = document.importNode(tmplRow.content, true)) => {
        row.querySelector('tr').setAttribute('ID', data.ID);
        fields.forEach(field => {
            processCell(row, field, data);
        });
        row.querySelector('[action="delete"]').addEventListener('click', e => {
            const id = Number(e.target.closest('tr').getAttribute('ID'));
            if (id > 0) {
                conn.send(JSON.stringify({
                    Jsonrpc: "2.0",
                    Method: 'delete',
                    Params: {
                        ID: id
                    },
                    context: {
                        source
                    }
                }));
            }
        }, { once: true });
        return row;
    }

    insertButton.insertingNew = false;
    insertButton.addEventListener('click', () => {
        if (!insertButton.insertingNew) {
            tbody.appendChild(processRow({}));
        }
        insertButton.insertingNew = true;
    });

    const onmessage = (data) => {
        const result = data.Result;
        if (data.Method === 'delete') {
            let row = tbody.querySelector(`tr[id="${result.ID}"]`);
            if (row) {
                row.remove();
            }
            return
        }
        if (data.ID === `id-for-${source}`) {
            result.forEach(element => tbody.appendChild(processRow(element)));
        } else {
            let row = tbody.querySelector(`tr[id="${result.ID}"]`);
            if (!row) {
                row = tbody.querySelector(`tr[id="undefined"]`);
                if (!row) {
                    return tbody.appendChild(processRow(result));
                }
                row.setAttribute('ID', result.ID)
                insertButton.insertingNew = false;
            }
            row.removeEventListener('click', blockEdit, true);
            row.removeAttribute('disabled');
            Object.keys(result).filter(key => key != 'ID').forEach(key => {
                const cell = row.querySelector(`[key="${key}"]`);
                cell.innerHTML = '';
                processCell(row, key, result);
            });
        }
    }
    conn.messageHandlers[source] = onmessage;

    const message = {
        Jsonrpc: '2.0',
        Method: 'read',
        ID: `id-for-${source}`,
        context: {
            source
        }
    };
    conn.send(JSON.stringify(message));
}

conn.onopen = () => {
    setupTable('goods', 'good-row', 'Goods', ['Description', 'Price']);
    setupTable('users', 'user-row', 'Users', ['Name', 'Email']);
}
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
function setupTable(tableid, rowid, source, fields, convertFn) {
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
            const converted = Object.keys(data).reduce((acc, key) => {
                acc[key] = convertFn[key](data[key]);
                return acc;
            }, {});
            if (data.ID.toString() != 'undefined') {
                span.textContent = data[key] ? data[key] : '-';
                fd = {
                    Method: 'update',
                    Params: converted
                };
            } else {
                fd = {
                    Method: 'insert',
                    Params: converted
                };
            }
            console.log({...fd,
                Context: {
                    source
                }
            });
            /*
            conn.send(JSON.stringify({...fd,
                Context: {
                    source
                }
            }));
            */
        });

        td.appendChild(cell);
    };

    const processRow = (data, row = document.importNode(tmplRow.content, true)) => {
        row.querySelector('tr').setAttribute('ID', data.ID);
        fields.forEach(field => {
            processCell(row, field, data);
        });
        row.querySelector('[action="delete"]').addEventListener('click', e => {
            const id = e.target.closest('tr').getAttribute('ID');
            if (id > 0) {
                conn.send(JSON.stringify({
                    Method: 'delete',
                    Params: {
                        ID: id
                    },
                    Context: {
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

    conn.messageHandlers[source] = (data) => {
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
    };
    conn.send(JSON.stringify({
        Method: 'read',
        ID: `id-for-${source}`,
        Context: {
            source
        }
    }));
}

conn.onopen = () => {
    setupTable(
        'AAU',
        'AAU-row',
        'AAU',
        ['ID', 'Parent', 'Expand', 'Description', 'Qop'],
        {
            "ID": String,
            "Parent": String,
            "Expand": Boolean,
            "Description": String,
            "Qop": Number,
        }
    );
}

document.querySelector('#subir').addEventListener('click', e => {
    conn.send(JSON.stringify({
        Method: "insert",
        Params: {
          Description: "Una descripcion",
          Parent: "-"
        },
        Context: {
          source: "AAU"
        }
    }));
})
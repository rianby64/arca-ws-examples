'use strict';
const conn = new WebSocket("ws://" + document.location.host + "/ws");
conn.messageHandlers = {};

conn.onmessage = (e) => {
    const data = JSON.parse(e.data);
    const handler = conn.messageHandlers[data.Context.Source] || function() {};
    handler(data);
}

const blockEdit = (e) => {
    e.preventDefault();
    e.stopPropagation();
    e.stopImmediatePropagation();
};

// Need to implement some Redux here in this thing...
function setupTable(tableid, rowid, Source, fields, convertFn) {
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
            conn.send(JSON.stringify({...fd,
                Context: {
                    Source
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
            const id = e.target.closest('tr').getAttribute('ID');
            if (convertFn["ID"](id) > 0) {
                conn.send(JSON.stringify({
                    Method: 'delete',
                    Params: {
                        ID: convertFn["ID"](id)
                    },
                    Context: {
                        Source
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

    conn.messageHandlers[Source] = (data) => {
        const result = data.Result;
        if (data.Method === 'delete') {
            let row = tbody.querySelector(`tr[id="${result.ID}"]`);
            if (row) {
                row.remove();
            }
            return
        }
        if (data.ID === `id-for-${Source}`) {
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
            Object.keys(result).filter(key => key != 'ID' && key != 'CreatedAt').forEach(key => {
                const cell = row.querySelector(`[key="${key}"]`);
                cell.innerHTML = '';
                processCell(row, key, result);
            });
        }
    };
    conn.send(JSON.stringify({
        Method: 'read',
        ID: `id-for-${Source}`,
        Context: {
            Source
        }
    }));
}

conn.onopen = () => {
    setupTable('Table1', 'Table1-row', 'Table1',
        ['ID', 'Num1', 'Num2'],
        {
            "ID": Number,
            "Num1": Number,
            "Num2": Number,
        }
    );
    setupTable('Table2', 'Table2-row', 'Table2',
        ['ID', 'Num3', 'Num4'],
        {
            "ID": Number,
            "Num3": Number,
            "Num4": Number,
        }
    );
    setupTable('ViewSum1', 'ViewSum1-row', 'ViewSum1',
        ['ID', 'Table1Num1', 'Table2Num3', 'Sum13'],
        {
            "ID": String,
            "Table1Num1": Number,
            "Table2Num3": Number,
            "Sum13": Number,
        }
    );
    setupTable('ViewSum2', 'ViewSum2-row', 'ViewSum2',
        ['ID', 'Table1Num2', 'Table2Num4', 'Sum24'],
        {
            "ID": String,
            "Table1Num2": Number,
            "Table2Num4": Number,
            "Sum24": Number,
        }
    );
}
'use strict';
const conn = new WebSocket("ws://" + document.location.host + "/ws");

// Need to implement some Redux here in this thing...
(() => {
    const table = document.querySelector('#mytable');
    const tbody = document.querySelector('tbody');
    const insertButton = table.querySelector('[action="insert"]');
    const tmplRow = document.querySelector('[id="user-row"]');
    const tmplCell = document.querySelector('[id="cell"]');

    const processCell = (row, key, data) => {
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

            span.classList.add('disabled');
            span.removeEventListener('click', toggleSpanToForm);

            const fd = {
                Jsonrpc: "2.0",
                Id: 'whatever',
                Method: 'updateUser',
                Params: new FormData(e.target).toJSON()
            };
            fd.Params.ID = Number(fd.Params.ID);
            span.textContent = fd.Params[key] ? fd.Params[key] : '-';
            conn.send(JSON.stringify(fd));
        });

        row.querySelector(`[key="${key}"]`).appendChild(cell);
    };

    const processRow = (data) => {
        const row = document.importNode(tmplRow.content, true);
        processCell(row, 'Name', data);
        processCell(row, 'Email', data);
        return row;
    }

    insertButton.insertingNew = false;
    insertButton.addEventListener('click', () => {
        if (!insertButton.insertingNew) {
            tbody.appendChild(processRow({}));
        }
        insertButton.insertingNew = true;
    });

    function onmessage(e) {
        const data = JSON.parse(e.data);
        const result = data.Result;
        result.forEach(element => tbody.appendChild(processRow(element)));
    }

    conn.onmessage = onmessage;
})();
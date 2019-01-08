'use strict';
const fs = require('fs');
const Ini = require('ini');
const express = require('express');
const proxy = require('http-proxy-middleware');
const app = express();
const config = Ini.parse(fs.readFileSync('./config.ini', 'utf-8'));



const arcaws = proxy('/arca-node', config.arcaws);
app.use('/arca-node', arcaws);

//const sockjs = proxy('/sockjs-node', config.sockjs);
//app.use(sockjs);
//app.use('/sockjs-node', sockjs);

const staticws = proxy(['!/arca-node', '/**'], config.static);
app.use(staticws);

console.log(`Listening to http://localhost:${config.port}`);
const server = app.listen(Number(config.port));
//server.on('upgrade', staticws.upgrade);
server.on('upgrade', arcaws.upgrade);
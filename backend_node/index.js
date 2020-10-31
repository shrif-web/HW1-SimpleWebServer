const express = require('express');
const app = express();
const cors = require("cors");
const port = 3000;


app.use(express.json())
app.use(cors());

var crypto = require('crypto');
var hash = crypto.createHash('sha256');

app.get("/nodejs/write", (req, res) => {
    var myline = req.query.line;
    if (myline < 1 || myline > 100){
        return res.status(400).end('the value is invalid!');
    }
    //TODO: changing the directory to a flexible one
    const nthline = require('nthline')
    , filePath = '/Users/alismac/Documents/CE_T7/Web\ Programming/HWs/1\ -\ SimpleWebServer/backend_node/test.txt'
    , rowNumber = myline - 1
    nthline(rowNumber, filePath)
    .then(line => res.end(line))
  });

  app.post("/nodejs/sha256", (req, res) => {
    var sum = req.body["1"] + req.body["2"];
    var code = crypto.createHash('sha256').update(sum.toString()).digest('hex');
    console.log(code);
    res.json({
        "result" : code
    });
  });
  

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);

});
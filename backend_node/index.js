const express = require('express');
const app = express();
const cors = require("cors");
const port = 3000;


app.use(express.json())
app.use(cors());

var crypto = require('crypto');
var hash = crypto.createHash('sha256');

app.get("/nodejs/write", (req, res) => {
    console.log(req.query.line);
    res.end('Hello World!');
  });

  app.post("/nodejs/sha256", (req, res) => {
    //console.log(req.body["1"]);
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
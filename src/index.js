const express = require('express');
const app = express();

const instanceId = Math.floor(Math.random() * 100000).toString(16);

app.get('/', function (req, res) {
  res.send(`Hello world from ${instanceId}!`);
});

const port = parseInt(process.env["PORT"]) || 3000;
app.listen(port, function () {
  console.log(`Instance ${instanceId} Listening on port ${port}...`);
});

["SIGINT", "SIGTERM"].forEach((s) => {
    process.on(s, function() {
    console.log(`Shutting down ${instanceId}...`);
    process.exit(0);
  });
});

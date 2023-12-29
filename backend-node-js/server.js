const express = require('express');
const bodyParser = require('body-parser');

const app = express();
const port = 3000;

app.use(bodyParser.json());

function performOperation(req, res, operationFn) {
  const { num1, num2 } = req.body;
  const result = operationFn(num1, num2);
  res.json({ result });
}

// Version endpoint
app.get('/version', (req, res) => {
  revision = require('child_process')
  .execSync('git rev-parse HEAD')
  .toString().trim()
  res.send("app_version: "+ revision)
})

// Healthcheck endpoint
app.get('/healthz', (req, res) => {
  const message = "The Node.js service is healthy";
  res.send(message);
});

app.post('/api/add', (req, res) => {
  performOperation(req, res, (a, b) => a + b);
});

app.post('/api/subtract', (req, res) => {
  performOperation(req, res, (a, b) => a - b);
});

app.post('/api/multiply', (req, res) => {
  performOperation(req, res, (a, b) => a * b);
});

app.post('/api/divide', (req, res) => {
  const { num1, num2 } = req.body;
  if (num2 !== 0) {
    performOperation(req, res, (a, b) => a / b);
  } else {
    res.status(400).json({ error: 'Cannot divide by zero' });
  }
});

app.listen(port, () => {
  console.log(`Node.js backend listening at http://localhost:${port}`);
});

module.exports = app
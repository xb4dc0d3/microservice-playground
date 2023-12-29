// server.test.js
const request = require('supertest');
const app = require('../server'); // Import the Node.js app

describe('Healthcheck Endpoint', () => {
  it('should return healthcheck string', function (done){
      request(app)
      .get('/healthz')
      .expect("The Node.js service is healthy")
      .end(done);
  });
});

describe('Arithmetic Operations', () => {
  it('should add two numbers', function (done){
      request(app)
      .post('/api/add')
      .send({ num1: 3, num2: 5 })
      .expect({result: 8})
      .end(done);
  });
});

describe('Arithmetic Operations', () => {
  it('should subtract two numbers', function (done){
      request(app)
      .post('/api/subtract')
      .send({ num1: 3, num2: 5 })
      .expect({result: -2})
      .end(done);
  });
});

describe('Arithmetic Operations', () => {
  it('should multiply two numbers', function (done){
      request(app)
      .post('/api/multiply')
      .send({ num1: 3, num2: 5 })
      .expect({result: 15})
      .end(done);
  });
});

describe('Arithmetic Operations', () => {
  it('should add divide two numbers', function (done){
      request(app)
      .post('/api/divide')
      .send({ num1: 3, num2: 5 })
      .expect({result: 0.6})
      .end(done);
});
});

const { gql } = require("apollo-server-express");
const user = require("./user");
const post = require("./post");

const base = gql`
  type PingResponse {
    message: String
    _version: Int
  }

  type Query {
    ping: PingResponse!
  }
`;

module.exports = [base, user, post];

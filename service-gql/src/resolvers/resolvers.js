const _ = require("lodash");
const user = require("./user/user");
const post = require("./post/post");

const base = {
  Query: {
    ping: () => ({ message: "ping", _version: 1 }),
  },
};

module.exports = _.merge(base, user, post);

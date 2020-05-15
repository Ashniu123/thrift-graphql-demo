const thrift = require("thrift");
const UserService = require("../thrift/user/UserService");
const logger = require("../utils/logger");

const SERVER_HOST =
  process.env.SERVICE_USER_HOST ||
  process.env.SERVICE_USER_CLUSTER_IP_SERVICE_SERVICE_HOST;
const SERVER_PORT = parseInt(process.env.SERVICE_USER_PORT);

logger.info(`userClient: ${SERVER_HOST} ${SERVER_PORT}`);

const thriftOptions = {
  transport: thrift.TBufferedTransport,
  protocol: thrift.TCompactProtocol,
};

const connection = thrift.createConnection(
  SERVER_HOST,
  SERVER_PORT,
  thriftOptions
);
let client;

connection.on("error", (err) => {
  logger.error(`userClient: ${JSON.stringify(err)}`);
});

connection.on("connect", () => {
  client = thrift.createClient(UserService, connection);
  logger.info("userClient: Connected to thrift server!");
});

connection.on("close", () => {
  logger.info("userClient: Disconnected from thrift server!");
  // process.exit(1);
});

process.on("SIGTERM", connection.end);

/**
 * thrift user client
 * @param {string} func thrift function to call
 * @param {object[]} params params to pass to the thrift function
 */
const userClient = (func, params) =>
  new Promise((resolve, reject) => {
    client[func](...params)
      .then(resolve)
      .fail(reject);
  });
// Apache uses 'Q' promise library
// See https://issues.apache.org/jira/browse/THRIFT-2376

module.exports = userClient;

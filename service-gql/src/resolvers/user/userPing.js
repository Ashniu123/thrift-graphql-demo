const userClient = require("../../thriftClients/user");

module.exports = async () => {
  try {
    const { message, version: _version } = await userClient("ping", []);
    return { message, _version };
  } catch (e) {
    throw e;
  }
};

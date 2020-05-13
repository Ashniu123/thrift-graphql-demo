const postClient = require("../../thriftClients/post");

module.exports = async () => {
  try {
    const { message, version: _version } = await postClient("ping", []);
    return { message, _version };
  } catch (e) {
    throw e;
  }
};

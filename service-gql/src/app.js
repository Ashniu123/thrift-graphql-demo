require("dotenv/config");
const express = require("express");
const compression = require("compression");
const bodyParser = require("body-parser");
const helmet = require("helmet");
const { ApolloServer } = require("apollo-server-express");
const typeDefs = require("./typeDefs/typeDefs");
const resolvers = require("./resolvers/resolvers");
const logger = require("./utils/logger");

const app = express();

app.set("port", process.env.GRAPHQL_PORT);
app.use(compression());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(helmet());

app.get("/healthz", (_req, res) => {
  res.json({
    health: "ok",
    version: 1,
  });
});

const apolloServer = new ApolloServer({
  typeDefs,
  resolvers,
  formatResponse: (response) => {
    logger.info(JSON.stringify(response));
    return response;
  },
  formatError: (error) => {
    logger.error(JSON.stringify(error));
    return error;
  },
});

apolloServer.applyMiddleware({ app, path: "/" });

/** Start Express server. */
const server = app.listen(app.get("port"), () => {
  console.log(
    "  App is running at http://localhost:%d in %s mode",
    app.get("port"),
    app.get("env")
  );
  console.log("  Press CTRL-C to stop\n");
});

module.exports = { app, server };

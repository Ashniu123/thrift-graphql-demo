const { createLogger, format, transports } = require("winston");

const { combine, timestamp, printf, label } = format;

/** Custom Formatter */
const myFormat = printf(
  (info) => `(${info.label}) [${info.timestamp}] ${info.level}: ${info.message}`
);

const logger = createLogger({
  format: combine(label({ label: process.pid }), timestamp(), myFormat),
  transports: [new transports.Console()],
  exitOnError: false,
  exceptionHandlers: [new transports.Console()],
});

module.exports = logger;

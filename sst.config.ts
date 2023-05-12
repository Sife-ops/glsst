import { SSTConfig } from "sst";
import { Api, Function, Table } from "sst/constructs";

const { BOT_PUBLIC_KEY, BOT_APP_ID } = process.env;

export default {
  config(_input) {
    return {
      name: "glsst",
      region: "us-east-1",
    };
  },
  stacks(app) {
    app.setDefaultFunctionProps({
      runtime: "go1.x",
    });
    app.stack(function Stack({ stack }) {
      const table = new Table(stack, "table", {
        fields: {
          pk: "string",
          sk: "string",
          gsi1pk: "string",
          gsi1sk: "string",
          gsi2pk: "string",
          gsi2sk: "string",
          gsi3pk: "string",
          gsi3sk: "string",
          gsi4pk: "string",
          gsi4sk: "string",
          gsi5pk: "string",
          gsi5sk: "string",
        },
        primaryIndex: {
          partitionKey: "pk",
          sortKey: "sk",
        },
        globalIndexes: {
          gsi1: {
            partitionKey: "gsi1pk",
            sortKey: "gsi1sk",
          },
          gsi2: {
            partitionKey: "gsi2pk",
            sortKey: "gsi2sk",
          },
          gsi3: {
            partitionKey: "gsi3pk",
            sortKey: "gsi3sk",
          },
          gsi4: {
            partitionKey: "gsi4pk",
            sortKey: "gsi4sk",
          },
          gsi5: {
            partitionKey: "gsi5pk",
            sortKey: "gsi5sk",
          },
        },
      });

      const consumerFn = new Function(stack, "consumerFn", {
        handler: "functions/bot/consumer/main.go",
        bind: [table],
        environment: {
          BOT_APP_ID: BOT_APP_ID || "REE",
          BOT_PUBLIC_KEY: BOT_PUBLIC_KEY || "REE",
          TABLE_NAME: table.tableName,
        },
      });

      const api = new Api(stack, "api", {
        routes: {
          "POST /bot": {
            function: {
              handler: "functions/bot/receiver/main.go",
              bind: [consumerFn, table],
              environment: {
                BOT_PUBLIC_KEY: BOT_PUBLIC_KEY || "REE",
                CONSUMER_FN: consumerFn.functionName,
              },
            },
          },
        },
      });

      stack.addOutputs({
        ApiEndpoint: api.url,
        Table: table.tableName,
      });
    });
  },
} satisfies SSTConfig;

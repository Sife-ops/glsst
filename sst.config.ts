import { SSTConfig } from "sst";
import { Api, Function, Table, StaticSite } from "sst/constructs";

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

      const importFn = new Function(stack, "importFn", {
        handler: "functions/import/main.go",
        bind: [table],
        environment: {
          TABLE_NAME: table.tableName,
        },
      });

      const mnemonicFn = new Function(stack, "mnemonicFn", {
        handler: "workspaces/functions/mnemonic.handler",
        runtime: "nodejs18.x",
      });

      const api = new Api(stack, "api", {
        routes: {
          "POST /user": {
            function: {
              handler: "functions/rest/user/main.go",
              bind: [table],
              environment: {
                TABLE_NAME: table.tableName,
              },
            },
          },
        },
      });

      const site = new StaticSite(stack, "site", {
        path: "workspaces/web",
        buildCommand: "npm run build",
        buildOutput: "dist",
        environment: {
          VITE_API_URL: api.url,
        },
      });

      const consumerFn = new Function(stack, "consumerFn", {
        handler: "functions/bot/consumer/main.go",
        bind: [table, mnemonicFn],
        environment: {
          BOT_APP_ID: BOT_APP_ID || "REE",
          BOT_PUBLIC_KEY: BOT_PUBLIC_KEY || "REE",
          MNEMONIC_FN: mnemonicFn.functionName,
          SITE_URL: site.url || "http://localhost:5173",
          TABLE_NAME: table.tableName,
        },
      });

      api.addRoutes(stack, {
        "POST /bot": {
          function: {
            handler: "functions/bot/receiver/main.go",
            bind: [consumerFn, table],
            environment: {
              BOT_PUBLIC_KEY: BOT_PUBLIC_KEY || "REE",
              CONSUMER_FN: consumerFn.functionName,
              SITE_URL: site.url || "http://localhost:5173",
              TABLE_NAME: table.tableName,
            },
          },
        },
      });

      stack.addOutputs({
        ApiEndpoint: api.url,
        Table: table.tableName,
        Site: site.url,
      });
    });
  },
} satisfies SSTConfig;

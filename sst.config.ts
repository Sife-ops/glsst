import { SSTConfig } from "sst";
import { Api } from "sst/constructs";

const { BOT_PUBLIC_KEY } = process.env;

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
      const api = new Api(stack, "api", {
        routes: {
          "POST /": {
            function: {
              handler: "functions/lambda/main.go",
              environment: {
                BOT_PUBLIC_KEY: BOT_PUBLIC_KEY || "REEEEEEEEEE"
              }
            },
          },
        },
      });
      stack.addOutputs({
        ApiEndpoint: api.url,
      });
    });
  },
} satisfies SSTConfig;

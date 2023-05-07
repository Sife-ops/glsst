import { SSTConfig } from "sst";
import { Api, Function } from "sst/constructs";

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

      const consumerFn = new Function(stack, "consumerFn", {
        handler: "functions/lambda/consumer/consumer.go",
        // bi
      })

      const api = new Api(stack, "api", {
        routes: {
          "POST /": {
            function: {
              handler: "functions/lambda/main/main.go",
              bind: [consumerFn],
              environment: {
                BOT_PUBLIC_KEY: BOT_PUBLIC_KEY || "REEEEEEEEEE",
                CONSUMER_FN: consumerFn.functionName
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

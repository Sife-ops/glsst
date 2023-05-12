export const prediction = {
  name: "create",
  description: "Make a prediction.",
  options: [
    {
      name: "condition",
      description: "The condition(s) for the prediction.",
      type: 3,
      required: true,
    },
    {
      name: "judge",
      description: "The default judge for the prediction.",
      type: 6,
      required: true,
    },
  ],
};

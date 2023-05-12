import { faker } from "@faker-js/faker";

export const handler = async (): Promise<string> => {
  const a = faker.word.adjective();
  const b = faker.word.adjective();
  const c = faker.word.noun();
  return `${a}-${b}-${c}`;
};

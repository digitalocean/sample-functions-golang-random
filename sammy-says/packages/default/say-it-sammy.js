const hello = (name) => `Hello ${name}! I am Sammy!`;
const water = (name) => `Oh ${name} I sure love the ocean`;
const servers = (name) => `Isn't it fun to not think about servers`;
const bored = (name) => `Swimming in the ocean is never boring`;
const prey = (name) =>
  `Did you know that tiger sharks are ecologically important predators of sea turtles and snakes`;

function main(args) {
  const phraseMakers = [hello, water, servers, bored, prey];
  var chosenPhraseMaker =
    phraseMakers[Math.floor(Math.random() * phraseMakers.length)];
  const name = args.name || "stranger";
  const greeting = chosenPhraseMaker(name);
  return { body: { greeting } };
}

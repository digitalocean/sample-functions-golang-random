const random = require('random')

const hello = (name) => `Hello ${name}! I am Sammy!`
const water = (name) => `Oh ${name} I sure love the ocean`
const servers = (name) => `Isn't it fun to not think about servers`
const bored = (name) => `Swimming in the ocean is never boring`
const prey = (name) => `Did you know that tiger sharks are ecologically important predators of sea turtles and snakes`

function main(args) {
  const phraseMakers = [hello, water, servers, bored, prey]
  // https://github.com/transitive-bullshit/random/blob/e11a840a1cfe0f5bd9c43640f9645a0b28f61406/src/random.js#L136-L148
  // random.Int has an inclusive upper-bound
  var chosenPhraseMaker = phraseMakers[random.int(0, phraseMakers.length-1)]
  const name = args.name || 'stranger'
  const greeting = chosenPhraseMaker(name)
  return { "body": { greeting } }
}

exports.main = main;

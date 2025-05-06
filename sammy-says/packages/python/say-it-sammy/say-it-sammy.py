import random

def main(args):
  quotes = [
    "I am Sammy!",
    "I sure love the ocean",
    "Isn't it fun to not think about servers",
    "Swimming in the ocean is never boring",
    "Did you know that tiger sharks are ecologically important predators of sea turtles and snakes"
  ]
  quote = quotes[random.randint(0, len(quotes)-1)]
  name = args.get("name", "stranger")
  greeting = "Hello " + name + "!" + " " + quote
  print(greeting)
  return {"body": greeting}

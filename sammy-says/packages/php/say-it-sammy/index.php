<?php

function main(array $args) : array
{
  $quotes = [
    "I am Sammy!",
    "I sure love the ocean",
    "Isn't it fun to not think about servers",
    "Swimming in the ocean is never boring",
    "Did you know that tiger sharks are ecologically important predators of sea turtles and snakes"
  ];
  $quote = $quotes[rand(0, count($quotes)-1)];

  $name = $args["name"] ?? "stranger";

  $greeting = "Hello {$name}! {$quote}.";
  echo $greeting;

  return [
      'body' => $greeting,
  ];
}

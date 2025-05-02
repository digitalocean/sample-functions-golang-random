package main

import "math/rand"

func Main(args map[string]interface{}) map[string]interface{} {
	quotes := []string{
		"I am Sammy!",
		"I sure love the ocean",
		"Isn't it fun to not think about servers",
		"Swimming in the ocean is never boring",
		"Did you know that tiger sharks are ecologically important predators of sea turtles and snakes",
	}

	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!" + " " + quotes[rand.Intn(len(quotes))]

	return msg
}

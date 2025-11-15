package main

import (
	perceptron "github.com/yourusername/myproject/Perceptron"
)




func main() {
	w1,w2,bias:=perceptron.SimplePerceptron()
	perceptron.TestPerceptron(w1,w2,bias)

	
}
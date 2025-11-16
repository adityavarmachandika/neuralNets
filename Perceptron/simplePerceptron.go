	package perceptron



	import (
		"fmt"
	)

	//activation function
	func binaryStepFunction(value float64)(float64){
		if value>=0{
			value=1;
		}else{
			value=0;
		}

		return value;
	}


	//prediction calculation function
	func simpleWeightedSum(x1 float64, x2 float64, w1 float64, w2 float64,bias float64)(float64){

		res:=float64(x1)*w1+float64(x2)*w2+bias;
		return res;
	}


	//weights learning calculation
	func simpleLearningRule(weight float64,err float64, learningRate float64, input int)(float64){
		weight+=err*learningRate*float64(input)
		return weight
	}
	
	func SimplePerceptron()(w1 float64,w2 float64,bias float64){
		epoch:=0;

		input1:=[4]int{1,1,0,0};
		input2:=[4]int{1,0,1,0};
		output:=[4]int{1,0,0,0};
		lr:=0.1
		isupdated:=false
		trained:=false

		for !trained{
			epoch++;
			isupdated=false
			for  i:=0;i<4;i++{

				pred:=simpleWeightedSum(float64(input1[i]),float64(input2[i]),w1,w2,bias)
				pred=binaryStepFunction(pred)
				if float64(output[i])!=pred{

					diff:=float64(output[i])-pred

					w1=simpleLearningRule(w1,diff,lr,input1[i])
					w2=simpleLearningRule(w2,diff,lr,input2[i])
					bias+=diff*lr;

					fmt.Printf("epoch %d, w1 %f, w2 %f, bias %f\n", epoch,w1,w2,bias)
					isupdated=true
					break;
				}
				
			}

			if !isupdated{
				trained=true
			}

		}


		fmt.Println(w1)
		fmt.Println(w2)
		fmt.Println(bias)
		return 
	}

	func TestPerceptron(w1, w2, bias float64) {

		tests := [4][2]float64{
			{1, 1},
			{1, 0},
			{0, 1},
			{0, 0},
		}

		expected := [4]float64{1, 0, 0, 0}

		fmt.Println("\nTesting perceptron...")

		for i := 0; i < 4; i++ {
			x1 := tests[i][0]
			x2 := tests[i][1]

			sum := x1*w1 + x2*w2 + bias

			
			pred := 0.0
			if sum >= 0 {
				pred = 1
			}

			status := "WRONG"
			if pred == expected[i] {
				status = "CORRECT"
			}

			fmt.Printf("Input (%v, %v) → Predicted: %v  Expected: %v  → %s\n",
				x1, x2, pred, expected[i], status)
		}
	}
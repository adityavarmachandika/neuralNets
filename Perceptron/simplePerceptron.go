package perceptron



import (
	"fmt"
)



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
			pred:=float64(input1[i])*w1+float64(input2[i])*w2+bias;
			if pred>=0{
				pred=1;
			}else{
				pred=0;
			}

			if float64(output[i])!=pred{

				diff:=float64(output[i])-pred

				w1+=diff*lr*float64(input1[i]);
				w2+=diff*lr*float64(input2[i]);
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
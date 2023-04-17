package b_oops.constructor.chaning;

/*To */

public class ConstructorChaining3 {

	static class Print {

		Print(int number) {
			System.out.println("Numbers addition is  = " + number);
		}

	}

	static class MathOperation extends Print {

		MathOperation(int number1, int number2) {
			super(number1 + number2);
		}

	}

	public static void main(String[] args) {

		MathOperation operation = new MathOperation(2, 4);

	}

}

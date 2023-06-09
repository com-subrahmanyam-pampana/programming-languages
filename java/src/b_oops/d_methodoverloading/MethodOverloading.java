package b_oops.d_methodoverloading;

/*Method overloading by changing datatype of argument*/
public class MethodOverloading {

	static int add(int a, int b) {
		return a + b;
	}

	static double add(double a, double b) {
		return a + b;
	}

	public static void main(String[] args) {
		System.out.println(add(11, 11));
		System.out.println(add(12.3, 12.6));
	}

}

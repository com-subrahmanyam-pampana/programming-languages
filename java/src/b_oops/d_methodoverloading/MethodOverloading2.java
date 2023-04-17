package b_oops.d_methodoverloading;


/*Method overloading by changing no of argumentd*/
public class MethodOverloading2 {

	static int add(int a, int b) {
		return a + b;
	}

	static int add(int a, int b, int c) {
		return a + b + c;
	}

	public static void main(String[] args) {
		System.out.println(add(11, 11));
		System.out.println(add(11, 11, 11));
	}

}